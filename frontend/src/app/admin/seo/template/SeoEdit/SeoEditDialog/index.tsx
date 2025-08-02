'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'

import { updateCreator, updateCreatorLogo } from '@/apis/creator'
import { Dialog } from '@/components/bases/Dialog'
import { FileInput } from '@/components/bases/FileInput'
import { Image } from '@/components/bases/Image'
import { TextArea } from '@/components/bases/TextArea'
import { ICreator } from '@/features/creator/type'
import { SeoEditSchema } from '@/features/seo/schema'
import { ISeoEdit } from '@/features/seo/type'

import styles from './styles.module.scss'

interface Props {
    creator: ICreator
    isOpen: boolean
    onClose: () => void
    onUpdate: () => Promise<void>
}

export const SeoEditDialog = ({ creator, isOpen, onClose, onUpdate }: Props) => {
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [logoFile, setLogoFile] = useState<File | null>(null)
    const [previewUrl, setPreviewUrl] = useState<string | null>(null)

    const {
        formState: { errors },
        register,
        watch,
    } = useForm<ISeoEdit>({
        defaultValues: {
            introduction: creator.introduction,
        },
        resolver: zodResolver(SeoEditSchema),
    })

    const watchedIntroduction = watch('introduction')

    const handleSubmit = async () => {
        setIsLoading(true)
        try {
            // テキスト情報を更新
            await updateCreator({ name: creator.name, introduction: watchedIntroduction || '' })

            // ロゴファイルがある場合はロゴも更新
            if (logoFile) {
                await updateCreatorLogo(logoFile)
            }

            await onUpdate()
            toast.success('SEO情報を更新しました')
            onClose()
            setLogoFile(null)
            if (previewUrl) {
                URL.revokeObjectURL(previewUrl)
                setPreviewUrl(null)
            }
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

    const handleClose = () => {
        onClose()
        setLogoFile(null)
        if (previewUrl) {
            URL.revokeObjectURL(previewUrl)
            setPreviewUrl(null)
        }
    }

    return (
        <Dialog
            cancelOption={{
                disabled: isLoading,
                label: 'キャンセル',
                onClick: handleClose,
            }}
            confirmOption={{
                disabled: isLoading,
                label: isLoading ? '更新中...' : '更新',
                onClick: handleSubmit,
            }}
            isOpen={isOpen}
            onClose={handleClose}
            title="SEO情報編集"
            wide
        >
            <div className={styles['form']}>
                <div className={styles['field']}>
                    <label className={styles['label']}>サイトロゴ</label>
                    {(previewUrl || creator.apiPath) && (
                        <div className={styles['logo-preview']}>
                            <div className={styles['logo-image']}>
                                <Image alt={creator.name} src={previewUrl || creator.apiPath} />
                            </div>
                        </div>
                    )}
                    <FileInput accept="image/*" onChange={handleLogoChange} value={logoFile} />
                </div>
                <div className={styles['field']}>
                    <TextArea error={errors.introduction?.message} label="サイト説明" required rows={8} {...register('introduction')} />
                </div>
            </div>
        </Dialog>
    )
}
