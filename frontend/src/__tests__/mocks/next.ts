import { vi } from 'vitest'

// Next.jsのモジュールをモック
export const mockRouter = {
    push: vi.fn(),
    replace: vi.fn(),
    back: vi.fn(),
    forward: vi.fn(),
    refresh: vi.fn(),
    prefetch: vi.fn(),
}

export const mockUseRouter = vi.fn(() => mockRouter)

// Next.jsのnavigationモジュールをモック
vi.mock('next/navigation', () => ({
    useRouter: mockUseRouter,
}))
