import { Construct } from 'constructs'
import { App, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import { compileForLambdaFunction } from './libs/compile'
import path = require('path')
import * as dotenv from 'dotenv'
// import * as fs from 'fs'
import { getDateString } from './libs/date'
import { VPC_CIDR_BLOCK } from './constants/vpc'
import { SUBNET_CIDR_BLOCK } from './constants/subnet'
// import { API_CLUSTER_NAME, ECS_TASK_ASSUME_ROLE_POLICY } from './constants/ecs'
// import { getUserData } from './resources/ec2/userData'

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

        // VPCの作成
        const vpc = new aws.vpc.Vpc(this, `${name}-vpc`, {
            cidrBlock: VPC_CIDR_BLOCK,
            tags: {
                Name: `${name}-vpc`,
            },
        })

        // Subnetの作成
        // const publicSubnetA = new aws.subnet.Subnet(this, `${name}-public-subnet-a`, {
        new aws.subnet.Subnet(this, `${name}-public-subnet-a`, {
            cidrBlock: SUBNET_CIDR_BLOCK.PublicA,
            vpcId: vpc.id,
            tags: {
                Name: `${name}-public-subnet-a`,
            },
        })

        new aws.subnet.Subnet(this, `${name}-public-subnet-c`, {
            cidrBlock: SUBNET_CIDR_BLOCK.PublicC,
            vpcId: vpc.id,
            tags: {
                Name: `${name}-public-subnet-c`,
            },
        })

        // Internet Gatewayの作成
        const internetGateway = new aws.internetGateway.InternetGateway(this, `${name}-igw`, {
            vpcId: vpc.id,
            tags: {
                Name: `${name}-igw`,
            },
        })

        // Route Tableの作成
        new aws.routeTable.RouteTable(this, 'public-rt', {
            vpcId: vpc.id,
            route: [{ cidrBlock: '0.0.0.0/0', gatewayId: internetGateway.id }],
            tags: {
                Name: 'public-rt',
            },
        })

        // Security Groupの作成
        const ecsSecurityGroup = new aws.securityGroup.SecurityGroup(this, 'public-ecs-sg', {
            description: 'sg for public ecs',
            egress: [
                {
                    cidrBlocks: ['0.0.0.0/0'],
                    fromPort: 0,
                    protocol: '-1',
                    toPort: 0,
                },
            ],
            ingress: [
                {
                    cidrBlocks: ['0.0.0.0/0'],
                    fromPort: 443,
                    protocol: 'tcp',
                    toPort: 443,
                },
                {
                    cidrBlocks: ['0.0.0.0/0'],
                    fromPort: 8080,
                    protocol: 'tcp',
                    toPort: 8080,
                },
                {
                    cidrBlocks: ['0.0.0.0/0'],
                    fromPort: 80,
                    protocol: 'tcp',
                    toPort: 80,
                },
                {
                    cidrBlocks: [process.env.HOME_IP!],
                    fromPort: 22,
                    protocol: 'tcp',
                    toPort: 22,
                },
            ],
            tags: {
                Name: 'public-ecs-sg',
            },
        })

        new aws.securityGroup.SecurityGroup(this, 'db-sg', {
            description: 'sg for db',
            egress: [
                {
                    cidrBlocks: ['0.0.0.0/0'],
                    fromPort: 0,
                    protocol: '-1',
                    toPort: 0,
                },
            ],
            ingress: [
                {
                    cidrBlocks: [],
                    fromPort: 3306,
                    protocol: 'tcp',
                    toPort: 3306,
                    securityGroups: [ecsSecurityGroup.id],
                },
                {
                    cidrBlocks: [process.env.HOME_IP!],
                    fromPort: 22,
                    protocol: 'tcp',
                    toPort: 22,
                },
            ],
            tags: {
                Name: 'db-sg',
            },
        })

        // // TODO: 空のECSクラスターを作成する
        // const apiCluster = new aws.ecsCluster.EcsCluster(this, `${name}-ecs-api-cluster`, {
        //     name: API_CLUSTER_NAME,
        //     setting: [
        //         {
        //             name: 'containerInsights',
        //             value: 'disabled',
        //         },
        //     ],
        // })

        // // TODO: ECSで使用する用のEC2を作成して、作成したECSクラスターと関連付ける
        // const apiInstance = new aws.instance.Instance(this, `${name}-api-instance`, {
        //     ami: 'ami-01e9b1393f6f885a6',
        //     associatePublicIpAddress: true,
        //     availabilityZone: 'ap-northeast-1a',
        //     iamInstanceProfile: 'ecsInstanceRole',
        //     instanceType: 't2.small',
        //     keyName: `${name}_rsa`,
        //     userData: Buffer.from(getUserData(API_CLUSTER_NAME)).toString('base64'),
        //     vpcSecurityGroupIds: [ecsSecurityGroup.id],
        //     instanceMarketOptions: {
        //         marketType: 'spot',
        //         spotOptions: {
        //             maxPrice: '0.030400',
        //             spotInstanceType: 'one-time',
        //         },
        //     },
        //     monitoring: false,
        //     tags: {
        //         Name: `test-instance`,
        //     },
        //     subnetId: publicSubnetA.id,
        // })

        // new aws.eip.Eip(this, `${name}-api-eip`, {
        //     domain: 'vpc',
        //     instance: apiInstance.id,
        //     tags: {
        //         Name: 'test-api-eip',
        //     },
        // })

        // // ecsのタスク実行用のroleを作成
        // const ecsTaskExecRole = new aws.iamRole.IamRole(this, `${name}-ecs-task-exec-role`, {
        //     name: `${name}-ecs-task-exec-role`,
        //     assumeRolePolicy: JSON.stringify(ECS_TASK_ASSUME_ROLE_POLICY),
        // })

        // const policyArnsIterator = TerraformIterator.fromList([
        //     'arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy',
        //     // secret managerの読み取り権限も追加
        //     'arn:aws:iam::aws:policy/SecretsManagerReadWrite',
        // ])

        // new aws.iamRolePolicyAttachment.IamRolePolicyAttachment(this, `${name}-ecs-task-managed-policy`, {
        //     forEach: policyArnsIterator,
        //     policyArn: policyArnsIterator.value,
        //     role: ecsTaskExecRole.name,
        // })

        // const { containerDefinitions: apiContainerDefinition }: { containerDefinitions: string } = JSON.parse(
        //     fs.readFileSync(path.join(__dirname, 'resources', 'ecs', 'tasks', 'api.json'), 'utf-8')
        // )

        // // TODO: タスクを作成する
        // const apiTask = new aws.ecsTaskDefinition.EcsTaskDefinition(this, `${name}-api-task`, {
        //     family: `${name}-api`,
        //     containerDefinitions: JSON.stringify(apiContainerDefinition),
        //     executionRoleArn: ecsTaskExecRole.arn,
        //     taskRoleArn: ecsTaskExecRole.arn,
        // })

        // new aws.ecsService.EcsService(this, `${name}-api-service`, {
        //     name: 'test-api-prod',
        //     cluster: apiCluster.id,
        //     taskDefinition: apiTask.arn,
        //     desiredCount: 1,
        //     orderedPlacementStrategy: [
        //         {
        //             field: 'attribute:ecs.availability-zone',
        //             type: 'spread',
        //         },
        //         {
        //             field: 'instanceId',
        //             type: 'spread',
        //         },
        //     ],
        // })

        // const { containerDefinitions: httpsContainerDefinition }: { containerDefinitions: string } = JSON.parse(
        //     fs.readFileSync(path.join(__dirname, 'resources', 'ecs', 'tasks', 'https.json'), 'utf-8')
        // )

        // // TODO: タスクを作成する
        // const httpsTask = new aws.ecsTaskDefinition.EcsTaskDefinition(this, `${name}-https-task`, {
        //     family: `${name}-https`,
        //     containerDefinitions: JSON.stringify(httpsContainerDefinition),
        //     executionRoleArn: ecsTaskExecRole.arn,
        //     taskRoleArn: ecsTaskExecRole.arn,
        // })

        // new aws.ecsService.EcsService(this, `${name}-https-service`, {
        //     name: 'test-https-prod',
        //     cluster: apiCluster.id,
        //     taskDefinition: httpsTask.arn,
        //     desiredCount: 1,
        //     orderedPlacementStrategy: [
        //         {
        //             field: 'attribute:ecs.availability-zone',
        //             type: 'spread',
        //         },
        //         {
        //             field: 'instanceId',
        //             type: 'spread',
        //         },
        //     ],
        // })

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

    // コード上でスケジューラを追加する
}

const app = new App()
new TkuStack(app, 'tku', { region: 'ap-northeast-1' })
app.synth()
