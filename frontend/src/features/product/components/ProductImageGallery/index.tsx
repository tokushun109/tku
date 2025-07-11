'use client'

import { Image } from '@/components/bases/Image'

import { useState } from 'react'

import { IProduct } from '@/features/product/type'

import styles from './styles.module.scss'

type Props = {
    product: IProduct
}

export const ProductImageGallery = ({ product }: Props) => {
    const [selectedImageIndex, setSelectedImageIndex] = useState(0)

    // 商品画像を順序でソート
    const sortedImages = [...product.productImages].sort((a, b) => a.order - b.order)

    const hasImages = sortedImages.length > 0

    const currentImage = hasImages ? sortedImages[selectedImageIndex] : null

    return (
        <div className={styles['container']}>
            <div className={styles['main-image-area']}>
                <Image alt={product.name} src={currentImage?.apiPath || ''} />
            </div>

            {hasImages && sortedImages.length > 1 && (
                <div className={styles['thumbnail-area']}>
                    {sortedImages.map((image, index) => (
                        <div
                            className={`${styles['thumbnail']} ${index === selectedImageIndex ? styles['thumbnail--active'] : ''}`}
                            key={image.uuid}
                            onClick={() => setSelectedImageIndex(index)}
                        >
                            <Image alt={`${product.name} ${index + 1}`} src={image.apiPath} />
                        </div>
                    ))}
                </div>
            )}
        </div>
    )
}
