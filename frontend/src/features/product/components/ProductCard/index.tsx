import { Image } from '@/components/bases/Image'
import { IProduct } from '@/features/product/type'

import styles from './styles.module.scss'

interface Props {
    admin?: boolean
    onDelete?: (_product: IProduct) => void
    onEdit?: (_product: IProduct) => void
    product: IProduct
}

export const ProductCard = ({ product, onEdit, onDelete, admin = false }: Props) => {
    const handleEdit = () => {
        onEdit?.(product)
    }

    const handleDelete = () => {
        onDelete?.(product)
    }

    return (
        <div
            className={`${styles['product-card']} ${product.isActive ? styles['is-active'] : ''} ${
                product.isActive && product.isRecommend ? styles['is-recommend'] : ''
            }`}
        >
            <div className={styles['product-card-wrapper']}>
                <div className={styles['product-card-header']}>
                    <div className={styles['product-name']}>{product.name}</div>
                    <div className={styles['product-status']}>
                        {admin && product.isRecommend && <span className={styles['chip-recommend']}>おすすめ</span>}
                        {admin && !product.isActive && <span className={styles['chip-inactive']}>展示</span>}
                    </div>
                </div>
                <div className={styles['product-card-image-container']}>
                    {product.productImages.length > 0 ? (
                        <Image alt={product.name} src={product.productImages[0].apiPath} />
                    ) : (
                        <Image alt="no-image" src="/image/gray-image.png" />
                    )}
                    {product.category.uuid && (
                        <div className={styles['product-category']}>
                            <span className={styles['chip-category']}>{product.category.name}</span>
                        </div>
                    )}
                    {product.target.uuid && (
                        <div className={styles['product-target']}>
                            <span className={styles['chip-target']}>{product.target.name}</span>
                        </div>
                    )}
                </div>
                <div className={styles['product-card-footer']}>
                    <div className={styles['product-card-footer-content']}>
                        {admin && (
                            <div className={styles['admin-actions']}>
                                <button className={styles['edit-button']} onClick={handleEdit}>
                                    編集
                                </button>
                                <button className={styles['delete-button']} onClick={handleDelete}>
                                    削除
                                </button>
                            </div>
                        )}
                        <div className={styles['price']}>
                            ¥{product.price.toLocaleString()}
                            <span className={styles['tax-label']}>税込</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
