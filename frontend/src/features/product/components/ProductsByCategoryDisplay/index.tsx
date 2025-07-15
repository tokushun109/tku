import classNames from 'classnames'
import { useState } from 'react'

import { Slide } from '@/components/animations/Slide'
import { AnimatedButton } from '@/components/bases/AnimatedButton'
import ProductThumbnail from '@/features/product/components/ProductThumbnail'
import { mainFontFace } from '@/utils/font'

import styles from './styles.module.scss'
import { IProductsByCategory } from '../../type'

// 初期表示の商品数
const INIT_DISPLAY_LIMIT = 4
// 「もっと見る」で追加表示する商品数
const MORE_DISPLAY_LIMIT = 8

type Props = {
    productsByCategory: IProductsByCategory
}

export const ProductsByCategoryDisplay = ({ productsByCategory }: Props) => {
    const [displayCount, setDisplayCount] = useState<number>(INIT_DISPLAY_LIMIT)

    /** 表示する商品のリスト(設定された件数まで表示) */
    const displayProducts = productsByCategory.products.slice(0, displayCount)

    const onClickMoreButton = () => {
        setDisplayCount((prev) => prev + MORE_DISPLAY_LIMIT)
    }

    // 商品がなければ表示しない
    if (productsByCategory.products.length === 0) {
        return undefined
    }

    return (
        <div className={styles['container']}>
            <div className={classNames(styles['category-name'], mainFontFace.className)}>{productsByCategory.category.name}</div>
            <div className={styles['product-list']}>
                {displayProducts.map((v) => (
                    <div className={styles['product-thumbnail']} key={v.uuid}>
                        <Slide>
                            <ProductThumbnail item={{ product: v, apiPath: v.productImages?.[0]?.apiPath || '' }} />
                        </Slide>
                    </div>
                ))}
            </div>
            {displayCount < productsByCategory.products.length && (
                <div className={styles['more-button']}>
                    <AnimatedButton onClick={onClickMoreButton}>もっと見る</AnimatedButton>
                </div>
            )}
        </div>
    )
}
