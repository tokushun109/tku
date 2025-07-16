import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCreator } from '@/apis/creator'
import { getCarouselImages } from '@/apis/product'
import Home from '@/app/(home)/page'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/creator')
vi.mock('@/apis/product')

const mockGetCreator = vi.mocked(getCreator)
const mockGetCarouselImages = vi.mocked(getCarouselImages)

describe('Home Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('ホームページが正常に表示される', async () => {
        // モックデータの設定
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetCarouselImages.mockResolvedValue([
            {
                apiPath: '/image/carousel1.jpg',
                product: {
                    uuid: 'earrings-women-1',
                    name: '女性向けイヤリング1',
                    price: 1500,
                    description: '女性向けイヤリング1の詳細',
                    isActive: true,
                    isRecommend: true,
                    category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                    target: { name: '女性', uuid: 'women-uuid' },
                    tags: [],
                    productImages: [],
                    siteDetails: [],
                },
            },
        ])

        // コンポーネントをレンダリング
        render(await Home())

        // カルーセルの表示を確認
        await waitFor(() => {
            expect(screen.getByTestId('carousel')).toBeInTheDocument()
        })

        // 商品情報の表示を確認
        await waitFor(() => {
            expect(screen.getAllByText('女性向けイヤリング1')).toHaveLength(4)
            expect(screen.getAllByText('イヤリング')).toHaveLength(4)
        })
    })

    it('カルーセル画像が空の場合でもエラーにならない', async () => {
        // モックデータの設定
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetCarouselImages.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await Home())

        // エラーが発生しないことを確認
        await waitFor(() => {
            expect(screen.queryByTestId('error-message')).not.toBeInTheDocument()
        })
    })

    it('API呼び出しが失敗した場合のエラーハンドリング', async () => {
        // モックでエラーを発生させる
        mockGetCreator.mockRejectedValue(new Error('API Error'))
        mockGetCarouselImages.mockRejectedValue(new Error('API Error'))

        // エラーが投げられることを確認
        await expect(Home()).rejects.toThrow('API Error')
    })

    it('クリエイター情報がない場合のフォールバック', async () => {
        // モックデータの設定（クリエイター情報なし）
        mockGetCarouselImages.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await Home())

        // フォールバックが適用されることを確認
        await waitFor(() => {
            expect(screen.queryByTestId('error-message')).not.toBeInTheDocument()
        })
    })

    it('カルーセル画像が複数ある場合の表示', async () => {
        // モックデータの設定
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetCarouselImages.mockResolvedValue([
            {
                apiPath: '/image/carousel1.jpg',
                product: {
                    uuid: 'earrings-women-1',
                    name: '女性向けイヤリング1',
                    price: 1500,
                    description: '女性向けイヤリング1の詳細',
                    isActive: true,
                    isRecommend: true,
                    category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                    target: { name: '女性', uuid: 'women-uuid' },
                    tags: [],
                    productImages: [],
                    siteDetails: [],
                },
            },
            {
                apiPath: '/image/carousel2.jpg',
                product: {
                    uuid: 'rings-men-1',
                    name: '男性向けリング1',
                    price: 3000,
                    description: '男性向けリング1の詳細',
                    isActive: true,
                    isRecommend: false,
                    category: { name: 'リング', uuid: 'rings-uuid' },
                    target: { name: '男性', uuid: 'men-uuid' },
                    tags: [],
                    productImages: [],
                    siteDetails: [],
                },
            },
        ])

        // コンポーネントをレンダリング
        render(await Home())

        // カルーセルの表示を確認
        await waitFor(() => {
            expect(screen.getByTestId('carousel')).toBeInTheDocument()
        })
    })
})
