import { Construct } from 'constructs'
import { App, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import { compileForLambdaFunction } from './libs/compile'
import path = require('path')
import * as dotenv from 'dotenv'
import { getDateString } from './libs/date'
import { NetworkResource } from './resources/network'
import { EcsClusterResource, EcsTaskRole, ServiceEnum } from './resources/ecs'
import { SecretsManagerResource } from './resources/asm'

interface OptionType {
    region: string
}

dotenv.config()
class TkuStack extends TerraformStack {
    constructor(scope: Construct, name: string, option: OptionType) {
        super(scope, name)

        const { region } = option
        new aws.provider.AwsProvider(this, `${name}-aws-provider`, {
            region,
        })

        // ネットワーク関連のリソースを作成
        const { subnetIds, securityGroupIds } = new NetworkResource(this, name)

        // ECSのタスク実行用のroleを作成
        const { roleArn: ecsTaskRoleArn } = new EcsTaskRole(this, name)

        // APIのECSクラスターに関連するリソースを作成
        const { instance: apiInstance } = new EcsClusterResource(this, name, ServiceEnum.Api, securityGroupIds.api, subnetIds.a, ecsTaskRoleArn)

        // apiインスタンスをEIPに紐付け
        const apiEip = new aws.eip.Eip(this, `${name}-${ServiceEnum.Api}-eip`, {
            domain: 'vpc',
            instance: apiInstance.id,
            tags: {
                Name: `${name}-${ServiceEnum.Api}-eip`,
            },
        })

        // DBのECSクラスターに関連するリソースを作成
        const { instance: dbInstance } = new EcsClusterResource(this, name, ServiceEnum.Db, securityGroupIds.db, subnetIds.a, ecsTaskRoleArn)

        // DBインスタンスに対してのEIP
        // new aws.eip.Eip(this, `${name}-${ServiceEnum.Db}-eip`, {
        //     domain: 'vpc',
        //     instance: dbInstance.id,
        //     tags: {
        //         Name: `${name}-${ServiceEnum.Db}-eip`,
        //     },
        // })

        // Amazon Secrets Managerのリソースを作成
        new SecretsManagerResource(this, name, dbInstance.privateIp, apiEip.publicIp)

        // lambda関数用のハンドラをコンパイルする
        const lambda = new compileForLambdaFunction(this, name, {
            path: path.join(__dirname, 'resources', 'lambda', 'healthCheck', 'handlers'),
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

        new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, 'lambda-managed-policy', {
            policyArn,
            role: role.name,
        })

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
        new aws.lambdaFunction.LambdaFunction(this, `lambda-function-${name}-health-check`, {
            functionName: `${name}-health-check`,
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
                },
            },
        })
    }
}

const app = new App()
new TkuStack(app, 'tku', { region: 'ap-northeast-1' })
app.synth()
