resource "aws_iam_openid_connect_provider" "github_actions" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = ["sts.amazonaws.com"]
}

data "aws_iam_policy_document" "github_actions_plan_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }

    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:sub"
      values   = ["repo:tokushun109/tku:pull_request"]
    }

    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.github_actions.arn]
    }
  }
}

resource "aws_iam_role" "github_actions_opentofu_plan" {
  name               = "tku-github-actions-opentofu-plan"
  assume_role_policy = data.aws_iam_policy_document.github_actions_plan_assume_role.json
}

data "aws_iam_policy_document" "github_actions_opentofu_plan" {
  statement {
    sid = "ListOpenTofuProductionState"

    actions = ["s3:ListBucket"]

    resources = [aws_s3_bucket.state.arn]

    condition {
      test     = "StringLike"
      variable = "s3:prefix"
      values = [
        "tku/production/terraform.tfstate",
        "tku/production/terraform.tfstate.tflock",
      ]
    }
  }

  statement {
    sid = "ReadOpenTofuProductionState"

    actions = [
      "s3:GetObject",
      "s3:GetObjectVersion",
    ]

    resources = ["${aws_s3_bucket.state.arn}/tku/production/terraform.tfstate"]
  }

  statement {
    sid = "LockOpenTofuProductionState"

    actions = [
      "s3:DeleteObject",
      "s3:GetObject",
      "s3:PutObject",
    ]

    resources = ["${aws_s3_bucket.state.arn}/tku/production/terraform.tfstate.tflock"]
  }

  statement {
    sid = "ReadManagedS3Buckets"

    actions = [
      "s3:GetBucketAcl",
      "s3:GetBucketEncryption",
      "s3:GetBucketLocation",
      "s3:GetBucketOwnershipControls",
      "s3:GetBucketPolicy",
      "s3:GetBucketPolicyStatus",
      "s3:GetBucketPublicAccessBlock",
      "s3:GetBucketTagging",
      "s3:GetBucketVersioning",
      "s3:ListBucket",
    ]

    resources = [
      "arn:aws:s3:::tku-api-ck57lb-prod",
      "arn:aws:s3:::tku-health-check-lambda-archive-bucket",
      "arn:aws:s3:::tku-warmup-lambda-archive-bucket",
    ]
  }

  statement {
    sid = "ReadManagedInfrastructure"

    actions = [
      "amplify:GetApp",
      "amplify:GetBranch",
      "amplify:GetDomainAssociation",
      "amplify:ListTagsForResource",
      "events:DescribeRule",
      "events:ListTargetsByRule",
      "events:ListTagsForResource",
      "iam:GetRole",
      "iam:GetRolePolicy",
      "iam:ListAttachedRolePolicies",
      "iam:ListRolePolicies",
      "iam:ListRoleTags",
      "lambda:GetFunction",
      "lambda:GetPolicy",
      "lambda:ListTags",
    ]

    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "github_actions_opentofu_plan" {
  name   = "tku-github-actions-opentofu-plan"
  role   = aws_iam_role.github_actions_opentofu_plan.id
  policy = data.aws_iam_policy_document.github_actions_opentofu_plan.json
}
