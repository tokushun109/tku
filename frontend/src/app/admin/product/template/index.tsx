'use client'

import { Add } from '@mui/icons-material'
import { useEffect, useState } from 'react'

import { getProducts } from '@/apis/product'
import { Button } from '@/components/bases/Button'
import { ProductCard } from '@/features/product/components/ProductCard'
import { IProduct } from '@/features/product/type'

import styles from './styles.module.scss'

export const AdminProductTemplate = () => {
    const [products, setProducts] = useState<IProduct[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
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

        fetchProducts()
    }, [])

    return (
        <div className={styles['product-container']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>商品一覧</h1>
                <div className={styles['header-actions']}>
                    <div className={styles['product-count']}>{products.length}件の商品</div>
                    <Button
                        onClick={() => {
                            // TODO: 新規作成機能を実装
                        }}
                    >
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
                                    <ProductCard
                                        admin
                                        key={product.uuid}
                                        onDelete={(_product) => {
                                            // TODO: 削除機能を実装
                                        }}
                                        onEdit={(_product) => {
                                            // TODO: 編集機能を実装
                                        }}
                                        product={product}
                                    />
                                ))}
                            </div>
                        )}
                    </div>
                )}
            </div>
        </div>
    )
}
