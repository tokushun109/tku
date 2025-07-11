import { ColorCode, ColorType, FontSizeType, FontSizeValue } from '@/types'

import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    color: ColorType
    fontColor?: string
    fontSize?: FontSizeType
}

export const Chip = ({ color, fontColor = '#ffffff', fontSize = FontSizeType.Medium, children }: Props) => {
    return (
        <span
            className={styles.container}
            style={{
                background: ColorCode[color],
                color: fontColor,
                fontSize: `${FontSizeValue[fontSize]}px`,
            }}
        >
            {children}
        </span>
    )
}
