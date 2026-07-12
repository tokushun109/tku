'use client'

import { Add, Close, Search } from '@mui/icons-material'
import { type FormEvent, useState } from 'react'
import { toast } from 'sonner'

import { getCategories } from '@/apis/category'
import {
    ADMIN_PRODUCT_PAGE_LIMIT,
    createProduct,
    deleteProduct,
    duplicateProductFromCreema,
    getProducts,
    type IGetProductsParams,
    updateProduct,
    uploadProductImages,
} from '@/apis/product'
import { getSalesSiteList } from '@/apis/salesSite'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { Input } from '@/components/bases/Input'
import { MultiSelectForm } from '@/components/bases/MultiSelectForm'
import { Pagination } from '@/components/bases/Pagination'
import { SelectForm } from '@/components/bases/SelectForm'
import { IClassification } from '@/features/classification/type'
import { ProductCard } from '@/features/product/components/ProductCard'
import { ProductFormDialog } from '@/features/product/components/ProductFormDialog'
import { EXISTING_PRODUCT_IMAGE_ID_PREFIX } from '@/features/product/constants'
import { ICreemaDuplicateForm } from '@/features/product/product/type'
import { IProduct, IProductForm, IProductList } from '@/features/product/type'
import { ISite } from '@/features/site/type'

import styles from './styles.module.scss'

interface Props {
    categories: IClassification[]
    initialProductList: IProductList
    salesSites: ISite[]
    tags: IClassification[]
    targets: IClassification[]
}

const ProductActiveStatus = {
    All: 'all',
    Active: 'active',
    Inactive: 'inactive',
} as const

const ProductRecommendStatus = {
    All: 'all',
    Recommended: 'recommended',
    NotRecommended: 'not_recommended',
} as const

type ProductActiveStatus = (typeof ProductActiveStatus)[keyof typeof ProductActiveStatus]
type ProductRecommendStatus = (typeof ProductRecommendStatus)[keyof typeof ProductRecommendStatus]

interface ProductSearchFilters {
    activeStatus: ProductActiveStatus
    category: string
    maxPrice: string
    minPrice: string
    recommendStatus: ProductRecommendStatus
    tagUuids: string[]
}

const defaultSearchFilters: ProductSearchFilters = {
    activeStatus: ProductActiveStatus.All,
    category: 'all',
    maxPrice: '',
    minPrice: '',
    recommendStatus: ProductRecommendStatus.All,
    tagUuids: [],
}

export const AdminProductTemplate = ({
    categories: initialCategories,
    initialProductList,
    salesSites: initialSalesSites,
    tags: initialTags,
    targets: initialTargets,
}: Props) => {
    const [products, setProducts] = useState<IProduct[]>(initialProductList.products)
    const [pageInfo, setPageInfo] = useState(initialProductList.pageInfo)
    const [categories, setCategories] = useState<IClassification[]>(initialCategories)
    const [targets, setTargets] = useState<IClassification[]>(initialTargets)
    const [tags, setTags] = useState<IClassification[]>(initialTags)
    const [salesSites, setSalesSites] = useState<ISite[]>(initialSalesSites)
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [submitError, setSubmitError] = useState<string | null>(null)
    const [updateItem, setUpdateItem] = useState<IProduct | null>(null)
    const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState<boolean>(false)
    const [deleteTargetItem, setDeleteTargetItem] = useState<IProduct | null>(null)
    const [searchText, setSearchText] = useState<string>('')
    const [keyword, setKeyword] = useState<string>('')
    const [searchFilters, setSearchFilters] = useState<ProductSearchFilters>(defaultSearchFilters)
    const [appliedFilters, setAppliedFilters] = useState<ProductSearchFilters>(defaultSearchFilters)

    const buildProductListParams = (page: number, nextKeyword: string, nextFilters: ProductSearchFilters): IGetProductsParams => {
        const params: IGetProductsParams = {
            mode: 'all',
            category: nextFilters.category,
            limit: ADMIN_PRODUCT_PAGE_LIMIT,
            page,
            target: 'all',
        }

        if (nextKeyword !== '') params.keyword = nextKeyword
        if (nextFilters.activeStatus !== ProductActiveStatus.All) params.activeStatus = nextFilters.activeStatus
        if (nextFilters.recommendStatus !== ProductRecommendStatus.All) params.recommendStatus = nextFilters.recommendStatus
        if (nextFilters.minPrice !== '') params.minPrice = Number(nextFilters.minPrice)
        if (nextFilters.maxPrice !== '') params.maxPrice = Number(nextFilters.maxPrice)
        if (nextFilters.tagUuids.length > 0) params.tagUuids = nextFilters.tagUuids

        return params
    }

    const fetchProducts = async (page: number, nextKeyword: string = keyword, nextFilters: ProductSearchFilters = appliedFilters) => {
        try {
            setIsLoading(true)
            const productList = await getProducts(buildProductListParams(page, nextKeyword, nextFilters))

            setProducts(productList.products)
            setPageInfo(productList.pageInfo)
        } catch (error) {
            console.error('商品リストの取得に失敗しました:', error)
        } finally {
            setIsLoading(false)
        }
    }

    const fetchData = async (page: number = pageInfo.page, nextKeyword: string = keyword, nextFilters: ProductSearchFilters = appliedFilters) => {
        try {
            setIsLoading(true)
            const [productList, categoriesData, targetsData, tagsData, salesSitesData] = await Promise.all([
                getProducts(buildProductListParams(page, nextKeyword, nextFilters)),
                getCategories({ mode: 'all' }),
                getTargets({ mode: 'all' }),
                getTags(),
                getSalesSiteList(),
            ])
            setProducts(productList.products)
            setPageInfo(productList.pageInfo)
            setCategories(categoriesData)
            setTargets(targetsData)
            setTags(tagsData)
            setSalesSites(salesSitesData)
        } catch (error) {
            console.error('データの取得に失敗しました:', error)
        } finally {
            setIsLoading(false)
        }
    }

    const handleCreate = () => {
        setUpdateItem(null)
        setSubmitError(null)
        setIsDialogOpen(true)
    }

    const handleEdit = (product: IProduct) => {
        setUpdateItem(product)
        setSubmitError(null)
        setIsDialogOpen(true)
    }

    const handleDelete = (product: IProduct) => {
        setDeleteTargetItem(product)
        setIsDeleteDialogOpen(true)
    }

    const handleCloseDeleteDialog = () => {
        setIsDeleteDialogOpen(false)
        setDeleteTargetItem(null)
    }

    const handleConfirmDelete = async () => {
        if (!deleteTargetItem) return

        try {
            await deleteProduct(deleteTargetItem.uuid)
            toast.success(`商品「${deleteTargetItem.name}」を削除しました`)
            const nextPage = products.length === 1 && pageInfo.page > 1 ? pageInfo.page - 1 : pageInfo.page
            await fetchData(nextPage, keyword, appliedFilters)
            handleCloseDeleteDialog()
        } catch (error) {
            console.error('商品の削除に失敗しました:', error)
            toast.error('商品の削除に失敗しました。もう一度お試しください。')
        }
    }

    const handleCloseDialog = () => {
        setIsDialogOpen(false)
        setUpdateItem(null)
        setSubmitError(null)
    }

    const handleSubmit = async (data: IProductForm) => {
        try {
            setIsSubmitting(true)
            setSubmitError(null)

            // フォームデータを商品データに変換
            const productData: Omit<IProduct, 'uuid'> = {
                name: data.name,
                description: data.description || '',
                price: data.price,
                isActive: data.isActive,
                isRecommend: data.isRecommend,
                category: data.categoryUuid
                    ? ({ uuid: data.categoryUuid, name: '' } as IClassification)
                    : ({ uuid: '', name: '' } as IClassification),
                target: data.targetUuid ? ({ uuid: data.targetUuid, name: '' } as IClassification) : ({ uuid: '', name: '' } as IClassification),
                tags: data.tagUuids ? data.tagUuids.map((uuid) => ({ uuid, name: '' }) as IClassification) : [],
                productImages: updateItem?.productImages || [],
                siteDetails: data.siteDetails
                    ? data.siteDetails.map((detail) => ({
                          uuid: '',
                          detailUrl: detail.detailUrl,
                          salesSite: { uuid: detail.salesSiteUuid, name: '' },
                      }))
                    : [],
            }

            let productUuid: string = ''

            if (updateItem) {
                // 編集時：まず基本データを更新、その後画像順序更新（必要な場合）
                let productDataToUpdate = { ...productData, uuid: updateItem.uuid }

                if (data.imageItems) {
                    // モーダル上で削除・並び替え後に残っている既存画像だけを更新対象にする。
                    // PUTに含めなかった既存画像はバックエンド側で削除される。
                    const existingItems = data.imageItems.filter((item) => !item.isNewUpload && item.id.startsWith(EXISTING_PRODUCT_IMAGE_ID_PREFIX))
                    const existingItemMap = new Map(existingItems.map((item) => [item.id.replace(EXISTING_PRODUCT_IMAGE_ID_PREFIX, ''), item]))

                    productDataToUpdate.productImages = updateItem.productImages
                        .filter((image) => existingItemMap.has(image.uuid))
                        .map((image) => {
                            const imageItem = existingItemMap.get(image.uuid)
                            if (data.isImageOrderChanged && imageItem?.displayOrder) {
                                return {
                                    ...image,
                                    displayOrder: 100 - (imageItem.displayOrder - 1),
                                }
                            }
                            return image
                        })
                }

                await updateProduct(updateItem.uuid, productDataToUpdate)
                productUuid = updateItem.uuid
            } else {
                // 新規作成
                const result = await createProduct(productData)
                productUuid = result.uuid
            }

            // 画像をアップロードする場合の処理
            if (data.uploadImages && data.uploadImages.length > 0) {
                const existingImagesCount = data.imageItems
                    ? data.imageItems.filter((item) => !item.isNewUpload && item.id.startsWith(EXISTING_PRODUCT_IMAGE_ID_PREFIX)).length
                    : updateItem?.productImages?.length || 0
                const hasOrderChanged = data.isImageOrderChanged || false

                // 新規画像の優先順位を計算
                let uploadFileDisplayOrder: { [key: number]: number } = {}

                if (hasOrderChanged && data.imageItems) {
                    // 並び替えが行われた場合、全体の順序から新規画像の順序を計算
                    const newUploadItems = data.imageItems.filter((item) => item.isNewUpload)

                    // data.uploadImagesの順序に合わせてuploadFileOrderを作成
                    data.uploadImages.forEach((file, uploadIndex) => {
                        // ファイル名でマッチングして対応するimageItemを見つける
                        const matchingItem = newUploadItems.find((item) => {
                            // ObjectURLの場合はファイル名で比較できないため、順序で推定
                            const expectedIndex = newUploadItems.indexOf(item)
                            return expectedIndex === uploadIndex
                        })

                        if (matchingItem && matchingItem.displayOrder) {
                            // 並び替え後の位置（1から始まる）を100から降順に変換
                            uploadFileDisplayOrder[uploadIndex] = 100 - (matchingItem.displayOrder - 1)
                        } else {
                            // フォールバック: 既存画像数を考慮した順序
                            uploadFileDisplayOrder[uploadIndex] = 100 - existingImagesCount - uploadIndex
                        }
                    })
                } else {
                    // 通常の場合、既存画像より低い優先順位を設定
                    data.uploadImages.forEach((_, index) => {
                        uploadFileDisplayOrder[index] = 100 - existingImagesCount - index
                    })
                }

                const displayOrderParams = {
                    isChanged: hasOrderChanged,
                    displayOrder: uploadFileDisplayOrder,
                }

                await uploadProductImages(productUuid, data.uploadImages, displayOrderParams)
            }

            setIsDialogOpen(false)
            setUpdateItem(null)

            // 成功通知
            if (updateItem) {
                toast.success(`商品「${data.name}」を更新しました`)
            } else {
                toast.success(`商品「${data.name}」を追加しました`)
            }

            await fetchData(updateItem ? pageInfo.page : 1, keyword, appliedFilters)
        } catch (error) {
            console.error('商品の保存に失敗しました:', error)
            const errorMessage = '商品の保存に失敗しました。もう一度お試しください。'
            setSubmitError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setIsSubmitting(false)
        }
    }

    const handleCreemaDuplicate = async (data: ICreemaDuplicateForm) => {
        try {
            setIsSubmitting(true)
            setSubmitError(null)

            await duplicateProductFromCreema({ url: data.creemaUrl })

            setIsDialogOpen(false)
            toast.success('Creemaから商品を複製しました')
            await fetchData(1, keyword, appliedFilters)
        } catch (error) {
            console.error('Creemaからの商品複製に失敗しました:', error)
            const errorMessage = 'Creemaからの商品複製に失敗しました。もう一度お試しください。'
            setSubmitError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setIsSubmitting(false)
        }
    }

    const handlePageChange = async (page: number) => {
        await fetchProducts(page, keyword, appliedFilters)
    }

    const handleSearchSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const nextKeyword = searchText.trim()
        const nextFilters = { ...searchFilters }
        setKeyword(nextKeyword)
        setAppliedFilters(nextFilters)
        await fetchData(1, nextKeyword, nextFilters)
    }

    const handleClearSearch = async () => {
        setSearchText('')
        setKeyword('')
        setSearchFilters(defaultSearchFilters)
        setAppliedFilters(defaultSearchFilters)
        await fetchData(1, '', defaultSearchFilters)
    }

    const handleFilterChange = (nextFilters: Partial<ProductSearchFilters>) => {
        setSearchFilters((currentFilters) => ({
            ...currentFilters,
            ...nextFilters,
        }))
    }

    const hasAppliedSearch =
        keyword !== '' ||
        appliedFilters.category !== 'all' ||
        appliedFilters.tagUuids.length > 0 ||
        appliedFilters.minPrice !== '' ||
        appliedFilters.maxPrice !== '' ||
        appliedFilters.activeStatus !== ProductActiveStatus.All ||
        appliedFilters.recommendStatus !== ProductRecommendStatus.All
    const hasDraftSearch =
        searchText !== '' ||
        searchFilters.category !== 'all' ||
        searchFilters.tagUuids.length > 0 ||
        searchFilters.minPrice !== '' ||
        searchFilters.maxPrice !== '' ||
        searchFilters.activeStatus !== ProductActiveStatus.All ||
        searchFilters.recommendStatus !== ProductRecommendStatus.All

    const categoryOptions = [
        { label: 'すべてのカテゴリ', value: 'all' },
        ...categories.map((category) => ({ label: category.name, value: category.uuid })),
    ]
    const tagOptions = tags.map((tag) => ({ label: tag.name, value: tag.uuid }))
    const activeStatusOptions = [
        { label: '公開状態すべて', value: ProductActiveStatus.All },
        { label: '公開中', value: ProductActiveStatus.Active },
        { label: '非公開', value: ProductActiveStatus.Inactive },
    ]
    const recommendStatusOptions = [
        { label: 'おすすめ状態すべて', value: ProductRecommendStatus.All },
        { label: 'おすすめ', value: ProductRecommendStatus.Recommended },
        { label: 'おすすめ以外', value: ProductRecommendStatus.NotRecommended },
    ]

    const emptyMessage = hasAppliedSearch ? '該当する商品がありません' : '登録されていません'
    const isClearDisabled = isLoading || (!hasAppliedSearch && !hasDraftSearch)

    return (
        <div className={styles['product-container']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>商品一覧</h1>
                <div className={styles['header-actions']}>
                    <div className={styles['product-count']}>{pageInfo.total}件の商品</div>
                    <Button onClick={handleCreate}>
                        <div className={styles['add-button-content']}>
                            <Add className={styles['add-icon']} fontSize="small" />
                            追加
                        </div>
                    </Button>
                </div>
            </div>
            <form className={styles['search-form']} onSubmit={handleSearchSubmit}>
                <div className={styles['search-grid']}>
                    <Input
                        aria-label="商品名で検索"
                        className={styles['search-input']}
                        onChange={(event) => {
                            setSearchText(event.target.value)
                        }}
                        placeholder="商品名で検索"
                        value={searchText}
                    />
                    <SelectForm
                        id="admin-product-category-filter"
                        onChange={(value) => {
                            handleFilterChange({ category: value ?? 'all' })
                        }}
                        options={categoryOptions}
                        placeholder="カテゴリ"
                        value={searchFilters.category}
                    />
                    <MultiSelectForm
                        id="admin-product-tag-filter"
                        onChange={(value) => {
                            handleFilterChange({ tagUuids: value })
                        }}
                        options={tagOptions}
                        placeholder="タグ"
                        value={searchFilters.tagUuids}
                    />
                    <Input
                        aria-label="最低価格"
                        inputMode="numeric"
                        min={0}
                        onChange={(event) => {
                            handleFilterChange({ minPrice: event.target.value })
                        }}
                        placeholder="最低価格"
                        type="number"
                        value={searchFilters.minPrice}
                    />
                    <Input
                        aria-label="最高価格"
                        inputMode="numeric"
                        min={0}
                        onChange={(event) => {
                            handleFilterChange({ maxPrice: event.target.value })
                        }}
                        placeholder="最高価格"
                        type="number"
                        value={searchFilters.maxPrice}
                    />
                    <SelectForm
                        id="admin-product-active-status-filter"
                        onChange={(value) => {
                            handleFilterChange({ activeStatus: value ?? ProductActiveStatus.All })
                        }}
                        options={activeStatusOptions}
                        placeholder="公開状態"
                        value={searchFilters.activeStatus}
                    />
                    <SelectForm
                        id="admin-product-recommend-status-filter"
                        onChange={(value) => {
                            handleFilterChange({ recommendStatus: value ?? ProductRecommendStatus.All })
                        }}
                        options={recommendStatusOptions}
                        placeholder="おすすめ状態"
                        value={searchFilters.recommendStatus}
                    />
                </div>
                <div className={styles['search-actions']}>
                    <Button className={styles['search-button']} disabled={isLoading} type="submit">
                        <div className={styles['search-button-content']}>
                            <Search className={styles['search-icon']} fontSize="small" />
                            検索
                        </div>
                    </Button>
                    <Button className={styles['clear-button']} contrast disabled={isClearDisabled} onClick={handleClearSearch} type="button">
                        <div className={styles['search-button-content']}>
                            <Close className={styles['search-icon']} fontSize="small" />
                            クリア
                        </div>
                    </Button>
                </div>
            </form>
            <div className={styles['product-content']}>
                {isLoading ? (
                    <div className={styles['loading']}>読み込み中...</div>
                ) : (
                    <div className={styles['product-list']}>
                        {products.length === 0 ? (
                            <div className={styles['empty-message']}>{emptyMessage}</div>
                        ) : (
                            <div className={styles['product-grid']}>
                                {products.map((product) => (
                                    <ProductCard admin key={product.uuid} onDelete={handleDelete} onEdit={handleEdit} product={product} />
                                ))}
                            </div>
                        )}
                    </div>
                )}
                <Pagination currentPage={pageInfo.page} disabled={isLoading} onPageChange={handlePageChange} totalPages={pageInfo.totalPages} />
            </div>

            <ProductFormDialog
                categories={categories}
                isOpen={isDialogOpen}
                isSubmitting={isSubmitting}
                onClose={handleCloseDialog}
                onCreemaDuplicate={handleCreemaDuplicate}
                onSubmit={handleSubmit}
                salesSites={salesSites}
                submitError={submitError}
                tags={tags}
                targets={targets}
                updateItem={updateItem}
            />

            <Dialog
                cancelOption={{
                    label: 'キャンセル',
                    onClick: handleCloseDeleteDialog,
                }}
                confirmOption={{
                    label: '削除',
                    onClick: handleConfirmDelete,
                }}
                isOpen={isDeleteDialogOpen}
                onClose={handleCloseDeleteDialog}
                title="削除確認"
            >
                {deleteTargetItem && (
                    <>
                        <p>商品「{deleteTargetItem.name}」を削除しますか？</p>
                        <p>この操作は取り消せません。</p>
                    </>
                )}
            </Dialog>
        </div>
    )
}
