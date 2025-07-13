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

## コーディング規約

### Sass/SCSS規約

#### @use を使用する

Dart Sass 3.0.0 で `@import` が削除される予定のため、`@use` を使用してください。

```scss
// ❌ 非推奨（廃止予定）
@import '@/styles/variables.scss';
@import '@/styles/mixins.scss';

// ✅ 推奨
@use '@/styles/variables.scss' as *;
@use '@/styles/mixins.scss' as *;
```

#### 変数・mixin の読み込み

- **variables**: `@use '@/styles/variables.scss' as *;`
- **mixins**: `@use '@/styles/mixins.scss' as *;`
- `as *` を使用することで名前空間なしで変数やmixinを使用可能

#### メディアクエリ mixins の使用

レスポンシブデザインには `@include media('breakpoint')` を使用してください。

```scss
// ✅ 推奨
.component {
  @include media('sm') {
    // スマートフォン向けスタイル
  }
  
  @include media('md') {
    // タブレット向けスタイル
  }
}

// ❌ 非推奨
.component {
  @include sm() {
    // 古い記法
  }
}
```

### CSS命名規約

#### ケバブケース（kebab-case）を使用

CSS クラス名はケバブケースで命名してください。

```scss
// ✅ 推奨
.error-wrapper {
  .site-title-area {
    .site-title {
      // スタイル
    }
  }
}
```

#### CSS Modules での参照

```tsx
// ✅ ケバブケースの参照方法
<div className={styles['error-wrapper']}>
  <div className={styles['site-title-area']}>
    <img className={styles['site-title']} />
  </div>
</div>
```

### TypeScript規約

#### 厳格な型定義

- `any` の使用を避け、適切な型定義を行う
- interfaceで明確な型を定義する
- プロパティの省略可能性を明示（`?:`）

```tsx
// ✅ 推奨
interface ErrorPageProps {
  errorMessage: React.ReactNode
  statusCode?: number
  showHomeButton?: boolean
}
```

### React/Next.js規約

#### コンポーネント設計

- **再利用性**: 既存のコンポーネントを優先的に使用
- **Props設計**: デフォルト値を適切に設定
- **命名**: コンポーネント名はPascalCase
- **スタイル上書き禁止**: 親コンポーネントから子コンポーネントのスタイルを`!important`で上書きしない

```tsx
// ✅ 推奨
const ErrorPage: React.FC<ErrorPageProps> = ({ 
  errorMessage, 
  statusCode, 
  showHomeButton = true 
}) => {
  // 実装
}

// ❌ 非推奨 - 子コンポーネントのスタイル上書き
.parent-component {
  .child-component {
    padding: 10px !important; // 禁止
  }
}

// ✅ 推奨 - コンポーネント自体でプロパティを提供
<ChildComponent size="small" />
```

#### インポート順序

ESLintルールに従った順序でインポートを記述：

```tsx
// ✅ 推奨順序
// 1. React関連
import React from 'react'
import { useState } from 'react'

// 2. Next.js関連
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

// 3. 外部ライブラリ
import axios from 'axios'

// 4. 内部モジュール（@/で始まる）
import { Button } from '@/components/bases/Button'
import { IClientError } from '@/utils/error'

// 5. 相対パス
import styles from './styles.module.scss'
```

### SEO対策

#### インデックス制御

エラーページ、メンテナンスページなどはインデックスさせない：

```tsx
// ✅ 推奨
<Head>
  <meta name="robots" content="noindex, nofollow" />
</Head>
```

### エラーハンドリング

#### API呼び出し

全てのAPI関数にtry-catchエラーハンドリングを実装：

```tsx
// ✅ 推奨
export const getProducts = async (): Promise<Product[]> => {
  try {
    const response = await fetch('/api/products')
    if (!response.ok) throw new ApiError(response)
    return await response.json()
  } catch (error) {
    if (error instanceof ApiError) {
      throw error
    }
    throw new Error('商品一覧の取得に失敗しました')
  }
}
```
