# OpenTofu 移行基盤

CDKTF は 2025 年 12 月に非推奨化されたため、AWS・Amplify・Railway のIaCは OpenTofu のHCLで管理する。ここでは既存リソースを再作成せず、必ず import してから管理を開始する。

## このディレクトリの責務

- `bootstrap/`: OpenTofu State用S3バケットを作成する最初の1回だけの構成
- ルート: プロバイダー、共通タグ、環境変数、Stateバックエンドの契約
- 今後追加するモジュール: AWS runtime、Amplify、Railway

現在はリモートの棚卸し前であるため、リソース定義を意図的に追加していない。実体と異なる定義を先に適用して、Lambda・Amplify・Railwayを再作成する事故を防ぐためである。

## 初期化

Stateバケット作成前はローカルStateを使う。

```bash
cd infra/opentofu/bootstrap
tofu init
tofu apply -var='state_bucket_name=<globally-unique-name>'
```

作成後は、リポジトリ外に `backend.hcl` を作成して各環境のStateを初期化する。

```bash
cd infra/opentofu
# backend.hcl.example を参照して、リポジトリ外に backend.hcl を作成する
tofu init -backend-config=/absolute/path/to/backend.hcl
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

PRでは `fmt` と `validate` だけを実行する。認証情報を必要とする `plan` と `apply` は、AWS OIDCとRailway API Tokenの権限・State分離・import完了後に別ワークフローで有効化する。
