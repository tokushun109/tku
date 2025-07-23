'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { SubmitHandler, useForm } from 'react-hook-form'

import { Dialog } from '@/components/bases/Dialog'
import { Form } from '@/components/bases/Form'
import { Input } from '@/components/bases/Input'
import { Message, MessageType } from '@/components/bases/Message'
import { ClassificationLabel, ClassificationType } from '@/types'

import { ClassificationSchema } from '../../classification/schema'

import type { IClassificationForm } from '../../type'

interface Props {
    classificationType: ClassificationType
    isOpen: boolean
    isSubmitting: boolean
    onClose: () => void
    onSubmit: (_data: IClassificationForm) => Promise<void>
    submitError: string | null
}

export const ClassificationFormDialog = ({ isOpen, isSubmitting, onClose, onSubmit, submitError, classificationType }: Props) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<IClassificationForm>({
        resolver: zodResolver(ClassificationSchema),
    })

    const handleClose = () => {
        reset()
        onClose()
    }

    const handleFormSubmit: SubmitHandler<IClassificationForm> = async (_data) => {
        await onSubmit(_data)
        reset()
    }

    return (
        <Dialog
            confirmOption={{
                label: isSubmitting ? '送信中...' : '追加',
                onClick: handleSubmit(handleFormSubmit),
                disabled: isSubmitting,
            }}
            isOpen={isOpen}
            onClose={handleClose}
            title={`${ClassificationLabel[classificationType]}を追加`}
            wide
        >
            {submitError && <Message type={MessageType.Error}>{submitError}</Message>}
            <Form noValidate onSubmit={handleSubmit(handleFormSubmit)}>
                <Input
                    {...register('name')}
                    error={errors.name?.message}
                    id="name"
                    label={`${ClassificationLabel[classificationType]}名`}
                    placeholder={`テスト${ClassificationLabel[classificationType]}`}
                    required
                    type="text"
                />
            </Form>
        </Dialog>
    )
}
