import { afterEach, describe, expect, it, vi } from 'vitest'

import { createAPIHeaders } from '@/utils/cookie'

describe('cookie utils', () => {
    afterEach(() => {
        vi.restoreAllMocks()
    })

    it('ブラウザ環境ではCookieヘッダーを付けずにヘッダーを返す', async () => {
        const consoleErrorSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        const headers = await createAPIHeaders({
            'X-Test-Header': 'test-value',
        })

        expect(headers).toEqual({
            'Content-Type': 'application/json',
            'X-Test-Header': 'test-value',
        })
        expect(consoleErrorSpy).not.toHaveBeenCalled()
    })
})
