'use client'

import { Category, Label, People } from '@mui/icons-material'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

import { Tab, TabItem } from '@/components/bases/Tab'
import { CategoryList } from '@/features/classification/components/CategoryList'
import { TagList } from '@/features/classification/components/TagList'
import { TargetList } from '@/features/classification/components/TargetList'
import { IClassification } from '@/features/classification/type'
import { ClassificationType, ClassificationLabel } from '@/types'

import styles from './styles.module.scss'

interface Props {
    categories: IClassification[]
    tags: IClassification[]
    targets: IClassification[]
}

export const ClassificationTemplate = ({ categories, targets, tags }: Props) => {
    const [activeTab, setActiveTab] = useState<ClassificationType>(ClassificationType.Category)
    const router = useRouter()

    const tabItems: TabItem<ClassificationType>[] = [
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

    const handleUpdate = () => {
        router.refresh()
    }

    const renderContent = () => {
        switch (activeTab) {
            case ClassificationType.Category:
                return <CategoryList items={categories} onUpdate={handleUpdate} />
            case ClassificationType.Target:
                return <TargetList items={targets} onUpdate={handleUpdate} />
            case ClassificationType.Tag:
                return <TagList items={tags} onUpdate={handleUpdate} />
            default:
                return null
        }
    }

    return (
        <div className={styles['classification-container']}>
            <Tab activeKey={activeTab} items={tabItems} onTabChange={setActiveTab} />
            <div className={styles['tab-content']}>{renderContent()}</div>
        </div>
    )
}
