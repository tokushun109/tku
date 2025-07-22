'use client'

import { Add, Delete } from '@mui/icons-material'
import { useState } from 'react'
import { Virtuoso } from 'react-virtuoso'

import { getCategories, postCategory } from '@/apis/category'
import { getTags, postTag } from '@/apis/tag'
import { getTargets, postTarget } from '@/apis/target'
import { Button } from '@/components/bases/Button'
import { IClassification } from '@/features/classification/type'
import { ClassificationType } from '@/types'

import { ClassificationFormDialog } from '../ClassificationFormDialog'
import styles from './styles.module.scss'

import type { IClassificationForm } from '../../type'

interface Props {
    initialItems: IClassification[]
    type: ClassificationType
}

export const ClassificationList = ({ initialItems, type }: Props) => {
    const [items, setItems] = useState<IClassification[]>(initialItems)
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [submitError, setSubmitError] = useState<string | null>(null)

    const handleOpenDialog = () => {
        setIsOpen(true)
        setSubmitError(null)
    }

    const handleCloseDialog = () => {
        setIsOpen(false)
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

    const handleFormSubmit = async (data: IClassificationForm) => {
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
            <ClassificationFormDialog
                isOpen={isOpen}
                isSubmitting={isSubmitting}
                onClose={handleCloseDialog}
                onSubmit={handleFormSubmit}
                submitError={submitError}
                type={type}
            />
        </div>
    )
}
