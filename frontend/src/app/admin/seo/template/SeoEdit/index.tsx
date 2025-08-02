'use client'

import { ChangeEvent, useState } from 'react'

import { updateCreator, updateCreatorLogo } from '@/apis/creator'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { Image } from '@/components/bases/Image'
import { ICreator } from '@/features/creator/type'

import { SeoEditForm } from './SeoEditForm'
import styles from './styles.module.scss'

interface Props {
    creator: ICreator
    onUpdate: (_creator: ICreator) => void
}

export const SeoEdit = ({ creator, onUpdate }: Props) => {
    const [isEditDialogOpen, setIsEditDialogOpen] = useState<boolean>(false)
    const [isLoading, setIsLoading] = useState<boolean>(false)

    const handleFileChange = async (event: ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0]
        if (!file) return

        setIsLoading(true)
        try {
            const updatedCreator = await updateCreatorLogo(file)
            onUpdate(updatedCreator)
        } catch (error) {
            console.error('ロゴの更新に失敗しました:', error)
            alert('ロゴの更新に失敗しました')
        } finally {
            setIsLoading(false)
        }
    }

    const handleTextUpdate = async (name: string, introduction: string) => {
        setIsLoading(true)
        try {
            const updatedCreator = await updateCreator({ name, introduction })
            onUpdate(updatedCreator)
            setIsEditDialogOpen(false)
        } catch (error) {
            console.error('作者情報の更新に失敗しました:', error)
            alert('作者情報の更新に失敗しました')
        } finally {
            setIsLoading(false)
        }
    }

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
                    <div className={styles['logo-upload']}>
                        <input
                            accept="image/*"
                            className={styles['file-input']}
                            disabled={isLoading}
                            id="logo-upload"
                            onChange={handleFileChange}
                            type="file"
                        />
                        <label className={styles['file-label']} htmlFor="logo-upload">
                            {isLoading ? 'アップロード中...' : 'ロゴを変更'}
                        </label>
                    </div>
                </div>
            </div>

            <div className={styles['section']}>
                <h2 className={styles['section-title']}>サイト説明</h2>
                <div className={styles['description-section']}>
                    <div className={styles['description-display']}>
                        <pre className={styles['description-text']}>{creator.introduction}</pre>
                    </div>
                    <Button disabled={isLoading} onClick={() => setIsEditDialogOpen(true)}>
                        編集
                    </Button>
                </div>
            </div>

            <Dialog isOpen={isEditDialogOpen} onClose={() => setIsEditDialogOpen(false)} title="SEO情報編集" wide>
                <SeoEditForm creator={creator} isLoading={isLoading} onCancel={() => setIsEditDialogOpen(false)} onSubmit={handleTextUpdate} />
            </Dialog>
        </div>
    )
}
