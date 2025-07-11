import { X, Facebook } from '@mui/icons-material'
import ReplyIcon from '@mui/icons-material/Reply'

import { Icon } from '@/components/bases/Icon'
import { ColorEnum } from '@/types'

import styles from './styles.module.scss'

export const ShareButtons = () => {
    return (
        <div className={styles['container']}>
            <div className={styles['message']}>
                <div className={styles['reply']}>
                    <ReplyIcon />
                </div>
                <div>Share This Page♪</div>
            </div>
            <div className={styles['icon-area']}>
                <div>
                    <Icon color={ColorEnum.Primary} contrast shadow={false} size={40}>
                        <X />
                    </Icon>
                </div>
                <div>
                    <Icon color={ColorEnum.Primary} contrast shadow={false} size={40}>
                        <Facebook />
                    </Icon>
                </div>
            </div>
        </div>
    )
}
