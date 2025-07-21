'use client'

import { Add, Delete } from '@mui/icons-material'
import { useState } from 'react'

import { createCategory, deleteCategory, updateCategory } from '@/apis/category'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { Input } from '@/components/bases/Input'
import { IClassification } from '@/features/classification/type'

import styles from './styles.module.scss'

interface Props {
    items: IClassification[]
    onUpdate: () => void
}

const ExecutionType = {
    Create: 'create',
    Edit: 'edit',
    Delete: 'delete',
} as const
type ExecutionType = (typeof ExecutionType)[keyof typeof ExecutionType]

export const CategoryList = ({ items, onUpdate }: Props) => {
    const [dialogOpen, setDialogOpen] = useState<boolean>(false)
    const [executionType, setExecutionType] = useState<ExecutionType>(ExecutionType.Create)
    const [selectedItem, setSelectedItem] = useState<IClassification | null>(null)
    const [name, setName] = useState<string>('')
    const [nameError, setNameError] = useState<string>('')
    const [loading, setLoading] = useState<boolean>(false)

    const openDialog = (type: ExecutionType, item?: IClassification) => {
        setExecutionType(type)
        setSelectedItem(item || null)
        setName(item?.name || '')
        setNameError('')
        setDialogOpen(true)
    }

    const closeDialog = () => {
        setDialogOpen(false)
        setExecutionType(ExecutionType.Create)
        setSelectedItem(null)
        setName('')
        setNameError('')
    }

    const validateName = (value: string): boolean => {
        if (!value.trim()) {
            setNameError('名前を入力してください')
            return false
        }
        if (value.length > 20) {
            setNameError('20文字以内で入力してください')
            return false
        }
        setNameError('')
        return true
    }

    const handleSubmit = async () => {
        if (executionType === ExecutionType.Delete) {
            await handleDelete()
        } else if (executionType === ExecutionType.Edit) {
            await handleEdit()
        } else {
            await handleCreate()
        }
    }

    const handleCreate = async () => {
        if (!validateName(name)) return

        setLoading(true)
        try {
            await createCategory({ name })
            closeDialog()
            onUpdate()
        } catch (error) {
            console.error('カテゴリの作成エラー:', error)
        } finally {
            setLoading(false)
        }
    }

    const handleEdit = async () => {
        if (!selectedItem || !validateName(name)) return

        setLoading(true)
        try {
            await updateCategory(selectedItem.uuid, { name })
            closeDialog()
            onUpdate()
        } catch (error) {
            console.error('カテゴリの編集エラー:', error)
        } finally {
            setLoading(false)
        }
    }

    const handleDelete = async () => {
        if (!selectedItem) return

        setLoading(true)
        try {
            await deleteCategory(selectedItem.uuid)
            closeDialog()
            onUpdate()
        } catch (error) {
            console.error('カテゴリの削除エラー:', error)
        } finally {
            setLoading(false)
        }
    }

    const getDialogTitle = () => {
        switch (executionType) {
            case ExecutionType.Create:
                return 'カテゴリーを追加'
            case ExecutionType.Edit:
                return 'カテゴリーを編集'
            case ExecutionType.Delete:
                return 'カテゴリーを削除'
            default:
                return ''
        }
    }

    return (
        <div className={styles['classification-list']}>
            <div className={styles['list-content']}>
                {items.length === 0 ? (
                    <div className={styles['empty-message']}>登録されていません</div>
                ) : (
                    items.map((item) => (
                        <div className={styles['list-item']} key={item.uuid} onClick={() => openDialog(ExecutionType.Edit, item)}>
                            <div className={styles['item-content']}>
                                <span className={styles['item-name']}>{item.name}</span>
                                <div className={styles['item-actions']}>
                                    <Delete
                                        className={styles['icon-button']}
                                        onClick={(e) => {
                                            e.stopPropagation()
                                            openDialog(ExecutionType.Delete, item)
                                        }}
                                    />
                                </div>
                            </div>
                        </div>
                    ))
                )}
            </div>
            <div className={styles['add-button-container']}>
                <Button onClick={() => openDialog(ExecutionType.Create)}>
                    <div className={styles['add-button-content']}>
                        <Add className={styles['add-icon']} />
                        追加
                    </div>
                </Button>
            </div>

            <Dialog
                cancelOption={{
                    label: 'キャンセル',
                    onClick: loading ? () => {} : closeDialog,
                }}
                confirmOption={{
                    label: loading
                        ? '処理中...'
                        : executionType === ExecutionType.Delete
                          ? '削除'
                          : executionType === ExecutionType.Edit
                            ? '更新'
                            : '追加',
                    onClick: loading || (executionType !== ExecutionType.Delete && !name.trim()) ? () => {} : handleSubmit,
                }}
                isOpen={dialogOpen}
                onClose={loading ? () => {} : closeDialog}
                title={getDialogTitle()}
            >
                {executionType === ExecutionType.Delete ? (
                    <p>「{selectedItem?.name}」を削除しますか？</p>
                ) : (
                    <Input
                        error={nameError}
                        label="名前"
                        onChange={(e) => setName(e.target.value)}
                        placeholder="名前を入力してください"
                        required
                        value={name}
                    />
                )}
            </Dialog>
        </div>
    )
}
