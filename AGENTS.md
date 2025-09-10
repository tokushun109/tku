# AGENTS.md

このドキュメントは、本リポジトリでAIコーディングエージェント（Codex CLI / Claude Code / Copilot 等）や人間の開発者が、一貫して安全・高品質に作業するための行動規範と運用ルールを定めます。詳細設計・各スタック固有の規約は既存の CLAUDE.md 群を参照してください。

## 適用範囲と優先順位

- スコープ: 本ファイルは配置ディレクトリ（リポジトリルート）配下の全ツリーに適用されます。
- 優先順位: 直接のユーザー/開発者指示 > より深い階層の AGENTS.md > 本ファイル（ルート AGENTS.md）> CLAUDE.md 群。
- CLAUDE.md は詳細リファレンスです。本書では重複を避け、要点化とリンクで参照します。

## 共通ルール（全エージェント共通）

- 言語: 回答・説明・コミットメッセージは原則として日本語。
- 変更は最小限・局所的に。無関係なリファクタや抱き合わせ変更は避ける。
- `.gitignore` 厳守。`git add -f` 等による強制追加は禁止。除外対象（`node_modules/`、ビルド成果物、`.env*`、`.vscode/settings.json` 他）をコミットしない。
- 機密情報を生成・記録・開示しない。環境差異に依存する値は環境変数で扱う。
- ライセンスや著作権ヘッダーを無断で追加しない。
- 破壊的操作（大規模削除、履歴改変等）は明示の依頼がない限り実施しない。

## 作業プロセス（Codex CLI 準拠）

- プレアンブル: まとまった処理単位ごとに「今から何をするか」を 1–2 文で簡潔に共有。
- 計画管理: マルチステップ作業は `update_plan` を用いて短い TODO を管理（常に 1 つが `in_progress`）。
- 変更適用: `apply_patch` で最小差分パッチを作成。不要なファイル変更や整形だけの差分は避ける。
- 検証: 変更範囲に応じてビルド/テスト/リンタを実行し、失敗の早期発見を優先。
- 出力スタイル: 箇条書き中心、コマンド/パス/識別子はバッククォート。過剰な装飾や長文は避ける。

## セキュリティ / 承認

- 秘密情報（鍵・トークン・パスワード・顧客データ等）は扱わない/生成しない/露出しない。
- ネットワークアクセスや破壊的変更が必要な操作は、事前に根拠・影響・ロールバック方針を提示し、承認を得てから実行。
- ハーネスのサンドボックス/承認ポリシーに従い、制約下で代替手段を検討。

## リポジトリ概要（要点）

- フロントエンド: Next.js 15 + TypeScript、Sass(CSS Modules)、App Router、SSR/SPA ハイブリッド、Zod、Storybook。MUI は原則アイコンのみ使用。
- バックエンド: Go + Gorilla Mux、GORM、MySQL、S3 画像保存（UUID）、Railway デプロイ、golang-migrate によるスキーマ管理。
- インフラ: CDK for Terraform（TypeScript）、AWS（VPC/ECS/RDS/S3/ALB/Route53/CloudFront 等）。

詳細: `./CLAUDE.md`、`./frontend/CLAUDE.md`、`./backend/CLAUDE.md`、`./infra/CLAUDE.md` を参照。

## ディレクトリ別の要点（詳細は各 CLAUDE.md）

### frontend/

- Sass は `@use` を使用。`next.config.ts` の `sassOptions.prependData` と重複読み込み禁止。
- CSS クラスはケバブケース。CSS Modules の参照は `styles['class-name']` 形式。
- MUI は基本的にアイコンのみ使用。UI コンポーネントは自前実装を優先。
- 型安全: `useState<Type>()` で型明示。Zod + React Hook Form でスキーマ検証。
- SEO: `generateMetadata` を活用。`app/admin/` 配下は必ず `noindex, nofollow`。
- インポート順序: perfectionist/sort-imports に準拠。変更時は Storybook も更新。

参照: `frontend/CLAUDE.md`

### backend/

- 既存の構造（controllers/models/db/config/utils）に準拠。新規フレームワーク導入は避ける。
- エラーハンドリングと HTTP ステータスの整合性を重視。構造化ログを推奨。
- マイグレーションは golang-migrate。S3 連携/UUID 前提を崩さない。

参照: `backend/CLAUDE.md`

### infra/

- フロー: `plan → review → apply → 検証` を厳守。State 管理の健全性を保つ。
- 環境分離（development / production）、最小権限、暗号化、ネットワーク分離、コスト最適化。

参照: `infra/CLAUDE.md`

## ローカル実行/検証コマンド（代表）

- ルート / Docker
  - `docker-compose up`
- frontend
  - `pnpm install` / `pnpm dev` / `pnpm build` / `pnpm lint` / `pnpm test` / `pnpm sb`
- backend
  - `air`（ホットリロード）/ `go run main.go` / `go build`
- infra
  - `pnpm install` / `pnpm build` / `pnpm synth` / `cdktf plan` / `cdktf apply`

各詳細は CLAUDE.md 群を参照。

## 禁止事項・要承認事項

- 禁止
  - `.gitignore` 対象のコミット、`git add -f` による強制追加。
  - 機密情報の生成・記録・露出。大規模な破壊的変更や履歴改変。
  - 安易な依存の大規模更新、MUI の乱用（frontend）。
- 要承認
  - 依存のメジャー更新、破壊的スキーマ変更、SEO/運用に影響する設定変更。
  - インフラコスト/構成に影響を与える変更。ネットワークアクセスを伴う操作。

## 変更提案の出し方（PR テンプレ）

- 目的/背景
- 変更点（要点箇条書き）
- 影響範囲（UI/SEO/DB/インフラ）
- 検証手順（実行コマンドと観点）
- リスクとロールバック
- 参照資料（関連する CLAUDE.md 節、Issue など）

## 参考

- ルート: `./CLAUDE.md`
- フロントエンド: `./frontend/CLAUDE.md`
- バックエンド: `./backend/CLAUDE.md`
- インフラ: `./infra/CLAUDE.md`

