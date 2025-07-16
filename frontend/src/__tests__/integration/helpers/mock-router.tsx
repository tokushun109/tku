import { useRouter } from 'next/navigation'
import { vi } from 'vitest'

// Next.jsルーターのモック
export const mockRouter = {
    push: vi.fn(),
    replace: vi.fn(),
    refresh: vi.fn(),
    back: vi.fn(),
    forward: vi.fn(),
    prefetch: vi.fn(),
}

// useRouterフック全体をモック
export const mockUseRouter = () => {
    vi.mocked(useRouter).mockReturnValue(mockRouter as any)
}

// ルーターのモック状態をリセット
export const resetRouterMock = () => {
    Object.values(mockRouter).forEach((fn) => {
        if (vi.isMockFunction(fn)) {
            fn.mockClear()
        }
    })
}
