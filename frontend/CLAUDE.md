# CLAUDE.md - Frontend

このファイルは、frontendディレクトリでのコード開発における Claude Code (claude.ai/code)向けのガイドラインです。

## プロジェクト概要

Next.js 15.4.1 + TypeScript による最新のフロントエンドアプリケーション。従来の Vue.js 2 + Nuxt.js 2 から移行中で、現代的な開発ベストプラクティスを適用。

## 技術スタック

- **フレームワーク**: Next.js 15.4.1（App Router）、React 19.1.0
- **言語**: TypeScript 5.8.3
- **スタイリング**: Sass 1.89.2 + CSS Modules
- **UI ライブラリ**: Material-UI 7.2.0、Emotion
- **アニメーション**: Framer Motion 12.23.6、React Transition Group
- **フォーム**: React Hook Form 7.60.0 + Zod 4.0.5
- **テスト**: Vitest 3.2.4、Testing Library、Playwright、MSW 2.10.4
- **開発ツール**: Storybook 9.0.17、ESLint 9.31.0、Prettier 3.6.2

## 開発コマンド

```bash
# 依存関係インストール
pnpm install

# 開発サーバー
pnpm dev

# 開発サーバー（ポート3001で起動）
pnpm dev:3001

# ビルド
pnpm build

# リント（警告ゼロで実行）
pnpm lint

# テスト（Unit）
pnpm test

# テスト（Integration）
pnpm test-integration

# テスト（Storybook）
pnpm test-storybook

# Storybook起動
pnpm sb

# コードフォーマット
pnpm fmt
```

## プロジェクト構造

```
frontend/
├── src/
│   ├── app/                       # App Router（Next.js 15）
│   │   ├── (home)/               # ホームページ用ルートグループ
│   │   ├── (contents)/           # コンテンツページ用ルートグループ
│   │   └── (maintenance)/        # メンテナンスページ用ルートグループ
│   ├── components/               # コンポーネント階層
│   │   ├── bases/               # 基本コンポーネント
│   │   │   ├── Button/          # ボタン関連
│   │   │   ├── Input/           # 入力関連
│   │   │   ├── Card/            # 表示関連
│   │   │   └── ...              # その他基本コンポーネント
│   │   ├── composites/          # 複合コンポーネント
│   │   │   ├── Carousel/        # カルーセル関連
│   │   │   ├── SlideShow/       # スライドショー
│   │   │   └── ...              # その他複合コンポーネント
│   │   ├── animations/          # アニメーションコンポーネント
│   │   └── layouts/             # レイアウトコンポーネント
│   │       ├── Header/          # ヘッダー
│   │       ├── Footer/          # フッター
│   │       └── ...              # その他レイアウト
│   ├── features/                # 機能別コンポーネント
│   │   ├── product/             # 商品関連機能
│   │   ├── contact/             # お問い合わせ機能
│   │   ├── creator/             # 作者情報機能
│   │   └── classification/      # 分類機能
│   ├── apis/                    # API関数
│   ├── types/                   # TypeScript型定義
│   ├── utils/                   # ユーティリティ関数
│   ├── styles/                  # グローバルスタイル
│   │   ├── variables.scss       # 変数定義
│   │   ├── mixins.scss          # Mixin定義
│   │   └── layouts.scss         # レイアウト用スタイル
│   └── __tests__/              # テスト関連
├── public/                      # 静的ファイル
├── .storybook/                  # Storybook設定
└── storybook-static/           # Storybookビルド成果物
```

## 開発環境の特徴

### テスト環境（3層構造）

1. **Unit Test**: 個別コンポーネント・関数テスト
2. **Integration Test**: ページレベル統合テスト
3. **Storybook Test**: Storybookコンポーネントテスト

### Storybook

- **バージョン**: 9.0.17系
- **アドオン**: a11y、docs、controls、vitest統合
- **カバレッジ**: 全基本コンポーネントで充実したStories
- **品質**: 複数状態・バリエーション、実用的ユースケース対応

### 品質管理

- **Husky + lint-staged**: コミット前自動品質チェック
- **厳格なESLint**: perfectionist等による統一コード規約
- **型安全性**: Zodによるスキーマ検証
- **最大警告数**: 0（`--max-warnings=0`）

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

#### 全体読み込み設定との重複を避ける

`next.config.ts`の`prependData`で全体に読み込まれているモジュールは、各ファイルで重複して読み込まないでください。

```typescript
// next.config.ts での設定
sassOptions: {
  prependData: '@use "sass:color"; @use "@/styles/variables.scss" as *; @use "@/styles/mixins.scss" as *; @use "@/styles/layouts.scss" as *;',
},
```

```scss
// ❌ 非推奨 - next.config.tsで既に読み込まれている
@use 'sass:color';
@use '@/styles/variables.scss' as *;
@use '@/styles/mixins.scss' as *;
@use '@/styles/layouts.scss' as *;

.my-component {
    color: $primary;
    @include media('sm') {
        // スタイル
    }
}

// ✅ 推奨 - 必要に応じて追加のモジュールのみ読み込み
@use 'sass:math';

.my-component {
    color: $primary; // variables.scss から利用可能
    @include media('sm') {
        // mixins.scss から利用可能
        // スタイル
    }
}
```

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
- **useState使用時は必ず型を明示**: `useState<型>(初期値)` の形式で型を記述する
- **ユニオンタイプはオブジェクトリテラルで定義**: 文字列リテラルのユニオンタイプではなく、オブジェクトリテラルを使用

```tsx
// ✅ 推奨
interface ErrorPageProps {
    errorMessage: React.ReactNode
    statusCode?: number
    showHomeButton?: boolean
}

// ✅ 推奨 - useState with types
const [isLoading, setIsLoading] = useState<boolean>(false)
const [data, setData] = useState<User | null>(null)
const [items, setItems] = useState<string[]>([])

// ❌ 非推奨 - useState without types
const [isLoading, setIsLoading] = useState(false)
const [data, setData] = useState(null)

// ❌ 非推奨 - ユニオンタイプ
type ButtonVariant = 'primary' | 'secondary' | 'danger'

// ✅ 推奨 - オブジェクトリテラル
export const ButtonVariant = {
    Primary: 'primary',
    Secondary: 'secondary',
    Danger: 'danger',
} as const
export type ButtonVariant = (typeof ButtonVariant)[keyof typeof ButtonVariant]
```

### React/Next.js規約

#### 関数定義

frontend配下では**アロー関数を使用して統一**してください。

```tsx
// ✅ 推奨 - アロー関数
const MyComponent: React.FC<Props> = ({ prop1, prop2 }) => {
    return <div>{prop1}</div>
}

const handleClick = (event: React.MouseEvent) => {
    // 処理
}

// ❌ 非推奨 - function宣言
function MyComponent({ prop1, prop2 }: Props) {
    return <div>{prop1}</div>
}
```

#### コンポーネント設計

- **再利用性**: 既存のコンポーネントを優先的に使用
- **Props設計**: デフォルト値を適切に設定
- **命名**: コンポーネント名はPascalCase
- **スタイル上書き禁止**: 親コンポーネントから子コンポーネントのスタイルを`!important`で上書きしない

#### App Router活用

- **ルートグループ**: `(group)`記法による論理的なルーティング
- **Layout階層**: 適切なlayout.tsx設計
- **generateMetadata**: 動的SEO対応

```tsx
// ✅ 推奨 - generateMetadata使用
export const generateMetadata = async ({ params }: Props): Promise<Metadata> => {
    const product = await getProduct(params.id)
    return {
        title: product.name,
        description: product.description,
    }
}
```

#### Storybook作成・更新規約

コンポーネントを**新規作成**または**プロパティを追加・修正**した場合は、必ずStorybookを作成・更新してください。

```tsx
// ✅ 推奨 - 新規コンポーネント作成時
import { Button } from '.'
import type { Meta, StoryObj } from '@storybook/nextjs'

const meta: Meta<typeof Button> = {
    component: Button,
    args: {
        children: 'ボタン',
        disabled: false,
    },
    argTypes: {
        disabled: {
            control: { type: 'boolean' },
        },
    },
}

export default meta
type Story = StoryObj<typeof Button>

export const Default: Story = {}
export const Disabled: Story = {
    args: { disabled: true },
}
```

**必須ストーリー:**

- `Default`: 基本状態
- **各プロパティのバリエーション**: disabled、error、required等の状態
- **実用例**: FormExampleなど実際の使用例

#### インポート順序

ESLintルール（perfectionist/sort-imports）に従った順序でインポートを記述：

```tsx
// ✅ 推奨順序
// 1. React関連
import React from 'react'
import { useState } from 'react'

// 2. Next.js関連
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

// 3. 外部ライブラリ（alphabetical）
import { Button } from '@mui/material'
import classNames from 'classnames'

// 4. 内部モジュール（@/で始まる、alphabetical）
import { Button as CustomButton } from '@/components/bases/Button'
import { ProductType } from '@/types/product'

// 5. 相対パス
import styles from './styles.module.scss'
```

### フォーム処理

#### React Hook Form + Zod

```tsx
// ✅ 推奨 - Zodスキーマ定義
const ContactSchema = z.object({
    name: z.string().min(1, '名前を入力してください'),
    email: z.string().email('正しいメールアドレスを入力してください'),
    message: z.string().min(10, 'メッセージは10文字以上で入力してください'),
})

type ContactForm = z.infer<typeof ContactSchema>

// フォーム使用例
const {
    register,
    handleSubmit,
    formState: { errors },
} = useForm<ContactForm>({
    resolver: zodResolver(ContactSchema),
})
```

### SEO対策

#### generateMetadata活用

```tsx
// ✅ 推奨 - 動的SEO
export const generateMetadata = async ({ params }: Props): Promise<Metadata> => {
    return {
        title: `${product.name} | tocoriri`,
        description: product.description,
        openGraph: {
            title: product.name,
            description: product.description,
            images: [product.imageUrl],
        },
    }
}
```

#### インデックス制御

```tsx
// ✅ 推奨 - エラーページ等
export const metadata: Metadata = {
    robots: 'noindex, nofollow',
}
```

### エラーハンドリング

#### API呼び出し

```tsx
// ✅ 推奨 - Zodによる型安全なAPI
export const getProducts = async (): Promise<Product[]> => {
    try {
        const response = await fetch('/api/products')
        if (!response.ok) throw new ApiError(response)
        const data = await response.json()
        return ProductArraySchema.parse(data) // Zod検証
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('商品一覧の取得に失敗しました')
    }
}
```

## 状態管理

現在は**ローカル状態管理**中心：

- **React Hook Form**: フォーム状態
- **useState/useEffect**: コンポーネント状態
- **将来**: Redux Toolkit または Zustand を検討

## パフォーマンス最適化

- **Next.js 15**: 最新機能活用
- **画像最適化**: next/image使用
- **Bundle分析**: 定期的なサイズ監視
- **Dynamic Import**: 必要に応じてコード分割

## デプロイ

- **ビルド**: `pnpm build`
- **静的エクスポート**: 必要に応じて設定
- **CI/CD**: GitHub Actions連携

## 移行方針

### client/からの移行

- **UI/UX**: 元アプリの忠実な再現
- **機能**: 段階的な機能移植
- **参考実装**: client/のコードを参考にしながら新規実装
- **品質向上**: テスト・型安全性・パフォーマンスの向上
