# CLAUDE.md - API (バックエンド)

このファイルは、apiディレクトリでのコード開発における Claude Code (claude.ai/code)向けのガイドラインです。

## プロジェクト概要

Go REST API with GORM によるバックエンドアプリケーション。フレームワークなしの Go + Gorilla Mux で構築。

## 開発コマンド

```bash
# ホットリロード開発（air使用）
go install github.com/cosmtrek/air@v1.40.0
air

# 直接実行
go run main.go

# ビルド
go build -o bin/main main.go
```

## プロジェクト構造

- **Controllers**: `/app/controllers/` - HTTP リクエストハンドラーとルーティング
- **Models**: `/app/models/` - GORM データベースモデル
- **Database**: `/app/db/` - マイグレーションファイルと DB 設定
- **Config**: `/config/` - アプリケーション設定管理

## 主要機能

- **商品管理**: ハンドメイドアクセサリーの CRUD 操作
- **画像アップロード**: S3 連携による商品画像管理
- **管理ダッシュボード**: CSV 一括編集機能付きコンテンツ管理
- **Web スクレイピング**: Creema マーケットプレイス連携による商品データ取得
- **お問い合わせフォーム**: SendGrid 連携によるメール通知
- **SEO**: SNS シェア用動的 OGP 生成

## データベース

### マイグレーション

- **場所**: `/app/db/migrations/`
- **管理**: golang-migrate による管理
- **本番環境**: ECS タスクによる自動実行
- **手動実行**: マイグレーション用 Docker コンテナ使用

### 設定

- **ローカル**: `.env`ファイルと`config.ini`使用
- **本番**: AWS サービスからの環境変数
- **データベース設定**: `/config/config.go`で設定

## 技術スタック

- **言語**: Go
- **ルーター**: Gorilla Mux
- **ORM**: GORM
- **データベース**: MySQL
- **クラウド**: AWS ECS にデプロイ
- **画像保存**: S3 (UUID ベース)

## デプロイ

- **CI/CD**: GitHub Actions（`.github/workflows/ci.yml`）
- **方法**: Docker イメージを Amazon ECR → ECS デプロイ
- **インフラ**: CDK for Terraform が ECS、RDS などを管理

## 開発方針

### アーキテクチャ

- **設計**: RESTful API
- **認証**: 必要に応じて実装
- **エラーハンドリング**: 適切なHTTPステータスコードとエラーレスポンス
- **ログ**: 構造化ログの実装

### コーディング規約

- **パッケージ構成**: 機能別にパッケージを分割
- **命名**: Go の慣例に従う（CamelCase、省略形の適切な使用）
- **エラーハンドリング**: Go らしいエラーハンドリングの実装
- **テスト**: 重要な機能にはユニットテストを実装