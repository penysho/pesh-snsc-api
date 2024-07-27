import { aws_apigateway as apigw, aws_lambda as lambda } from "aws-cdk-lib";
import { Construct } from "constructs";
import { AppProps } from "../postApp";

export interface ApiGatewayProps extends AppProps {
  lambda: lambda.Function;
}

export class ApiGateway extends Construct {
  public readonly apiGateway: apigw.RestApi;

  constructor(scope: Construct, id: string, props: ApiGatewayProps) {
    super(scope, id);

    // APIGateway
    this.apiGateway = new apigw.RestApi(this, `ApiGateway`, {
      restApiName: `${props.projectName}-${props.appName}-${props.deployEnvironment}`,
      description: `API Gateway to control the ${props.projectName}-${props.appName} API.`,
      deployOptions: {
        loggingLevel: apigw.MethodLoggingLevel.INFO,
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

    const app = this.apiGateway.root.addResource(props.appName);
    const courseSearchIntegration = new apigw.LambdaIntegration(props.lambda);
    app.addMethod("GET", courseSearchIntegration);
  }
}
