import { render, RenderOptions } from '@testing-library/react'
import { ReactElement } from 'react'

// App Routerコンポーネント用のテストラッパー
const TestWrapper = ({ children }: { children: React.ReactNode }) => {
    // eslint-disable-next-line react/jsx-no-useless-fragment
    return <>{children}</>
}

// Next.jsアプリケーションのテストヘルパー関数
export const renderWithProviders = (ui: ReactElement, options?: Omit<RenderOptions, 'wrapper'>) => {
    return render(ui, {
        wrapper: TestWrapper,
        ...options,
    })
}

// 一般的なテストヘルパー関数をre-export
export * from '@testing-library/react'
export { renderWithProviders as render }
export { expect, describe, it, beforeEach, afterEach, vi } from 'vitest'
