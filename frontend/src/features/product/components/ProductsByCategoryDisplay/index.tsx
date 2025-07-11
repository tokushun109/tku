import classNames from 'classnames'
import { useState } from 'react'

import { Slide } from '@/components/animations/Slide'
import { Button } from '@/components/bases/Button'
import ProductThumbnail from '@/features/product/components/ProductThumbnail'
import { mainFontFace } from '@/utils/font'

import styles from './styles.module.scss'
import { IProductsByCategory } from '../../type'

// 初期表示の商品の上限数
const INIT_DISPLAY_LIMIT = 4

type Props = {
    productsByCategory: IProductsByCategory
}

export const ProductsByCategoryDisplay = ({ productsByCategory }: Props) => {
    const [isAllDisplayed, setIsAllDisplayed] = useState<boolean>(false)

    /** 表示する商品のリスト(もっと見るを押す前は4件まで、押した後は全て表示) */
    const displayProducts = isAllDisplayed ? productsByCategory.products : productsByCategory.products.slice(0, INIT_DISPLAY_LIMIT)

    const onClickMoreButton = () => {
        setIsAllDisplayed(true)
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
            {!isAllDisplayed && productsByCategory.products.length > 4 && (
                <div className={styles['more-button']}>
                    <Button onClick={onClickMoreButton}>もっと見る</Button>
                </div>
            )}
        </div>
    )
}
