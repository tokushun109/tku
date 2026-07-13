locals {
  lambda_functions = {
    warmup = {
      archive_bucket = "tku-warmup-lambda-archive-bucket"
      archive_key    = "tku/warmup/20250827203532/archive.zip"
      function_name  = "tku-warmup"
      schedule       = "cron(*/5 21-23,0-13 * * ? *)"
    }
    health_check = {
      archive_bucket = "tku-health-check-lambda-archive-bucket"
      archive_key    = "tku/health-check/20250827203532/archive.zip"
      function_name  = "tku-health-check"
      schedule       = "cron(0 21-23,0-11/1 * * ? *)"
    }
  }
}

resource "aws_s3_bucket" "lambda_archive" {
  for_each = local.lambda_functions

  bucket = each.value.archive_bucket

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_public_access_block" "lambda_archive" {
  for_each = local.lambda_functions

  bucket = aws_s3_bucket.lambda_archive[each.key].id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_server_side_encryption_configuration" "lambda_archive" {
  for_each = local.lambda_functions

  bucket = aws_s3_bucket.lambda_archive[each.key].id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_iam_role" "lambda" {
  for_each = local.lambda_functions

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
  name = "${each.value.function_name}-lambda-role"
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  for_each = local.lambda_functions

  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda[each.key].name
}

resource "aws_lambda_function" "scheduled" {
  for_each = local.lambda_functions

  architectures = ["x86_64"]
  function_name = each.value.function_name
  handler       = "index.handler"
  memory_size   = 512
  role          = aws_iam_role.lambda[each.key].arn
  runtime       = "nodejs22.x"
  s3_bucket     = each.value.archive_bucket
  s3_key        = each.value.archive_key
  timeout       = 30

  lifecycle {
    prevent_destroy = true
    # 既存の環境変数と配布済みアーカイブは、専用のシークレット／デプロイ移行後に管理する。
    ignore_changes = [environment, publish, s3_bucket, s3_key, source_code_hash]
  }
}

resource "aws_cloudwatch_event_rule" "lambda_schedule" {
  for_each = local.lambda_functions

  name                = "${each.value.function_name}-event-rule"
  schedule_expression = each.value.schedule
  state               = "ENABLED"
}

resource "aws_cloudwatch_event_target" "lambda_schedule" {
  for_each = local.lambda_functions

  arn       = aws_lambda_function.scheduled[each.key].arn
  rule      = aws_cloudwatch_event_rule.lambda_schedule[each.key].name
  target_id = "terraform-20250825214459${each.key == "warmup" ? "952200000002" : "880600000001"}"
}

resource "aws_lambda_permission" "eventbridge" {
  for_each = local.lambda_functions

  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.scheduled[each.key].function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.lambda_schedule[each.key].arn
  statement_id  = "AllowEventBridgeInvokeLambda"
}
