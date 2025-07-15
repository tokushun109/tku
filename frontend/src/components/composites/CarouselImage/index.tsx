import classNames from 'classnames'
import { useRouter } from 'next/navigation'

import { Chip } from '@/components/bases/Chip'
import { Image } from '@/components/bases/Image'
import { IThumbnail } from '@/features/product/type'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

type Props = {
    item?: IThumbnail
    shadow?: boolean
}

export const CarouselImage = ({ item, shadow = true }: Props) => {
    const router = useRouter()

    if (!item || !item.product) {
        return null
    }

    const handleClick = () => {
        if (item.product?.uuid) {
            router.push(`/product/${item.product.uuid}`)
        }
    }

    return (
        <div className={classNames(styles['container'], shadow && styles['shadow'])} onClick={handleClick} style={{ cursor: 'pointer' }}>
            {item.product.category?.uuid && (
                <div className={classNames(styles['chip'], styles['category'])}>
                    <Chip color={ColorType.Accent} fontSize={FontSizeType.SmMd}>
                        {item.product.category.name}
                    </Chip>
                </div>
            )}
            <Image alt={item.product.name} key={item.product.name} src={item.apiPath} />
            <div className={classNames(styles['chip'], styles['name'])}>
                <Chip color={ColorType.Accent} fontSize={FontSizeType.SmMd}>
                    {item.product.name}
                </Chip>
            </div>
        </div>
    )
}
