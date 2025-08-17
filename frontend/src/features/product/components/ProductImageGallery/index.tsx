'use client'

import { useState } from 'react'

import { Image } from '@/components/bases/Image'
import { IProduct } from '@/features/product/type'

import styles from './styles.module.scss'

type Props = {
    product: IProduct
}

export const ProductImageGallery = ({ product }: Props) => {
    const [selectedImageIndex, setSelectedImageIndex] = useState(0)

    // 管理画面で設定済みの順序で取得されるため、ソートは不要
    const images = product.productImages || []

    const hasImages = images.length > 0

    const currentImage = hasImages ? images[selectedImageIndex] : null

    return (
        <div className={styles['container']}>
            <div className={styles['main-image-area']}>
                <Image alt={product.name} src={currentImage?.apiPath || ''} />
            </div>

            {hasImages && images.length > 1 && (
                <div className={styles['thumbnail-area']}>
                    {images.map((image, index) => (
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
