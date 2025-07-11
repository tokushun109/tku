'use client'

import Image from 'next/image'
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

    // 画像がない場合のデフォルト画像
    const noImagePath = '/image/product/no-image.png'
    const hasImages = sortedImages.length > 0

    const currentImage = hasImages ? sortedImages[selectedImageIndex] : null

    return (
        <div className={styles['container']}>
            <div className={styles['main-image-area']}>
                <Image
                    alt={product.name}
                    className={styles['main-image']}
                    fill
                    src={currentImage?.apiPath || noImagePath}
                    style={{ objectFit: 'cover' }}
                />
            </div>

            {hasImages && sortedImages.length > 1 && (
                <div className={styles['thumbnail-area']}>
                    {sortedImages.map((image, index) => (
                        <button
                            className={`${styles['thumbnail']} ${index === selectedImageIndex ? styles['thumbnail--active'] : ''}`}
                            key={image.uuid}
                            onClick={() => setSelectedImageIndex(index)}
                            type="button"
                        >
                            <Image
                                alt={`${product.name} ${index + 1}`}
                                className={styles['thumbnail-image']}
                                fill
                                src={image.apiPath}
                                style={{ objectFit: 'cover' }}
                            />
                        </button>
                    ))}
                </div>
            )}
        </div>
    )
}
