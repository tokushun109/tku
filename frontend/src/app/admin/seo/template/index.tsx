'use client'

import { useState } from 'react'

import { getCreator } from '@/apis/creator'
import { ICreator } from '@/features/creator/type'

import { SeoEdit } from './SeoEdit'
import styles from './styles.module.scss'

interface Props {
    creator: ICreator | null
}

export const AdminSeoTemplate = ({ creator: initialCreator }: Props) => {
    const [creator, setCreator] = useState<ICreator | null>(initialCreator)

    const handleCreatorUpdate = async () => {
        const creator = await getCreator()
        setCreator(creator)
    }

    return (
        <div className={styles['seo-page']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>SEO設定</h1>
            </div>
            <div className={styles['divider']} />
            <div className={styles['content']}>
                {creator ? (
                    <SeoEdit creator={creator} onUpdate={handleCreatorUpdate} />
                ) : (
                    <div className={styles['error-message']}>作者情報の読み込みに失敗しました</div>
                )}
            </div>
        </div>
    )
}
