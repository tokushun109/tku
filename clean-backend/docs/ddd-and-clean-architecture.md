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
        repository.go
        service.go
      category/
        entity.go
        repository.go
      target/
        entity.go
        repository.go
      tag/
        entity.go
        repository.go
      creator/
        entity.go
        repository.go
      contact/
        entity.go
        repository.go
      user/
        entity.go
        repository.go
      session/
        entity.go
        repository.go
      sales_site/
        entity.go
        repository.go
      skill_market/
        entity.go
        repository.go
      sns/
        entity.go
        repository.go
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
          gorm.go
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

## 依存ルール（重要）

- `domain` は外部に依存しない（純粋なルールとモデルのみ）
- `usecase` は `domain` に依存してよい
- `interface` は `usecase` に依存してよい
- `infra` は `domain` / `usecase` に依存してよい
- 依存の向きは常に「外側 → 内側」

## Interface の配置ルール

- Usecase の IF は `internal/usecase/<domain>/usecase.go`
- Repository の IF は `internal/domain/<domain>/repository.go`
- Repository の実装は `internal/infra/db/mysql/repository`

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

| 現状処理 | 目的/ルール | 移動先レイヤ | 移動先の例 | 補足 |
| --- | --- | --- | --- | --- |
| 重複チェック（Category/Target/Tag） | 名称の重複防止 | Usecase (+ Domain Service 任意) | `internal/usecase/*` | DB参照が必要。Usecaseで`repo.ExistsByName`を使う。 |
| 削除時の関連整理 | 関連テーブルの整合性維持 | Usecase | `internal/usecase/*` | Repositoryで更新・削除。Usecaseでトランザクション制御。 |
| スクレイピング複製 | 外部サイトから商品作成 | Usecase + Infra | Usecase: `internal/usecase/product` / Infra: `internal/infra/marketplace/creema` | HTTP/HTML解析はInfra。 |
| CSV更新 | CSVからの更新処理 | Usecase | `internal/usecase/product` | CSV DTO は `interface` 層に置く。 |
| 画像削除/ファイル操作 | DB更新 + ファイル削除 | Usecase + Infra | Usecase: `internal/usecase/product` / Infra: `internal/infra/storage` | I/OはInfra。順序制御はUsecase。 |
| セッション判定 | 有効性チェック | Usecase | `internal/usecase/session` | 将来期限が入る場合はDomainにルール化。 |
| パスワードハッシュ | 暗号化 | Domain Service or Infra | `internal/domain/user/password.go` or `internal/infra/crypto` | 実装差し替えを見込むならInfra。 |
| 初期データ投入 | seed/migration | Infra (Seed/Migration) | `db/seed` or `internal/infra/seed` | 起動時に行わない。 |
| ロゴ更新時の旧画像削除 | ストレージ操作 | Usecase + Infra | `internal/usecase/creator` + `internal/infra/storage` | 削除はStorageの責務。 |

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

- `internal/domain/<domain>/errors.go`
- `var ErrCategoryNameDuplicate = errors.New(\"category name is duplicate\")`

### Usecase エラーの種別（例）

- `ErrInvalidInput`
- `ErrNotFound`
- `ErrConflict`
- `ErrInternal`

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
    id := id.New() // internal/shared/id
    r.Header.Set("X-Request-ID", id)
    return id
}
```

### UUID の配置

- UUID 生成は `internal/shared/id` に置く
- 例: `internal/shared/id/uuid.go` に `func New() string` を用意

## マイグレーション

- `db/migrations` をルート直下に置く（DB切り替え予定なしのため固定）

## 進め方の原則

- まずドメイン単位の最小移行を行い、構成の正当性を確認する
- `product` は依存が多いので後半に回しても良い
- すべての新規実装はこの構成に沿って行う
