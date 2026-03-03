# CLAUDE.md - backend

このファイルは、`backend/` 配下で作業する開発者とエージェント向けの補助ガイドです。

## プロジェクト概要

- Go 製の REST API
- Gorilla Mux を使った HTTP ルーティング
- MySQL + sqlx ベースの永続化
- DDD + Clean Architecture を前提に `internal/` 配下へ責務分離
- `internal/app/di` を Composition Root として依存関係を組み立てる
- 商品機能は Query / Command を分離して実装している

## 開発コマンド

```bash
# ホットリロード開発
sh ./docker/api/script/local/command.sh

# 直接実行
go run ./cmd/api

# ビルド
go build -o ./bin/main ./cmd/api

# テスト
go test ./...
```

## 主要な提供機能

- 商品、カテゴリ、タグ、対象、販売サイト、スキルマーケットの管理 API
- 作家情報（プロフィール / ロゴ画像）の更新 API
- 商品画像のアップロード / 配信 API
- 商品 CSV のエクスポート / インポート
- Creema の商品情報を利用した商品複製補助
- ユーザー認証、セッション管理、管理者権限チェック
- お問い合わせ保存 + SendGrid による通知
- ヘルスチェック API

## ディレクトリ構成

```text
backend/
├── cmd/api/                 # アプリケーションのエントリーポイント
├── db/migrations/           # golang-migrate 用マイグレーション
├── docker/                  # ローカル開発 / DB / migrate 用設定
├── docs/                    # 設計メモと移行ドキュメント
├── internal/app/di/         # 依存注入
├── internal/domain/         # ドメインモデル
├── internal/usecase/        # ユースケース
├── internal/interface/http/ # handler / presenter / request / response / router
├── internal/infra/          # DB / mail / storage / marketplace などの実装
└── internal/shared/         # ロガーや共通 ID など
```

## レイヤ構成

### `internal/domain`

- Entity / Value Object / Repository interface を配置
- ドメインルールを外部実装から切り離して管理する

### `internal/usecase`

- アプリケーション固有のユースケースを配置
- product は `query` と `command` を分離している
- 外部依存は interface に抽象化し、infra で実装する

### `internal/interface/http`

- `handler`: HTTP エンドポイント
- `request`: リクエスト DTO
- `presenter` / `response`: レスポンス整形
- `middleware`: 認証、管理者権限、Origin 検証、ロギング、CORS
- `router`: ルーティング定義

### `internal/infra`

- `config`: 環境変数ロード
- `db/mysql`: DB 接続、Repository 実装、Query 実装、Tx 管理
- `storage/s3`: S3 連携
- `mail/sendgrid`: メール送信
- `marketplace/creema`: 商品複製用のスクレイピング実装
- `uuid` / `clock` / `crypto`: 補助的な実装

## ローカル開発メモ

- ローカルの Docker 起動は `backend/.env` を前提にする
- `docker-compose` では MySQL、migrate、MinIO、API が連携する
- `AWS_ENDPOINT_URL_S3` を設定すると、S3 互換ストレージ向けに PathStyle を有効化する
- API の起動スクリプトは `docker/api/script/local/command.sh`

## 実装上の注意点

- 新しい機能追加時は、既存のレイヤ分離を崩さない
- handler にビジネスロジックを書かず、usecase に寄せる
- 直接 DB を読む実装が必要でも、まず `usecase` 側の責務を確認する
- 商品関連は Query / Command のどちらに属するかを先に整理する
- 管理系の変更では認証・権限・Origin 制御への影響を確認する

## 運用メモ

- Railway では `backend` を Root Directory として扱う
- 詳細な移行背景は `backend/docs/` 配下のドキュメントを参照する
