'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useEffect } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'

import { Dialog } from '@/components/bases/Dialog'
import { Form } from '@/components/bases/Form'
import { Input } from '@/components/bases/Input'
import { Message, MessageType } from '@/components/bases/Message'
import { SiteLabel, SiteType } from '@/types'

import { SiteSchema } from '../../schema'
import { ISite, ISiteForm } from '../../type'

interface Props {
    isOpen: boolean
    isSubmitting: boolean
    onClose: () => void
    onSubmit: (_data: ISiteForm) => Promise<void>
    siteType: SiteType
    submitError: string | null
    updateItem: ISite | null
}

export const SiteFormDialog = ({ isOpen, isSubmitting, onClose, onSubmit, submitError, siteType, updateItem }: Props) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<ISiteForm>({
        resolver: zodResolver(SiteSchema),
        defaultValues: {
            name: '',
            url: '',
        },
    })

    // updateItemが変更されたときにフォームをリセット
    useEffect(() => {
        if (isOpen) {
            reset({
                name: updateItem?.name || '',
                url: updateItem?.url || '',
            })
        }
    }, [updateItem, isOpen, reset])

    const handleClose = () => {
        reset()
        onClose()
    }

    const handleFormSubmit: SubmitHandler<ISiteForm> = async (_data) => {
        await onSubmit(_data)
        reset()
    }

    const isEdit = updateItem !== null

    return (
        <Dialog
            confirmOption={{
                label: isSubmitting ? '送信中...' : isEdit ? '更新' : '追加',
                onClick: handleSubmit(handleFormSubmit),
                disabled: isSubmitting,
            }}
            isOpen={isOpen}
            onClose={handleClose}
            title={`${SiteLabel[siteType]}を${isEdit ? '編集' : '追加'}`}
            wide
        >
            {submitError && <Message type={MessageType.Error}>{submitError}</Message>}
            <Form noValidate onSubmit={handleSubmit(handleFormSubmit)}>
                <Input
                    {...register('name')}
                    error={errors.name?.message}
                    id="name"
                    label={`${SiteLabel[siteType]}名`}
                    placeholder={`テスト${SiteLabel[siteType]}`}
                    required
                    type="text"
                />
                <Input {...register('url')} error={errors.url?.message} id="url" label="URL" placeholder="https://example.com" required type="url" />
            </Form>
        </Dialog>
    )
}
