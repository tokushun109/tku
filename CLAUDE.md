# CLAUDE.md

このファイルは、このリポジトリでコードを扱う際の Claude Code (claude.ai/code)向けのガイドラインです。

## プロジェクト概要

ハンドメイドアクセサリー作家「tku」の商品販売サイト（tocoriri.com）です。プロジェクト構成：

- **フロントエンド**: Next.js with TypeScript への移行中（従来の Vue.js 2 + Nuxt.js 2 から置き換え）
- **バックエンド**: Go REST API with GORM、AWS ECS にデプロイ
- **データベース**: MySQL with golang-migrate によるスキーマ管理
- **インフラ**: CDK for Terraform による AWS リソース管理

## 開発コマンド

### ローカル開発（Docker）

```bash
# 全サービス起動（API、Client、Database）
docker-compose up

# サービスアクセス:
# - フロントエンド: http://localhost:3000
# - バックエンドAPI: http://localhost:8080
# - データベース: localhost:3306
```

### フロントエンド（新規開発：frontend/）

```bash
# 依存関係インストール
yarn install

# 開発サーバー
yarn dev

# ビルド
yarn build

# コードリント
yarn lint
```

### レガシーフロントエンド（client/）

```bash
# 依存関係インストール
yarn install

# 開発サーバー
yarn dev

# ビルド＆AWS Lambdaデプロイ
yarn deploy

# コードリント
yarn lint
```

### バックエンド（api/）

```bash
# ホットリロード開発（air使用）
go install github.com/cosmtrek/air@v1.40.0
air

# 直接実行
go run main.go

# ビルド
go build -o bin/main main.go
```

### インフラ（infra/）

```bash
# 依存関係インストール
yarn install

# インフラ変更計画
cdktf plan

# インフラ変更適用
cdktf apply

# Terraform設定生成
cdktf synth
```

## プロジェクトアーキテクチャ

### 新規フロントエンド構造（frontend/）

- **Pages/App Router**: Next.js App Router による自動ルーティング
- **Components**: 機能別に整理された React コンポーネント
- **管理画面**: 自作の管理者向けコンテンツ管理画面
- **Types**: TypeScript 型定義
- **Store**: 状態管理（Redux Toolkit または Zustand）

### レガシーフロントエンド構造（client/）

- **Pages**: `/pages/` - Nuxt.js 自動ルーティング対応ページ
- **Components**: `/components/` - 機能別に整理された Vue.js コンポーネント
- **管理画面**: `/pages/admin/` - 自作の管理者向けコンテンツ管理画面
- **Types**: `/types/` - TypeScript 型定義
- **Store**: `/store/` - Vuex 状態管理

### バックエンド構造

- **Controllers**: `/app/controllers/` - HTTP リクエストハンドラーとルーティング
- **Models**: `/app/models/` - GORM データベースモデル
- **Database**: `/app/db/` - マイグレーションファイルと DB 設定
- **Config**: `/config/` - アプリケーション設定管理

### 主要機能

- **商品管理**: ハンドメイドアクセサリーの CRUD 操作
- **画像アップロード**: S3 連携による商品画像管理
- **管理ダッシュボード**: CSV 一括編集機能付きコンテンツ管理
- **Web スクレイピング**: Creema マーケットプレイス連携による商品データ取得
- **お問い合わせフォーム**: SendGrid 連携によるメール通知
- **SEO**: SNS シェア用動的 OGP 生成

### データベースマイグレーション

- 場所: `/api/app/db/migrations/`
- golang-migrate による管理
- 本番環境では ECS タスクによる自動実行
- 手動実行: マイグレーション用 Docker コンテナ使用

### 設定

- **ローカル**: `.env`ファイルと`config.ini`使用
- **本番**: AWS サービスからの環境変数
- **データベース**: `/api/config/config.go`で設定

### デプロイ

- **CI/CD**: GitHub Actions（`.github/workflows/ci.yml`）
- **フロントエンド**: Serverless Framework → AWS Lambda + API Gateway
- **バックエンド**: Docker イメージを Amazon ECR → ECS デプロイ
- **インフラ**: CDK for Terraform が VPC、ECS、RDS などを管理

## 重要な注意事項

- 日本語の EC サイト（コンテンツは日本語）
- **フロントエンド移行中**: Next.js で新規開発（frontend/）、レガシーコードは client/ に残存
- 新規実装では元の client/ アプリの外観と動作を忠実に再現することが目標
- バックエンドはフレームワークなしの Go + Gorilla Mux
- データベースマイグレーションはバージョン管理され ECS タスクで実行
- 画像は S3 に UUID ベースで保存

## 開発方針

### フロントエンド移行について

- **移行対象**: client/ ディレクトリの Vue.js 2 + Nuxt.js 2 アプリケーション
- **移行先**: frontend/ ディレクトリの Next.js + TypeScript アプリケーション
- **移行目標**: 元のアプリケーションの UI/UX を忠実に再現し、機能を完全に移植
- **並行開発**: レガシーコードを参照しながら新規実装を進める
