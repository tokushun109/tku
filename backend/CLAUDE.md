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

```
backend/
├── api/                     # Goアプリケーションのメインディレクトリ
│   ├── app/                # アプリケーションコア
│   │   ├── controllers/    # HTTP リクエストハンドラーとルーティング
│   │   ├── models/         # GORM データベースモデル
│   │   └── db/             # マイグレーションファイルと DB 設定
│   ├── config/             # アプリケーション設定管理
│   ├── utils/              # ユーティリティ関数
│   ├── docker/             # Docker関連ファイル
│   ├── main.go             # アプリケーションエントリーポイント
│   ├── go.mod              # Go modules設定
│   └── go.sum              # Go modules依存関係
└── CLAUDE.md               # 本ドキュメント
```

### 詳細構成

- **Controllers**: `/api/app/controllers/` - HTTP リクエストハンドラーとルーティング
- **Models**: `/api/app/models/` - GORM データベースモデル
- **Database**: `/api/app/db/` - マイグレーションファイルと DB 設定
- **Config**: `/api/config/` - アプリケーション設定管理
- **Utils**: `/api/utils/` - UUID生成、ログ、ディレクトリ操作等のユーティリティ

## 主要機能

- **商品管理**: ハンドメイドアクセサリーの CRUD 操作
- **画像アップロード**: S3 連携による商品画像管理
- **管理ダッシュボード**: CSV 一括編集機能付きコンテンツ管理
- **Web スクレイピング**: Creema マーケットプレイス連携による商品データ取得
- **お問い合わせフォーム**: SendGrid 連携によるメール通知
- **SEO**: SNS シェア用動的 OGP 生成

## データベース

### マイグレーション

- **場所**: `/api/app/db/migrations/`
- **管理**: golang-migrate による管理
- **本番環境**: Railway でマイグレーション実行
- **手動実行**: マイグレーション用 Docker コンテナ使用

### 設定

- **ローカル**: `.env`ファイルと`config.ini`使用
- **本番**: Railway の環境変数
- **データベース設定**: `/api/config/config.go`で設定

## 技術スタック

- **言語**: Go
- **ルーター**: Gorilla Mux
- **ORM**: GORM
- **データベース**: MySQL
- **デプロイ**: Railway にデプロイ
- **画像保存**: S3 (UUID ベース)

## デプロイ

- **プラットフォーム**: Railway にデプロイ
- **CI/CD**: GitHub Actions（`.github/workflows/ci.yml`）
- **方法**: Git Push によるデプロイ（Railway 自動ビルド）
- **データベース**: Railway MySQL インスタンス使用

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