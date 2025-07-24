import { useState } from 'react'
import { toast } from 'sonner'

import { deleteCategory, getCategories, postCategory, putCategory } from '@/apis/category'
import { deleteTag, getTags, postTag, putTag } from '@/apis/tag'
import { deleteTarget, getTargets, postTarget, putTarget } from '@/apis/target'
import { IClassification } from '@/features/classification/type'
import { ClassificationType, ClassificationLabel } from '@/types'

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
    const [updateItem, setUpdateItem] = useState<IClassification | null>(null)
    const [isDeleting, setIsDeleting] = useState<boolean>(false)

    const classificationName = ClassificationLabel[classificationType]

    const handleOpenDialog = (item: IClassification | null) => {
        setIsOpen(true)
        setSubmitError(null)
        setUpdateItem(item)
    }

    const handleCloseDialog = () => {
        setIsOpen(false)
        setSubmitError(null)
    }

    // typeに応じてAPIを切り替える関数（追加）
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

    // typeに応じてAPIを切り替える関数（更新）
    const putClassification = async (data: IClassificationForm, uuid: string) => {
        switch (classificationType) {
            case ClassificationType.Category:
                return await putCategory({ form: data, uuid })
            case ClassificationType.Target:
                return await putTarget({ form: data, uuid })
            case ClassificationType.Tag:
                return await putTag({ form: data, uuid })
            default:
                throw new Error('不正なタイプです')
        }
    }

    // typeに応じてAPIを切り替える関数（削除）
    const deleteClassification = async (uuid: string) => {
        switch (classificationType) {
            case ClassificationType.Category:
                return await deleteCategory({ uuid })
            case ClassificationType.Target:
                return await deleteTarget({ uuid })
            case ClassificationType.Tag:
                return await deleteTag({ uuid })
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

            if (updateItem) {
                // 更新処理
                await putClassification(data, updateItem.uuid)
                toast.success(`${classificationName}「${data.name}」を更新しました`)
            } else {
                // 追加処理
                await postClassification(data)
                toast.success(`${classificationName}「${data.name}」を追加しました`)
            }

            // 成功後、一覧を再取得して状態を更新
            const updatedItems = await fetchClassifications()
            setItems(updatedItems)

            handleCloseDialog()
        } catch {
            const errorMessage = updateItem
                ? `${classificationName}の更新に失敗しました。もう一度お試しください。`
                : `${classificationName}の追加に失敗しました。もう一度お試しください。`
            setSubmitError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setIsSubmitting(false)
        }
    }

    const handleDelete = async (item: IClassification) => {
        if (!window.confirm(`${classificationName}「${item.name}」を削除しますか？\nこの操作は取り消せません。`)) {
            return
        }

        try {
            setIsDeleting(true)
            await deleteClassification(item.uuid)
            toast.success(`${classificationName}「${item.name}」を削除しました`)

            // 成功後、一覧を再取得して状態を更新
            const updatedItems = await fetchClassifications()
            setItems(updatedItems)
        } catch {
            const errorMessage = `${classificationName}の削除に失敗しました。もう一度お試しください。`
            toast.error(errorMessage)
        } finally {
            setIsDeleting(false)
        }
    }

    return {
        items,
        isOpen,
        isSubmitting,
        submitError,
        updateItem,
        isDeleting,
        handleOpenDialog,
        handleCloseDialog,
        handleFormSubmit,
        handleDelete,
    }
}
