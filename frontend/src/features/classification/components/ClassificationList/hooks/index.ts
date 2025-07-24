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
    const [targetItem, setTargetItem] = useState<IClassification | null>(null)
    const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState<boolean>(false)

    const classificationName = ClassificationLabel[classificationType]

    const handleOpenDialog = (item: IClassification | null) => {
        setIsOpen(true)
        setSubmitError(null)
        setTargetItem(item)
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

            if (targetItem) {
                // 更新処理
                await putClassification(data, targetItem.uuid)
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
            const errorMessage = targetItem
                ? `${classificationName}の更新に失敗しました。もう一度お試しください。`
                : `${classificationName}の追加に失敗しました。もう一度お試しください。`
            setSubmitError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setIsSubmitting(false)
        }
    }

    const handleOpenDeleteDialog = (item: IClassification) => {
        setTargetItem(item)
        setIsDeleteDialogOpen(true)
    }

    const handleCloseDeleteDialog = () => {
        setIsDeleteDialogOpen(false)
        setTargetItem(null)
    }

    const handleConfirmDelete = async () => {
        if (!targetItem) return

        try {
            await deleteClassification(targetItem.uuid)
            toast.success(`${classificationName}「${targetItem.name}」を削除しました`)

            // 成功後、一覧を再取得して状態を更新
            const updatedItems = await fetchClassifications()
            setItems(updatedItems)

            handleCloseDeleteDialog()
        } catch {
            const errorMessage = `${classificationName}の削除に失敗しました。もう一度お試しください。`
            toast.error(errorMessage)
        }
    }

    return {
        items,
        isOpen,
        isSubmitting,
        submitError,
        targetItem,
        isDeleteDialogOpen,
        handleOpenDialog,
        handleCloseDialog,
        handleFormSubmit,
        handleOpenDeleteDialog,
        handleCloseDeleteDialog,
        handleConfirmDelete,
    }
}
