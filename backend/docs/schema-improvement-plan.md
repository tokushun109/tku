# テーブルスキーマ改善方針

本ドキュメントは、`backend/db/migrations` と `backend/internal/infra/db/mysql` の現状実装を確認した上で、
テーブルスキーマと SQL 呼び出しの責務分担を整理し、今後の改善方針をまとめたものです。

既存の `backend/docs/uuid-access-migration-plan.md` を前提としつつ、その後に残っている
「DB 側で担保すべき制約」と「アプリケーション側に寄りすぎている運用」を補強することを目的とします。

## 1. 目的

- アプリケーション実装に依存している整合性を、可能な限り DB 制約へ移す。
- 同時実行時でも破綻しにくいテーブル設計へ寄せる。
- 読み取りクエリの実態に合わせて、必要なインデックスを整理する。
- repository / usecase の責務を明確にし、トランザクション境界を一貫させる。

## 2. 現状認識

- UUID ベースの参照へ移行する方針自体は概ね完了している。
- 一方で、名称の重複防止や 1 ユーザー 1 セッションのような業務ルールは、
  一部がアプリケーションの事前チェックに依存している。
- 現在の repository 実装では、`getExecutor` + `txctx` によるトランザクション共有の仕組みがある一方で、
  一部 repository が独自に `BeginTxx` を行っており、境界が混在している。
- `product_to_tag` / `site_detail` などの関連更新は「全削除して再作成」の実装が残っており、
  データ量増加時のコストが読みにくい。

## 3. 優先度高の修正方針

### 3.1 名称の一意性を DB で担保する

対象:

- `category.name`
- `tag.name`
- `target.name`
- `sales_site.name`

現状:

- アプリケーションでは `ExistsByName` や `FindByName` により重複しない前提で扱っている。
- ただし、DB 側には一意制約がないため、同時実行時には重複が混入しうる。
- 重複が入ると、`FindByName` が任意の 1 件を返す構造になり、挙動が不安定になる。

方針:

- 業務上「同名は同一マスタ」とみなす前提で、DB 側にも一意制約を追加する。
- 既存データに重複がある場合は、migration 前に棚卸しと統合作業を行う。
- 論理削除を維持したまま「有効レコードのみ一意」を担保したい場合は、
  generated column などを使った一意制約を検討する。
- そこまでの対応をすぐに行わない場合でも、少なくとも `name` を軸にした検索用 index は付与する。

補足:

- `skill_market` / `sns` は現時点では単一名称 lookup に依存していないため、優先度は一段下げてよい。
- ただし、将来的に `FindByName` ベースの運用を追加するなら同様の制約を検討する。

### 3.2 `session` を 1 ユーザー 1 レコードで担保する

現状:

- ログイン時は `DeleteByUserUUID` の後に `Create` しており、実装上は 1 ユーザー 1 セッションを意図している。
- ただし、DB 側は `user_uuid` の通常 index のみで、同時ログイン時に複数行が残る余地がある。

方針:

- `session.user_uuid` に `UNIQUE` 制約を追加する。
- session 作成は `INSERT ... ON DUPLICATE KEY UPDATE` へ寄せるか、
  重複エラーを前提にしたリトライ可能な実装へ変更する。
- 1 ユーザー複数セッションを今後認める要件が出るまでは、DB でも単一セッションを保証する。

### 3.3 トランザクション境界を usecase 層に統一する

現状:

- `TxManager.WithinTransaction` により、context 経由で tx を共有できる。
- 一方で、`category` / `tag` / `target` / `sales_site` の一部削除処理は repository 内で独自に tx を開始している。

方針:

- 複数 SQL を束ねる責務は usecase 層に寄せ、repository は「与えられた executor を使うだけ」に統一する。
- repository 内で `BeginTxx` / `Commit` / `Rollback` を直接持つ実装は段階的に解消する。
- これにより、親ユースケースの tx に自然に参加できる状態を標準にする。

期待効果:

- ネストした更新で一部だけ先に commit される事故を防ぎやすくなる。
- unit test でも transaction 境界の前提が読みやすくなる。

## 4. 優先度中の修正方針

### 4.1 関連テーブルの全削除・全再作成を見直す

対象:

- `product_to_tag`
- `site_detail`

現状:

- 更新時は既存レコードをすべて削除し、入力内容を再挿入している。
- 実装は単純だが、件数増加時に不要な delete / insert が増える。
- `created_at` を監査目的で見たい場合、更新のたびに履歴が失われる。

方針:

- まずは `site_detail` の逐次 `INSERT` を bulk insert 化し、SQL 発行回数を減らす。
- 中長期では差分更新へ寄せ、必要な行のみ `INSERT` / `UPDATE` / `DELETE` する。
- `product_to_tag` は純粋な中間テーブルとしてよいが、将来的に件数が増える場合は差分化を検討する。

補足:

- 単純性を優先するなら、すぐに完全差分更新へは進めず、
  まずは bulk insert 化だけでも十分に価値がある。

### 4.2 実際の読み取りクエリに合わせて複合 index を追加する

現状:

- UUID ベース FK 用の単独 index は揃っている。
- ただし、一覧系のクエリは `deleted_at`、`is_active`、`is_recommend`、`created_at`、
  `display_order` などを組み合わせて使っている。

方針:

- 追加する index は、必ず `EXPLAIN` を取りながら決める。
- 初期候補は以下とする。

候補:

- `product (deleted_at, is_active, is_recommend, created_at, id)`
- `product_image (product_uuid, deleted_at, display_order, id)`
- `contact (deleted_at, created_at)`
- 必要に応じて `category (name, deleted_at)`、`tag (name, deleted_at)`、`target (name, deleted_at)`、
  `sales_site (name, deleted_at)`

注意:

- index は追加しすぎると write コストが増えるため、推測だけで増やしすぎない。
- まずは管理画面や公開 API の主要クエリから優先的に観測する。

### 4.3 timestamp カラムの nullable を減らす

対象:

- 主に `created_at`
- 主に `updated_at`

現状:

- 多くのテーブルで `created_at` / `updated_at` が nullable になっている。
- 実装上は `UTC_TIMESTAMP()` を入れており、通常運用で null を許す必要性は低い。

方針:

- 新規作成するテーブル、または今後定義を触るテーブルから、
  `created_at DATETIME(6) NOT NULL`、`updated_at DATETIME(6) NOT NULL` を標準とする。
- 既存テーブルは、null データがないことを確認した上で段階的に `NOT NULL` 化する。
- `deleted_at` は論理削除列として nullable のままでよい。

補足:

- ミリ秒以上の精度が必要ない場合でも、`DATETIME(6)` に揃えておくと将来の差分解析がしやすい。

## 5. 優先度低の修正方針

### 5.1 `INSERT` 直後の再読込を減らす

対象:

- `contact`
- `session`

現状:

- `INSERT` 後に `created_at` を再取得するため、追加で `SELECT` を発行している。

方針:

- DB 時刻に依存しなくてよいユースケースは、アプリ側の `clock.Now().UTC()` で時刻を確定して保存する。
- もしくは「戻り値に DB 側 `created_at` を必須で含めない」設計に寄せ、即時再取得を減らす。

補足:

- ただし、認証や監査で DB 時刻を厳密に信頼したい要件があるなら、優先度は低い。

### 5.2 中間テーブルの主キー設計は将来再検討する

対象:

- `product_to_tag`

現状:

- 現在は `id` を持ったまま、`product_uuid + tag_uuid` に複合一意制約を持たせている。

方針:

- 現時点では `id` を残したままで問題ない。
- ただし、純粋な join テーブルとして整理する段階では、
  `id` を持たず複合主キーへ寄せるかどうかを別途再検討する。

## 6. migration 実施時の共通ルール

- 既存データを壊さないよう、制約追加前に重複・不整合データを必ず洗い出す。
- 破壊的変更は、原則として以下の順で行う。

1. 現状データを棚卸しする。
2. 重複や不整合を解消する migration または手当てを先に入れる。
3. index / unique / foreign key を追加する。
4. アプリケーション実装を制約前提へ切り替える。
5. 不要になった暫定コードや旧運用を削除する。

- `UNIQUE` 追加時は、既存アプリが重複エラーを適切に扱える状態にしてから deploy する。
- データ削除を伴う migration は、ロールバック方法を事前に定義してから実施する。

## 7. SQL 実装側のルール

- repository は「単一テーブル操作」または「単一責務の SQL 実行」に寄せる。
- 複数 repository をまたぐ整合性は usecase 層の `TxManager` でまとめる。
- `FindByName` を提供するテーブルは、業務上の一意性が本当に必要かを先に確認する。
- 業務上一意な lookup を持つなら、アプリ側の事前チェックだけに依存しない。
- 一覧系 query は、実際の `WHERE` / `ORDER BY` とセットで index を見直す。

## 8. 推奨する実施順

### Phase 1: 安全性の補強

- `session.user_uuid` の一意制約追加
- `category` / `tag` / `target` / `sales_site` の重複データ調査
- repository 内独自 tx の棚卸し

### Phase 2: 制約の導入

- 名称一意性のための制約または補助 index の追加
- 重複エラーを正しく返すよう、usecase / handler のエラーハンドリング整理

### Phase 3: 性能と保守性の改善

- `site_detail` の bulk insert 化
- 必要な複合 index の追加
- `created_at` / `updated_at` の `NOT NULL` 化

## 9. 保留事項

- `skill_market` / `sns` の名称一意性をどこまで業務ルールとして扱うか。
- `product_to_tag` を将来的に複合主キーへ寄せるか。
- `created_at` の時刻ソースを DB とアプリのどちらに寄せるか。

## 10. このドキュメントの扱い

- 本書は「今すぐ一括変更する一覧」ではなく、段階的に進めるための優先順位付き方針書とする。
- 実際の修正は、影響範囲ごとに小さな migration と実装変更へ分けて進める。
- 各変更の着手時には、本書の該当節を参照しながら個別の実装計画を起こす。
