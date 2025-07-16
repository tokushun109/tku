import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getProduct } from '@/apis/product'
import ProductDetail from '@/app/(contents)/product/[uuid]/page'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/product')

const mockGetProduct = vi.mocked(getProduct)

describe('Product Detail Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('商品詳細ページが正常に表示される', async () => {
        // モックデータの設定
        mockGetProduct.mockResolvedValue({
            uuid: 'earrings-women-1',
            name: '女性向けイヤリング1',
            price: 1500,
            description: '女性向けイヤリング1の詳細\n日常使いにも特別な日にもお使いいただけます。',
            isActive: true,
            isRecommend: true,
            category: { name: 'イヤリング', uuid: 'earrings-uuid' },
            target: { name: '女性', uuid: 'women-uuid' },
            tags: [
                { name: 'ハンドメイド', uuid: 'handmade-uuid' },
                { name: 'シンプル', uuid: 'simple-uuid' },
            ],
            productImages: [
                {
                    apiPath: '/image/earrings-women-1.jpg',
                    name: 'earrings-women-1.jpg',
                    order: 1,
                    uuid: 'earrings-women-1-image-uuid',
                },
                {
                    apiPath: '/image/earrings-women-1-2.jpg',
                    name: 'earrings-women-1-2.jpg',
                    order: 2,
                    uuid: 'earrings-women-1-2-image-uuid',
                },
            ],
            siteDetails: [
                {
                    uuid: 'site-detail-1',
                    detailUrl: 'https://example.com/product/1',
                    salesSite: { name: 'ECサイト1', uuid: 'site-1-uuid' },
                },
            ],
        })

        // コンポーネントをレンダリング
        const params = Promise.resolve({ uuid: 'earrings-women-1' })
        render(await ProductDetail({ params }))

        // 商品情報の表示を確認
        await waitFor(() => {
            expect(screen.getAllByText('女性向けイヤリング1')).toHaveLength(2)
            expect(
                screen.getAllByText((_, element) => {
                    return element?.textContent === '￥1,500(税込)'
                }),
            ).toHaveLength(2)
            expect(screen.getByText(/女性向けイヤリング1の詳細/)).toBeInTheDocument()
            expect(screen.getByText(/日常使いにも特別な日にもお使いいただけます。/)).toBeInTheDocument()
        })

        // カテゴリーとターゲットの表示を確認
        await waitFor(() => {
            expect(screen.getByText('イヤリング')).toBeInTheDocument()
            expect(screen.getByText('女性')).toBeInTheDocument()
        })

        // タグの表示を確認
        await waitFor(() => {
            expect(screen.getByText('ハンドメイド')).toBeInTheDocument()
            expect(screen.getByText('シンプル')).toBeInTheDocument()
        })

        // 画像の表示を確認
        await waitFor(() => {
            const images = screen.getAllByRole('img')
            expect(images.length).toBeGreaterThanOrEqual(2)
        })
    })

    it('商品が見つからない場合のエラーハンドリング', async () => {
        // モックで404エラーを発生させる
        mockGetProduct.mockRejectedValue(new Error('Product not found'))

        // エラーが投げられることを確認
        const params = Promise.resolve({ uuid: 'non-existent-uuid' })
        await expect(ProductDetail({ params })).rejects.toThrow('Product not found')
    })

    it('画像がない商品の場合でも正常に表示される', async () => {
        // モックデータの設定（画像なし）
        mockGetProduct.mockResolvedValue({
            uuid: 'earrings-women-1',
            name: '女性向けイヤリング1',
            price: 1500,
            description: '女性向けイヤリング1の詳細',
            isActive: true,
            isRecommend: true,
            category: { name: 'イヤリング', uuid: 'earrings-uuid' },
            target: { name: '女性', uuid: 'women-uuid' },
            tags: [],
            productImages: [], // 画像なし
            siteDetails: [],
        })

        // コンポーネントをレンダリング
        const params = Promise.resolve({ uuid: 'earrings-women-1' })
        render(await ProductDetail({ params }))

        // 商品情報の表示を確認
        await waitFor(() => {
            expect(screen.getAllByText('女性向けイヤリング1')).toHaveLength(2)
            expect(
                screen.getAllByText((_, element) => {
                    return element?.textContent === '￥1,500(税込)'
                }),
            ).toHaveLength(2)
        })

        // エラーが発生しないことを確認
        await waitFor(() => {
            expect(screen.queryByTestId('error-message')).not.toBeInTheDocument()
        })
    })

    it('販売サイトのリンクが正常に表示される', async () => {
        // モックデータの設定
        mockGetProduct.mockResolvedValue({
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
            siteDetails: [
                {
                    uuid: 'site-detail-1',
                    detailUrl: 'https://creema.jp/item/1',
                    salesSite: { name: 'Creema', uuid: 'creema-uuid' },
                },
                {
                    uuid: 'site-detail-2',
                    detailUrl: 'https://minne.com/item/1',
                    salesSite: { name: 'minne', uuid: 'minne-uuid' },
                },
            ],
        })

        // コンポーネントをレンダリング
        const params = Promise.resolve({ uuid: 'earrings-women-1' })
        render(await ProductDetail({ params }))

        // 販売サイトリンクの表示を確認
        await waitFor(() => {
            expect(screen.getByText('Creema')).toBeInTheDocument()
            expect(screen.getByText('minne')).toBeInTheDocument()
        })

        // ボタンが正しく表示されていることを確認
        await waitFor(() => {
            const creemaButton = screen.getByRole('button', { name: /Creema/i })
            const minneButton = screen.getByRole('button', { name: /minne/i })
            expect(creemaButton).toBeInTheDocument()
            expect(minneButton).toBeInTheDocument()
        })
    })

    it('タグが多い場合でも正常に表示される', async () => {
        // モックデータの設定（多数のタグ）
        mockGetProduct.mockResolvedValue({
            uuid: 'earrings-women-1',
            name: '女性向けイヤリング1',
            price: 1500,
            description: '女性向けイヤリング1の詳細',
            isActive: true,
            isRecommend: true,
            category: { name: 'イヤリング', uuid: 'earrings-uuid' },
            target: { name: '女性', uuid: 'women-uuid' },
            tags: [
                { name: 'ハンドメイド', uuid: 'handmade-uuid' },
                { name: 'シンプル', uuid: 'simple-uuid' },
                { name: '上品', uuid: 'elegant-uuid' },
                { name: 'カジュアル', uuid: 'casual-uuid' },
                { name: 'ビジネス', uuid: 'business-uuid' },
            ],
            productImages: [],
            siteDetails: [],
        })

        // コンポーネントをレンダリング
        const params = Promise.resolve({ uuid: 'earrings-women-1' })
        render(await ProductDetail({ params }))

        // すべてのタグが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('ハンドメイド')).toBeInTheDocument()
            expect(screen.getByText('シンプル')).toBeInTheDocument()
            expect(screen.getByText('上品')).toBeInTheDocument()
            expect(screen.getByText('カジュアル')).toBeInTheDocument()
            expect(screen.getByText('ビジネス')).toBeInTheDocument()
        })
    })
})
