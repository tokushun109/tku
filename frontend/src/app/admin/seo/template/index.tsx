'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'

import { getCreator, updateCreator, updateCreatorLogo } from '@/apis/creator'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { FileInput } from '@/components/bases/FileInput'
import { Image } from '@/components/bases/Image'
import { TextArea } from '@/components/bases/TextArea'
import { ICreator } from '@/features/creator/type'
import { SeoEditSchema } from '@/features/seo/schema'
import { ISeoEdit } from '@/features/seo/type'

import styles from './styles.module.scss'

interface Props {
    creator: ICreator | null
}

export const AdminSeoTemplate = ({ creator: initialCreator }: Props) => {
    const [creator, setCreator] = useState<ICreator | null>(initialCreator)
    const [isEditDialogOpen, setIsEditDialogOpen] = useState<boolean>(false)
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [logoFile, setLogoFile] = useState<File | null>(null)
    const [previewUrl, setPreviewUrl] = useState<string | null>(null)

    const {
        formState: { errors },
        register,
        watch,
    } = useForm<ISeoEdit>({
        defaultValues: {
            introduction: creator?.introduction || '',
        },
        resolver: zodResolver(SeoEditSchema),
    })

    const watchedIntroduction = watch('introduction')

    const handleCreatorUpdate = async () => {
        const creator = await getCreator()
        setCreator(creator)
    }

    const handleSubmit = async () => {
        if (!creator) return

        setIsLoading(true)
        try {
            // テキスト情報を更新
            await updateCreator({ name: creator.name, introduction: watchedIntroduction || '' })

            // ロゴファイルがある場合はロゴも更新
            if (logoFile) {
                await updateCreatorLogo(logoFile)
            }

            await handleCreatorUpdate()
            toast.success('SEO情報を更新しました')
            handleDialogClose()
        } catch (error) {
            console.error('作者情報の更新に失敗しました:', error)
            toast.error('SEO情報の更新に失敗しました。もう一度お試しください。')
        } finally {
            setIsLoading(false)
        }
    }

    const handleLogoChange = (file: File | null) => {
        setLogoFile(file)
        if (file) {
            const url = URL.createObjectURL(file)
            setPreviewUrl(url)
        } else {
            setPreviewUrl(null)
        }
    }

    const handleDialogClose = () => {
        setIsEditDialogOpen(false)
        setLogoFile(null)
        if (previewUrl) {
            URL.revokeObjectURL(previewUrl)
            setPreviewUrl(null)
        }
    }

    return (
        <div className={styles['seo-page']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>SEO設定</h1>
            </div>
            <div className={styles['divider']} />
            <div className={styles['content']}>
                {creator ? (
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

                        <Dialog
                            cancelOption={{
                                disabled: isLoading,
                                label: 'キャンセル',
                                onClick: handleDialogClose,
                            }}
                            confirmOption={{
                                disabled: isLoading,
                                label: isLoading ? '更新中...' : '更新',
                                onClick: handleSubmit,
                            }}
                            isOpen={isEditDialogOpen}
                            onClose={handleDialogClose}
                            title="SEO情報編集"
                            wide
                        >
                            <div className={styles['form']}>
                                <div className={styles['field']}>
                                    <label className={styles['label']}>サイトロゴ</label>
                                    {(previewUrl || creator.apiPath) && (
                                        <div className={styles['dialog-logo-preview']}>
                                            <div className={styles['dialog-logo-image']}>
                                                <Image alt={creator.name} src={previewUrl || creator.apiPath} />
                                            </div>
                                        </div>
                                    )}
                                    <FileInput accept="image/*" onChange={handleLogoChange} value={logoFile} />
                                </div>
                                <div className={styles['field']}>
                                    <TextArea
                                        error={errors.introduction?.message}
                                        label="サイト説明"
                                        required
                                        rows={8}
                                        {...register('introduction')}
                                    />
                                </div>
                            </div>
                        </Dialog>
                    </div>
                ) : (
                    <div className={styles['error-message']}>作者情報の読み込みに失敗しました</div>
                )}
            </div>
        </div>
    )
}
