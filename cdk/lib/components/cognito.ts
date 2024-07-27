import { aws_cognito as cognito } from "aws-cdk-lib";
import { Construct } from "constructs";
import { AppProps } from "../postApp";

export interface CognitoProps extends AppProps {}

export class Cognito extends Construct {
  public readonly userPool: cognito.UserPool;

  constructor(scope: Construct, id: string, props: CognitoProps) {
    super(scope, id);

    // https://note.com/hiroyu0510/n/n629ccace5591
    this.userPool = new cognito.UserPool(this, "UserPool", {
      selfSignUpEnabled: true,
      userVerification: {
        emailSubject: "Verify your email for our app!",
        emailBody:
          "Thanks for signing up to our app! Your verification code is {####}",
        emailStyle: cognito.VerificationEmailStyle.CODE,
      },
      signInAliases: {
        email: true,
      },
    });

    const userPoolClient = this.userPool.addClient("UserPoolClient", {
      authFlows: {
        adminUserPassword: true,
        userPassword: true,
      },
    });

    this.userPool.addDomain("UserPoolDomain", {
      cognitoDomain: {
        domainPrefix: "pesh-igpjt",
      },
    });
  }
}
