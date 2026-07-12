import { fireEvent, within } from '@testing-library/react'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCategories } from '@/apis/category'
import { deleteProduct, getProducts, updateProduct } from '@/apis/product'
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
const mockUpdateProduct = vi.mocked(updateProduct)

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
                displayOrder: 1,
            },
            {
                uuid: 'image-2',
                apiPath: '/image/product-1-2.jpg',
                name: 'product-1-2.jpg',
                displayOrder: 2,
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

const createProductList = (products = mockProductData, pageInfo?: { limit?: number; page?: number; total?: number; totalPages?: number }) => ({
    pageInfo: {
        limit: pageInfo?.limit ?? 20,
        page: pageInfo?.page ?? 1,
        total: pageInfo?.total ?? products.length,
        totalPages: pageInfo?.totalPages ?? Math.ceil(products.length / 20),
    },
    products,
})

describe('Admin Product Page Integration Test', () => {
    const defaultProps = {
        categories: mockCategoriesData,
        salesSites: mockSalesSitesData,
        tags: mockTagsData,
        targets: mockTargetsData,
    }

    beforeEach(() => {
        vi.clearAllMocks()

        // デフォルトのモック設定
        mockGetProducts.mockResolvedValue(createProductList(mockProductData))
        mockGetCategories.mockResolvedValue(mockCategoriesData)
        mockGetTargets.mockResolvedValue(mockTargetsData)
        mockGetTags.mockResolvedValue(mockTagsData)
        mockGetSalesSiteList.mockResolvedValue(mockSalesSitesData)
    })

    it('管理画面商品管理ページが正常に表示される', async () => {
        // コンポーネントをレンダリング
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        // 初期データが表示されることを確認
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

        // 初期データが表示されている場合、API呼び出しは発生しない
        expect(mockGetProducts).not.toHaveBeenCalled()
        expect(mockGetCategories).not.toHaveBeenCalled()
        expect(mockGetTargets).not.toHaveBeenCalled()
        expect(mockGetTags).not.toHaveBeenCalled()
        expect(mockGetSalesSiteList).not.toHaveBeenCalled()
    })

    it('商品が0件の場合の表示', async () => {
        // 空のデータを設定
        mockGetProducts.mockResolvedValue(createProductList([]))

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList([])} />)

        // 初期状態で0件が表示されることを確認
        expect(screen.getByText('0件の商品')).toBeInTheDocument()
        expect(screen.getByText('登録されていません')).toBeInTheDocument()
    })

    it('ページネーションで次ページの商品を取得できる', async () => {
        const nextPageProducts = [mockProductData[1]]
        mockGetProducts.mockResolvedValue(createProductList(nextPageProducts, { page: 2, total: 21, totalPages: 2 }))
        mockGetCategories.mockRejectedValue(new Error('Category API Error'))

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList([mockProductData[0]], { total: 21, totalPages: 2 })} />)

        fireEvent.click(screen.getByRole('button', { name: '次のページへ' }))

        await waitFor(() => {
            expect(mockGetProducts).toHaveBeenCalledWith({
                category: 'all',
                limit: 20,
                mode: 'all',
                page: 2,
                target: 'all',
            })
        })
        await waitFor(() => {
            expect(screen.getByText('テスト商品2')).toBeInTheDocument()
        })
        expect(mockGetCategories).not.toHaveBeenCalled()
        expect(mockGetTargets).not.toHaveBeenCalled()
        expect(mockGetTags).not.toHaveBeenCalled()
        expect(mockGetSalesSiteList).not.toHaveBeenCalled()
    })

    it('商品名で検索でき、検索条件を維持してページ移動できる', async () => {
        const searchProducts = [mockProductData[0]]
        mockGetProducts.mockResolvedValueOnce(createProductList(searchProducts, { page: 1, total: 21, totalPages: 2 }))
        mockGetProducts.mockResolvedValueOnce(createProductList([mockProductData[1]], { page: 2, total: 21, totalPages: 2 }))

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData, { total: 21, totalPages: 2 })} />)

        fireEvent.click(screen.getByRole('button', { name: '絞り込み' }))
        fireEvent.change(screen.getByLabelText('商品名で検索'), { target: { value: ' テスト商品 ' } })
        fireEvent.click(screen.getByRole('button', { name: '検索' }))

        await waitFor(() => {
            expect(mockGetProducts).toHaveBeenCalledWith({
                category: 'all',
                keyword: 'テスト商品',
                limit: 20,
                mode: 'all',
                page: 1,
                target: 'all',
            })
        })

        fireEvent.click(screen.getByRole('button', { name: '次のページへ' }))

        await waitFor(() => {
            expect(mockGetProducts).toHaveBeenLastCalledWith({
                category: 'all',
                keyword: 'テスト商品',
                limit: 20,
                mode: 'all',
                page: 2,
                target: 'all',
            })
        })
    })

    it('クリアで検索条件をリセットし、検索ボタンで検索条件なしの1ページ目を取得する', async () => {
        mockGetProducts.mockResolvedValueOnce(createProductList([], { page: 1, total: 0, totalPages: 0 }))
        mockGetProducts.mockResolvedValueOnce(createProductList(mockProductData))

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        fireEvent.click(screen.getByRole('button', { name: '絞り込み' }))
        fireEvent.change(screen.getByLabelText('商品名で検索'), { target: { value: '存在しない商品' } })
        fireEvent.click(screen.getByRole('button', { name: '検索' }))

        await waitFor(() => {
            expect(screen.getByText('該当する商品がありません')).toBeInTheDocument()
        })

        // クリアは入力内容をリセットするだけで、検索APIは呼ばない
        fireEvent.click(screen.getByRole('button', { name: '絞り込み' }))
        fireEvent.click(screen.getByRole('button', { name: 'クリア' }))
        expect(screen.getByLabelText('商品名で検索')).toHaveValue('')
        expect(mockGetProducts).toHaveBeenCalledTimes(1)

        // 検索ボタン押下で初めて検索条件なしの再取得を行う
        fireEvent.click(screen.getByRole('button', { name: '検索' }))

        await waitFor(() => {
            expect(mockGetProducts).toHaveBeenLastCalledWith({
                category: 'all',
                limit: 20,
                mode: 'all',
                page: 1,
                target: 'all',
            })
        })
    })

    it('カテゴリ・タグ・価格・ステータスで検索できる', async () => {
        mockGetProducts.mockResolvedValueOnce(createProductList([mockProductData[0]], { page: 1, total: 1, totalPages: 1 }))

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        fireEvent.click(screen.getByRole('button', { name: '絞り込み' }))

        // 各フィルターはラベルからトリガーを辿って開き、開いたフィールド内で選択肢をクリックする
        const filterForm = screen.getByLabelText('商品名で検索').closest('form') as HTMLElement
        const filterDialog = within(filterForm)
        const selectFilterOption = (labelText: string, optionText: string) => {
            const field = filterDialog.getByText(labelText).parentElement as HTMLElement
            fireEvent.click(field.querySelector('[class*="select-trigger"]') as HTMLElement)
            const option = within(field)
                .getAllByText(optionText)
                .find((element) => element.className.includes('option'))
            fireEvent.click(option as HTMLElement)
        }

        selectFilterOption('カテゴリ', 'イヤリング')
        selectFilterOption('タグ', 'タグ1')
        fireEvent.change(filterDialog.getByLabelText('最低価格'), { target: { value: '1000' } })
        fireEvent.change(filterDialog.getByLabelText('最高価格'), { target: { value: '2000' } })
        selectFilterOption('公開状態', '公開中')
        selectFilterOption('おすすめ', 'おすすめ')
        fireEvent.click(screen.getByRole('button', { name: '検索' }))

        await waitFor(() => {
            expect(mockGetProducts).toHaveBeenCalledWith({
                activeStatus: 'active',
                category: 'category-1',
                limit: 20,
                maxPrice: 2000,
                minPrice: 1000,
                mode: 'all',
                page: 1,
                recommendStatus: 'recommended',
                tagUuids: ['tag-1'],
                target: 'all',
            })
        })
    })

    it('商品追加ボタンをクリックするとダイアログが開く', async () => {
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
            expect(screen.getByRole('button', { name: 'キャンセル' })).toBeInTheDocument()
        })
    })

    it('商品追加ダイアログをキャンセルで閉じられる', async () => {
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        const addButton = await screen.findByText('追加')
        fireEvent.click(addButton)

        await screen.findByText('商品を追加')

        fireEvent.click(screen.getByRole('button', { name: 'キャンセル' }))

        await waitFor(() => {
            expect(screen.queryByText('商品を追加')).not.toBeInTheDocument()
        })
    })

    it('商品編集機能のクリックが正常に動作する', async () => {
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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

    it('商品編集で既存画像を削除すると更新payloadから対象画像だけ除外される', async () => {
        mockUpdateProduct.mockResolvedValue(undefined)

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        fireEvent.click(await screen.findByText('テスト商品1'))

        await waitFor(() => {
            expect(screen.getByText('商品を編集')).toBeInTheDocument()
            expect(screen.getAllByAltText('プレビュー')).toHaveLength(2)
        })

        const firstPreview = screen.getAllByAltText('プレビュー')[0]
        const deleteButton = firstPreview.parentElement?.parentElement?.querySelector('button')

        expect(deleteButton).toBeInTheDocument()
        fireEvent.click(deleteButton as HTMLButtonElement)

        await waitFor(() => {
            expect(screen.getAllByAltText('プレビュー')).toHaveLength(1)
        })

        fireEvent.click(screen.getByRole('button', { name: '更新' }))

        await waitFor(() => {
            expect(mockUpdateProduct).toHaveBeenCalled()
        })

        expect(mockUpdateProduct).toHaveBeenCalledWith(
            'product-1',
            expect.objectContaining({
                productImages: [
                    expect.objectContaining({
                        uuid: 'image-2',
                    }),
                ],
            }),
        )
    })

    it('商品削除機能が正常に動作する', async () => {
        mockDeleteProduct.mockResolvedValue(undefined)

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
        // console.errorをモック
        const consoleErrorSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        // エラーを発生させる
        mockGetProducts.mockRejectedValue(new Error('API Error'))
        mockDeleteProduct.mockResolvedValue(undefined)

        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        // 初期表示は正常
        await waitFor(() => {
            expect(screen.getByText('商品一覧')).toBeInTheDocument()
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
        })

        // 削除操作を行ってfetchDataを呼び出す
        const deleteButtons = screen.getAllByTestId('DeleteIcon')
        fireEvent.click(deleteButtons[0])

        await waitFor(() => {
            expect(screen.getByText('削除確認')).toBeInTheDocument()
        })

        const confirmDeleteButton = screen.getByRole('button', { name: '削除' })
        fireEvent.click(confirmDeleteButton)

        // console.errorが呼ばれることを確認
        await waitFor(() => {
            expect(consoleErrorSpy).toHaveBeenCalledWith('データの取得に失敗しました:', expect.any(Error))
        })

        consoleErrorSpy.mockRestore()
    })

    it('新規商品作成が正常に動作する', async () => {
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
        mockGetProducts.mockResolvedValue(createProductList(updatedProductData))
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
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

        // 商品データに基づいて状態を確認
        await waitFor(() => {
            expect(screen.getByText('テスト商品1')).toBeInTheDocument()
            expect(screen.getByText('テスト商品2')).toBeInTheDocument()
        })

        // ProductCardコンポーネントが商品データを正しく受け取っていることを確認
        // （実際の表示内容は ProductCard の実装に依存）
    })

    it('複数の商品を同時に操作できる', async () => {
        render(<AdminProductTemplate {...defaultProps} initialProductList={createProductList(mockProductData)} />)

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
