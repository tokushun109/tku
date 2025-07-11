import classNames from 'classnames'
import { ReactNode } from 'react'

import styles from './styles.module.scss'
import { ColorType } from '@/types'

type Props = {
    children: ReactNode
    className?: string
    color?: ColorType
}

export const Tag = ({ children, color = ColorType.Primary, className }: Props) => {
    return <span className={classNames(styles['tag'], styles[`tag--${color}`], className)}>{children}</span>
}
