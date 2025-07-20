'use client'

import { useState } from 'react'

import { ClassificationList } from '@/features/classification/components/ClassificationList'
import { IClassification } from '@/features/classification/type'
import { ClassificationType, ClassificationLabel } from '@/types'

import styles from './styles.module.scss'

interface Props {
    categories: IClassification[]
    tags: IClassification[]
    targets: IClassification[]
}

export const ClassificationTemplate = ({ categories, targets, tags }: Props) => {
    const [activeTab, setActiveTab] = useState<string>(ClassificationType.Category)

    const handleTabChange = (newTab: string) => {
        setActiveTab(newTab)
    }

    return (
        <div className={styles['classification-container']}>
            <div className={styles['tab-header']}>
                <div className={styles['tab-buttons']}>
                    <button
                        className={`${styles['tab-button']} ${activeTab === ClassificationType.Category ? styles['active'] : ''}`}
                        onClick={() => handleTabChange(ClassificationType.Category)}
                    >
                        {ClassificationLabel[ClassificationType.Category]}
                    </button>
                    <button
                        className={`${styles['tab-button']} ${activeTab === ClassificationType.Target ? styles['active'] : ''}`}
                        onClick={() => handleTabChange(ClassificationType.Target)}
                    >
                        {ClassificationLabel[ClassificationType.Target]}
                    </button>
                    <button
                        className={`${styles['tab-button']} ${activeTab === ClassificationType.Tag ? styles['active'] : ''}`}
                        onClick={() => handleTabChange(ClassificationType.Tag)}
                    >
                        {ClassificationLabel[ClassificationType.Tag]}
                    </button>
                </div>
            </div>
            <div className={styles['tab-content']}>
                <div className={activeTab === ClassificationType.Category ? styles['tab-panel-active'] : styles['tab-panel-hidden']}>
                    <ClassificationList items={categories} />
                </div>
                <div className={activeTab === ClassificationType.Target ? styles['tab-panel-active'] : styles['tab-panel-hidden']}>
                    <ClassificationList items={targets} />
                </div>
                <div className={activeTab === ClassificationType.Tag ? styles['tab-panel-active'] : styles['tab-panel-hidden']}>
                    <ClassificationList items={tags} />
                </div>
            </div>
        </div>
    )
}
