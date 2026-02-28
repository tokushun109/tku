# CLAUDE.md - backend

このファイルは、`backend/` 配下で作業する開発者とエージェント向けの補助ガイドです。

## プロジェクト概要

- Go 製の REST API
- Gorilla Mux を使った HTTP ルーティング
- MySQL + sqlx ベースの永続化
- DDD + Clean Architecture を前提に `internal/` 配下へ責務分離

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
└── internal/infra/          # DB / mail / storage / marketplace などの実装
```

## 運用メモ

- Railway では `backend` を Root Directory として扱う
- ローカルの Docker 起動も `backend/.env` を前提にする
- 詳細な移行背景は `backend/docs/` 配下のドキュメントを参照する
