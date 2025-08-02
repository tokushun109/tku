import { beforeEach, describe, expect, it, vi } from 'vitest'

import AdminProductPage, { metadata } from '@/app/admin/product/page'

import { render, screen } from '../helpers'

describe('Admin Product Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('管理画面商品管理ページが正常に表示される', () => {
        // コンポーネントをレンダリング
        render(<AdminProductPage />)

        // 商品管理画面の表示を確認
        expect(screen.getByText('商品管理')).toBeInTheDocument()
    })

    it('ページメタデータでnoindex, nofollowが設定されている', () => {
        // メタデータの確認
        expect(metadata.title).toBe('商品管理 | admin')
        expect(metadata.robots).toEqual({
            index: false,
            follow: false,
        })
    })
})
