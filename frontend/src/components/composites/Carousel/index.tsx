import { CarouselImage } from '@/components/composites/CarouselImage'
import { IThumbnail } from '@/features/product/type'

import styles from './styles.module.scss'

type Props = {
    items: IThumbnail[]
}

export const Carousel = ({ items }: Props) => {
    return (
        <div className={styles['container']}>
            <div className={styles['wrapper']}>
                {[...items, ...items].map((v, index) => (
                    <div className={styles['product-image']} key={index}>
                        <CarouselImage item={v} />
                    </div>
                ))}
            </div>
        </div>
    )
}
