import classNames from 'classnames'

import { Slide } from '@/components/animations/Slide'
import { AnimatedButton } from '@/components/bases/AnimatedButton'
import ProductThumbnail from '@/features/product/components/ProductThumbnail'
import { mainFontFace } from '@/utils/font'

import styles from './styles.module.scss'
import { IProductsByCategory } from '../../type'

type Props = {
    hasMore: boolean
    isLoadingMore?: boolean
    onClickMoreButton?: () => Promise<void> | void
    productsByCategory: IProductsByCategory
}

export const ProductsByCategoryDisplay = ({ productsByCategory, hasMore, isLoadingMore = false, onClickMoreButton }: Props) => {
    // 商品がなければ表示しない
    if (productsByCategory.products.length === 0) {
        return undefined
    }

    return (
        <div className={styles['container']}>
            <div className={classNames(styles['category-name'], mainFontFace.className)}>{productsByCategory.category.name}</div>
            <div className={styles['product-list']}>
                {productsByCategory.products.map((v) => (
                    <div className={styles['product-thumbnail']} key={v.uuid}>
                        <Slide>
                            <ProductThumbnail item={{ product: v, apiPath: v.productImages?.[0]?.apiPath || '' }} />
                        </Slide>
                    </div>
                ))}
            </div>
            {hasMore && (
                <div className={styles['more-button']}>
                    <AnimatedButton
                        enabledAnimation={!isLoadingMore}
                        onClick={
                            isLoadingMore || !onClickMoreButton
                                ? undefined
                                : () => {
                                      void onClickMoreButton()
                                  }
                        }
                    >
                        {isLoadingMore ? '読み込み中...' : 'もっと見る'}
                    </AnimatedButton>
                </div>
            )}
        </div>
    )
}
