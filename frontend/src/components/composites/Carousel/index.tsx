import { CarouselImage } from '@/components/composites/CarouselImage'
import { IThumbnail } from '@/features/product/type'

import styles from './styles.module.scss'

import type { CSSProperties } from 'react'

type Props = {
    items: IThumbnail[]
}

type CarouselStyle = CSSProperties & {
    '--animation-duration': string
    '--item-length': number
}

const REPEAT_COUNT = 3

export const Carousel = ({ items }: Props) => {
    const wrapperStyle: CarouselStyle = {
        '--animation-duration': `${Math.max(items.length, 1) * 6}s`,
        '--item-length': items.length,
    }
    const repeatedItems = Array.from({ length: REPEAT_COUNT }, (_, repeatIndex) =>
        items.map((item, itemIndex) => ({
            item,
            key: `${item.product?.uuid ?? item.apiPath}-${repeatIndex}-${itemIndex}`,
        })),
    ).flat()

    return (
        <div className={styles['container']} data-testid="carousel">
            <div className={styles['wrapper']} style={wrapperStyle}>
                {repeatedItems.map(({ item, key }) => (
                    <div className={styles['product-image']} key={key}>
                        <CarouselImage item={item} />
                    </div>
                ))}
            </div>
        </div>
    )
}
