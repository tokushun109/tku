'use client'

import { useRouter } from 'next/navigation'

import { Chip } from '@/components/bases/Chip'
import { Image } from '@/components/bases/Image'
import { ColorEnum } from '@/types'
import { numToPrice } from '@/utils/convert'

import styles from './styles.module.scss'
import { IThumbnail } from '../../type'

type Props = {
    item: IThumbnail
}

const ProductThumbnail = ({ item }: Props) => {
    const router = useRouter()

    const handleClick = () => {
        router.push(`/product/${item.product.uuid}`)
    }

    return (
        <div className={styles['container']} onClick={handleClick} style={{ cursor: 'pointer' }}>
            <div className={styles['image-container']}>
                <Image alt={item.product.name} src={item.apiPath} />
                <div className={styles['chip']}>
                    <Chip color={ColorEnum.Secondary} fontSize={12}>
                        {item.product.target.name}
                    </Chip>
                </div>
            </div>
            <div className={styles['name']}>{item.product.name}</div>
            <div className={styles['price']}>
                {numToPrice(item.product.price)} <span>(税込)</span>
            </div>
        </div>
    )
}

export default ProductThumbnail
