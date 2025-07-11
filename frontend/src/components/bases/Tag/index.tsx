import classNames from 'classnames'
import { ReactNode } from 'react'

import styles from './styles.module.scss'

type Color = 'primary' | 'secondary' | 'accent'

type Props = {
    children: ReactNode
    className?: string
    color?: Color
}

export const Tag = ({ children, color = 'primary', className }: Props) => {
    return <span className={classNames(styles['tag'], styles[`tag--${color}`], className)}>{children}</span>
}
