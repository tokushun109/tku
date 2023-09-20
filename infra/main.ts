import { Construct } from "constructs";
import { App, TerraformStack } from "cdktf";
import * as aws from '@cdktf/provider-aws/lib'
import { compileForLambdaFunction } from "./libs/compile";
import path = require("path");
import * as dotenv from 'dotenv'

interface OptionType {
  region: string
}

// 日付をYYYYMMddhhmmssの文字列で取得
const getDateString = (date: Date | null = null): string => {
  const d = date || new Date()

  return (
    d.getFullYear() +
    String(d.getMonth() + 1).padStart(2, '0') +
    String(d.getDate()).padStart(2, '0') +
    String(d.getHours()).padStart(2, '0') +
    String(d.getMinutes()).padStart(2, '0') +
    String(d.getSeconds()).padStart(2, '0')
  )
}

dotenv.config()
class TkuStack extends TerraformStack {
  constructor(scope: Construct, name: string, option: OptionType) {
    super(scope, name);

    const { region } = option
    new aws.provider.AwsProvider(this, `${name}-aws-provider`, {
      region
    })

    // lambda関数用のハンドラをコンパイルする
    const lambda = new compileForLambdaFunction(this, name, {
      path: path.join(__dirname, 'lambda', 'healthCheck', 'handlers'),
    })

    const lambdaAssumeRolePolicy = {
      Version: '2012-10-17',
      Statement: [
        {
          Action: 'sts:AssumeRole',
          Principal: {
            Service: 'lambda.amazonaws.com',
          },
          Effect: 'Allow',
        },
      ],
    }

    // lambda実行用のroleを作成
    const role = new aws.iamRole.IamRole(this, 'lambda-exec', {
      name: `${name}-lambda-role`,
      assumeRolePolicy: JSON.stringify(lambdaAssumeRolePolicy),
    })

    // CloudWatchログへの書き込み権限を追加
    const policyArn = 'arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole'

    new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(
      this,
      'lambda-managed-policy',
      {
        policyArn,
        role: role.name,
      }
    )

    // lambdaのarchiveを保存するS3を作成
    const bucket = new aws.s3Bucket.S3Bucket(this, `${name}-lambda-archive`, {
      bucket: `${name}-lambda-archive`,
    })

    // S3にlambda関数をアップロード
    const lambdaArchive = new aws.s3Object.S3Object(this, 'lambda-archive', {
      bucket: bucket.bucket,
      key: `${name}/${getDateString()}/${lambda.asset.fileName}`,
      source: lambda.asset.path,
    })

    // lambdaを作成
    new aws.lambdaFunction.LambdaFunction(
      this,
      `lambda-function-${name}-health-check`,
      {
        functionName: `${name}-health-check`,
        s3Bucket: bucket.bucket,
        s3Key: lambdaArchive.key,
        handler: 'index.handler',
        runtime: 'nodejs18.x',
        memorySize: 512,
        role: role.arn,
        environment: {
          variables: {
            LINE_HEALTH_CHECK_TOKEN: process.env.LINE_HEALTH_CHECK_TOKEN!
          }
        }
      }
    )
  }

  // コード上でスケジューラを追加する
}

const app = new App();
new TkuStack(app, "tku", { region: 'ap-northeast-1' });
app.synth();
