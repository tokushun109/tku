import { NextRequest } from 'next/server'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCurrentUser } from '@/apis/auth'
import { healthCheck } from '@/apis/healthCheck'
import { middleware } from '@/middleware'

vi.mock('@/apis/auth', () => ({
    getCurrentUser: vi.fn(),
}))

vi.mock('@/apis/healthCheck', () => ({
    healthCheck: vi.fn(),
}))

const mockedGetCurrentUser = vi.mocked(getCurrentUser)
const mockedHealthCheck = vi.mocked(healthCheck)

function createRequest(pathname: string, headers?: Record<string, string>) {
    return new NextRequest(new URL(pathname, 'https://tocoriri.com'), { headers })
}

describe('middleware', () => {
    beforeEach(() => {
        vi.clearAllMocks()
        delete process.env.ENV
        delete process.env.MY_IP_ADDRESS

        mockedHealthCheck.mockResolvedValue({})
        mockedGetCurrentUser.mockResolvedValue(null)
    })

    it('MY_IP_ADDRESSと異なるIPからadminへアクセスしたならNotFoundページへrewriteする', async () => {
        process.env.MY_IP_ADDRESS = '203.0.113.1'

        const response = await middleware(
            createRequest('/admin/product', {
                'x-forwarded-for': '198.51.100.1',
            }),
        )

        expect(response.status).toBe(404)
        expect(response.headers.get('x-middleware-rewrite')).toBe('https://tocoriri.com/not-found')
        expect(mockedHealthCheck).not.toHaveBeenCalled()
        expect(mockedGetCurrentUser).not.toHaveBeenCalled()
    })

    it('MY_IP_ADDRESSと同じIPからadminへアクセスしたなら既存の認証チェックを行う', async () => {
        process.env.MY_IP_ADDRESS = '203.0.113.1'

        const response = await middleware(
            createRequest('/admin/product', {
                'x-forwarded-for': '203.0.113.1, 198.51.100.1',
            }),
        )

        expect(response.status).toBe(307)
        expect(response.headers.get('location')).toBe('https://tocoriri.com/admin/login')
        expect(mockedHealthCheck).toHaveBeenCalledTimes(1)
        expect(mockedGetCurrentUser).not.toHaveBeenCalled()
    })

    it('x-forwarded-forの末尾だけが許可IPならNotFoundページへrewriteする', async () => {
        process.env.MY_IP_ADDRESS = '203.0.113.1'

        const response = await middleware(
            createRequest('/admin/product', {
                'x-forwarded-for': '198.51.100.1, 203.0.113.1',
            }),
        )

        expect(response.status).toBe(404)
        expect(response.headers.get('x-middleware-rewrite')).toBe('https://tocoriri.com/not-found')
        expect(mockedHealthCheck).not.toHaveBeenCalled()
    })

    it('ENVがlocalならMY_IP_ADDRESSが設定済みでもadminのIP制限を行わない', async () => {
        process.env.ENV = 'local'
        process.env.MY_IP_ADDRESS = '203.0.113.1'

        const response = await middleware(createRequest('/admin/product'))

        expect(response.status).toBe(307)
        expect(response.headers.get('location')).toBe('https://tocoriri.com/admin/login')
        expect(mockedHealthCheck).toHaveBeenCalledTimes(1)
    })

    it('local以外でMY_IP_ADDRESSが未設定ならNotFoundページへrewriteする', async () => {
        const response = await middleware(
            createRequest('/admin/product', {
                'x-forwarded-for': '198.51.100.1',
            }),
        )

        expect(response.status).toBe(404)
        expect(response.headers.get('x-middleware-rewrite')).toBe('https://tocoriri.com/not-found')
        expect(mockedHealthCheck).not.toHaveBeenCalled()
    })
})
