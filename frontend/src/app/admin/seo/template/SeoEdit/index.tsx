'use client'

import { useState } from 'react'

import { Button } from '@/components/bases/Button'
import { Image } from '@/components/bases/Image'
import { TextArea } from '@/components/bases/TextArea'
import { ICreator } from '@/features/creator/type'

import { SeoEditDialog } from './SeoEditDialog'
import styles from './styles.module.scss'

interface Props {
    creator: ICreator
    onUpdate: () => Promise<void>
}

export const SeoEdit = ({ creator, onUpdate }: Props) => {
    const [isEditDialogOpen, setIsEditDialogOpen] = useState<boolean>(false)

    return (
        <div className={styles['seo-edit']}>
            <div className={styles['section']}>
                <h2 className={styles['section-title']}>サイトロゴ</h2>
                <div className={styles['logo-section']}>
                    {creator.apiPath && (
                        <div className={styles['logo-preview']}>
                            <div className={styles['logo-image']}>
                                <Image alt={creator.name} src={creator.apiPath} />
                            </div>
                        </div>
                    )}
                </div>
            </div>

            <div className={styles['section']}>
                <h2 className={styles['section-title']}>サイト説明</h2>
                <div className={styles['description-section']}>
                    <TextArea disabled rows={6} value={creator.introduction} />
                    <Button className={styles['edit-button']} onClick={() => setIsEditDialogOpen(true)}>
                        編集
                    </Button>
                </div>
            </div>

            <SeoEditDialog creator={creator} isOpen={isEditDialogOpen} onClose={() => setIsEditDialogOpen(false)} onUpdate={onUpdate} />
        </div>
    )
}
