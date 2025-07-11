import classNames from 'classnames'
import { useRouter } from 'next/navigation'

import { Chip } from '@/components/bases/Chip'
import { Image } from '@/components/bases/Image'
import { IThumbnail } from '@/features/product/type'
import { ColorType } from '@/types'

import styles from './styles.module.scss'

type Props = {
    item: IThumbnail
    shadow?: boolean
}

export const CarouselImage = ({ item, shadow = true }: Props) => {
    const router = useRouter()

    const handleClick = () => {
        router.push(`/product/${item.product.uuid}`)
    }

    return (
        <div className={classNames(styles['container'], shadow && styles['shadow'])} onClick={handleClick} style={{ cursor: 'pointer' }}>
            {item.product.category.uuid && (
                <div className={classNames(styles['chip'], styles['category'])}>
                    <Chip color={ColorType.Accent} fontSize={12}>
                        {item.product.category.name}
                    </Chip>
                </div>
            )}
            <Image alt={item.product.name} key={item.product.name} src={item.apiPath} />
            <div className={classNames(styles['chip'], styles['name'])}>
                <Chip color={ColorType.Accent} fontSize={12}>
                    {item.product.name}
                </Chip>
            </div>
        </div>
    )
}
