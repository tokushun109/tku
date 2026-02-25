# 商品系置き換え方針（`backend/api` → `clean-backend`）

本ドキュメントは、商品関連APIの置き換え作業を安全に進めるための実装方針をまとめたものです。  
`clean-backend/docs/ddd-and-clean-architecture.md` の方針を前提とし、商品ドメイン特有の複雑性（画像・CSV・複数テーブル参照）に対する具体案を定義します。

## 1. 前提と目的

- 現在稼働中APIは `backend/api` 側であり、`clean-backend` は置き換え作業中。
- 本作業では、以下の商品関連ルートを `clean-backend` へ移行する。
  - `/api/product` (GET/POST)
  - `/api/product/{product_uuid}` (GET/PUT/DELETE)
  - `/api/csv/product` (GET/POST)
  - `/api/category/product` (GET)
  - `/api/carousel_image` (GET)
  - `/api/product_image/{product_image_uuid}/blob` (GET)
  - `/api/product/{product_uuid}/product_image` (POST)
  - `/api/product/{product_uuid}/product_image/{product_image_uuid}` (DELETE)
- `/api/product/duplicate` は初回リリース対象外とし、後続フェーズで移行する。

## 2. 合意済みの重要方針

### 2.1 商品画像の `apiPath`

- `productImages[].apiPath` は **常に presigned URL** を返す。
- `creator/logo` 実装と同様に、`usecase.Storage` 経由で `PresignGet` を利用する。
- local 環境でも MinIO を使用し、ローカルファイル直接参照はしない。

### 2.2 Command / Query の分離

- 商品一覧・詳細・カテゴリ別一覧・カルーセル等の読み取りは、Query専用実装に切り出す。
- 更新系（作成・更新・削除・画像登録削除・CSV更新）は Command 側で扱う。
- 目的:
  - SQL回数の削減
  - 将来の検索・ページネーション追加を容易にする
  - 読み取りモデル最適化とドメイン保護の両立

### 2.3 `duplicate` の扱い

- 初回リリースから除外する。
- 後続フェーズで `infra` 側に外部依存（スクレイピング/HTTP）を隔離して実装する。

### 2.4 Entity の `id` 利用方針（商品系）

- 商品系Entityは原則として `id`（DB内部識別子）を持つ
- API入力/出力の識別子は `uuid` を使い、内部関連では `id` を使う
  - 例: `product_to_tag`, `site_detail`, `product_image` などの関連テーブル連携
- DB からの復元は `Rebuild(...)` を使う
  - `New(...)` は新規作成（未永続化）専用
  - `Rebuild(...)` は `id` を受け取り、`id == 0` は不正として扱う

## 3. ディレクトリ構成方針

### 3.1 Domain（Entity/VO）

`product` と `product_image` は同一ディレクトリ内でファイル分割する。

```text
internal/domain/product/
  product_entity.go
  product_image_entity.go
  product_error.go

  product_name_vo.go
  product_description_vo.go
  product_price_vo.go
  product_is_active_vo.go
  product_is_recommend_vo.go

  product_image_name_vo.go
  product_image_mime_type_vo.go
  product_image_path_vo.go
  product_image_order_vo.go

  product_repo.go
  product_image_repo.go
```

補足:

- UUID は既存方針どおり `internal/domain/primitive/uuid_vo.go` を利用する。
- Product Entity は `id` と `uuid` の両方を保持する。
- Repository の read 実装では `SELECT id, uuid, ...` で取得し、`Rebuild(id, uuid, ...)` で復元する。
- `site_detail` は商品配下で扱うが、必要に応じて Value Object を追加する（detail URL 等）。

### 3.2 Usecase（Command / Query）

```text
internal/usecase/product/
  usecase.go                     # Command中心（Create/Update/Delete など）
  dto.go                         # Command用DTO（必要な場合）

internal/usecase/product/query/
  reader.go                      # Queryインターフェース
  model.go                       # Read Model
```

`reader.go` 例:

```go
type ProductQueryReader interface {
    ListProducts(ctx context.Context, q ListProductsQuery) (*ProductPage, error)
    GetProduct(ctx context.Context, productUUID string) (*ProductDetail, error)
    ListCategoryProducts(ctx context.Context, q ListCategoryProductsQuery) ([]*CategoryProducts, error)
    ListCarouselItems(ctx context.Context, q ListCarouselQuery) ([]*CarouselItem, error)
    ExportProductsCSV(ctx context.Context, q ExportProductsCSVQuery) ([]*ProductCSVRow, error)
}
```

### 3.3 Infra（MySQL Query 実装）

```text
internal/infra/db/mysql/
  repository/
    product.go                   # Command側リポジトリ実装
    product_image.go             # Command側リポジトリ実装
  query/
    product_query.go             # Query側実装(sqlx)
```

意図:

- `repository` はドメイン整合性と更新責務に集中
- `query` は画面要件に最適化したSQL（JOIN/集約/ページング）を担当

### 3.4 VO制約とフロントバリデーションの統一値

商品系では、バックエンドVOとフロント（Zod / Input属性）の制約を以下に統一する。

| 対象                            | 統一値                         | 備考 |
| ------------------------------- | ------------------------------ | ---- |
| 商品名 (`product.name`)         | `min=1`, `max=255`             | フロントの `ProductSchema` と Input `maxLength` を `255` に統一する。 |
| 商品価格 (`product.price`)      | `min=1`, `max=1000000`         | 既存方針を維持する。 |
| 分類名 (`category/target/tag`)  | `min=1`, `max=30`              | フロント分類フォーム（現状20）を30へ引き上げる。 |
| サイト名 (`sales_site` など)    | `min=1`, `max=30`              | 既存VOと合わせ、必要に応じてフロント側を30へ統一する。 |

補足:

- `description` はフロントで optional であり、上限は現時点で固定しない。
- URL系（`siteDetails[].detailUrl`）はフロント同様にURL形式を必須とする。
- 画像ファイルは `image/*` を受け付け、MIMEはバックエンドで再判定する。

## 4. API別の実装責務

### 4.1 Query側

- `GET /api/product`
  - 管理画面向け一覧
  - 一般的なページネーション（`page`, `perPage`）
- `GET /api/product/{product_uuid}`
  - 商品詳細
- `GET /api/category/product`
  - 公開画面向けカテゴリ別一覧
  - 「さらに読み込む」時に `limit` を増やす方式
- `GET /api/carousel_image`
  - おすすめ優先 + 不足時は新着補完
- `GET /api/csv/product`
  - CSVダウンロード用投影データ

### 4.2 Command側

- `POST /api/product`
- `PUT /api/product/{product_uuid}`
- `DELETE /api/product/{product_uuid}`
- `POST /api/product/{product_uuid}/product_image`
- `DELETE /api/product/{product_uuid}/product_image/{product_image_uuid}`
- `GET /api/product_image/{product_image_uuid}/blob`
- `POST /api/csv/product`

## 5. ページネーションと将来拡張

### 5.1 `/api/product`（管理画面）

- 初期実装は一般的なページネーション:
  - `page`（1始まり）
  - `perPage`
  - `totalCount`
- 初回実装では cursor pagination は採用しない。
- 将来検索向けに `keyword`（任意）を事前に受け取れる設計にしておく。

推奨レスポンス:

```json
{
  "items": [],
  "page": 1,
  "perPage": 20,
  "totalCount": 100
}
```

### 5.2 `/api/category/product`（公開画面）

- 「さらに読み込む」要件に合わせ、`limit` を増やして再取得する方式を採用。
- クエリ例:
  - `mode=active&category=all&target=all&limit=20`
  - 次回 `limit=40` のように増加
- 初回実装では cursor pagination は採用しない。
- この方式でも Query 分離と相性が良く、必要になれば後で cursor 方式へ切り替える余地を残せる。

### 5.3 将来の検索機能

- Query実装に検索条件を追加するだけで拡張可能。
- 先行して以下を型に入れておくと後続タスクが容易:
  - `keyword`
  - `sort`（ホワイトリスト）
  - `page/perPage` または `limit`

## 6. 商品画像の詳細方針（creator/logo踏襲）

- アップロード:
  1. バイナリ受領
  2. MIME判定（`http.DetectContentType`）
  3. `ProductImageMimeType` VO で許可判定
  4. 保存キー生成（例: `img/product/{a}/{b}/{uuid}.{ext}`）
  5. `storage.Put`
  6. DB登録
- 失敗時補償:
  - DB更新失敗時は `storage.Delete` でロールバックを試行
- 旧画像削除:
  - DB更新成功後に best-effort 削除（失敗は WARN ログ）
- 取得:
  - `apiPath` は `PresignGet` で生成
  - blob API は `storage.Get` + `Content-Type` 返却

## 7. トランザクション方針

- 複数テーブル更新があるユースケースは `TxManager.WithinTransaction` を使用する。
- 代表例:
  - 商品更新時の `product_to_tag` 再構築
  - `site_detail` 再構築
  - 商品削除時の関連削除
  - CSVアップロード時のカテゴリ/ターゲット補完 + 商品更新

## 8. 初回リリース対象とフェーズ分離

### フェーズ1（初回リリース対象）

- 商品一覧/詳細/カテゴリ別一覧/カルーセル
- 商品CRUD
- 商品画像CRUD + blob
- CSVダウンロード/アップロード

### フェーズ2（後続）

- `POST /api/product/duplicate`
  - `internal/infra/marketplace/creema` 等へ外部依存を隔離して実装
- 商品検索機能
- ページネーションの高度化（必要であれば cursor へ拡張）

## 9. 実装時の互換性ルール

- 既存フロントのレスポンス契約を崩さない:
  - `productImages`, `siteDetails`, `tags`, `category`, `target`, `isActive`, `isRecommend`
- エラー形式は既存 `clean-backend` の統一レスポンスに合わせる:
  - `{ "message": "..." }`
- 管理系の更新APIは既存どおり認証必須。

## 10. 受け入れ条件（Definition of Done）

- 置き換え対象ルートが `clean-backend` 側で仕様どおり動作する
- `productImages[].apiPath` が local/prod 問わず presigned URL になる
- 一覧系で不必要なN+1を回避し、Query実装で一括取得できる
- 将来の検索/ページネーション追加時に、Command側へ大きな変更を入れず拡張できる
- `duplicate` がフェーズ2項目としてドキュメントに明記されている
