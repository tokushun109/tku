import { setupServer } from 'msw/node'

import { handlers } from './handlers'

// MSWサーバーをセットアップ
export const server = setupServer(...handlers)
