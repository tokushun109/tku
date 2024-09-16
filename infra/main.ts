import { Construct } from 'constructs'
import { App, TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'
import { compileForLambdaFunction } from './libs/compile'
import path = require('path')
import * as dotenv from 'dotenv'
import { getDateString } from './libs/date'
import { VPC_CIDR_BLOCK } from './constants/vpc'
import { SUBNET_CIDR_BLOCK } from './constants/subnet'

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
