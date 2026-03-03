# CLAUDE.md

このファイルは、このリポジトリで作業するエージェント向けのルートガイドです。  
詳細な規約は各ディレクトリ配下の `CLAUDE.md` を参照し、このファイルでは全体像と作業導線をまとめます。

## 基本設定

- 回答は日本語で行う
- `/fix-issue` のカスタムコマンドが実行された場合は `../.claude/commands/fix-issue.md` を確認する

## プロジェクト概要

ハンドメイドアクセサリー作家「tku」の商品販売サイト（`tocoriri.com`）です。

- **frontend**: Next.js + TypeScript による公開サイト / 管理画面
- **backend**: Go 製 REST API
- **database**: MySQL（`golang-migrate` でスキーマ管理）
- **infra**: CDK for Terraform による運用系 Lambda / EventBridge 管理

## リポジトリ構成

```text
tku/
├── frontend/   # Next.js アプリ
├── backend/    # Go API
└── infra/      # 定期実行タスクなどの IaC
```

## ローカル開発

### 全体起動

```bash
docker-compose up
```

起動対象:

- フロントエンド: `http://localhost:3000`
- バックエンド API: `http://localhost:8080`
- MySQL: `localhost:3306`
- MinIO: `http://localhost:9000`

### 補足

- backend のローカル開発では `backend/.env` を利用する
- Docker 構成には `db` / `migrate` / `minio` / `api` / `frontend` が含まれる
- ローカルでも S3 互換ストレージを使った検証ができる

## アーキテクチャの要点

### frontend

- App Router ベースの SSR / SPA ハイブリッド
- 公開サイトと管理画面を同一アプリ内で運用
- Storybook / Vitest を使った UI 品質管理

### backend

- フレームワークに依存しすぎない Go + Gorilla Mux
- `domain / usecase / interface / infra` で責務分離
- `internal/app/di` を Composition Root として依存関係を集約
- product は Query / Command を分離
- S3 / SendGrid / Creema 連携を usecase から抽象化

### infra

- CDK for Terraform による運用タスク管理
- フロントエンド warmup 用 Lambda
- API ヘルスチェック用 Lambda

## 主要機能

- 商品管理（CRUD）
- 商品画像アップロード
- CSV による商品データ入出力
- Creema 商品情報を使った複製補助
- 管理者ログイン / セッション認証
- お問い合わせ保存 + メール通知
- 動的 OGP を含む SEO 対応

## デプロイ / 運用

- フロントエンド: AWS Amplify
- バックエンド: Railway
- 商品画像: AWS S3
- お問い合わせ通知: SendGrid
- 定期実行: AWS Lambda + EventBridge

## ディレクトリ別ガイド

- `frontend/CLAUDE.md`: フロントエンドの開発規約
- `backend/CLAUDE.md`: バックエンドの構成・実装ルール
- `infra/CLAUDE.md`: インフラ変更時の注意点

## 作業時の注意事項

- 日本語サイトのため、コンテンツや UI 文言は日本語前提で扱う
- `.gitignore` 対象は絶対にコミットしない
- `git add -f` は使用しない
- `.env*`、ビルド成果物、個人設定ファイルはコミットしない
- 大きな構成変更を行う場合は、各ディレクトリ配下の `CLAUDE.md` と `docs/` を先に確認する
