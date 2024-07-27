import {
  Duration,
  aws_ec2 as ec2,
  aws_iam as iam,
  aws_lambda as lambda,
} from "aws-cdk-lib";
import { Construct } from "constructs";
import { AppProps } from "../postApp";
import path = require("path");

const appBuildPath = "./cmd/aws/lambda/apigw/main.go";

export interface LambdaProps extends AppProps {
  vpc: ec2.IVpc;
}

export class Lambda extends Construct {
  public readonly lambdaApp: lambda.Function;

  constructor(scope: Construct, id: string, props: LambdaProps) {
    super(scope, id);
    const vpc = props.vpc;

    // IAM
    const lambdaServiceRole = new iam.Role(this, `LambdaServiceRole`, {
      description: `Generally used in ${props.projectName}-${props.deployEnvironment} Lambda deployed in VPC.`,
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaVPCAccessExecutionRole"
        ),
      ],
    });

    // SG
    const lambdaSecurityGroup = new ec2.SecurityGroup(
      this,
      `LambdaSecurityGroup`,
      {
        vpc,
        description: `Generally used in ${props.projectName}-${props.deployEnvironment} Lambda deployed in VPC.`,
        allowAllOutbound: true,
      }
    );

    // Lambda post-app
    // 将来的に以下に移行を想定
    // https://docs.aws.amazon.com/cdk/api/v2/docs/aws-lambda-go-alpha-readme.html
    this.lambdaApp = new lambda.Function(this, `LambdaApp`, {
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
      tracing: lambda.Tracing.ACTIVE,
      environment: {
        // DB_HOST: SecretValue.secretsManager(
        //   `pesh-snsc-api-tst/rds/admin-secret`,
        //   {
        //     jsonField: `host`,
        //   }
        // ).unsafeUnwrap(),
        // DB_PORT: SecretValue.secretsManager(
        //   `pesh-snsc-api-tst/rds/admin-secret`,
        //   {
        //     jsonField: `port`,
        //   }
        // ).unsafeUnwrap(),
        // DB_USER: SecretValue.secretsManager(
        //   `pesh-snsc-api-tst/rds/admin-secret`,
        //   {
        //     jsonField: `username`,
        //   }
        // ).unsafeUnwrap(),
        // DB_PASSWORD: SecretValue.secretsManager(
        //   `pesh-snsc-api-tst/rds/admin-secret`,
        //   {
        //     jsonField: `password`,
        //   }
        // ).unsafeUnwrap(),
        DB_NAME: "postgres",
      },
      code: lambda.Code.fromAsset(
        path.join(__dirname, `../../../${props.appName}`),
        {
          // https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-package.html
          // https://iret.media/79515
          bundling: {
            image: lambda.Runtime.PROVIDED_AL2023.bundlingImage,
            command: [
              "bash",
              "-c",
              // デフォルトではルートディレクトリにGOCACHEとGOPATHを作成してパーミッションではじかれるため変更
              `export GOCACHE=/tmp/go-cache && export GOPATH=/tmp/go-path && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o /asset-output/bootstrap ${appBuildPath}`,
            ],
          },
        }
      ),
    });
  }
}
