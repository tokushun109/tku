# CLAUDE.md

このファイルは、このリポジトリでコードを扱う際の Claude Code (claude.ai/code)向けのガイドラインです。

## Claude Code設定

**重要**: Claude の回答は日本語で行ってください。

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

### 各ディレクトリ固有の開発コマンド

詳細は各ディレクトリのCLAUDE.mdファイルを参照：
- **frontend/**: `/frontend/CLAUDE.md` - 新規開発用Next.js
- **client/**: `/client/CLAUDE.md` - レガシーVue.js
- **api/**: `/api/CLAUDE.md` - Go REST API
- **infra/**: `/infra/CLAUDE.md` - CDK for Terraform

## プロジェクトアーキテクチャ

詳細なプロジェクト構造は各ディレクトリのCLAUDE.mdファイルを参照：
- **frontend/**: `/frontend/CLAUDE.md` - Next.js + TypeScript 構造
- **client/**: `/client/CLAUDE.md` - Vue.js + Nuxt.js 構造  
- **api/**: `/api/CLAUDE.md` - Go REST API 構造
- **infra/**: `/infra/CLAUDE.md` - AWS インフラ構造

### 主要機能

- **商品管理**: ハンドメイドアクセサリーの CRUD 操作
- **画像アップロード**: S3 連携による商品画像管理
- **管理ダッシュボード**: CSV 一括編集機能付きコンテンツ管理
- **Web スクレイピング**: Creema マーケットプレイス連携による商品データ取得
- **お問い合わせフォーム**: SendGrid 連携によるメール通知
- **SEO**: SNS シェア用動的 OGP 生成

### デプロイ

- **CI/CD**: GitHub Actions（`.github/workflows/ci.yml`）
- **インフラ**: CDK for Terraform が VPC、ECS、RDS などを管理
- 詳細は各ディレクトリのCLAUDE.mdを参照

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
## 各ディレクトリ固有の詳細情報

各ディレクトリでの開発における詳細なガイドラインは、それぞれのCLAUDE.mdファイルを参照してください：

- **frontend/**: `/frontend/CLAUDE.md` - Next.js + TypeScript の開発規約とガイドライン
- **client/**: `/client/CLAUDE.md` - Vue.js + Nuxt.js のレガシーアプリ情報
- **api/**: `/api/CLAUDE.md` - Go REST API の開発ガイドライン
- **infra/**: `/infra/CLAUDE.md` - CDK for Terraform のインフラ開発ガイドライン
