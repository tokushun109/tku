import React from 'react'

import styles from './styles.module.scss'

export interface TabItem {
    icon?: React.ReactNode
    key: string
    label: string
}

interface Props {
    activeKey: string
    items: TabItem[]
    onTabChange: (_key: string) => void
}

export const Tab = ({ items, activeKey, onTabChange }: Props) => {
    return (
        <div className={styles['tab-container']}>
            <div className={styles['tab-buttons']}>
                {items.map((item) => (
                    <button
                        className={`${styles['tab-button']} ${activeKey === item.key ? styles['active'] : ''}`}
                        key={item.key}
                        onClick={() => onTabChange(item.key)}
                    >
                        {item.icon && <span className={styles['tab-icon']}>{item.icon}</span>}
                        {item.label}
                    </button>
                ))}
            </div>
        </div>
    )
}
