# DDD + Clean Architecture 移行方針

本ドキュメントは、`tku/backend/api` を DDD + Clean Architecture で再構築するための合意事項をまとめたものです。
ここに記載された方針を優先し、実装中の迷いを減らします。

## 目的

- MVC で混在している責務を分離し、変更に強い構成にする
- ドメイン単位で理解できる構成にして移行を進めやすくする
- 依存方向を統一して、テスト可能性を高める

## モジュール配置

- `go.mod` と `go.sum` は `tku/clean-backend` のルートに置く
- `internal/` は非公開パッケージの格納先（モジュールルートにはしない）

## ディレクトリ構成（確定版）

```
tku/clean-backend/
  go.mod
  go.sum
  cmd/
    api/
      main.go
  internal/
    app/
      di/                  # 依存注入・組み立て（Composition Root）
    domain/
      product/
        entity.go
        product_repo.go
        service.go
      category/
        entity.go
        category_repo.go
      target/
        entity.go
        target_repo.go
      tag/
        entity.go
        tag_repo.go
      creator/
        entity.go
        creator_repo.go
      contact/
        entity.go
        contact_repo.go
      user/
        entity.go
        user_repo.go
      session/
        entity.go
        session_repo.go
      sales_site/
        entity.go
        sales_site_repo.go
      skill_market/
        entity.go
        skill_market_repo.go
      sns/
        entity.go
        sns_repo.go
    usecase/
      product/
        usecase.go
      category/
        usecase.go
      target/
        usecase.go
      tag/
        usecase.go
      creator/
        usecase.go
      contact/
        usecase.go
      user/
        usecase.go
      session/
        usecase.go
      sales_site/
        usecase.go
      skill_market/
        usecase.go
      sns/
        usecase.go
    interface/
      http/
        handler/
        presenter/
        request/
        response/
        router/
      jobs/                # バッチやCLIがあれば
    infra/
      config/
        config.go
      db/
        mysql/
          db.go
          repository/
      storage/
        s3/
      mail/
        sendgrid/
      logger/
    shared/
      id/
      timeutil/
      errors/
  db/
    migrations/            # 本プロジェクトはここで管理
  docs/
```

※上記の `entity.go` 等は構成例です。実際のファイル名は命名規則に合わせて `category_entity.go` のようにドメイン名を明示します。

## 依存ルール（重要）

- `domain` は外部に依存しない（純粋なルールとモデルのみ）
- `usecase` は `domain` に依存してよい
- `interface` は `usecase` に依存してよい
- `infra` は `domain` / `usecase` に依存してよい
- 依存の向きは常に「外側 → 内側」

## Entity / VO の扱い方と命名規則

- Entity はドメインの中心モデル。`internal/domain/<domain>` に置く
- Value Object（VO）も同じドメイン配下に置き、同一パッケージ内で完結させる
- UUID は **共通の VO を 1 つだけ用意**して使い回す
  - `internal/domain/primitive/uuid.go` に `primitive.UUID` を定義
  - 正規表現でフォーマット検証のみを行う（外部依存なし）
- UUID の生成は **infra 層**（例: google/uuid）で行い、VO に変換して usecase に渡す
  - usecase 側は `UUIDGenerator` インターフェースに依存
  - infra 側で `UUIDGenerator` を実装して DI する
- 文字列や数値の制約は VO に寄せる（Entity に生データを持たせない）

### 命名規則（ファイル名）

- Entity: `<domain>_entity.go`
  - 例: `category_entity.go`, `session_entity.go`
- VO: `<domain>_<value>_vo.go`
  - 例: `category_name_vo.go`
- 共通 VO: `internal/domain/primitive/<value>.go`
  - 例: `internal/domain/primitive/uuid.go`
- テスト: 対象ファイル名 + `_test.go`
  - 例: `category_entity_test.go`, `category_name_vo_test.go`, `uuid_test.go`

### 命名規則（型名）

- Entity: `Category`, `Session` のようにドメイン名そのもの
- VO: `CategoryName` など（UUID は共通の `primitive.UUID` を使用）

### VO 例（UUID）

- `internal/domain/primitive/uuid.go` に共通 VO を定義
- 生成は infra で行い、usecase で VO を受け取る

```
// internal/domain/primitive/uuid.go
type UUID string
```

### VO 例（Name）

- 文字列の長さ・形式などの制約は VO に閉じ込める
- VO 内の定数は private にする

```
// internal/domain/category/category_name_vo.go
const (
    nameMinLen = 1
    nameMaxLen = 20
)
```

## Interface の配置ルール

- Usecase の IF は `internal/usecase/<domain>/usecase.go`
- Repository の IF は `internal/domain/<domain>/<domain>_repo.go`
- Repository の実装は `internal/infra/db/mysql/repository`

## HTTP Router の責務

- `internal/interface/http/router` はルーティング定義のみを持つ
- handler の紐付けと middleware の適用順をここで定義する
- ビジネスロジックは置かない
- 認証が必要なルートは `AuthMiddleware` を付与する

## HTTP Request の役割

- `internal/interface/http/request` は HTTP 入力の DTO を置く
- JSON ボディ・クエリ・パスパラメータの受け取り型を定義する
- ドメインエンティティやビジネスロジックは置かない

### 例

```go
type CreateProductRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int    `json:"price"`
    CategoryID  string `json:"category_id"`
}

type ListProductsQuery struct {
    Mode     string `json:"mode"`
    Category string `json:"category"`
    Target   string `json:"target"`
}
```

## HTTP Response の役割

- `internal/interface/http/response` は HTTP 出力の DTO とヘルパーを置く
- レスポンスの形を統一する（`success` / `message` など）
- Usecase の結果を HTTP 用に整形する処理のみを持つ

### 例

```go
type ProductResponse struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int    `json:"price"`
}
```

## Presenter の役割

- `internal/interface/http/presenter` は **Usecase 出力 → Response DTO** の変換を担当
- 変換ロジックを handler から分離するために置く
- ドメインのルールは持たない

### 例

```go
func ToProductResponse(p *product.Product) *response.ProductResponse {
    return &response.ProductResponse{
        ID:          p.ID,
        Name:        p.Name,
        Description: p.Description,
        Price:       p.Price,
    }
}
```

### CORS の扱い

- CORS は `internal/interface/http/middleware` に実装する
- ルーターで `Use()` して全ルートに適用する
- 許可 Origin は `CLIENT_URL` のみ
- `AllowCredentials` は `true` 固定

#### CORS ミドルウェア例

```go
func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
    c := cors.New(cors.Options{
        AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
        AllowedOrigins:   allowedOrigins,
        AllowCredentials: true,
    })
    return c.Handler
}
```

## DI（Composition Root）

- 依存注入は `internal/app/di` に集約する
- `cmd/api/main.go` は `di` のエントリを呼ぶだけにする
- ドメイン追加時は `di` の組み立てを増やす

## Domain Service の扱い

- エンティティ単体に閉じないルールは `domain/<domain>/service.go` に置く
- DBアクセスが必要なルールは Repository IF 経由で行う
- 不要なら `service.go` を置かなくてよい

## 層割り当て（現行実装からの移行ルール）

この章は「どの処理をどこに移すべきか」の判断基準です。

| 現状処理                            | 目的/ルール              | 移動先レイヤ                    | 移動先の例                                                                       | 補足                                                    |
| ----------------------------------- | ------------------------ | ------------------------------- | -------------------------------------------------------------------------------- | ------------------------------------------------------- |
| 重複チェック（Category/Target/Tag） | 名称の重複防止           | Usecase (+ Domain Service 任意) | `internal/usecase/*`                                                             | DB参照が必要。Usecaseで`repo.ExistsByName`を使う。      |
| 削除時の関連整理                    | 関連テーブルの整合性維持 | Usecase                         | `internal/usecase/*`                                                             | Repositoryで更新・削除。Usecaseでトランザクション制御。 |
| スクレイピング複製                  | 外部サイトから商品作成   | Usecase + Infra                 | Usecase: `internal/usecase/product` / Infra: `internal/infra/marketplace/creema` | HTTP/HTML解析はInfra。                                  |
| CSV更新                             | CSVからの更新処理        | Usecase                         | `internal/usecase/product`                                                       | CSV DTO は `interface` 層に置く。                       |
| 画像削除/ファイル操作               | DB更新 + ファイル削除    | Usecase + Infra                 | Usecase: `internal/usecase/product` / Infra: `internal/infra/storage`            | I/OはInfra。順序制御はUsecase。                         |
| セッション判定                      | 有効性チェック           | Usecase                         | `internal/usecase/session`                                                       | 将来期限が入る場合はDomainにルール化。                  |
| パスワードハッシュ                  | 暗号化                   | Domain Service or Infra         | `internal/domain/user/password.go` or `internal/infra/crypto`                    | 実装差し替えを見込むならInfra。                         |
| 初期データ投入                      | seed/migration           | Infra (Seed/Migration)          | `db/seed` or `internal/infra/seed`                                               | 起動時に行わない。                                      |
| ロゴ更新時の旧画像削除              | ストレージ操作           | Usecase + Infra                 | `internal/usecase/creator` + `internal/infra/storage`                            | 削除はStorageの責務。                                   |

## 外部I/Oの整理方針

- S3, SendGrid, HTTP などの外部サービスはすべて `infra` に配置
- `controllers` 直下に外部パッケージを置かない

## エラーの扱い

- Domain は業務ルールのエラーのみを返す（`errors.New` のセンチネルでOK）
- Usecase は Domain のエラーを受け取り、アプリケーションの結果に変換して返す
- Interface（HTTP）は Usecase のエラーを HTTP ステータスと JSON に変換する
- 想定外のエラーは `internal error` に統一し、詳細はログのみ出す
- `http.Error` は使わず、JSON の統一レスポンスにする

### バリデーションの方針

- エンティティに集約する（単体で完結するルールのみ）
- 例: 必須/文字数/フォーマット/値の範囲など
- DB参照が必要な重複チェックは Usecase / Domain Service 側で扱う

### Domain エラー例（イメージ）

- `internal/domain/<domain>/<domain>_error.go`
- `var ErrCategoryNameDuplicate = errors.New(\"category name is duplicate\")`
- 共通プリミティブのエラーは `internal/domain/primitive/<name>_error.go` に置く

### Usecase エラーの種別（例）

- `ErrInvalidInput`
- `ErrNotFound`
- `ErrConflict`
- `ErrUnauthorized`
- `ErrInternal`

### HTTP ステータス対応（例）

- `ErrInvalidInput` → 400
- `ErrNotFound` → 404
- `ErrConflict` → 409
- `ErrUnauthorized` → 401
- `ErrInternal` → 500

### HTTP のエラーレスポンス（code なし）

```
{
  \"message\": \"category name is duplicate\"
}
```

### エラーレスポンスのヘルパー（設計方針）

- `internal/interface/http/response` に共通ヘルパーを置く
- Usecase エラーを `status` に変換して返す

## ログの出し方

- ログは Interface 層（middleware/handler）で集約する
- Usecase/Domain は原則ログしない
- Infra は外部I/O失敗時のみ補助ログを許可
- 基本は「1リクエスト1ログ」（開始+終了、もしくは終了のみ）

### レベルの基準

- `INFO`: 正常終了
- `WARN`: 想定されるエラー（バリデーション/重複/NotFound など）
- `ERROR`: 想定外のエラー（DB/外部API/パニック）

### ログ項目（最低限）

- `request_id`
- `method`, `path`
- `status`
- `latency_ms`
- `error`（ある場合のみ）

### 実装メモ

- ResponseWriter をラップして `status` を取得する
- handler が `WriteHeader` した結果を middleware が参照する
- `request_id` は `X-Request-ID` ヘッダーを優先し、なければ UUID を生成する
- `request_id` はレスポンスヘッダーにも付与する

### ログ用ミドルウェアの雛形

```go
// internal/interface/http/middleware/logger.go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        reqID := getOrCreateRequestID(r)

        // ResponseWriterをラップして status を取得
        rw := NewResponseWriter(w)
        rw.Header().Set("X-Request-ID", reqID)

        next.ServeHTTP(rw, r)

        latency := time.Since(start).Milliseconds()
        status := rw.Status()

        level := "INFO"
        if status >= 500 {
            level = "ERROR"
        } else if status >= 400 {
            level = "WARN"
        }

        log.Printf("[%s] request_id=%s method=%s path=%s status=%d latency_ms=%d",
            level, reqID, r.Method, r.URL.Path, status, latency,
        )
    })
}
```

### ResponseWriter ラッパー（ステータス取得用）

```go
type ResponseWriter struct {
    http.ResponseWriter
    status int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
    return &ResponseWriter{ResponseWriter: w, status: http.StatusOK}
}

func (rw *ResponseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Status() int {
    return rw.status
}
```

### request_id 生成ルール（方針）

- 受け入れヘッダーは `X-Request-ID`
- ない場合は UUID v4 を生成
- 生成/受け入れた `request_id` はレスポンスヘッダーに付与する

### request_id 補助関数（例）

```go
func getOrCreateRequestID(r *http.Request) string {
    if v := r.Header.Get("X-Request-ID"); v != "" {
        return v
    }
    id := id.GenerateUUID() // internal/shared/id
    r.Header.Set("X-Request-ID", id)
    return id
}
```

### UUID の配置

- `request_id` 用の UUID 生成は `internal/shared/id` に置く
- 例: `internal/shared/id/uuid.go` に `func GenerateUUID() string` を用意
- **エンティティの UUID** は共通 VO（`internal/domain/primitive/uuid.go`）を使用
- エンティティ UUID の生成は infra 層で行い、usecase に DI する

## マイグレーション

- `db/migrations` をルート直下に置く（DB切り替え予定なしのため固定）

## Docker 配置方針

- `tku/clean-backend/docker/` に新設して集約する
- 既存 `tku/backend/api/docker` の内容をベースに移行して管理する
- 重複期間は両方のディレクトリが存在してよい

### 想定構成

```
tku/clean-backend/
  docker/
    api/
    db/
    migrations/
```

## DB 接続（MySQL + sqlx）

- 既存の GORM から sqlx へ移行する（マイグレーションは golang-migrate を継続使用）
- GORM の `gorm.DeletedAt` は **自動で `deleted_at IS NULL` を付与**するが、sqlx では **明示的に条件を追加**する必要がある
  - 取得系: `WHERE deleted_at IS NULL`
  - 削除系: `DELETE` ではなく `UPDATE ... SET deleted_at = NOW()` を使う（論理削除）

- 接続初期化は `internal/infra/db/mysql/db.go` に集約する
- 設定値は `internal/infra/config` から受け取る
- 本番/開発で同一設定の想定（当面は分けない）

### 接続実装例

```go
func NewDB(cfg *config.Config) (*sqlx.DB, error) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo",
        cfg.DBUser,
        cfg.DBPass,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )

    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(30 * time.Minute)

    return db, nil
}
```

## Repository 実装例（sqlx）

- Repository IF は `internal/domain/<domain>/<domain>_repo.go`
- 実装は `internal/infra/db/mysql/repository` に置く

### Domain IF 例

```go
type Repository interface {
    Create(ctx context.Context, p *Product) error
    FindByID(ctx context.Context, id string) (*Product, error)
    ExistsByName(ctx context.Context, name string) (bool, error)
    Update(ctx context.Context, p *Product) error
    Delete(ctx context.Context, id string) error
}
```

### sqlx 実装例

```go
type ProductRepository struct {
    db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *product.Product) error {
    _, err := r.db.ExecContext(
        ctx,
        `INSERT INTO product (uuid, name, price, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`,
        p.ID, p.Name, p.Price,
    )
    return err
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*product.Product, error) {
    var p product.Product
    if err := r.db.GetContext(ctx, &p, `SELECT uuid, name, price FROM product WHERE uuid = ?`, id); err != nil {
        return nil, err
    }
    return &p, nil
}

func (r *ProductRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
    var count int64
    if err := r.db.GetContext(ctx, &count, `SELECT COUNT(1) FROM product WHERE name = ?`, name); err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *ProductRepository) Update(ctx context.Context, p *product.Product) error {
    _, err := r.db.ExecContext(
        ctx,
        `UPDATE product SET name = ?, price = ?, updated_at = NOW() WHERE uuid = ?`,
        p.Name, p.Price, p.ID,
    )
    return err
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
    _, err := r.db.ExecContext(ctx, `DELETE FROM product WHERE uuid = ?`, id)
    return err
}
```

## SQL 実装の注意（sqlx）

- **プレースホルダ（`?`）でパラメータを渡す**（文字列連結はしない）
- `ORDER BY` / `LIMIT` / カラム名など動的SQLは **ホワイトリストで制限**
- `IN` 句は `sqlx.In` を使って安全に組み立てる

### 例（安全な書き方）

```go
db.GetContext(ctx, &row, `SELECT * FROM user WHERE email = ?`, email)
```

### 例（IN 句）

```go
query, args, _ := sqlx.In(`SELECT * FROM t WHERE id IN (?)`, ids)
query = db.Rebind(query)
db.SelectContext(ctx, &rows, query, args...)
```

## テスト方針（配置・命名・モック）

### 配置場所

- Domain: `internal/domain/<domain>/*_test.go`
- Usecase: `internal/usecase/<domain>/*_test.go`
- HTTP Handler: `internal/interface/http/handler/*_test.go`
- Repository（結合/必要箇所のみ）: `internal/infra/db/mysql/repository/*_test.go`

### 命名規則

- 基本: `<対象>_test.go`
- テスト関数: `Test<構造体/関数>_<条件>`（例: `TestCreateProduct_DuplicateName`）
- Table test を推奨（`name` フィールド付き）

### モックの置き場所

- Usecase テスト用の Repository モック:
  - `internal/usecase/<domain>/mock_repository_test.go`
- Handler テスト用の Usecase モック:
  - `internal/interface/http/handler/mocks_test.go`
- 共有したい場合のみ `internal/shared/testutil` に集約

### 優先順位

1. Domain（最優先）
2. Usecase
3. Handler（薄く）
4. Repository（必要な箇所のみ）

### testdata の扱い

- 大きな入力データは `testdata/` ディレクトリに置く
- 参照先は各テストファイルの近くに配置（例: `internal/usecase/product/testdata`）
- JSON/CSV などはテストで読み込み、固定値のベタ書きを避ける
- Go の `testing` では `testdata` 配下はビルド対象外になる

### DB 結合テストの起動方法（方針）

- Repository の結合テストのみ対象
- 起動方法は「Docker で MySQL を立ち上げる前提」に統一する
- テスト用の接続情報は `.env.test` などで管理し、CI でも再利用する

#### 例: 起動フロー（概念）

1. `docker-compose -f docker-compose.test.yml up -d`
2. `go test ./internal/infra/db/mysql/repository -count=1`
3. `docker-compose -f docker-compose.test.yml down`

#### ルール

- テスト用DBは本番/開発と完全に分離する
- 破壊的なマイグレーションは結合テスト専用DBでのみ実行する

## 進め方の原則

- まずドメイン単位の最小移行を行い、構成の正当性を確認する
- `product` は依存が多いので後半に回しても良い
- すべての新規実装はこの構成に沿って行う

## 実装フロー（フェーズ方針）

### フェーズ1: 最低限のAPIを完成形で実装

- GET/POST + Session を対象に「完成形」で実装する
- TODO は残さず、Domain/Usecase/Interface/Infra/DI/テストまで一通り行う
- 対象例: `category`（一覧取得 + 作成）

### フェーズ2: API実装の横展開

- フェーズ1の型を踏襲し、`target` / `tag` / `sales_site` などへ展開する

### フェーズ3: 複雑なドメインへの対応

- `product` など複雑な機能（画像/CSV/S3/スクレイピング）を段階的に実装

## フェーズ1（category）具体チェックリスト

### API 仕様の確定

- [ ] `GET /api/category?mode=all|used` を確定
- [ ] `POST /api/category` のレスポンスは `{"success": true}` で統一
- [ ] `POST /api/category` は session 認証必須
- [ ] `name` のバリデーションは 1〜20

### Domain

- [ ] `internal/domain/category/entity.go` を作成
- [ ] エンティティのバリデーションを実装（name 1〜20）
- [ ] `internal/domain/category/category_error.go` を作成
- [ ] `internal/domain/category/category_repo.go` を作成

### Usecase

- [ ] `internal/usecase/category/usecase.go` を作成（IF）
- [ ] `ListCategories(mode)` を実装
- [ ] `CreateCategory(name)` を実装
- [ ] Domain error → Usecase error 変換

### Interface（HTTP）

- [ ] `internal/interface/http/request/category.go` を作成
- [ ] `internal/interface/http/response/category.go` を作成
- [ ] `internal/interface/http/presenter/category.go` を作成
- [ ] `internal/interface/http/handler/category.go` を作成
- [ ] `internal/interface/http/router` にルーティング追加

### Session

- [ ] `internal/domain/session` を作成（entity/repository）
- [ ] `internal/usecase/session` を作成（認証判定）
- [ ] `internal/infra/db/mysql/repository/session.go` を作成
- [ ] router で `AuthMiddleware` を付与

### Infra

- [ ] `internal/infra/db/mysql/repository/category.go` を作成
- [ ] `internal/infra/config` に必要項目を追加

### DI

- [ ] `internal/app/di` に Category/Session の組み立てを追加

### テスト

- [ ] Domain テスト（Category）
- [ ] Usecase テスト（Category）
- [ ] Handler テスト（Category）
