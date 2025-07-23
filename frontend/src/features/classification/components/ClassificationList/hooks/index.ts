import { useState } from 'react'

import { getCategories, postCategory } from '@/apis/category'
import { getTags, postTag } from '@/apis/tag'
import { getTargets, postTarget } from '@/apis/target'
import { IClassification } from '@/features/classification/type'
import { ClassificationType } from '@/types'

import type { IClassificationForm } from '../../../type'

interface UseClassificationListProps {
    classificationType: ClassificationType
    initialItems: IClassification[]
}

export const useClassificationList = ({ initialItems, classificationType }: UseClassificationListProps) => {
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
        switch (classificationType) {
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
        switch (classificationType) {
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

    return {
        items,
        isOpen,
        isSubmitting,
        submitError,
        handleOpenDialog,
        handleCloseDialog,
        handleFormSubmit,
    }
}
