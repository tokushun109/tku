import { TerraformStack } from 'cdktf'
import * as aws from '@cdktf/provider-aws/lib'

export class SecretsManagerResource {
    constructor(private stack: TerraformStack, private name: string, private apiPublicIp: string, private dbPrivateIp: string = 'Not Found') {
        const secretsManager = new aws.secretsmanagerSecret.SecretsmanagerSecret(this.stack, `${this.name}-secrets-manager`, {
            name: `${this.name}-secrets-manager`,
            description: `${this.name}-secrets-manager`,
        })

        new aws.secretsmanagerSecretVersion.SecretsmanagerSecretVersion(this.stack, `${this.name}-secrets-manager-version`, {
            secretId: secretsManager.id,
            secretString: JSON.stringify({
                API_BASE_URL: process.env.API_BASE_URL,
                DB_PASS: process.env.DB_PASS,
                MYSQL_ROOT_PASSWORD: process.env.MYSQL_ROOT_PASSWORD,
                DB_USER: process.env.DB_USER,
                CREATOR_NAME: process.env.CREATOR_NAME,
                DB_NAME: process.env.DB_NAME,
                ENV: 'prod',
                MYSQL_HOST: this.dbPrivateIp,
                AWS_SECRET_ACCESS_KEY: process.env.AWS_SECRET_ACCESS_KEY,
                AWS_ACCESS_KEY_ID: process.env.AWS_ACCESS_KEY_ID,
                AWS_REGION: process.env.AWS_REGION,
                API_BUCKET_NAME: process.env.API_BUCKET_NAME,
                SEND_GRID_API_KEY: process.env.SEND_GRID_API_KEY,
                CLIENT_URL: process.env.CLIENT_URL,
                DOMAINS: `api.tocoriri.com -> http://${this.apiPublicIp}:8080`,
                STAGE: process.env.STAGE,
            }),
        })
    }
}
