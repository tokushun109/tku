'use client'

import { Add } from '@mui/icons-material'
import { useEffect, useState } from 'react'

import { createProduct, deleteProduct, getProducts, updateProduct } from '@/apis/product'
import { Button } from '@/components/bases/Button'
import { IClassification } from '@/features/classification/type'
import { ProductCard } from '@/features/product/components/ProductCard'
import { ProductFormDialog } from '@/features/product/components/ProductFormDialog'
import { IProduct, IProductForm } from '@/features/product/type'

import styles from './styles.module.scss'

export const AdminProductTemplate = () => {
    const [products, setProducts] = useState<IProduct[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [submitError, setSubmitError] = useState<string | null>(null)
    const [updateItem, setUpdateItem] = useState<IProduct | null>(null)

    const fetchProducts = async () => {
        try {
            setIsLoading(true)
            const fetchedProducts = await getProducts({
                mode: 'all',
                category: 'all',
                target: 'all',
            })
            setProducts(fetchedProducts)
        } catch (error) {
            console.error('商品の取得に失敗しました:', error)
        } finally {
            setIsLoading(false)
        }
    }

    useEffect(() => {
        fetchProducts()
    }, [])

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

    const handleDelete = async (product: IProduct) => {
        if (!confirm(`${product.name}を削除します。よろしいですか？`)) {
            return
        }

        try {
            await deleteProduct(product.uuid)
            await fetchProducts()
        } catch (error) {
            console.error('商品の削除に失敗しました:', error)
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

            if (updateItem) {
                // 編集
                await updateProduct(updateItem.uuid, { ...productData, uuid: updateItem.uuid })
            } else {
                // 新規作成
                await createProduct(productData)
            }

            setIsDialogOpen(false)
            setUpdateItem(null)
            await fetchProducts()
        } catch (error) {
            console.error('商品の保存に失敗しました:', error)
            setSubmitError('商品の保存に失敗しました。もう一度お試しください。')
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
                isOpen={isDialogOpen}
                isSubmitting={isSubmitting}
                onClose={handleCloseDialog}
                onSubmit={handleSubmit}
                submitError={submitError}
                updateItem={updateItem}
            />
        </div>
    )
}
