# UUIDベースアクセス移行方針

本ドキュメントは、`backend` のスキーマ・API・CSV を `id` 中心から `uuid` 中心へ移行するための合意事項をまとめたものです。  
`backend/docs/product-replacement-plan.md` の商品系初期方針を補完し、現在の移行方針を定義します。

## 1. 目的

- 外部公開する識別子（API 入出力、CSV 更新キー、関連付け）は `uuid` に統一する。
- 目標は「`id` の物理削除」ではなく、「`id` を外部識別子として廃止する」こととする。
- `id` は当面、内部実装用の補助キー（安定ソートの tie-break、互換維持など）として保持してよい。
- ただし、新しい外部キー参照やアプリケーション上の主識別には `id` を使わない。

## 2. 共通方針

### 2.1 UUID の扱い

- UUID は lowercase のみを使用する。
- DB に保存する UUID も lowercase を前提とする。
- UUID 比較・バリデーション・バックフィル時も lowercase を維持する。

### 2.2 切り替え手順

`*_id` から `*_uuid` へ切り替えるテーブルは、原則として以下の順序で移行する。

1. migration で `*_uuid` 列を追加する。
2. 既存の `*_id` から `*_uuid` をバックフィルする。
3. `*_uuid` にインデックスを付与し、必要に応じて `FOREIGN KEY (..._uuid) REFERENCES parent(uuid)` を追加する。
4. アプリケーションを deploy し、読み書きを UUID ベースへ切り替える。
5. 安定確認後、旧 `*_id` 参照コードを削除する。
6. 必要に応じて旧 `*_id` 列や旧 FK を掃除する。

補足:

- DB migration 完了直後に旧コードが動いていても問題ないよう、アプリ切り替えは deploy 後に行う。
- UUID ベース切り替え後は、`*_uuid` を source of truth とする。

### 2.3 `id` の位置づけ

- `id` は外部向けレスポンス、API パラメータ、CSV の識別子としては使わない。
- `id` を残す場合も、内部専用の補助キーとしてのみ扱う。
- 並び順の tie-break が必要な箇所では、当面 `id` を使ってよい。
- ただし、純粋な中間テーブルは `id` 自体を持たせない設計を優先する。

## 3. テーブルごとの方針

### 3.1 UUID を追加して統一するテーブル

- `creator`
  - `uuid` 列を追加し、既存レコードには migration で UUID を採番する。
  - 当面は単一プロフィール前提のまま扱う。
  - `creator` と `user` の所有者関連は今回の移行対象外とする。
- `contact`
  - `uuid` 列を追加し、既存レコードには migration で UUID を採番する。
  - API レスポンスは `id` ではなく `uuid` を返す。

### 3.2 UUID ベース参照へ切り替えるテーブル

- `product`
- `category`
- `target`
- `tag`
- `product_image`
- `sales_site`
- `site_detail`
- `user`
- `session`

上記テーブルでは、外部キーやアプリケーション上の参照は `*_uuid` ベースへ段階的に移行する。

### 3.3 純粋な中間テーブル

- `product_to_tag`
  - `product_uuid + tag_uuid` の複合キーを採用する。
  - 行自体を表す `uuid` は追加しない。
  - 純粋な join テーブルとして扱うため、最終的には物理削除を前提とする。
  - `deleted_at` は不要になった段階で削除候補とする。

## 4. 並び順の扱い

### 4.1 `product_to_tag`

- タグの並び自体に業務上の意味は持たせない。
- 取得順は以下で固定する。

```sql
ORDER BY ptt.created_at ASC, ptt.tag_uuid ASC
```

- `created_at` が同値でも `tag_uuid` を第2キーにすることで安定順を確保する。
- `product_to_tag.id` は並び順のためには使わない。

### 4.2 その他のテーブル

- `id` を現在 tie-break として使っている箇所は、必要に応じて当面残してよい。
- ただし、外部仕様や関連付けの識別子に `id` を使い続けないことを優先する。
- 将来的に `id` が不要になったテーブルは、個別に削除を再検討する。

## 5. API / CSV 方針

### 5.1 API

- 外部向け API の識別子は `uuid` に統一する。
- `contact` を含め、外部向けレスポンスで `id` は返さない方針とする。
- 既存の `id` ベース DTO / response は UUID ベースへ置き換える。

### 5.2 CSV

- 商品 CSV の更新キーは `id` ではなく `uuid` に切り替える。
- 想定フローは `export -> 内容を編集 -> 再 upload` とする。
- この運用では、識別子を新規手入力しないため、`uuid` へ切り替えても実務上の手間はほぼ増えない。
- CSV 上の UUID 列は「編集しないキー」として扱う。
- 以下はバリデーションエラーとする。
  - UUID 空欄
  - UUID 重複
  - 未知の UUID

## 6. `user` / `session` の切り替え方針

- `user` / `session` は認証に直結するため、商品系とは分けて慎重に移行する。
- 主な対象は `session.user_id` から `session.user_uuid` への移行とする。
- 移行手順は共通方針に従い、
  1. `session.user_uuid` 追加
  2. `session.user_id` からバックフィル
  3. `FOREIGN KEY (user_uuid) REFERENCES user(uuid)` 追加
  4. アプリを UUID ベースへ切り替え
  5. 安定確認後に旧 `user_id` 参照を掃除
  の順で行う。

補足:

- セッションは再ログインで再生成できるため、データとしての永続重要度は低い。
- ただし、切り替え途中の不整合はログイン / 認証不具合につながるため、移行タイミングはまとめて扱う。
- 運用上許容できる場合は、切り替え時に全セッションを無効化して再ログインへ寄せる案も選択肢とする。

## 7. 中間テーブルの削除方針

- 純粋な中間テーブルは、最終的に物理削除ベースへ統一する。
- 論理削除を維持する明確な要件がない限り、`deleted_at` を持たない構成を優先する。
- 物理削除へ統一する対象は、個別テーブルごとに段階的に切り替える。

## 8. 現時点での保留事項

- `creator` の「常に 1 レコード」前提を DB / アプリのどこで担保するかは別途検討とする。
- 内部補助キーとして残した `id` を、将来的にどこまで物理削除するかは UUID 移行完了後に再判断する。

## 9. 現状の実装状況

### 9.1 実装済み

- UUID ベース移行の初期 migration として `db/migrations/000021_add_uuid_access_columns` を追加済み。
- `creator` / `contact` に `uuid` 列を追加する前提で、ドメイン・repository・response を UUID 対応済み。
- `product` の `category` / `target` 参照は、アプリケーション上は `category_uuid` / `target_uuid` を主に使う実装へ切り替え済み。
- `product_to_tag` は、アプリケーション上は `product_uuid` / `tag_uuid` を使って関連更新する実装へ切り替え済み。
- `product_image` / `site_detail` は、アプリケーション上は `product_uuid` / `sales_site_uuid` を使って関連更新する実装へ切り替え済み。
- `user` / `session` は、アプリケーション上は `session.user_uuid` を主に使う実装へ切り替え済み。
- 商品 CSV は `id` ではなく `uuid` を更新キーに使う実装へ切り替え済み。
- 商品一覧・CSV 出力などの query 側も、UUID 参照を優先する JOIN / 読み取りへ切り替え済み。
- repository / query の読み取りは、旧 `*_id` へのフォールバックを外し、UUID のみを参照する実装へ切り替え済み。
- `product` / `user` repository に残っていた未使用の `FindByID` は削除済み。
- 2026-03-01 時点で、`backend` 配下の `go test ./...` は通過済み。

### 9.2 実装時の補足

- 既存データの UUID バックフィルは migration 内で実施する。
- `creator.uuid` / `contact.uuid` の既存データ採番は、アプリ側の Go 実装ではなく MySQL の `LOWER(UUID())` を使用する。
- 形式としてはアプリ側と同じ lowercase UUID を採用するが、生成アルゴリズム自体は MySQL の `UUID()` に依存する。
- 現段階では、旧 `id` 列および旧 `*_id` 列は物理削除していない。

## 10. 残作業

### 10.1 運用 / 適用

- 各環境で `000021_add_uuid_access_columns` を適用する。
- migration 適用前に、対象テーブルのバックアップ取得手順を確認する。
- migration 適用後に、既存データの `*_uuid` バックフィル結果を確認する。
- migration 適用後に、`creator.uuid` / `contact.uuid` / `session.user_uuid` などの NULL 残存有無を確認する。

### 10.2 動作確認

- 商品の作成 / 更新 / 削除で、UUID ベースの関連更新が正しく動くかを確認する。
- 商品 CSV の export / 編集 / upload の一連フローを確認する。
- 画像アップロード、タグ更新、販売サイト詳細更新など、商品の関連データ更新を確認する。
- ログイン / ログアウト / セッション解決を含む `user` / `session` の認証導線を確認する。
- `contact` 一覧 / 作成、`creator` 取得 / 更新 / ロゴ更新の API 動作を確認する。

### 10.3 後続の整理

- 必要に応じて、旧 `*_id` 列・旧 FK・不要なインデックスの削除を別 migration で行う。
- `product_to_tag` など中間テーブルの物理削除方針に合わせて、`deleted_at` の削除要否を再確認する。
- `creator` の「常に 1 レコード」前提をどう担保するかを別途確定する。
- 将来的に `id` を内部補助キーとして残し続けるか、さらに削減するかを再判断する。
