import { Button } from '@/components/bases/Button'
import { Chip } from '@/components/bases/Chip'
import { ProductImageGallery } from '@/features/product/components/ProductImageGallery'
import { IProduct } from '@/features/product/type'
import { ColorType, FontSizeType } from '@/types'
import { formatPrice } from '@/utils/price'

import styles from './styles.module.scss'

type Props = {
    product: IProduct
}

const handleOpenExternal = (url: string) => {
    window.open(url, '_blank', 'noopener')
}

export const ProductDetail = ({ product }: Props) => {
    return (
        <div className={styles['container']}>
            <h1 className={styles['product-name']}>{product.name}</h1>
            <div className={styles['detail-area']}>
                <div className={styles['image-area']}>
                    <ProductImageGallery product={product} />
                </div>
                <div className={styles['info-area']}>
                    <div className={styles['description-area']}>
                        <pre className={styles['description']}>{product.description}</pre>
                    </div>

                    {product.target.uuid && (
                        <div className={styles['target-area']}>
                            <p className={styles['label']}>対象</p>
                            <div className={styles['content']}>
                                <Chip color="secondary" fontSize={FontSizeType.SmMd}>
                                    {product.target.name}
                                </Chip>
                            </div>
                        </div>
                    )}

                    {product.category.uuid && (
                        <div className={styles['category-area']}>
                            <p className={styles['label']}>カテゴリー</p>
                            <div className={styles['content']}>
                                <Chip color="secondary" fontSize={FontSizeType.SmMd}>
                                    {product.category.name}
                                </Chip>
                            </div>
                        </div>
                    )}

                    {product.tags.length > 0 && (
                        <div className={styles['tag-area']}>
                            <p className={styles['label']}>タグ</p>
                            <div className={styles['tag-content']}>
                                {product.tags.map((tag) => (
                                    <Chip color="secondary" fontSize={FontSizeType.SmMd} key={tag.uuid}>
                                        {tag.name}
                                    </Chip>
                                ))}
                            </div>
                        </div>
                    )}

                    {product.siteDetails.length > 0 && (
                        <div className={styles['sales-site-area']}>
                            <p className={styles['label']}>販売サイト</p>
                            <div className={styles['site-buttons']}>
                                {product.siteDetails.map((siteDetail) => (
                                    <Button
                                        colorType={ColorType.Accent}
                                        key={siteDetail.uuid}
                                        onClick={() => handleOpenExternal(siteDetail.detailUrl)}
                                    >
                                        {siteDetail.salesSite.name}
                                    </Button>
                                ))}
                            </div>
                        </div>
                    )}

                    <div className={styles['price-area']}>
                        <p className={styles['price']}>
                            ￥{formatPrice(product.price)}
                            <span className={styles['tax-label']}>(税込)</span>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    )
}
