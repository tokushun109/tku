# backend 置き換えランブック

## 目的

旧 `backend/api` 構成の API サーバーを廃止し、現在の `backend/` を正式なバックエンド実装として運用する。

このドキュメントは、リポジトリ側で適用した変更と、切り替え時に手作業で確認すべき項目を整理するためのランブックである。

## 現在の前提

- Railway の build / start は Railpack の自動検出を使う
- Railway の `Build Command` と `Start Command` は未設定
- Railway の `RAILWAY_DOCKERFILE_PATH` は未設定
- ローカル開発用の `.env` は `backend/.env` を利用する
- API 切り替え時にはダウンタイムを許容する

## リポジトリ側の反映内容

### ディレクトリ構成

- `backend/` 直下を新実装のルートとして扱う
- エントリーポイントは `backend/cmd/api/main.go`
- 旧 `backend/api` を前提にした参照は `backend` 基準へ置き換える

### Go module / import path

- Go module 名は `github.com/tokushun109/tku/backend`
- import path も `github.com/tokushun109/tku/backend/...` に統一する

### ローカル開発

- `docker-compose.yml` の `db` / `migrate` / `minio` / `api` はすべて `./backend` を参照する
- API は `backend/.env` を読み込み、`backend/docker/api/script/local/command.sh` で起動する

### CI / デプロイ

- GitHub Actions の Go job は `./backend` を対象にする
- Railway デプロイワークフローは `backend/**` の変更を監視する
- VS Code ワークスペースは `./backend` を直接開く

### ガイド類

- `backend/CLAUDE.md` を新構成向けに作り直す
- `backend/docs/` 配下の設計メモは `backend` 基準のパス表記へ更新する

## Railway で手作業が必要な項目

### API サービス

- `Root Directory` を `backend` に設定する

そのままでよい設定:

- `Build Command`
- `Start Command`
- `RAILWAY_DOCKERFILE_PATH`

### migrate サービス

- migrate サービスも同じリポジトリのサブディレクトリを参照している場合は、`Root Directory` を `backend` に合わせる
- 既に別設定になっている場合は、その設定を維持してよい

## 切り替え時の確認ポイント

- Railway のデプロイログで Go アプリが自動検出される
- API が `PORT` 環境変数で待ち受ける
- migrate サービスが新しい `backend/db/migrations` を参照できる
- `docker-compose up` で frontend から API に接続できる
- CI の backend job が `./backend` で成功する

## 検証コマンド

```bash
cd backend
go build ./...
go test ./...
```

## 備考

- root 直下の `.backup/` は作業用退避ディレクトリであり、Git 管理対象外
- 追加の運用メモは `backend/CLAUDE.md` と `backend/docs/` を参照する
