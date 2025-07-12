'use client'

import { Breadcrumbs } from '@/components/bases/Breadcrumbs'
import { ProductDetail } from '@/features/product/components/ProductDetail'
import { IProduct } from '@/features/product/type'

import styles from './styles.module.scss'

type Props = {
    product: IProduct
}

const ProductDetailTemplate = ({ product }: Props) => {
    return (
        <div className={styles['container']}>
            <div className={styles['product-detail-area']}>
                <ProductDetail product={product} />
            </div>
            <Breadcrumbs
                breadcrumbs={[
                    {
                        label: 'トップページ',
                        link: '/',
                    },
                    {
                        label: '商品一覧',
                        link: '/product',
                    },
                    {
                        label: product.name,
                    },
                ]}
            />
        </div>
    )
}

export default ProductDetailTemplate
