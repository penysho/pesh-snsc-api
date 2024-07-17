import {
  Stack,
  StackProps,
  aws_iam as iam,
  aws_lambda as lambda,
} from "aws-cdk-lib";

import { Construct } from "constructs";
import path = require("path");

export interface Parameters extends StackProps {
  projectName: string;
  appName: string;
  deployEnvironment: string;
}

export class PostApp extends Stack {
  constructor(scope: Construct, id: string, props: Parameters) {
    super(scope, id, props);

    // IAM
    const lambdaServiceRole = new iam.Role(this, "lambdaServiceRole", {
      roleName: `${props.projectName}-${props.deployEnvironment}-lambda-role`,
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaVPCAccessExecutionRole"
        ),
      ],
    });

    // Lambda post-app
    const fn = new lambda.Function(this, "lambdaApp", {
      functionName: `${props.projectName}-${props.appName}-${props.deployEnvironment}`,
      runtime: lambda.Runtime.PROVIDED_AL2023,
      handler: `${props.appName}.main`,
      code: lambda.Code.fromAsset(
        path.join(__dirname, `../../${props.appName}`)
      ),
    });
  }
}
