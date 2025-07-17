# CLAUDE.md - Client (レガシーフロントエンド)

このファイルは、clientディレクトリでのコード開発における Claude Code (claude.ai/code)向けのガイドラインです。

## プロジェクト概要

Vue.js 2 + Nuxt.js 2 によるレガシーフロントエンドアプリケーション。現在は frontend/ ディレクトリの Next.js + TypeScript への移行中のため、新規開発は frontend/ で行ってください。

## 開発コマンド

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

## プロジェクト構造

- **Pages**: `/pages/` - Nuxt.js 自動ルーティング対応ページ
- **Components**: `/components/` - 機能別に整理された Vue.js コンポーネント
- **管理画面**: `/pages/admin/` - 自作の管理者向けコンテンツ管理画面
- **Types**: `/types/` - TypeScript 型定義
- **Store**: `/store/` - Vuex 状態管理

## 開発方針

### 移行について

- **現状**: Vue.js 2 + Nuxt.js 2 のレガシーアプリケーション
- **移行先**: frontend/ ディレクトリの Next.js + TypeScript
- **新規開発**: frontend/ ディレクトリで実施
- **参考実装**: 新規実装時はこのディレクトリの既存コードを参考にしてUI/UXを忠実に再現

### メンテナンス方針

- **bug修正**: 必要最小限に留める
- **新機能**: frontend/ ディレクトリで実装
- **リファクタリング**: 移行完了まで原則実施しない

## デプロイ

- **方法**: Serverless Framework → AWS Lambda + API Gateway
- **コマンド**: `yarn deploy`
- **CI/CD**: GitHub Actions（`.github/workflows/ci.yml`）で自動デプロイ