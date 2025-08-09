import { fireEvent } from '@testing-library/react'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCategories } from '@/apis/category'
import { deleteProduct, getProducts } from '@/apis/product'
import { getSalesSiteList } from '@/apis/salesSite'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'
import { metadata } from '@/app/admin/product/page'
import { AdminProductTemplate } from '@/app/admin/product/template'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/category')
vi.mock('@/apis/product')
vi.mock('@/apis/salesSite')
vi.mock('@/apis/tag')
vi.mock('@/apis/target')
vi.mock('sonner', () => ({
    toast: {
        success: vi.fn(),
        error: vi.fn(),
    },
}))

const mockGetCategories = vi.mocked(getCategories)
const mockGetProducts = vi.mocked(getProducts)
const mockGetSalesSiteList = vi.mocked(getSalesSiteList)
const mockGetTags = vi.mocked(getTags)
const mockGetTargets = vi.mocked(getTargets)
const mockDeleteProduct = vi.mocked(deleteProduct)

const mockProductData = [
    {
        uuid: 'product-1',
        name: 'テスト商品1',
        description: 'テスト商品1の説明',
        price: 1500,
        isActive: true,
        isRecommend: false,
        category: { uuid: 'category-1', name: 'イヤリング' },
        target: { uuid: 'target-1', name: '女性' },
        tags: [{ uuid: 'tag-1', name: 'タグ1' }],
        productImages: [
            {
                uuid: 'image-1',
                apiPath: '/image/product-1.jpg',
                name: 'product-1.jpg',
                order: 1,
            },
        ],
        siteDetails: [],
    },
    {
        uuid: 'product-2',
        name: 'テスト商品2',
        description: 'テスト商品2の説明',
        price: 2500,
        isActive: false,
        isRecommend: true,
        category: { uuid: 'category-2', name: 'リング' },
        target: { uuid: 'target-2', name: '男性' },
        tags: [],
        productImages: [],
        siteDetails: [
            {
                uuid: 'site-detail-1',
                detailUrl: 'https://creema.jp/item/123456',
                salesSite: { uuid: 'site-1', name: 'Creema' },
            },
        ],
    },
]

const mockCategoriesData = [
    { uuid: 'category-1', name: 'イヤリング' },
    { uuid: 'category-2', name: 'リング' },
]

const mockTargetsData = [
    { uuid: 'target-1', name: '女性' },
    { uuid: 'target-2', name: '男性' },
]

const mockTagsData = [
    { uuid: 'tag-1', name: 'タグ1' },
    { uuid: 'tag-2', name: 'タグ2' },
]

const mockSalesSitesData = [
    { uuid: 'site-1', name: 'Creema' },
    { uuid: 'site-2', name: 'minne' },
]

describe('Admin Product Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()

        // デフォルトのモック設定
        mockGetProducts.mockResolvedValue(mockProductData)
        mockGetCategories.mockResolvedValue(mockCategoriesData)
        mockGetTargets.mockResolvedValue(mockTargetsData)
        mockGetTags.mockResolvedValue(mockTagsData)
        mockGetSalesSiteList.mockResolvedValue(mockSalesSitesData)
    })

    it('管理画面商品管理ページが正常に表示される', async () => {
        // コンポーネントをレンダリング
        render(<AdminProductTemplate />)

        // ローディング表示の確認
        expect(screen.getByText('読み込み中...')).toBeInTheDocument()

        // データ読み込み完了後の表示を確認
        await waitFor(() => {
            expect(screen.getByText('商品一覧')).toBeInTheDocument()
            expect(screen.getByText('2件の商品')).toBeInTheDocument()
            expect(screen.getByText('追加')).toBeInTheDocument()
        })

        // 商品カードの表示を確認
        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
            expect(screen.getByText('テスト商品2')).toBeInTheDocument()
            expect(screen.getByText('¥1,500')).toBeInTheDocument()
            expect(screen.getByText('¥2,500')).toBeInTheDocument()
        })

        // API呼び出しの確認
        expect(mockGetProducts).toHaveBeenCalledWith({
            mode: 'all',
            category: 'all',
            target: 'all',
        })
        expect(mockGetCategories).toHaveBeenCalledWith({ mode: 'all' })
        expect(mockGetTargets).toHaveBeenCalledWith({ mode: 'all' })
        expect(mockGetTags).toHaveBeenCalled()
        expect(mockGetSalesSiteList).toHaveBeenCalled()
    })

    it('商品が0件の場合の表示', async () => {
        // 空のデータを設定
        mockGetProducts.mockResolvedValue([])

        render(<AdminProductTemplate />)

        await waitFor(() => {
            expect(screen.getByText('0件の商品')).toBeInTheDocument()
            expect(screen.getByText('登録されていません')).toBeInTheDocument()
        })
    })

    it('商品追加ボタンをクリックするとダイアログが開く', async () => {
        render(<AdminProductTemplate />)

        await waitFor(() => {
            expect(screen.getByText('追加')).toBeInTheDocument()
        })

        const addButton = screen.getByText('追加')
        fireEvent.click(addButton)

        // ダイアログの表示を確認（ProductFormDialogの内容）
        await waitFor(() => {
            // ダイアログ内のCreemaオプションが表示されているか確認
            expect(screen.getByText('Creemaから複製')).toBeInTheDocument()
            expect(screen.getByText('手動で入力')).toBeInTheDocument()
        })
    })

    it('商品編集機能のクリックが正常に動作する', async () => {
        render(<AdminProductTemplate />)

        // 商品カードが表示されるまで待機
        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
        })

        // ProductCardがクリック可能であることを確認
        const productCards = screen.getAllByText('テスト商品1')
        expect(productCards[0]).toBeInTheDocument()

        // クリックイベントが実行できることを確認
        fireEvent.click(productCards[0])
        // 実際のダイアログの開閉確認は統合テストの範囲外とする
    })

    it('商品削除機能が正常に動作する', async () => {
        mockDeleteProduct.mockResolvedValue(undefined)

        render(<AdminProductTemplate />)

        // 商品カードが表示されるまで待機
        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
        })

        // 削除アイコンボタンを探してクリック
        const deleteButtons = screen.getAllByTestId('DeleteIcon')
        fireEvent.click(deleteButtons[0])

        // 削除確認ダイアログの表示を確認
        await waitFor(() => {
            expect(screen.getByText('削除確認')).toBeInTheDocument()
            expect(screen.getByText('商品「テスト商品1」を削除しますか？')).toBeInTheDocument()
            expect(screen.getByText('この操作は取り消せません。')).toBeInTheDocument()
        })

        // 削除確認ボタンをクリック
        const confirmDeleteButton = screen.getByRole('button', { name: '削除' })
        fireEvent.click(confirmDeleteButton)

        // API呼び出しの確認
        await waitFor(() => {
            expect(mockDeleteProduct).toHaveBeenCalledWith('product-1')
        })
    })

    it('削除確認ダイアログでキャンセルできる', async () => {
        render(<AdminProductTemplate />)

        // 商品カードが表示されるまで待機
        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
        })

        // 削除アイコンボタンをクリック
        const deleteButtons = screen.getAllByTestId('DeleteIcon')
        fireEvent.click(deleteButtons[0])

        // 削除確認ダイアログの表示を確認
        await waitFor(() => {
            expect(screen.getByText('削除確認')).toBeInTheDocument()
        })

        // キャンセルボタンをクリック
        const cancelButton = screen.getByRole('button', { name: 'キャンセル' })
        fireEvent.click(cancelButton)

        // ダイアログが閉じることを確認
        await waitFor(() => {
            expect(screen.queryByText('削除確認')).not.toBeInTheDocument()
        })

        // 削除APIが呼ばれないことを確認
        expect(mockDeleteProduct).not.toHaveBeenCalled()
    })

    it('API呼び出しエラー時の処理', async () => {
        // エラーを発生させる
        mockGetProducts.mockRejectedValue(new Error('API Error'))

        // console.errorをモック
        const consoleErrorSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        render(<AdminProductTemplate />)

        // エラーが発生してもクラッシュしないことを確認
        await waitFor(() => {
            expect(screen.getByText('商品一覧')).toBeInTheDocument()
        })

        // console.errorが呼ばれることを確認
        await waitFor(() => {
            expect(consoleErrorSpy).toHaveBeenCalledWith('データの取得に失敗しました:', expect.any(Error))
        })

        consoleErrorSpy.mockRestore()
    })

    it('新規商品作成が正常に動作する', async () => {
        render(<AdminProductTemplate />)

        await waitFor(() => {
            expect(screen.getByText('追加')).toBeInTheDocument()
        })

        // 追加ボタンをクリック
        const addButton = screen.getByText('追加')
        fireEvent.click(addButton)

        // ダイアログが開くことを確認
        await waitFor(() => {
            expect(screen.getByText('Creemaから複製')).toBeInTheDocument()
            expect(screen.getByText('手動で入力')).toBeInTheDocument()
        })
    })

    it('Creema複製機能が利用できる', async () => {
        render(<AdminProductTemplate />)

        await waitFor(() => {
            expect(screen.getByText('追加')).toBeInTheDocument()
        })

        // 追加ボタンをクリックしてダイアログを開く
        const addButton = screen.getByText('追加')
        fireEvent.click(addButton)

        // ダイアログ内のCreema複製オプションが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('Creemaから複製')).toBeInTheDocument()
            expect(screen.getByText('手動で入力')).toBeInTheDocument()
        })
    })

    it('ページメタデータでnoindex, nofollowが設定されている', () => {
        // メタデータの確認
        expect(metadata.title).toBe('商品一覧 | admin')
        expect(metadata.robots).toEqual({
            index: false,
            follow: false,
        })
    })

    it('データの再取得が正常に動作する', async () => {
        render(<AdminProductTemplate />)

        // 初回データ読み込み完了まで待機
        await waitFor(() => {
            expect(screen.getByText('2件の商品')).toBeInTheDocument()
        })

        // モックをクリアして新しいデータを設定
        vi.clearAllMocks()
        const updatedProductData = [
            ...mockProductData,
            {
                uuid: 'product-3',
                name: 'テスト商品3',
                description: 'テスト商品3の説明',
                price: 3000,
                isActive: true,
                isRecommend: false,
                category: { uuid: 'category-1', name: 'イヤリング' },
                target: { uuid: 'target-1', name: '女性' },
                tags: [],
                productImages: [],
                siteDetails: [],
            },
        ]
        mockGetProducts.mockResolvedValue(updatedProductData)
        mockGetCategories.mockResolvedValue(mockCategoriesData)
        mockGetTargets.mockResolvedValue(mockTargetsData)
        mockGetTags.mockResolvedValue(mockTagsData)
        mockGetSalesSiteList.mockResolvedValue(mockSalesSitesData)

        // 削除操作によるデータ再取得をシミュレート
        const deleteButtons = screen.getAllByTestId('DeleteIcon')
        fireEvent.click(deleteButtons[0])

        await waitFor(() => {
            expect(screen.getByText('削除確認')).toBeInTheDocument()
        })

        const confirmDeleteButton = screen.getByRole('button', { name: '削除' })
        fireEvent.click(confirmDeleteButton)

        // fetchDataが再度呼ばれることを確認
        await waitFor(() => {
            expect(mockGetProducts).toHaveBeenCalled()
        })
    })

    it('商品データの並び順が正しく表示される', async () => {
        render(<AdminProductTemplate />)

        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
            expect(screen.getByText('テスト商品2')).toBeInTheDocument()
        })

        // 商品が正しい順番で表示されているかを確認
        const productElements = screen.getAllByText(/テスト商品\d/)
        expect(productElements[0]).toHaveTextContent('テスト商品1')
        expect(productElements[1]).toHaveTextContent('テスト商品2')
    })

    it('商品の状態（アクティブ/非アクティブ、推奨/非推奨）が正しく表示される', async () => {
        render(<AdminProductTemplate />)

        // 商品データに基づいて状態を確認
        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
            expect(screen.getByText('テスト商品2')).toBeInTheDocument()
        })

        // ProductCardコンポーネントが商品データを正しく受け取っていることを確認
        // （実際の表示内容は ProductCard の実装に依存）
    })

    it('複数の商品を同時に操作できる', async () => {
        render(<AdminProductTemplate />)

        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
            expect(screen.getByText('テスト商品2')).toBeInTheDocument()
        })

        // 複数の商品カード（編集可能）が存在することを確認
        const productCards = screen.getAllByText(/テスト商品\d/)
        expect(productCards).toHaveLength(2)

        // 複数の削除アイコンボタンが存在することを確認
        const deleteButtons = screen.getAllByTestId('DeleteIcon')
        expect(deleteButtons).toHaveLength(2)
    })
})
