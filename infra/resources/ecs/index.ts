import { TerraformIterator, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import { getUserData } from '../ec2/userData'
import { importTaskDefinition } from '../../libs/task'
import path = require('path')

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

export const ServiceEnum = {
    Api: 'api',
    Db: 'db',
} as const
export type ServiceEnum = (typeof ServiceEnum)[keyof typeof ServiceEnum]

export class EcsClusterResource {
    public instance: aws.instance.Instance

    constructor(
        private stack: TerraformStack,
        private name: string,
        private service: ServiceEnum,
        private securityGroupId: string,
        private subnetId: string,
        private taskRoleArn: string
    ) {
        const prefix = `${this.name}-${this.service}`
        const clusterName = `${prefix}-cluster`

        const cluster = new aws.ecsCluster.EcsCluster(this.stack, `${this.name}-ecs-${this.service}-cluster`, {
            name: clusterName,
            setting: [
                {
                    name: 'containerInsights',
                    value: 'disabled',
                },
            ],
        })

        this.instance = new aws.instance.Instance(this.stack, `${prefix}-instance`, {
            ami: 'ami-01e9b1393f6f885a6',
            associatePublicIpAddress: false,
            availabilityZone: 'ap-northeast-1a',
            iamInstanceProfile: 'ecsInstanceRole',
            instanceType: 't2.small',
            keyName: `${this.name}_rsa`,
            userData: Buffer.from(getUserData(clusterName)).toString('base64'),
            vpcSecurityGroupIds: [this.securityGroupId],
            instanceMarketOptions: {
                marketType: 'spot',
                spotOptions: {
                    maxPrice: '0.030400',
                    spotInstanceType: 'one-time',
                },
            },
            monitoring: false,
            tags: {
                Name: `${prefix}-instance`,
            },
            subnetId: this.subnetId,
            lifecycle: {
                ignoreChanges: ['associate_public_ip_address'],
            },
        })

        const task = new aws.ecsTaskDefinition.EcsTaskDefinition(this.stack, `${prefix}-task`, {
            family: `${prefix}`,
            containerDefinitions: importTaskDefinition(path.join(__dirname, 'tasks', `${this.service}.json`)),
            executionRoleArn: this.taskRoleArn,
            taskRoleArn: this.taskRoleArn,
            requiresCompatibilities: ['EC2'],
        })

        new aws.ecsService.EcsService(this.stack, `${prefix}-service`, {
            name: `${prefix}-service`,
            cluster: cluster.id,
            taskDefinition: task.arn,
            desiredCount: 1,
            orderedPlacementStrategy: [
                {
                    field: 'attribute:ecs.availability-zone',
                    type: 'spread',
                },
                {
                    field: 'instanceId',
                    type: 'spread',
                },
            ],
        })

        // apiサービスにはSSL通信用のコンテナも作成する
        if (this.service === ServiceEnum.Api) {
            const HTTPS_SERVICE = 'https'
            const https_prefix = `${this.name}-${HTTPS_SERVICE}`
            const httpsTask = new aws.ecsTaskDefinition.EcsTaskDefinition(this.stack, `${https_prefix}-task`, {
                family: `${https_prefix}`,
                containerDefinitions: importTaskDefinition(path.join(__dirname, 'tasks', `${HTTPS_SERVICE}.json`)),
                executionRoleArn: this.taskRoleArn,
                taskRoleArn: this.taskRoleArn,
                requiresCompatibilities: ['EC2'],
            })

            new aws.ecsService.EcsService(this.stack, `${https_prefix}-service`, {
                name: `${https_prefix}-service`,
                cluster: cluster.id,
                taskDefinition: httpsTask.arn,
                desiredCount: 1,
                orderedPlacementStrategy: [
                    {
                        field: 'attribute:ecs.availability-zone',
                        type: 'spread',
                    },
                    {
                        field: 'instanceId',
                        type: 'spread',
                    },
                ],
            })
        }
    }
}
