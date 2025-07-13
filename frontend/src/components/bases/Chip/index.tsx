import { ColorCode, ColorType, FontSizeType, FontSizeValue } from '@/types'

import styles from './styles.module.scss'

export const ChipSize = {
    Small: 'small',
    Medium: 'medium',
    Large: 'large',
} as const
export type ChipSize = (typeof ChipSize)[keyof typeof ChipSize]

type Props = {
    children: React.ReactNode
    color: ColorType
    fontColor?: string
    fontSize?: FontSizeType
    size?: ChipSize
}

const getSizeStyles = (size: ChipSize): { borderRadius: string; padding: string } => {
    switch (size) {
        case ChipSize.Small:
            return {
                borderRadius: '12px',
                padding: '4px 8px',
            }
        case ChipSize.Large:
            return {
                borderRadius: '32px',
                padding: '12px 16px',
            }
        case ChipSize.Medium:
        default:
            return {
                borderRadius: '24px',
                padding: '8px',
            }
    }
}

export const Chip = ({ color, fontColor = '#ffffff', fontSize = FontSizeType.Medium, size = ChipSize.Medium, children }: Props) => {
    const sizeStyles = getSizeStyles(size)

    return (
        <span
            className={styles.container}
            style={{
                background: ColorCode[color],
                color: fontColor,
                fontSize: `${FontSizeValue[fontSize]}px`,
                ...sizeStyles,
            }}
        >
            {children}
        </span>
    )
}
