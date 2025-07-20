import React from 'react'

import styles from './styles.module.scss'

export interface TabButton {
    key: string
    label: string
}

interface Props {
    activeKey: string
    items: TabButton[]
    onTabChange: (_key: string) => void
}

export const TabButtons = ({ items, activeKey, onTabChange }: Props) => {
    return (
        <div className={styles['tab-buttons']}>
            {items.map((item) => (
                <button
                    className={`${styles['tab-button']} ${activeKey === item.key ? styles['active'] : ''}`}
                    key={item.key}
                    onClick={() => onTabChange(item.key)}
                >
                    {item.label}
                </button>
            ))}
        </div>
    )
}
