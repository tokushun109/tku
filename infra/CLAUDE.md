# CLAUDE.md - Infrastructure

このファイルは、`infra/` ディレクトリで作業する開発者とエージェント向けの補助ガイドです。

## プロジェクト概要

`infra/` では、CDK for Terraform を使って運用系の AWS リソースを管理しています。  
現在の主な対象は、フロントエンド warmup 用 Lambda と、バックエンド API ヘルスチェック用 Lambda、およびそれらの定期実行設定です。

## 開発コマンド

```bash
# 依存関係インストール
pnpm install

# TypeScript コンパイル
pnpm build

# Terraform 設定生成
pnpm synth

# インフラ変更計画
cdktf plan

# インフラ変更適用
cdktf apply

# テスト実行
pnpm test
```

## 技術スタック

- **IaC**: CDK for Terraform
- **クラウド**: AWS
- **言語**: TypeScript
- **パッケージマネージャー**: pnpm

## 管理対象リソース

### Lambda

- **warmup**
  - フロントエンドの応答性維持を目的に定期実行する
- **health-check**
  - API の生存確認を目的に定期実行する

### EventBridge

- warmup Lambda のスケジュール実行
- health-check Lambda のスケジュール実行

## ディレクトリ構造

```text
infra/
├── main.ts                         # スタック定義
├── resources/
│   ├── eventBridge/               # EventBridge ルール定義
│   ├── lambda/                    # Lambda 定義とハンドラー
│   │   └── handlers/
│   │       ├── warmup/           # warmup 実処理
│   │       └── healthCheck/      # health-check 実処理
│   ├── ecs/                       # Lambda から参照する設定ファイル等
│   ├── asm/
│   ├── ec2/
│   └── network/
├── libs/                          # 補助ユーティリティ
├── cdktf.json
├── tsconfig.json
└── package.json
```

## 実装の要点

- `main.ts` では `ap-northeast-1` を対象リージョンとしてスタックを組み立てる
- warmup は日本時間の稼働帯を考慮し、5 分間隔で実行する
- health-check は日本時間の稼働帯を考慮し、1 時間間隔で実行する
- スケジュールは UTC ベースの cron で記述しているため、変更時は時差を必ず確認する

## デプロイフロー

1. `cdktf plan` で差分を確認する
2. コードレビューで意図した変更だけが含まれているか確認する
3. `cdktf apply` で適用する
4. Lambda とスケジュールの動作を確認する

## 注意事項

- 適用前に `cdktf plan` を必ず確認する
- cron を変更する場合は、日本時間と UTC の対応を確認する
- Lambda ハンドラー変更時は、対象 URL や必要な環境変数の影響も確認する
- 実装済みリソースとドキュメントの乖離を作らない
