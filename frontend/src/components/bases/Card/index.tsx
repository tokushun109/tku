import classNames from 'classnames'

import { ColorCodeEnum, ColorEnum, ColorType } from '@/types'

import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    color?: ColorType
    height?: string
    shadow?: boolean
    width?: string
}

export const Card = ({ color = ColorEnum.White, width = 'auto', height = 'auto', shadow = true, children }: Props) => {
    return (
        <div
            className={classNames(styles['container'], !shadow && styles['no-shadow'])}
            style={{
                width,
                height,
                background: ColorCodeEnum[color],
            }}
        >
            <div className={styles['content']}>{children}</div>
        </div>
    )
}
