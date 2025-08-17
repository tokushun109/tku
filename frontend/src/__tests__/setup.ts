import React from 'react'
import { beforeAll, afterEach, afterAll, vi } from 'vitest'
import '@testing-library/jest-dom'

import { server } from './mocks/server'

// MSW サーバーを起動
beforeAll(() => server.listen())

// 各テスト後にハンドラーをリセット
afterEach(() => server.resetHandlers())

// テスト終了後にサーバーを停止
afterAll(() => server.close())

// process.env の設定
process.env.API_BASE_URL = 'http://localhost:8080'

// React の global 設定
global.React = React

// IntersectionObserver のモック
global.IntersectionObserver = vi.fn().mockImplementation(() => ({
    observe: vi.fn(),
    unobserve: vi.fn(),
    disconnect: vi.fn(),
}))

// ResizeObserver のモック
global.ResizeObserver = vi.fn().mockImplementation(() => ({
    observe: vi.fn(),
    unobserve: vi.fn(),
    disconnect: vi.fn(),
}))

// matchMedia のモック
Object.defineProperty(window, 'matchMedia', {
    writable: true,
    value: vi.fn().mockImplementation((query) => ({
        matches: false,
        media: query,
        onchange: null,
        addListener: vi.fn(),
        removeListener: vi.fn(),
        addEventListener: vi.fn(),
        removeEventListener: vi.fn(),
        dispatchEvent: vi.fn(),
    })),
})

// Next.js フォント関数のモック
const mockFont = () => ({
    style: {
        fontFamily: '"Sawarabi Mincho", serif',
        fontWeight: 400,
    },
    className: 'mock-font',
})

// フォント関数のモック
vi.mock('next/font/google', () => ({
    Sawarabi_Mincho: mockFont,
    Lobster: mockFont,
}))

// Next.js コンポーネントのモック
vi.mock('next/image', () => ({
    default: (props: any) => {
        const { src, alt, width, height, fill, priority, ...restProps } = props
        return global.React.createElement('img', {
            src,
            alt,
            width: width || undefined,
            height: height || undefined,
            'data-fill': fill ? 'true' : undefined,
            'data-priority': priority ? 'true' : undefined,
            ...restProps,
        })
    },
}))

vi.mock('next/link', () => ({
    default: (props: any) => {
        const { href, children, ...restProps } = props
        return global.React.createElement(
            'a',
            {
                href,
                ...restProps,
            },
            children,
        )
    },
}))

// Next.js navigation のモック
vi.mock('next/navigation', () => ({
    useRouter: () => ({
        push: vi.fn(),
        replace: vi.fn(),
        refresh: vi.fn(),
        back: vi.fn(),
        forward: vi.fn(),
        prefetch: vi.fn(),
    }),
    useSearchParams: () => ({
        get: vi.fn(),
        getAll: vi.fn(),
        has: vi.fn(),
        keys: vi.fn(),
        values: vi.fn(),
        entries: vi.fn(),
        forEach: vi.fn(),
        toString: vi.fn(),
    }),
    usePathname: () => '/',
    useParams: () => ({}),
    notFound: vi.fn(),
    redirect: vi.fn(),
}))
