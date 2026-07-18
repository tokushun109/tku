resource "aws_amplify_app" "production" {
  provider = aws.singapore

  name                 = "tku"
  platform             = "WEB_COMPUTE"
  repository           = "https://github.com/tokushun109/tku"
  iam_service_role_arn = "arn:aws:iam::418549683327:role/service-role/AmplifySSRLoggingRole-8ac03953-f06a-45b5-bccb-e7f960c811dd"

  custom_rule {
    source = "/<*>"
    status = "404-200"
    target = "/index.html"
  }

  lifecycle {
    prevent_destroy = true
    # Console連携トークン・環境変数・ビルド設定は、シークレット管理方針を確定するまで維持する。
    ignore_changes = [access_token, build_spec, custom_headers, environment_variables, oauth_token]
  }
}

resource "aws_amplify_branch" "production" {
  provider = aws.singapore

  app_id                      = aws_amplify_app.production.id
  branch_name                 = "main"
  enable_auto_build           = true
  enable_basic_auth           = false
  enable_notification         = false
  enable_performance_mode     = false
  enable_pull_request_preview = false
  framework                   = "Next.js - SSR"
  stage                       = "PRODUCTION"

  lifecycle {
    ignore_changes = [basic_auth_credentials, environment_variables]
  }
}

resource "aws_amplify_domain_association" "production" {
  provider = aws.singapore

  app_id                 = aws_amplify_app.production.id
  domain_name            = "tocoriri.com"
  enable_auto_sub_domain = false
  wait_for_verification  = false

  sub_domain {
    branch_name = aws_amplify_branch.production.branch_name
    prefix      = ""
  }

  sub_domain {
    branch_name = aws_amplify_branch.production.branch_name
    prefix      = "www"
  }
}
