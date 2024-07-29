import * as cdk from "aws-cdk-lib";
import { Template } from "aws-cdk-lib/assertions";
import { postAppProps, rdsProps } from "../bin/cdk";
import { deployEnvironment, postAppName, projectName } from "../config/config";
import { PostApp } from "../lib/postApp";
import { Rds } from "../lib/rds";

describe("fine grained assertions test ", () => {
  test("rds", () => {
    const app = new cdk.App();

    const stack = new Rds(
      app,
      `${projectName}-rds-${deployEnvironment}`,
      rdsProps
    );

    const template = Template.fromStack(stack);

    template.resourceCountIs("AWS::RDS::DBInstance", 2);
  });

  test("post-app", () => {
    const app = new cdk.App();

    const stack = new PostApp(
      app,
      `${projectName}-${postAppName}-${deployEnvironment}`,
      postAppProps
    );

    const template = Template.fromStack(stack);

    template.resourceCountIs("AWS::Lambda::Function", 1);
  });
});
