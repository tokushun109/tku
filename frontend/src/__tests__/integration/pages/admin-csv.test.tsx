import { render, screen } from '@testing-library/react'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { metadata } from '@/app/admin/csv/page'
import { AdminCsvTemplate } from '@/app/admin/csv/template'

describe('Admin CSV Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('CSV操作ページが正常に表示される', () => {
        // コンポーネントをレンダリング
        render(<AdminCsvTemplate />)

        // ページタイトルの表示を確認
        expect(screen.getByText('商品レコード')).toBeInTheDocument()

        // ダウンロードボタンの表示を確認
        expect(screen.getByRole('button', { name: 'ダウンロード' })).toBeInTheDocument()

        // アップロードボタンの表示を確認
        expect(screen.getByRole('button', { name: 'アップロード' })).toBeInTheDocument()
    })

    it('ページメタデータでnoindex, nofollowが設定されている', () => {
        // メタデータの確認
        expect(metadata.title).toBe('CSV操作 | tocoriri')
        expect(metadata.description).toBe('商品レコードのCSVダウンロード・アップロード機能')
        expect(metadata.robots).toEqual({
            index: false,
            follow: false,
        })
    })
})
