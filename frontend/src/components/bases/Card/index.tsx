import classNames from 'classnames'

import { ColorCode, ColorType } from '@/types'

import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    color?: ColorType
    height?: string
    shadow?: boolean
    width?: string
}

export const Card = ({ color = ColorType.White, width = 'auto', height = 'auto', shadow = true, children }: Props) => {
    return (
        <div
            className={classNames(styles['container'], !shadow && styles['no-shadow'])}
            style={{
                width,
                height,
                background: ColorCode[color],
            }}
        >
            <div className={styles['content']}>{children}</div>
        </div>
    )
}
