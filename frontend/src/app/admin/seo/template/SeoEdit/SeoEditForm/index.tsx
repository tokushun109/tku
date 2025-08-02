'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'

import { Button } from '@/components/bases/Button'
import { Input } from '@/components/bases/Input'
import { TextArea } from '@/components/bases/TextArea'
import { ICreator } from '@/features/creator/type'
import { SeoEditSchema } from '@/features/seo/schema'
import { ISeoEdit } from '@/features/seo/type'

import styles from './styles.module.scss'

interface Props {
    creator: ICreator
    isLoading: boolean
    onCancel: () => void
    onSubmit: (_name: string, _introduction: string) => void
}

export const SeoEditForm = ({ creator, isLoading, onCancel, onSubmit }: Props) => {
    const {
        formState: { errors },
        handleSubmit,
        register,
    } = useForm<ISeoEdit>({
        defaultValues: {
            introduction: creator.introduction,
            name: creator.name,
        },
        resolver: zodResolver(SeoEditSchema),
    })

    const handleFormSubmit = (data: ISeoEdit) => {
        onSubmit(data.name, data.introduction)
    }

    return (
        <form className={styles['form']} onSubmit={handleSubmit(handleFormSubmit)}>
            <div className={styles['field']}>
                <Input error={errors.name?.message} label="サイト名" required {...register('name')} />
            </div>
            <div className={styles['field']}>
                <TextArea error={errors.introduction?.message} label="サイト説明" required rows={8} {...register('introduction')} />
            </div>
            <div className={styles['actions']}>
                <Button disabled={isLoading} onClick={onCancel} type="button">
                    キャンセル
                </Button>
                <Button disabled={isLoading} type="submit">
                    {isLoading ? '更新中...' : '更新'}
                </Button>
            </div>
        </form>
    )
}
