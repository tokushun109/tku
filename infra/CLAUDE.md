# CLAUDE.md - Infrastructure

このファイルは、infraディレクトリでのインフラ開発における Claude Code (claude.ai/code)向けのガイドラインです。

## プロジェクト概要

CDK for Terraform によるAWSリソース管理。VPC、ECS、RDS、Lambda などのクラウドインフラをコードで管理しています。

## 開発コマンド

```bash
# 依存関係インストール
pnpm install

# TypeScript コンパイル
pnpm build

# Terraform設定生成
pnpm synth

# インフラ変更計画
cdktf plan

# インフラ変更適用
cdktf apply

# テスト実行
pnpm test
```

## 技術スタック

- **IaC**: CDK for Terraform
- **クラウド**: AWS
- **言語**: TypeScript
- **パッケージマネージャー**: pnpm (package.jsonではnpmスクリプト使用)

## 管理対象リソース

### ネットワーク
- **VPC**: Virtual Private Cloud
- **サブネット**: Public/Private サブネット
- **インターネットゲートウェイ**: 外部接続
- **NAT ゲートウェイ**: Private サブネットからの外部接続
- **セキュリティグループ**: ファイアウォール設定

### コンピューティング
- **ECS**: Elastic Container Service
  - クラスター設定
  - サービス定義
  - タスク定義
- **ECR**: Elastic Container Registry
  - Docker イメージ保存

### データベース
- **RDS**: Relational Database Service
  - MySQL インスタンス
  - セキュリティ設定
  - バックアップ設定

### ストレージ
- **S3**: Simple Storage Service
  - 画像ファイル保存
  - 静的ファイル配信

### その他
- **ALB**: Application Load Balancer
- **Route53**: DNS 管理
- **CloudFront**: CDN
- **Certificate Manager**: SSL/TLS 証明書

## ディレクトリ構造

```
infra/
├── main.ts              # CDK for Terraform のメインエントリーポイント
├── resources/           # AWSリソースの定義
│   ├── asm/            # AWS Systems Manager関連
│   ├── ec2/            # EC2関連（userData等）
│   ├── ecs/            # ECS関連（クラスター、サービス、タスク定義）
│   ├── eventBridge/    # EventBridge関連
│   ├── lambda/         # Lambda関数定義とハンドラー
│   └── network/        # VPC、サブネット、セキュリティグループ等
├── libs/               # ユーティリティライブラリ
│   ├── compile.ts      # コンパイル関連ユーティリティ
│   ├── convert.ts      # データ変換ユーティリティ
│   ├── date.ts         # 日付操作ユーティリティ
│   └── task.ts         # タスク関連ユーティリティ
├── cdktf.json          # CDK for Terraform設定ファイル
├── tsconfig.json       # TypeScript設定
├── package.json        # パッケージ管理
└── cdktf.out/          # 生成されるTerraformファイル（ビルド成果物）
```

## 開発方針

### 環境管理
- **開発環境**: development
- **本番環境**: production
- 環境ごとの設定ファイルで管理

### セキュリティ
- **最小権限の原則**: 必要最小限のアクセス権限
- **暗号化**: データの暗号化（保存時・転送時）
- **ネットワーク分離**: Public/Private サブネットの適切な分離

### 運用
- **モニタリング**: CloudWatch による監視
- **ログ**: 構造化ログの実装
- **バックアップ**: 定期的なデータバックアップ
- **災害復旧**: Multi-AZ 構成

## デプロイフロー

1. **計画確認**: `cdktf plan` で変更内容確認
2. **レビュー**: インフラ変更のコードレビュー
3. **適用**: `cdktf apply` で変更適用
4. **検証**: デプロイ後の動作確認

## 注意事項

- **本番環境への変更**: 必ず staging 環境で事前検証
- **状態管理**: Terraform State の適切な管理
- **コスト最適化**: 使用していないリソースの定期的な見直し
- **セキュリティ**: 定期的なセキュリティ監査の実施