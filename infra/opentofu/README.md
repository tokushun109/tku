# OpenTofu 移行基盤

CDKTF は 2025 年 12 月に非推奨化されたため、AWS・Amplify・Railway のIaCは OpenTofu のHCLで管理する。ここでは既存リソースを再作成せず、必ず import してから管理を開始する。

## このディレクトリの責務

- `bootstrap/`: OpenTofu State用S3バケットを作成する最初の1回だけの構成
- ルート: プロバイダー、共通タグ、環境変数、Stateバックエンドの契約
- `environments/`: production / development ごとのState・変数設定テンプレート
- 今後追加するモジュール: AWS runtime、Amplify、Railway

Lambda・EventBridge・LambdaアーカイブS3については、読み取り専用の棚卸し結果をもとに既存のproductionリソースを定義している。import手順は [`IMPORT.md`](./IMPORT.md) を参照する。Amplify・Railway・商品画像S3は、詳細設定を棚卸ししてから追加する。

## 初期化

Stateバケット作成前はローカルStateを使う。

```bash
cd infra/opentofu/bootstrap
tofu init -backend=false
tofu apply -var='state_bucket_name=<globally-unique-name>'
```

作成後は、`bootstrap/backend.hcl.example` を参照してリポジトリ外に bootstrap 用 `backend.hcl` を作成し、ローカルStateをS3へ移行する。

```bash
cd infra/opentofu/bootstrap
tofu init -migrate-state -backend-config=/absolute/path/to/bootstrap/backend.hcl
```

作成後は、[`environments/`](./environments/) のテンプレートを使って、リポジトリ外に環境ごとの `backend.hcl` と `terraform.tfvars` を作成する。

```bash
cd infra/opentofu
# environments/production のテンプレートを参照して、リポジトリ外に設定を作成する
tofu init -reconfigure -backend-config=/absolute/path/to/production/backend.hcl
tofu plan -var-file=/absolute/path/to/production/terraform.tfvars
tofu fmt -check -recursive
tofu validate
```

`backend.hcl`、`*.tfvars`、Stateファイルはコミットしない。シークレット値はOpenTofu設定やStateに置かず、GitHub ActionsのSecrets、AWS Secrets Manager、Railwayのsealed variablesのいずれかで管理する。

## import 手順

1. AWS、Amplify、RailwayのリソースID、設定、依存関係を読み取り専用で棚卸しする。
2. 管理対象ごとにHCLを記述し、既存リソースを `tofu import` する。
3. 直後に `tofu plan` を実行し、意図しない作成・削除・置換がないことを確認する。
4. `plan` が連続して差分ゼロになってから、CIでproduction applyを有効化する。

### 管理順序

1. AWS: Lambda archive用S3、Lambda、EventBridge、IAM
2. Amplify: App、production branch、ビルド設定、redirect、domain association
3. Railway: project、environment、service、domain、非機密設定

Railwayのbuild/deploy設定は、現在のDashboard設定を確認してから `backend/railway.toml` に反映する。RailwayのConfig as Codeはコード側を優先するため、確認前に追加しない。

## CI

PRでは `fmt` と `validate` を実行する。GitHub Actions の設定完了後は、同一リポジトリからのPRで `plan` も実行する。`apply` は別タスクで有効化する。

同一リポジトリから作成されたPRでは、`opentofu-plan` ジョブも実行する。このジョブは GitHub Actions OIDC で `tku-github-actions-opentofu-plan` Role を引き受け、production State の読み取りと lock file の操作、および管理対象 AWS リソースの読み取りだけを許可する。`tofu apply` 権限は付与しない。

有効化前に、bootstrap を適用して次の GitHub Actions 設定を登録する。

- Variable `AWS_PLAN_ROLE_ARN`: `tofu -chdir=infra/opentofu/bootstrap output -raw github_actions_opentofu_plan_role_arn` の出力値
- Variable `TOFU_STATE_BUCKET`: OpenTofu State バケット名
- Secret `RAILWAY_TOKEN`: Railway Provider 用 Workspace API token

外部 fork のPRではこのジョブを実行しない。`pull_request_target` は使用しないため、外部 fork に Secrets を渡すことはない。

`AWS_PLAN_ROLE_ARN` または `TOFU_STATE_BUCKET` が未設定の間は、`opentofu-plan` ジョブをスキップする。Role の作成と GitHub Actions 設定が完了してから有効になるため、bootstrap 用の変更を含む最初のPRで認証エラーは発生しない。
