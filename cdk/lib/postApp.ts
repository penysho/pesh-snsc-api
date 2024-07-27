import { Fn, Stack, StackProps, aws_ec2 as ec2 } from "aws-cdk-lib";
import { ApiGateway } from "./components/apiGateway";
import { Lambda } from "./components/lambda";

import { Construct } from "constructs";
import { Cognito } from "./components/cognito";

export interface AppProps extends StackProps {
  projectName: string;
  appName: string;
  deployEnvironment: string;
}

export class PostApp extends Stack {
  constructor(scope: Construct, id: string, props: AppProps) {
    super(scope, id, props);

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

    const postApp = new Lambda(this, "PostApp", {
      vpc: vpc,
      ...props,
    });
    const cognito = new Cognito(this, "Cognito", {
      ...props,
    });
    new ApiGateway(this, "ApiGateway", {
      lambda: postApp.lambdaApp,
      userPool: cognito.userPool,
      ...props,
    });
  }
}
