# 環境別設定

OpenTofu の共通リソース定義は `infra/opentofu/` に置き、環境ごとの差分はこのディレクトリに置く。

- `production/`: 現行本番環境。State は `tku/production/terraform.tfstate`
- `development/`: 将来の開発環境。State は `tku/development/terraform.tfstate`

各 `*.example` をリポジトリ外の `backend.hcl` と `terraform.tfvars` にコピーして利用する。実行時は共通定義のディレクトリから、環境ごとの設定を明示して実行する。

```bash
cd infra/opentofu
tofu init -reconfigure -backend-config=/absolute/path/to/production/backend.hcl
tofu plan -var-file=/absolute/path/to/production/terraform.tfvars
```

開発環境を作成する前に、AWS のリソース名・Railway 環境・Amplify ブランチおよびドメインを production と分離する。
