import { ColorCode, ColorType } from '@/types'

import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    colorType?: ColorType
    onClick?: () => void
}

export const Button = ({ children, colorType = ColorType.Primary, onClick = () => {} }: Props) => {
    return (
        <button className={styles['container']} onClick={onClick} style={{ backgroundColor: ColorCode[colorType] }}>
            <span>{children}</span>
        </button>
    )
}
