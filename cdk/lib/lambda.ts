import {
  Duration,
  Fn,
  Stack,
  StackProps,
  aws_apigateway as apigw,
  aws_ec2 as ec2,
  aws_iam as iam,
  aws_lambda as lambda,
} from "aws-cdk-lib";

import { Construct } from "constructs";
import path = require("path");

export interface AppProps extends StackProps {
  projectName: string;
  appName: string;
  deployEnvironment: string;
  appBuildPath: string;
}

export class PostApp extends Stack {
  constructor(scope: Construct, id: string, props: AppProps) {
    super(scope, id, props);

    // IAM
    const lambdaServiceRole = new iam.Role(this, `LambdaServiceRole`, {
      roleName: `${props.projectName}-${props.deployEnvironment}-lambda-role`,
      description: `Generally used in ${props.projectName}-${props.deployEnvironment} Lambda deployed in VPC.`,
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaVPCAccessExecutionRole"
        ),
      ],
    });

    // VPC
    // https://docs.aws.amazon.com/cdk/v2/guide/tokens.html
    // Fn.importValueでは、返却されるのは値を指すトークンのため、fromVpcAttributesを利用する
    const vpc = ec2.Vpc.fromVpcAttributes(this, "vpc", {
      vpcId: Fn.importValue(`shared-vpc-${props.deployEnvironment}-Vpc`),
      // 動的に指定不可かつデプロイ環境ごとに共通のため直接指定する
      availabilityZones: ["ap-northeast-1a", "ap-northeast-1c"],
      publicSubnetIds: [
        Fn.importValue(`shared-vpc-${props.deployEnvironment}-PublicSubnet1`),
        Fn.importValue(`shared-vpc-${props.deployEnvironment}-PublicSubnet2`),
      ],
      privateSubnetIds: [
        Fn.importValue(`shared-vpc-${props.deployEnvironment}-PrivateSubnet1`),
        Fn.importValue(`shared-vpc-${props.deployEnvironment}-PrivateSubnet2`),
      ],
    });

    const lambdaSecurityGroup = new ec2.SecurityGroup(
      this,
      `LambdaSecurityGroup`,
      {
        vpc,
        securityGroupName: `${props.projectName}-${props.deployEnvironment}-lambda`,
        description: `Generally used in ${props.projectName}-${props.deployEnvironment} Lambda deployed in VPC.`,
        allowAllOutbound: true,
      }
    );

    // Lambda post-app
    // 将来的に以下に移行を想定
    // https://docs.aws.amazon.com/cdk/api/v2/docs/aws-lambda-go-alpha-readme.html
    const lambdaApp = new lambda.Function(this, `LambdaApp`, {
      functionName: `${props.projectName}-${props.appName}-${props.deployEnvironment}`,
      description: `Lambda that executes the ${props.projectName}-${props.appName}-${props.deployEnvironment} API.`,
      runtime: lambda.Runtime.PROVIDED_AL2023,
      memorySize: 512,
      timeout: Duration.seconds(30),
      role: lambdaServiceRole,
      logRetention: 180,
      vpc: vpc,
      securityGroups: [lambdaSecurityGroup],
      handler: "bootstrap",
      code: lambda.Code.fromAsset(
        path.join(__dirname, `../../${props.appName}`),
        {
          // https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-package.html
          // https://iret.media/79515
          bundling: {
            image: lambda.Runtime.PROVIDED_AL2023.bundlingImage,
            command: [
              "bash",
              "-c",
              // デフォルトではルートディレクトリにGOCACHEとGOPATHを作成してパーミッションではじかれるため変更
              `export GOCACHE=/tmp/go-cache && export GOPATH=/tmp/go-path && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o /asset-output/bootstrap ${props.appBuildPath}`,
            ],
          },
        }
      ),
    });

    // APIGateway
    const apiGateway = new apigw.RestApi(this, `ApiGateway`, {
      restApiName: `${props.projectName}-${props.appName}-${props.deployEnvironment}`,
      description: `API Gateway to control the ${props.projectName}-${props.appName} API.`,
      deployOptions: {
        loggingLevel: apigw.MethodLoggingLevel.INFO,
        dataTraceEnabled: true,
        metricsEnabled: true,
      },
    });
    const apiKey = apiGateway.addApiKey(`ApiKey`, {
      apiKeyName: `${props.projectName}-${props.appName}-${props.deployEnvironment}-key`,
      description: `API Key for ${props.projectName}-${props.appName}-${props.deployEnvironment}.`,
    });
    const usagePlan = apiGateway.addUsagePlan(`UsagePlan`, {
      name: `${props.projectName}-${props.appName}-${props.deployEnvironment}-plan`,
      description: `Usage Plan for ${props.projectName}-${props.appName}-${props.deployEnvironment}.`,
      apiStages: [
        {
          api: apiGateway,
          stage: apiGateway.deploymentStage,
        },
      ],
    });
    usagePlan.addApiKey(apiKey);

    const postApp = apiGateway.root.addResource(props.appName);
    const courseSearchIntegration = new apigw.LambdaIntegration(lambdaApp);
    postApp.addMethod("GET", courseSearchIntegration);
  }
}
