'use client'

import { Category, Label, People } from '@mui/icons-material'
import { useState } from 'react'

import { Tab, TabItem } from '@/components/bases/Tab'
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

    const tabItems: TabItem[] = [
        {
            key: ClassificationType.Category,
            label: ClassificationLabel[ClassificationType.Category],
            icon: <Category />,
        },
        {
            key: ClassificationType.Target,
            label: ClassificationLabel[ClassificationType.Target],
            icon: <People />,
        },
        {
            key: ClassificationType.Tag,
            label: ClassificationLabel[ClassificationType.Tag],
            icon: <Label />,
        },
    ]

    const getCurrentItems = () => {
        switch (activeTab) {
            case ClassificationType.Category:
                return categories
            case ClassificationType.Target:
                return targets
            case ClassificationType.Tag:
                return tags
            default:
                return []
        }
    }

    return (
        <div className={styles['classification-container']}>
            <Tab activeKey={activeTab} items={tabItems} onTabChange={setActiveTab} />
            <div className={styles['tab-content']}>
                <ClassificationList items={getCurrentItems()} />
            </div>
        </div>
    )
}
