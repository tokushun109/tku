# CLAUDE.md - Infrastructure

このファイルは、`infra/` ディレクトリで作業する開発者とエージェント向けの補助ガイドです。

## プロジェクト概要

`infra/opentofu/` では OpenTofu を使い、本番の AWS / Railway / Amplify リソースを管理しています。
`infra/lambda/` には、フロントエンド warmup 用とバックエンド API ヘルスチェック用 Lambda のコードを置きます。

## 開発コマンド

```bash
# OpenTofu の初期化・差分確認（本番 state）
cd opentofu
tofu init -reconfigure -backend-config=backend.hcl
tofu plan -var-file=production.tfvars

# Lambda コードのビルド
cd ../lambda
pnpm install
pnpm build
```

## 技術スタック

- **IaC**: OpenTofu
- **クラウド**: AWS
- **PaaS**: Railway
- **フロントエンドホスティング**: AWS Amplify
- **Lambda コード**: TypeScript + esbuild
- **パッケージマネージャー**: pnpm

## 管理対象リソース

### Lambda

- **warmup**
  - フロントエンドの応答性維持を目的に定期実行する
- **health-check**
  - API の生存確認を目的に定期実行する

### EventBridge

- warmup Lambda のスケジュール実行
- health-check Lambda のスケジュール実行

## ディレクトリ構造

```text
infra/
├── opentofu/                       # AWS / Railway / Amplify の IaC
│   └── bootstrap/                  # OpenTofu state 用 S3 の IaC
└── lambda/                         # Lambda コードと ZIP ビルド
    ├── handlers/
    │   ├── warmup/
    │   └── health-check/
    └── scripts/build.mjs
```

## 実装の要点

- AWS の通常リソースは `ap-northeast-1`、Amplify は `ap-southeast-1` を対象とする
- warmup は日本時間の稼働帯を考慮し、5 分間隔で実行する
- health-check は日本時間の稼働帯を考慮し、1 時間間隔で実行する
- スケジュールは UTC ベースの cron で記述しているため、変更時は時差を必ず確認する

## デプロイフロー

1. `tofu plan -var-file=production.tfvars` で差分を確認する
2. コードレビューで意図した変更だけが含まれているか確認する
3. `tofu apply -var-file=production.tfvars` で適用する
4. Lambda とスケジュールの動作を確認する

## 注意事項

- 適用前に `tofu plan` を必ず確認する
- cron を変更する場合は、日本時間と UTC の対応を確認する
- Lambda のコード更新では `infra/lambda` で ZIP をビルドし、AWS CLI で `update-function-code` を実行する。OpenTofu は Lambda の環境変数と配布済みアーカイブを変更しない。
- 実装済みリソースとドキュメントの乖離を作らない
