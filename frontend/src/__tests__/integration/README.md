# Integration Tests

このディレクトリには、ページ単位のIntegration Testが含まれています。

## 実行方法

```bash
# 統合テストのみ実行
yarn test-integration

# 全てのテストを実行
yarn test-all
```

## テスト構成

### helpers/

- `render.tsx` - レンダリング用のヘルパー関数
- `mock-router.tsx` - Next.js Router のモック
- `test-utils.ts` - テスト用のユーティリティ関数

### pages/

- `home.test.tsx` - ホームページのテスト
- `product.test.tsx` - 商品一覧ページのテスト
- `product-detail.test.tsx` - 商品詳細ページのテスト
- `about.test.tsx` - Aboutページのテスト
- `contact.test.tsx` - お問い合わせページのテスト

## テスト方針

1. **ページ単位でのテスト**: 各ページが正常に表示されることを確認
2. **API連携のテスト**: MSWを使用してAPIのモック化
3. **エラーハンドリングのテスト**: API呼び出し失敗時の挙動確認
4. **フォーム機能のテスト**: 入力値の検証、送信処理の確認
5. **レスポンシブデザインのテスト**: 異なる画面サイズでの表示確認

## モックデータ

`../mocks/handlers.ts` にAPIのモックデータが定義されています。
テスト用のデータは実際のAPIレスポンスに近い形で作成されています。

## 注意事項

- テストはjsdom環境で実行されます
- Next.jsのフォント関数やImageコンポーネントはモック化されています
- APIレスポンスは`../mocks/handlers.ts`で定義されたデータを使用します
- テスト実行時は`@testing-library/jest-dom`の拡張マッチャーが利用可能です
