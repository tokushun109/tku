import { Construct } from 'constructs'
import { App, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import * as dotenv from 'dotenv'
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
