import classNames from 'classnames'

import { Card } from '@/components/bases/Card'
import { MaterialIconType } from '@/types'
import { labelFontFace } from '@/utils/font'

import styles from './styles.module.scss'

type Props = {
    Icon: MaterialIconType
    isSelected?: boolean
    label: string
}

export const IconCard = ({ Icon, label, isSelected = false }: Props) => {
    return (
        <div className={classNames(styles['container'], isSelected && styles['selected'])}>
            <Card>
                <div className={styles['content']}>
                    <div className={styles['icon']}>
                        <Icon />
                    </div>
                    <div className={classNames(styles['label'], labelFontFace.className)}>{label}</div>
                </div>
            </Card>
        </div>
    )
}
