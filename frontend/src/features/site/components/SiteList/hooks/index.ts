import { useState } from 'react'
import { toast } from 'sonner'

import { deleteSalesSite, getSalesSiteList, postSalesSite, putSalesSite } from '@/apis/salesSite'
import { deleteSns, getSnsList, postSns, putSns } from '@/apis/sns'
import { SiteLabel, SiteType } from '@/types'

import { ISite, ISiteForm } from '../../../type'

interface UseSiteListProps {
    initialItems: ISite[]
    siteType: SiteType
}

export const useSiteList = ({ initialItems, siteType }: UseSiteListProps) => {
    const [items, setItems] = useState<ISite[]>(initialItems)
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [submitError, setSubmitError] = useState<string | null>(null)
    const [targetItem, setTargetItem] = useState<ISite | null>(null)
    const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState<boolean>(false)

    const siteName = SiteLabel[siteType]

    const handleOpenDialog = (item: ISite | null) => {
        setIsOpen(true)
        setSubmitError(null)
        setTargetItem(item)
    }

    const handleCloseDialog = () => {
        setIsOpen(false)
        setSubmitError(null)
    }

    // typeに応じてAPIを切り替える関数（追加）
    const postSite = async (data: ISiteForm) => {
        switch (siteType) {
            case SiteType.Sns:
                return await postSns({ form: data })
            case SiteType.SalesSite:
                return await postSalesSite({ form: data })
            default:
                throw new Error('不正なタイプです')
        }
    }

    // typeに応じてAPIを切り替える関数（更新）
    const putSite = async (data: ISiteForm, uuid: string) => {
        switch (siteType) {
            case SiteType.Sns:
                return await putSns({ form: data, uuid })
            case SiteType.SalesSite:
                return await putSalesSite({ form: data, uuid })
            default:
                throw new Error('不正なタイプです')
        }
    }

    // typeに応じてAPIを切り替える関数（削除）
    const deleteSite = async (uuid: string) => {
        switch (siteType) {
            case SiteType.Sns:
                return await deleteSns({ uuid })
            case SiteType.SalesSite:
                return await deleteSalesSite({ uuid })
            default:
                throw new Error('不正なタイプです')
        }
    }

    const fetchSites = async (): Promise<ISite[]> => {
        switch (siteType) {
            case SiteType.Sns:
                return await getSnsList()
            case SiteType.SalesSite:
                return await getSalesSiteList()
            default:
                throw new Error('不正なタイプです')
        }
    }

    const handleFormSubmit = async (data: ISiteForm) => {
        try {
            setIsSubmitting(true)
            setSubmitError(null)

            if (targetItem && targetItem.uuid) {
                // 更新処理
                await putSite(data, targetItem.uuid)
                toast.success(`${siteName}「${data.name}」を更新しました`)
            } else {
                // 追加処理
                await postSite(data)
                toast.success(`${siteName}「${data.name}」を追加しました`)
            }

            // 成功後、一覧を再取得して状態を更新
            const updatedItems = await fetchSites()
            setItems(updatedItems)

            handleCloseDialog()
        } catch {
            const errorMessage = targetItem
                ? `${siteName}の更新に失敗しました。もう一度お試しください。`
                : `${siteName}の追加に失敗しました。もう一度お試しください。`
            setSubmitError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setIsSubmitting(false)
        }
    }

    const handleOpenDeleteDialog = (item: ISite) => {
        setTargetItem(item)
        setIsDeleteDialogOpen(true)
    }

    const handleCloseDeleteDialog = () => {
        setIsDeleteDialogOpen(false)
        setTargetItem(null)
    }

    const handleConfirmDelete = async () => {
        if (!targetItem || !targetItem.uuid) return

        try {
            await deleteSite(targetItem.uuid)
            toast.success(`${siteName}「${targetItem.name}」を削除しました`)

            // 成功後、一覧を再取得して状態を更新
            const updatedItems = await fetchSites()
            setItems(updatedItems)

            handleCloseDeleteDialog()
        } catch {
            const errorMessage = `${siteName}の削除に失敗しました。もう一度お試しください。`
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
