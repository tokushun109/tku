import { ColorCodeEnum, ColorType } from '@/types'

import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    color: ColorType
    fontColor?: string
    fontSize?: number
}

export const Chip = ({ color, fontColor = '#ffffff', fontSize = 16, children }: Props) => {
    return (
        <span
            className={styles.container}
            style={{
                background: ColorCodeEnum[color],
                color: fontColor,
                fontSize: `${fontSize}px`,
            }}
        >
            {children}
        </span>
    )
}
