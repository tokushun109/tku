import * as aws from '@cdktf/provider-aws/lib'
import { TerraformStack } from 'cdktf'
import { compileForLambdaFunction } from '../../../../libs/compile'
import path = require('path')
import { getDateString } from '../../../../libs/date'

const LAMBDA_ASSUME_ROLE_POLICY = {
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

export class LambdaResource {
    public lambdaFunction: aws.lambdaFunction.LambdaFunction

    constructor(private stack: TerraformStack, private name: string, private lambdaFunctionName: string) {
        // lambda実行用のroleを作成
        const role = new aws.iamRole.IamRole(this.stack, `${this.name}-lambda-exec`, {
            name: `${this.name}-lambda-role`,
            assumeRolePolicy: JSON.stringify(LAMBDA_ASSUME_ROLE_POLICY),
        })

        // CloudWatchログへの書き込み権限を追加
        new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this.stack, `${this.name}-lambda-managed-policy`, {
            policyArn: 'arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole',
            role: role.name,
        })

        // lambda関数用のハンドラをコンパイルする
        const lambda = new compileForLambdaFunction(this.stack, this.name, {
            path: path.join(__dirname),
        })

        const prefix = `${this.name}-${this.lambdaFunctionName}`

        // lambdaのarchiveを保存するS3を作成
        const bucket = new aws.s3Bucket.S3Bucket(this.stack, `${prefix}-lambda-archive-bucket`, {
            bucket: `${prefix}-lambda-archive-bucket`,
        })

        // S3にlambda関数をアップロード
        const lambdaArchive = new aws.s3Object.S3Object(this.stack, `${prefix}-lambda-archive`, {
            bucket: bucket.bucket,
            key: `${this.name}/${this.lambdaFunctionName}/${getDateString()}/${lambda.asset.fileName}`,
            source: lambda.asset.path,
        })

        // lambdaを作成
        this.lambdaFunction = new aws.lambdaFunction.LambdaFunction(this.stack, `${prefix}-lambda-function`, {
            functionName: `${prefix}`,
            s3Bucket: bucket.bucket,
            s3Key: lambdaArchive.key,
            timeout: 30,
            handler: 'index.handler',
            runtime: 'nodejs18.x',
            memorySize: 512,
            role: role.arn,
            environment: {
                variables: {
                    LINE_HEALTH_CHECK_TOKEN: process.env.LINE_HEALTH_CHECK_TOKEN!,
                    LINE_HEALTH_CHECK_USER_ID: process.env.LINE_HEALTH_CHECK_USER_ID!,
                },
            },
        })
    }
}
