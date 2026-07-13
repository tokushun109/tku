# AWS runtime import 手順

この手順は既存リソースをOpenTofu Stateへ登録するだけで、AWSリソースを作成・変更しない。実行前に `tofu plan` の出力を確認し、置換または削除が表示された場合は中断する。

## 前提

- `infra/opentofu/bootstrap` によりリモートStateを作成済みであること
- `backend.hcl` と `production.tfvars` はリポジトリ外に保存すること
- `aws login --profile tku-terraform --region ap-northeast-1` が成功していること

## 初期化

```bash
cd infra/opentofu
tofu init -backend-config=/absolute/path/to/backend.hcl
export AWS_PROFILE=tku-terraform
```

## import

```bash
tofu import 'aws_s3_bucket.lambda_archive["warmup"]' tku-warmup-lambda-archive-bucket
tofu import 'aws_s3_bucket.lambda_archive["health_check"]' tku-health-check-lambda-archive-bucket
tofu import 'aws_s3_bucket_public_access_block.lambda_archive["warmup"]' tku-warmup-lambda-archive-bucket
tofu import 'aws_s3_bucket_public_access_block.lambda_archive["health_check"]' tku-health-check-lambda-archive-bucket
tofu import 'aws_s3_bucket_server_side_encryption_configuration.lambda_archive["warmup"]' tku-warmup-lambda-archive-bucket
tofu import 'aws_s3_bucket_server_side_encryption_configuration.lambda_archive["health_check"]' tku-health-check-lambda-archive-bucket
tofu import aws_s3_bucket.product_images tku-api-ck57lb-prod
tofu import aws_s3_bucket_ownership_controls.product_images tku-api-ck57lb-prod
tofu import aws_s3_bucket_versioning.product_images tku-api-ck57lb-prod
tofu import aws_s3_bucket_server_side_encryption_configuration.product_images tku-api-ck57lb-prod
tofu import 'aws_iam_role.lambda["warmup"]' tku-warmup-lambda-role
tofu import 'aws_iam_role.lambda["health_check"]' tku-health-check-lambda-role
tofu import 'aws_iam_role_policy_attachment.lambda_basic_execution["warmup"]' tku-warmup-lambda-role/arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
tofu import 'aws_iam_role_policy_attachment.lambda_basic_execution["health_check"]' tku-health-check-lambda-role/arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
tofu import 'aws_lambda_function.scheduled["warmup"]' tku-warmup
tofu import 'aws_lambda_function.scheduled["health_check"]' tku-health-check
tofu import 'aws_cloudwatch_event_rule.lambda_schedule["warmup"]' tku-warmup-event-rule
tofu import 'aws_cloudwatch_event_rule.lambda_schedule["health_check"]' tku-health-check-event-rule
tofu import 'aws_cloudwatch_event_target.lambda_schedule["warmup"]' tku-warmup-event-rule/terraform-20250825214459952200000002
tofu import 'aws_cloudwatch_event_target.lambda_schedule["health_check"]' tku-health-check-event-rule/terraform-20250825214459880600000001
tofu import 'aws_lambda_permission.eventbridge["warmup"]' tku-warmup/AllowEventBridgeInvokeLambda
tofu import 'aws_lambda_permission.eventbridge["health_check"]' tku-health-check/AllowEventBridgeInvokeLambda

# Amplify（ap-southeast-1）
tofu import aws_amplify_app.production d2q4f71iidth8s
tofu import aws_amplify_branch.production d2q4f71iidth8s/main
tofu import aws_amplify_domain_association.production d2q4f71iidth8s/tocoriri.com
```

## 確認

```bash
tofu plan -var-file=/absolute/path/to/production.tfvars
```

既存のLambda環境変数とアーカイブ更新は意図的に `ignore_changes` としている。これらを管理対象に移す前に、シークレットの保管先とLambda配布方法を決定する。
