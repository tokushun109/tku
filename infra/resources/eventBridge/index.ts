import { TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'

export class EventBridgeResource {
    constructor(
        private stack: TerraformStack,
        private name: string,
        /** 例: health-checkなど */
        private lambdaFunctionIdentifierName: string,
        private lambdaFunction: aws.lambdaFunction.LambdaFunction,
        private scheduleExpression: string
    ) {
        // EventBridgeのルールの作成（毎日6時から20時まで1時間ごと）
        const eventRule = new aws.cloudwatchEventRule.CloudwatchEventRule(
            this.stack,
            `${this.name}-${this.lambdaFunctionIdentifierName}-event-rule`,
            {
                name: `${name}-${this.lambdaFunctionIdentifierName}-event-rule`,
                scheduleExpression: this.scheduleExpression,
            }
        )

        // EventBridgeのターゲットとしてLambdaを設定
        new aws.cloudwatchEventTarget.CloudwatchEventTarget(this.stack, `${this.name}-${this.lambdaFunctionIdentifierName}-event-target`, {
            rule: eventRule.name,
            arn: this.lambdaFunction.arn,
        })

        // LambdaにEventBridgeからの実行権限を付与
        new aws.lambdaPermission.LambdaPermission(this.stack, `${this.name}-${this.lambdaFunctionIdentifierName}-allow-event-bridge-invoke-lambda`, {
            statementId: 'AllowEventBridgeInvokeLambda',
            action: 'lambda:InvokeFunction',
            functionName: this.lambdaFunction.functionName,
            principal: 'events.amazonaws.com',
            sourceArn: eventRule.arn,
        })
    }
}
