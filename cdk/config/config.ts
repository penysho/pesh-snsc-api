// プロジェクト設定
export const projectName: string = "pesh-snsc-api";
export const postAppName: string = "post-app";

// デプロイ環境設定
export const envCodes = ["dev", "tst", "prd"] as const;
export type EnvCode = (typeof envCodes)[number];

const getDeployEnvironment = () => {
  const env = process.env.DEPLOY_ENV;
  if (envCodes.includes(env as EnvCode)) {
    return env as EnvCode;
  }
  return "tst";
};

export const deployEnvironment: EnvCode = getDeployEnvironment();
