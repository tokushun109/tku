import React from 'react'

import styles from './styles.module.scss'

export interface TabItem<T = string> {
    icon?: React.ReactNode
    key: T
    label: string
}

interface Props<T extends React.Key> {
    activeKey: T
    items: TabItem<T>[]
    onTabChange: (_key: T) => void
}

export const Tab = <T extends React.Key>({ items, activeKey, onTabChange }: Props<T>) => {
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
