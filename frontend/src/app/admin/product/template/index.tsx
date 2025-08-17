'use client'

import { Add } from '@mui/icons-material'
import { useState } from 'react'
import { toast } from 'sonner'

import { getCategories } from '@/apis/category'
import { createProduct, deleteProduct, duplicateProductFromCreema, getProducts, updateProduct, uploadProductImage } from '@/apis/product'
import { getSalesSiteList } from '@/apis/salesSite'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { IClassification } from '@/features/classification/type'
import { ProductCard } from '@/features/product/components/ProductCard'
import { ProductFormDialog } from '@/features/product/components/ProductFormDialog'
import { ICreemaDuplicateForm } from '@/features/product/product/type'
import { IProduct, IProductForm } from '@/features/product/type'
import { ISite } from '@/features/site/type'

import styles from './styles.module.scss'

interface Props {
    categories: IClassification[]
    initialProducts: IProduct[]
    salesSites: ISite[]
    tags: IClassification[]
    targets: IClassification[]
}

export const AdminProductTemplate = ({
    categories: initialCategories,
    initialProducts,
    salesSites: initialSalesSites,
    tags: initialTags,
    targets: initialTargets,
}: Props) => {
    const [products, setProducts] = useState<IProduct[]>(initialProducts)
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

    const fetchData = async () => {
        try {
            setIsLoading(true)
            const [fetchedProducts, categoriesData, targetsData, tagsData, salesSitesData] = await Promise.all([
                getProducts({
                    mode: 'all',
                    category: 'all',
                    target: 'all',
                }),
                getCategories({ mode: 'all' }),
                getTargets({ mode: 'all' }),
                getTags(),
                getSalesSiteList(),
            ])
            setProducts(fetchedProducts)
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
            await fetchData()
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

                // 既存画像の順序更新処理（並び替えが行われた場合）
                if (data.isImageOrderChanged && data.imageItems) {
                    const existingItems = data.imageItems.filter((item) => !item.isNewUpload)
                    const updatedProductImages = updateItem.productImages.map((image) => {
                        const reorderedItem = existingItems.find((item) => item.src === image.apiPath)
                        if (reorderedItem && reorderedItem.order) {
                            // 並び替え後の位置に基づいて100から降順で計算
                            return {
                                ...image,
                                order: 100 - (reorderedItem.order - 1),
                            }
                        }
                        return {
                            ...image,
                            order: image.order, // 並び替えされていない場合は既存の値を維持
                        }
                    })

                    // 順序更新されたproductImagesを含める
                    productDataToUpdate.productImages = updatedProductImages
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
                const existingImagesCount = updateItem?.productImages?.length || 0
                const hasOrderChanged = data.isImageOrderChanged || false

                // 新規画像の優先順位を計算
                let uploadFileOrder: { [key: number]: number } = {}

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

                        if (matchingItem && matchingItem.order) {
                            // 並び替え後の位置（1から始まる）を100から降順に変換
                            uploadFileOrder[uploadIndex] = 100 - (matchingItem.order - 1)
                        } else {
                            // フォールバック: 既存画像数を考慮した順序
                            uploadFileOrder[uploadIndex] = 100 - existingImagesCount - uploadIndex
                        }
                    })
                } else {
                    // 通常の場合、既存画像より低い優先順位を設定
                    data.uploadImages.forEach((_, index) => {
                        uploadFileOrder[index] = 100 - existingImagesCount - index
                    })
                }

                const orderParams = {
                    isChanged: hasOrderChanged,
                    order: uploadFileOrder,
                }

                await uploadProductImage(productUuid, data.uploadImages, orderParams)
            }

            setIsDialogOpen(false)
            setUpdateItem(null)

            // 成功通知
            if (updateItem) {
                toast.success(`商品「${data.name}」を更新しました`)
            } else {
                toast.success(`商品「${data.name}」を追加しました`)
            }

            await fetchData()
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
            await fetchData()
        } catch (error) {
            console.error('Creemaからの商品複製に失敗しました:', error)
            const errorMessage = 'Creemaからの商品複製に失敗しました。もう一度お試しください。'
            setSubmitError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <div className={styles['product-container']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>商品一覧</h1>
                <div className={styles['header-actions']}>
                    <div className={styles['product-count']}>{products.length}件の商品</div>
                    <Button onClick={handleCreate}>
                        <div className={styles['add-button-content']}>
                            <Add className={styles['add-icon']} fontSize="small" />
                            追加
                        </div>
                    </Button>
                </div>
            </div>
            <div className={styles['product-content']}>
                {isLoading ? (
                    <div className={styles['loading']}>読み込み中...</div>
                ) : (
                    <div className={styles['product-list']}>
                        {products.length === 0 ? (
                            <div className={styles['empty-message']}>登録されていません</div>
                        ) : (
                            <div className={styles['product-grid']}>
                                {products.map((product) => (
                                    <ProductCard admin key={product.uuid} onDelete={handleDelete} onEdit={handleEdit} product={product} />
                                ))}
                            </div>
                        )}
                    </div>
                )}
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
