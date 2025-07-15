import { beforeAll, afterEach, afterAll } from 'vitest'

import { server } from './mocks/server'

// MSW サーバーを起動
beforeAll(() => server.listen())

// 各テスト後にハンドラーをリセット
afterEach(() => server.resetHandlers())

// テスト終了後にサーバーを停止
afterAll(() => server.close())

// process.env の設定
process.env.API_URL = 'http://localhost:8080'
