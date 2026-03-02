'use client'

import { KeyboardArrowDown } from '@mui/icons-material'

import { Breadcrumbs } from '@/components/bases/Breadcrumbs'
import { Select } from '@/components/bases/Select'
import { IClassification } from '@/features/classification/type'
import { ProductsByCategoryDisplay } from '@/features/product/components/ProductsByCategoryDisplay'
import { IProductsByCategory } from '@/features/product/type'

import { useProductTemplate } from './hooks'
import styles from './styles.module.scss'

type Props = {
    categories: IClassification[]
    productsByCategory: IProductsByCategory[]
    targets: IClassification[]
}

const ProductTemplate = ({ productsByCategory: initialProductsByCategory, categories, targets }: Props) => {
    const { isFetchingProducts, loadingCategoryUUIDs, onClickMore, onSelect, visibleProductsByCategory } = useProductTemplate({
        initialProductsByCategory,
    })

    return (
        <div className={styles['container']}>
            <div className={styles['search-area']}>
                <div className={styles['search-area__select']}>
                    <Select
                        onSelect={(option) => {
                            onSelect(option, 'category')
                        }}
                        options={categories.map((v) => ({ value: v.uuid, label: v.name }))}
                        suffix={<KeyboardArrowDown />}
                        title="Category"
                    />
                </div>
                <div className={styles['search-area__select']}>
                    <Select
                        onSelect={(option) => {
                            onSelect(option, 'target')
                        }}
                        options={targets.map((v) => ({ value: v.uuid, label: v.name }))}
                        suffix={<KeyboardArrowDown />}
                        title="Target"
                    />
                </div>
            </div>
            <div className={styles['product-area']}>
                {(() => {
                    if (!isFetchingProducts && visibleProductsByCategory.length === 0) {
                        return (
                            <div className={styles['product-area__no-product-message']}>
                                該当する商品が
                                <br className={styles['sm']} />
                                見つかりませんでした
                            </div>
                        )
                    } else {
                        return visibleProductsByCategory.map((v) => {
                            return (
                                <div className={styles['product-area__products-by-category']} key={v.category.uuid}>
                                    <ProductsByCategoryDisplay
                                        hasMore={v.pageInfo.hasMore}
                                        isLoadingMore={loadingCategoryUUIDs.includes(v.category.uuid)}
                                        onClickMoreButton={() => {
                                            void onClickMore(v.category.uuid)
                                        }}
                                        productsByCategory={v}
                                    />
                                </div>
                            )
                        })
                    }
                })()}
            </div>
            <Breadcrumbs
                breadcrumbs={[
                    {
                        label: 'トップページ',
                        link: '/',
                    },
                    {
                        label: '商品一覧',
                    },
                ]}
            />
        </div>
    )
}

export default ProductTemplate
