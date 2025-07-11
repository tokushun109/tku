import { Facebook, Reply, X } from '@mui/icons-material'

import { Icon } from '@/components/bases/Icon'
import { ColorType } from '@/types'

import styles from './styles.module.scss'

export const ShareButtons = () => {
    return (
        <div className={styles['container']}>
            <div className={styles['message']}>
                <div className={styles['reply']}>
                    <Reply />
                </div>
                <div>Share This Page♪</div>
            </div>
            <div className={styles['icon-area']}>
                <div>
                    <Icon color={ColorType.Primary} contrast shadow={false} size={40}>
                        <X />
                    </Icon>
                </div>
                <div>
                    <Icon color={ColorType.Primary} contrast shadow={false} size={40}>
                        <Facebook />
                    </Icon>
                </div>
            </div>
        </div>
    )
}
