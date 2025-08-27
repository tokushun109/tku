import * as aws from '@cdktf/provider-aws/lib'
import { TerraformStack } from 'cdktf'

const VPC_CIDR_BLOCK = '10.0.0.0/16'

const SUBNET_CIDR_BLOCK = {
    PublicA: '10.0.1.0/24',
    PublicC: '10.0.2.0/24',
} as const

interface SubnetIds {
    a: string
    c: string
}

interface SecurityGroupIds {
    api: string
    db: string
}

export class NetworkResource {
    public subnetIds: SubnetIds
    public securityGroupIds: SecurityGroupIds

    constructor(private stack: TerraformStack, private name: string) {
        // VPCの作成
        const vpc = new aws.vpc.Vpc(this.stack, `${this.name}-vpc`, {
            cidrBlock: VPC_CIDR_BLOCK,
            tags: {
                Name: `${this.name}-vpc`,
            },
        })

        // Subnetの作成
        const subnetA = new aws.subnet.Subnet(this.stack, `${name}-public-subnet-a`, {
            cidrBlock: SUBNET_CIDR_BLOCK.PublicA,
            vpcId: vpc.id,
            tags: {
                Name: `${name}-public-subnet-a`,
            },
        })

        const subnetC = new aws.subnet.Subnet(this.stack, `${name}-public-subnet-c`, {
            cidrBlock: SUBNET_CIDR_BLOCK.PublicC,
            vpcId: vpc.id,
            tags: {
                Name: `${name}-public-subnet-c`,
            },
        })

        this.subnetIds = {
            a: subnetA.id,
            c: subnetC.id,
        }

        // Internet Gatewayの作成
        const internetGateway = new aws.internetGateway.InternetGateway(this.stack, `${name}-igw`, {
            vpcId: vpc.id,
            tags: {
                Name: `${name}-igw`,
            },
        })

        // Route Tableの作成
        const routeTable = new aws.routeTable.RouteTable(this.stack, 'public-rt', {
            vpcId: vpc.id,
            route: [{ cidrBlock: '0.0.0.0/0', gatewayId: internetGateway.id }],
            tags: {
                Name: 'public-rt',
            },
        })

        // Route Table Associationの作成
        new aws.routeTableAssociation.RouteTableAssociation(this.stack, 'public-rta-a', {
            subnetId: subnetA.id,
            routeTableId: routeTable.id,
        })

        new aws.routeTableAssociation.RouteTableAssociation(this.stack, 'public-rta-c', {
            subnetId: subnetC.id,
            routeTableId: routeTable.id,
        })

        // Security Groupの作成
        const apiSecurityGroup = new aws.securityGroup.SecurityGroup(this.stack, 'public-ecs-sg', {
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

        const dbSecurityGroup = new aws.securityGroup.SecurityGroup(this.stack, 'db-sg', {
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
                    securityGroups: [apiSecurityGroup.id],
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

        this.securityGroupIds = {
            api: apiSecurityGroup.id,
            db: dbSecurityGroup.id,
        }
    }
}
