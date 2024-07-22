#!/usr/bin/env node
import * as cdk from "aws-cdk-lib";
import "source-map-support/register";
import { PostApp } from "../lib/lambda";
import { Rds } from "../lib/rds";

const app = new cdk.App();

const projectName: string = "pesh-snsc-api";
const postAppName: string = "post-app";
const deployEnvironment = process.env.DEPLOY_ENV
  ? process.env.DEPLOY_ENV
  : "tst";
const appBuildPath = "./cmd/aws/lambda/apigw/main.go";

new PostApp(app, `${projectName}-${postAppName}-${deployEnvironment}`, {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
  projectName: projectName,
  appName: postAppName,
  deployEnvironment: deployEnvironment,
  appBuildPath: appBuildPath,
});

new Rds(app, `${projectName}-rds-${deployEnvironment}`, {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
  projectName: projectName,
  deployEnvironment: deployEnvironment,
});
