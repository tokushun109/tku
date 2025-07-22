'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { Add, Delete } from '@mui/icons-material'
import { useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'
import { Virtuoso } from 'react-virtuoso'

import { getCategories, postCategory } from '@/apis/category'
import { getTags, postTag } from '@/apis/tag'
import { getTargets, postTarget } from '@/apis/target'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { Form } from '@/components/bases/Form'
import { Input } from '@/components/bases/Input'
import { Message, MessageType } from '@/components/bases/Message'
import { IClassification } from '@/features/classification/type'
import { ClassificationLabel, ClassificationType } from '@/types'

import styles from './styles.module.scss'
import { ClassificationSchema } from '../../classification/schema'
import { IClassificationForm } from '../../classification/type'

interface Props {
    initialItems: IClassification[]
    type: ClassificationType
}

export const ClassificationList = ({ initialItems, type }: Props) => {
    const [items, setItems] = useState<IClassification[]>(initialItems)
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [submitError, setSubmitError] = useState<string | null>(null)

    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<IClassificationForm>({
        mode: 'onChange',
        resolver: zodResolver(ClassificationSchema),
    })

    const handleOpenDialog = () => {
        setIsOpen(true)
        reset()
        setSubmitError(null)
    }

    const handleCloseDialog = () => {
        setIsOpen(false)
        reset()
        setSubmitError(null)
    }

    // typeに応じてAPIを切り替える関数
    const postClassification = async (data: IClassificationForm) => {
        switch (type) {
            case ClassificationType.Category:
                return await postCategory({ form: data })
            case ClassificationType.Target:
                return await postTarget({ form: data })
            case ClassificationType.Tag:
                return await postTag({ form: data })
            default:
                throw new Error('不正なタイプです')
        }
    }

    const fetchClassifications = async (): Promise<IClassification[]> => {
        switch (type) {
            case ClassificationType.Category:
                return await getCategories({ mode: 'all' })
            case ClassificationType.Target:
                return await getTargets({ mode: 'all' })
            case ClassificationType.Tag:
                return await getTags()
            default:
                throw new Error('不正なタイプです')
        }
    }

    const onSubmit: SubmitHandler<IClassificationForm> = async (data) => {
        try {
            setIsSubmitting(true)
            setSubmitError(null)

            await postClassification(data)

            // 追加成功後、一覧を再取得して状態を更新
            const updatedItems = await fetchClassifications()
            setItems(updatedItems)

            handleCloseDialog()
        } catch {
            setSubmitError('送信中にエラーが発生しました。もう一度お試しください。')
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <div className={styles['classification-list']}>
            <div className={styles['list-content']}>
                {items.length === 0 ? (
                    <div className={styles['empty-message']}>登録されていません</div>
                ) : (
                    <Virtuoso
                        computeItemKey={(_index, item) => item.uuid}
                        data={items}
                        itemContent={(_index, item) => (
                            <div className={styles['list-item']} onClick={() => {}}>
                                <div className={styles['item-content']}>
                                    <span className={styles['item-name']}>{item.name}</span>
                                    <div className={styles['item-actions']}>
                                        <Delete
                                            className={styles['icon-button']}
                                            onClick={(e) => {
                                                e.stopPropagation()
                                            }}
                                        />
                                    </div>
                                </div>
                            </div>
                        )}
                        style={{ height: '640px' }}
                    />
                )}
            </div>
            <div className={styles['add-button-container']}>
                <Button onClick={handleOpenDialog}>
                    <div className={styles['add-button-content']}>
                        <Add className={styles['add-icon']} />
                        追加
                    </div>
                </Button>
            </div>
            <Dialog
                confirmOption={{
                    label: isSubmitting ? '送信中...' : '追加',
                    onClick: handleSubmit(onSubmit),
                    disabled: isSubmitting,
                }}
                isOpen={isOpen}
                onClose={handleCloseDialog}
                title={`${ClassificationLabel[type]}を追加`}
                wide
            >
                {submitError && <Message type={MessageType.Error}>{submitError}</Message>}
                <Form noValidate onSubmit={handleSubmit(onSubmit)}>
                    <Input
                        {...register('name')}
                        error={errors.name?.message}
                        id="name"
                        label={`${ClassificationLabel[type]}名`}
                        placeholder={`テスト${ClassificationLabel[type]}`}
                        required
                        type="text"
                    />
                </Form>
            </Dialog>
        </div>
    )
}
