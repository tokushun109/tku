import { TerraformIterator, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'

export const ECS_TASK_ASSUME_ROLE_POLICY = {
    Version: '2012-10-17',
    Statement: [
        {
            Effect: 'Allow',
            Principal: {
                Service: 'ecs-tasks.amazonaws.com',
            },
            Action: 'sts:AssumeRole',
        },
    ],
}

export class EcsTaskRole {
    public roleArn: string

    constructor(private stack: TerraformStack, private name: string) {
        // ecsのタスク実行用のroleを作成
        const ecsTaskExecRole = new aws.iamRole.IamRole(this.stack, `${this.name}-ecs-task-exec-role`, {
            name: `${name}-ecs-task-exec-role`,
            assumeRolePolicy: JSON.stringify(ECS_TASK_ASSUME_ROLE_POLICY),
        })

        const policyArnsIterator = TerraformIterator.fromList([
            'arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy',
            // secret managerの読み取り権限も追加
            'arn:aws:iam::aws:policy/SecretsManagerReadWrite',
        ])

        new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this.stack, `${this.name}-ecs-task-managed-policy`, {
            forEach: policyArnsIterator,
            policyArn: policyArnsIterator.value,
            role: ecsTaskExecRole.name,
        })

        this.roleArn = ecsTaskExecRole.arn
    }
}
