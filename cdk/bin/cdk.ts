#!/usr/bin/env node
import * as cdk from "aws-cdk-lib";
import "source-map-support/register";
import { deployEnvironment, postAppName, projectName } from "../config/config";
import { AppProps, PostApp } from "../lib/postApp";
import { DbProps, Rds } from "../lib/rds";

const app = new cdk.App();

export const postAppProps: AppProps = {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
  projectName: projectName,
  appName: postAppName,
  deployEnvironment: deployEnvironment,
};

export const rdsProps: DbProps = {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
  projectName: projectName,
  deployEnvironment: deployEnvironment,
};

new PostApp(
  app,
  `${projectName}-${postAppName}-${deployEnvironment}`,
  postAppProps
);

new Rds(app, `${projectName}-rds-${deployEnvironment}`, rdsProps);
