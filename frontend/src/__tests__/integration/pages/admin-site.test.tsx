import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getSalesSiteList } from '@/apis/salesSite'
import { getSnsList } from '@/apis/sns'
import SitePage, { metadata } from '@/app/admin/site/page'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/salesSite')
vi.mock('@/apis/sns')

const mockGetSalesSiteList = vi.mocked(getSalesSiteList)
const mockGetSnsList = vi.mocked(getSnsList)

describe('Admin Site Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('管理画面サイト管理ページが正常に表示される', async () => {
        // モックデータの設定
        mockGetSnsList.mockResolvedValue([
            {
                uuid: 'sns-1',
                name: 'Instagram',
                url: 'https://instagram.com/example',
            },
        ])

        mockGetSalesSiteList.mockResolvedValue([
            {
                uuid: 'site-1',
                name: 'Creema',
                url: 'https://creema.jp/example',
            },
        ])

        // コンポーネントをレンダリング
        render(await SitePage())

        // サイト管理画面の表示を確認
        await waitFor(() => {
            expect(screen.getByText('SNS')).toBeInTheDocument()
            expect(screen.getByText('販売サイト')).toBeInTheDocument()
        })

        // API呼び出しの確認
        expect(mockGetSnsList).toHaveBeenCalledTimes(1)
        expect(mockGetSalesSiteList).toHaveBeenCalledTimes(1)
    })

    it('APIエラー時のフォールバック処理', async () => {
        // コンソールエラーをモック
        const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        // モックでエラーを発生させる
        mockGetSnsList.mockRejectedValue(new Error('SNS API Error'))
        mockGetSalesSiteList.mockRejectedValue(new Error('Sales Site API Error'))

        // コンポーネントをレンダリング
        render(await SitePage())

        // エラーログが出力されることを確認
        await waitFor(() => {
            expect(consoleSpy).toHaveBeenCalledWith('データの取得に失敗しました:', expect.any(Error))
        })

        // 空配列でテンプレートが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('SNS')).toBeInTheDocument()
            expect(screen.getByText('販売サイト')).toBeInTheDocument()
        })

        consoleSpy.mockRestore()
    })

    it('ページメタデータでnoindex, nofollowが設定されている', () => {
        // メタデータの確認
        expect(metadata.title).toBe('サイト管理 | admin')
        expect(metadata.robots).toEqual({
            index: false,
            follow: false,
        })
    })
})
