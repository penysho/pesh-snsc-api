import * as cdk from "aws-cdk-lib";
import {
  aws_apigateway as apigw,
  aws_cognito as cognito,
  aws_lambda as lambda,
  aws_logs as logs,
} from "aws-cdk-lib";
import { Construct } from "constructs";
import { AppProps } from "../postApp";

export interface ApiGatewayProps extends AppProps {
  lambda: lambda.Function;
  userPool: cognito.UserPool;
}

export class ApiGateway extends Construct {
  public readonly apiGateway: apigw.RestApi;

  constructor(scope: Construct, id: string, props: ApiGatewayProps) {
    super(scope, id);

    // Log Group
    const logGroup = new logs.LogGroup(this, "LogGroup", {
      logGroupName: `/aws/apigateway/${props.projectName}-${props.appName}-${props.deployEnvironment}`,
      retention: logs.RetentionDays.THREE_MONTHS,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });

    // API Gateway
    this.apiGateway = new apigw.RestApi(this, `ApiGateway`, {
      restApiName: `${props.projectName}-${props.appName}-${props.deployEnvironment}`,
      description: `API Gateway to control the ${props.projectName}-${props.appName} API.`,
      deployOptions: {
        stageName: "prod",
        loggingLevel: apigw.MethodLoggingLevel.INFO,
        accessLogDestination: new apigw.LogGroupLogDestination(logGroup),
        accessLogFormat: apigw.AccessLogFormat.clf(),
        dataTraceEnabled: true,
        metricsEnabled: true,
        tracingEnabled: true,
      },
    });

    const apiKey = this.apiGateway.addApiKey(`ApiKey`, {
      apiKeyName: `${props.projectName}-${props.appName}-${props.deployEnvironment}-key`,
      description: `API Key for ${props.projectName}-${props.appName}-${props.deployEnvironment}.`,
    });
    const usagePlan = this.apiGateway.addUsagePlan(`UsagePlan`, {
      name: `${props.projectName}-${props.appName}-${props.deployEnvironment}-plan`,
      description: `Usage Plan for ${props.projectName}-${props.appName}-${props.deployEnvironment}.`,
      apiStages: [
        {
          api: this.apiGateway,
          stage: this.apiGateway.deploymentStage,
        },
      ],
    });
    usagePlan.addApiKey(apiKey);

    const integration = new apigw.LambdaIntegration(props.lambda);

    const authorizer = new apigw.CognitoUserPoolsAuthorizer(
      this,
      "CognitoAuthorizer",
      {
        cognitoUserPools: [props.userPool],
      }
    );

    const app = this.apiGateway.root.addResource(props.appName);
    app.addMethod("GET", integration, {
      authorizationType: apigw.AuthorizationType.COGNITO,
      authorizer,
    });
  }
}
