import * as cdk from "aws-cdk-lib";
import { Template } from "aws-cdk-lib/assertions";
import { postAppProps, rdsProps } from "../bin/cdk";
import { deployEnvironment, postAppName, projectName } from "../config/config";
import { PostApp } from "../lib/postApp";
import { Rds } from "../lib/rds";

describe("snapshot test", () => {
  test("rds", () => {
    const app = new cdk.App();

    const stack = new Rds(
      app,
      `${projectName}-rds-${deployEnvironment}`,
      rdsProps
    );

    const template = Template.fromStack(stack);

    expect(template).toMatchSnapshot();
  });

  test("post-app", () => {
    const app = new cdk.App();

    const stack = new PostApp(
      app,
      `${projectName}-${postAppName}-${deployEnvironment}`,
      postAppProps
    );

    const template = Template.fromStack(stack).toJSON();

    template.Parameters = {};
    Object.values(template.Resources).forEach((resource: any) => {
      // Codeの変更は許容するため上書き
      if (resource?.Properties?.Code) {
        resource.Properties.Code = {};
      }
    });

    expect(template).toMatchSnapshot();
  });
});
