'use client'

import { ReactNode } from 'react'

import styles from './styles.module.scss'

interface Props {
    actions?: ReactNode
    children: ReactNode
    onClick?: () => void
    subItem?: ReactNode
}

export const ListItem = ({ children, subItem, actions, onClick }: Props) => {
    return (
        <div className={styles['list-item']} onClick={onClick}>
            <div className={styles['item-content']}>
                <div className={styles['item-main']}>
                    {children}
                    {subItem && <div className={styles['item-sub']}>{subItem}</div>}
                </div>
                {actions && <div className={styles['item-actions']}>{actions}</div>}
            </div>
        </div>
    )
}
