import { Construct } from 'constructs'
import { App, TerraformIterator, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import { compileForLambdaFunction } from './libs/compile'
import path = require('path')
import * as dotenv from 'dotenv'
import { getDateString } from './libs/date'
import { ECS_TASK_ASSUME_ROLE_POLICY, API_CLUSTER_NAME, DB_CLUSTER_NAME } from './constants/ecs'
import { getUserData } from './resources/ec2/userData'
import { importTaskDefinition } from './libs/task'
import { NetworkResource } from './resources/network'

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

        // ecsのタスク実行用のroleを作成
        const ecsTaskExecRole = new aws.iamRole.IamRole(this, `${name}-ecs-task-exec-role`, {
            name: `${name}-ecs-task-exec-role`,
            assumeRolePolicy: JSON.stringify(ECS_TASK_ASSUME_ROLE_POLICY),
        })

        const policyArnsIterator = TerraformIterator.fromList([
            'arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy',
            // secret managerの読み取り権限も追加
            'arn:aws:iam::aws:policy/SecretsManagerReadWrite',
        ])

        new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, `${name}-ecs-task-managed-policy`, {
            forEach: policyArnsIterator,
            policyArn: policyArnsIterator.value,
            role: ecsTaskExecRole.name,
        })

        const apiCluster = new aws.ecsCluster.EcsCluster(this, `${name}-ecs-api-cluster`, {
            name: API_CLUSTER_NAME,
            setting: [
                {
                    name: 'containerInsights',
                    value: 'disabled',
                },
            ],
        })

        const apiInstance = new aws.instance.Instance(this, `${name}-api-instance`, {
            ami: 'ami-01e9b1393f6f885a6',
            associatePublicIpAddress: true,
            availabilityZone: 'ap-northeast-1a',
            iamInstanceProfile: 'ecsInstanceRole',
            instanceType: 't2.small',
            keyName: `${name}_rsa`,
            userData: Buffer.from(getUserData(API_CLUSTER_NAME)).toString('base64'),
            vpcSecurityGroupIds: [securityGroupIds.api],
            instanceMarketOptions: {
                marketType: 'spot',
                spotOptions: {
                    maxPrice: '0.030400',
                    spotInstanceType: 'one-time',
                },
            },
            monitoring: false,
            tags: {
                Name: `${name}-api-instance`,
            },
            subnetId: subnetIds.a,
        })

        const apiEip = new aws.eip.Eip(this, `${name}-api-eip`, {
            domain: 'vpc',
            instance: apiInstance.id,
            tags: {
                Name: `${name}-api-eip`,
            },
        })

        const apiTask = new aws.ecsTaskDefinition.EcsTaskDefinition(this, `${name}-api-task`, {
            family: `${name}-api`,
            containerDefinitions: importTaskDefinition(path.join(__dirname, 'resources', 'ecs', 'tasks', 'api.json')),
            executionRoleArn: ecsTaskExecRole.arn,
            taskRoleArn: ecsTaskExecRole.arn,
            requiresCompatibilities: ['EC2'],
        })

        new aws.ecsService.EcsService(this, `${name}-api-service`, {
            name: `${name}-api-service`,
            cluster: apiCluster.id,
            taskDefinition: apiTask.arn,
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

        const httpsTask = new aws.ecsTaskDefinition.EcsTaskDefinition(this, `${name}-https-task`, {
            family: `${name}-https`,
            containerDefinitions: importTaskDefinition(path.join(__dirname, 'resources', 'ecs', 'tasks', 'https.json')),
            executionRoleArn: ecsTaskExecRole.arn,
            taskRoleArn: ecsTaskExecRole.arn,
            requiresCompatibilities: ['EC2'],
        })

        new aws.ecsService.EcsService(this, `${name}-https-service`, {
            name: `${name}-https-service`,
            cluster: apiCluster.id,
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

        const dbCluster = new aws.ecsCluster.EcsCluster(this, `${name}-ecs-db-cluster`, {
            name: DB_CLUSTER_NAME,
            setting: [
                {
                    name: 'containerInsights',
                    value: 'disabled',
                },
            ],
        })

        const dbInstance = new aws.instance.Instance(this, `${name}-db-instance`, {
            ami: 'ami-01e9b1393f6f885a6',
            associatePublicIpAddress: true,
            availabilityZone: 'ap-northeast-1a',
            iamInstanceProfile: 'ecsInstanceRole',
            instanceType: 't2.small',
            keyName: `${name}_rsa`,
            userData: Buffer.from(getUserData(DB_CLUSTER_NAME)).toString('base64'),
            vpcSecurityGroupIds: [securityGroupIds.db],
            instanceMarketOptions: {
                marketType: 'spot',
                spotOptions: {
                    maxPrice: '0.030400',
                    spotInstanceType: 'one-time',
                },
            },
            monitoring: false,
            tags: {
                Name: `${name}-db-instance`,
            },
            subnetId: subnetIds.a,
        })

        const dbTask = new aws.ecsTaskDefinition.EcsTaskDefinition(this, `${name}-db-task`, {
            family: `${name}-db`,
            containerDefinitions: importTaskDefinition(path.join(__dirname, 'resources', 'ecs', 'tasks', 'db.json')),
            executionRoleArn: ecsTaskExecRole.arn,
            taskRoleArn: ecsTaskExecRole.arn,
            requiresCompatibilities: ['EC2'],
        })

        new aws.ecsService.EcsService(this, `${name}-db-service`, {
            name: `${name}-db-service`,
            cluster: dbCluster.id,
            taskDefinition: dbTask.arn,
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

        const secretsManager = new aws.secretsmanagerSecret.SecretsmanagerSecret(this, `${name}-secrets-manager`, {
            name: `${name}-secrets-manager`,
            description: `${name}-secrets-manager`,
        })

        new aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion(this, `${name}-secrets-manager-version`, {
            secretId: secretsManager.id,
            secretString: JSON.stringify({
                API_BASE_URL: process.env.API_BASE_URL,
                DB_PASS: process.env.DB_PASS,
                MYSQL_ROOT_PASSWORD: process.env.MYSQL_ROOT_PASSWORD,
                DB_USER: process.env.DB_USER,
                CREATOR_NAME: process.env.CREATOR_NAME,
                DB_NAME: process.env.DB_NAME,
                ENV: 'prod',
                MYSQL_HOST: dbInstance.privateIp,
                AWS_SECRET_ACCESS_KEY: process.env.AWS_SECRET_ACCESS_KEY,
                AWS_ACCESS_KEY_ID: process.env.AWS_ACCESS_KEY_ID,
                AWS_REGION: process.env.AWS_REGION,
                API_BUCKET_NAME: process.env.API_BUCKET_NAME,
                SEND_GRID_API_KEY: process.env.SEND_GRID_API_KEY,
                LINE_CONTACT_TOKEN: process.env.LINE_CONTACT_TOKEN,
                CLIENT_URL: process.env.CLIENT_URL,
                DOMAINS: `api.tocoriri.com -> http://${apiEip.publicIp}:8080`,
                STAGE: process.env.STAGE,
            }),
        })

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

    // TODO: 冗長なコードの書き方をリファクタする
}

const app = new App()
new TkuStack(app, 'tku', { region: 'ap-northeast-1' })
app.synth()
