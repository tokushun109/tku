import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCategories } from '@/apis/category'
import { getProductsByCategory } from '@/apis/product'
import { getTargets } from '@/apis/target'
import Product from '@/app/(contents)/product/page'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/category')
vi.mock('@/apis/product')
vi.mock('@/apis/target')

const mockGetCategories = vi.mocked(getCategories)
const mockGetProductsByCategory = vi.mocked(getProductsByCategory)
const mockGetTargets = vi.mocked(getTargets)

describe('Product Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('商品一覧ページが正常に表示される', async () => {
        // モックデータの設定
        mockGetCategories.mockResolvedValue([
            { name: 'イヤリング', uuid: 'earrings-uuid' },
            { name: 'リング', uuid: 'rings-uuid' },
        ])

        mockGetTargets.mockResolvedValue([
            { name: '女性', uuid: 'women-uuid' },
            { name: '男性', uuid: 'men-uuid' },
        ])

        mockGetProductsByCategory.mockResolvedValue([
            {
                category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                products: [
                    {
                        uuid: 'earrings-women-1',
                        name: '女性向けイヤリング1',
                        price: 1500,
                        description: '女性向けイヤリング1の詳細',
                        isActive: true,
                        isRecommend: true,
                        category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                        target: { name: '女性', uuid: 'women-uuid' },
                        tags: [],
                        productImages: [
                            {
                                apiPath: '/image/earrings-women-1.jpg',
                                name: 'earrings-women-1.jpg',
                                order: 1,
                                uuid: 'earrings-women-1-image-uuid',
                            },
                        ],
                        siteDetails: [],
                    },
                ],
            },
            {
                category: { name: 'リング', uuid: 'rings-uuid' },
                products: [
                    {
                        uuid: 'rings-men-1',
                        name: '男性向けリング1',
                        price: 3000,
                        description: '男性向けリング1の詳細',
                        isActive: true,
                        isRecommend: false,
                        category: { name: 'リング', uuid: 'rings-uuid' },
                        target: { name: '男性', uuid: 'men-uuid' },
                        tags: [],
                        productImages: [
                            {
                                apiPath: '/image/rings-men-1.jpg',
                                name: 'rings-men-1.jpg',
                                order: 1,
                                uuid: 'rings-men-1-image-uuid',
                            },
                        ],
                        siteDetails: [],
                    },
                ],
            },
        ])

        // コンポーネントをレンダリング
        render(await Product())

        // カテゴリーの表示を確認
        await waitFor(() => {
            expect(screen.getAllByText('イヤリング')).toHaveLength(2)
            expect(screen.getAllByText('リング')).toHaveLength(2)
        })

        // 商品情報の表示を確認
        await waitFor(() => {
            expect(screen.getByText('女性向けイヤリング1')).toBeInTheDocument()
            expect(screen.getByText('男性向けリング1')).toBeInTheDocument()
            expect(screen.getByText('¥1,500')).toBeInTheDocument()
            expect(screen.getByText('¥3,000')).toBeInTheDocument()
        })
    })

    it('商品がない場合でもエラーにならない', async () => {
        // モックデータの設定（商品なし）
        mockGetCategories.mockResolvedValue([])
        mockGetTargets.mockResolvedValue([])
        mockGetProductsByCategory.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await Product())

        // エラーが発生しないことを確認
        await waitFor(() => {
            expect(screen.queryByTestId('error-message')).not.toBeInTheDocument()
        })
    })

    it('API呼び出しが失敗した場合のエラーハンドリング', async () => {
        // モックでエラーを発生させる
        mockGetCategories.mockRejectedValue(new Error('API Error'))
        mockGetProductsByCategory.mockRejectedValue(new Error('API Error'))
        mockGetTargets.mockRejectedValue(new Error('API Error'))

        // エラーが投げられることを確認
        await expect(Product()).rejects.toThrow('API Error')
    })

    it('特定のカテゴリーに商品が複数ある場合の表示', async () => {
        // モックデータの設定
        mockGetCategories.mockResolvedValue([{ name: 'イヤリング', uuid: 'earrings-uuid' }])

        mockGetTargets.mockResolvedValue([{ name: '女性', uuid: 'women-uuid' }])

        mockGetProductsByCategory.mockResolvedValue([
            {
                category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                products: [
                    {
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
                    {
                        uuid: 'earrings-women-2',
                        name: '女性向けイヤリング2',
                        price: 2000,
                        description: '女性向けイヤリング2の詳細',
                        isActive: true,
                        isRecommend: false,
                        category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                        target: { name: '女性', uuid: 'women-uuid' },
                        tags: [],
                        productImages: [],
                        siteDetails: [],
                    },
                ],
            },
        ])

        // コンポーネントをレンダリング
        render(await Product())

        // カテゴリーの表示を確認
        await waitFor(() => {
            expect(screen.getAllByText('イヤリング')).toHaveLength(2)
        })

        // 複数の商品情報が表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('女性向けイヤリング1')).toBeInTheDocument()
            expect(screen.getByText('女性向けイヤリング2')).toBeInTheDocument()
            expect(screen.getByText('¥1,500')).toBeInTheDocument()
            expect(screen.getByText('¥2,000')).toBeInTheDocument()
        })
    })

    it('フィルタリング機能が正常に動作する', async () => {
        // モックデータの設定
        mockGetCategories.mockResolvedValue([
            { name: 'イヤリング', uuid: 'earrings-uuid' },
            { name: 'リング', uuid: 'rings-uuid' },
        ])

        mockGetTargets.mockResolvedValue([
            { name: '女性', uuid: 'women-uuid' },
            { name: '男性', uuid: 'men-uuid' },
        ])

        mockGetProductsByCategory.mockResolvedValue([
            {
                category: { name: 'イヤリング', uuid: 'earrings-uuid' },
                products: [
                    {
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
                ],
            },
        ])

        // コンポーネントをレンダリング
        render(await Product())

        // フィルタリングオプションの表示を確認
        await waitFor(() => {
            expect(screen.getAllByText('女性')).toHaveLength(2)
            expect(screen.getAllByText('男性')).toHaveLength(1)
        })

        // 商品が表示されていることを確認
        await waitFor(() => {
            expect(screen.getByText('女性向けイヤリング1')).toBeInTheDocument()
        })
    })
})
