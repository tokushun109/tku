import { Construct } from 'constructs'
import { App, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import * as dotenv from 'dotenv'
import { NetworkResource } from './resources/network'
import { EcsClusterResource, EcsTaskRole, ServiceEnum } from './resources/ecs'
import { SecretsManagerResource } from './resources/asm'
import { EventBridgeResource } from './resources/eventBridge'
import { LambdaResource } from './resources/lambda'

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

        // // DBインスタンスに対してのEIP
        // new aws.eip.Eip(this, `${name}-${ServiceEnum.Db}-eip`, {
        //     domain: 'vpc',
        //     instance: dbInstance.id,
        //     tags: {
        //         Name: `${name}-${ServiceEnum.Db}-eip`,
        //     },
        // })

        // Amazon Secrets Managerのリソースを作成
        new SecretsManagerResource(this, name, apiEip.publicIp, dbInstance.privateIp)

        // フロントエンドのwarmup用のlambda関数を作成
        const { lambdaFunction: warmupLambdaFunction } = new LambdaResource(this, name, 'warmup')

        // warmupのlambda関数を6時〜22時の間で5分ごとに実行(cron式はUTCを日本時間に変換したものを設定)
        new EventBridgeResource(this, name, 'warmup', warmupLambdaFunction, 'cron(*/5 21-23,0-13 * * ? *)')

        // apiのヘルスチェック用のlambda関数を作成
        const { lambdaFunction: healthCheckLambdaFunction } = new LambdaResource(this, name, 'health-check')

        // ヘルスチェックのlambda関数を6時〜20時の間で1時間おきに実行(cron式はUTCを日本時間に変換したものを設定)
        new EventBridgeResource(this, name, 'health-check', healthCheckLambdaFunction, 'cron(0 21-23,0-11/1 * * ? *)')
    }
}

const app = new App()
new TkuStack(app, 'tku', { region: 'ap-northeast-1' })
app.synth()
