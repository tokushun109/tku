# CLAUDE.md

このファイルは、このリポジトリでコードを扱う際の Claude Code (claude.ai/code)向けのガイドラインです。

## Claude Code 設定

**重要**

- Claude の回答は日本語で行ってください。
- /fix-issue のカスタムコマンドが実行されたときは../.claude/commands/fix-issue.md の中身を全て確認してください。

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

詳細は各ディレクトリの CLAUDE.md ファイルを参照：

- **frontend/**: `/frontend/CLAUDE.md` - 新規開発用 Next.js
- **client/**: `/client/CLAUDE.md` - レガシー Vue.js
- **api/**: `/api/CLAUDE.md` - Go REST API
- **infra/**: `/infra/CLAUDE.md` - CDK for Terraform

## プロジェクトアーキテクチャ

詳細なプロジェクト構造は各ディレクトリの CLAUDE.md ファイルを参照：

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
- 詳細は各ディレクトリの CLAUDE.md を参照

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

各ディレクトリでの開発における詳細なガイドラインは、それぞれの CLAUDE.md ファイルを参照してください：

- **frontend/**: `/frontend/CLAUDE.md` - Next.js + TypeScript の開発規約とガイドライン
- **client/**: `/client/CLAUDE.md` - Vue.js + Nuxt.js のレガシーアプリ情報
- **api/**: `/api/CLAUDE.md` - Go REST API の開発ガイドライン
- **infra/**: `/infra/CLAUDE.md` - CDK for Terraform のインフラ開発ガイドライン

## Git操作における重要な注意事項

### .gitignoreの厳格な遵守

- **絶対禁止**: `.gitignore`で除外されているファイル・ディレクトリを**一切コミットしない**
- **強制追加禁止**: `git add -f`による強制追加は**絶対に使用しない**
- **対象ファイル例**:
  - `.vscode/settings.json`（個人設定ファイル）
  - `node_modules/`（依存関係）
  - `dist/`, `build/`（ビルド成果物）
  - `.env*`（環境変数ファイル）
  - ログファイル、一時ファイル

### 理由

- **個人設定の混入防止**: 開発者個人の設定が他の開発者に影響することを防ぐ
- **リポジトリサイズの制御**: 不要なファイルによるリポジトリ肥大化を防ぐ
- **セキュリティ**: 機密情報や環境固有の設定の漏洩を防ぐ
- **ビルド環境の一貫性**: ビルド成果物はCI/CDで生成し、環境間の一貫性を保つ
