'use client'

import { Add, Delete } from '@mui/icons-material'
import { useState } from 'react'

import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { Input } from '@/components/bases/Input'
import { IClassification } from '@/features/classification/type'
import { ClassificationType } from '@/types'

import styles from './styles.module.scss'

interface Props {
    items: IClassification[]
    onUpdate: () => void
    type: ClassificationType
}

type ExecutionType = 'create' | 'edit' | 'delete'

export const ClassificationList = ({ items, type, onUpdate }: Props) => {
    const [dialogOpen, setDialogOpen] = useState<boolean>(false)
    const [executionType, setExecutionType] = useState<ExecutionType>('create')
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
        setExecutionType('create')
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
        if (executionType === 'delete') {
            await handleDelete()
        } else if (executionType === 'edit') {
            await handleEdit()
        } else {
            await handleCreate()
        }
    }

    const handleCreate = async () => {
        if (!validateName(name)) return

        setLoading(true)
        try {
            if (type === 'category') {
                const { createCategory } = await import('@/apis/category')
                await createCategory({ name })
            } else if (type === 'target') {
                const { createTarget } = await import('@/apis/target')
                await createTarget({ name })
            } else if (type === 'tag') {
                const { createTag } = await import('@/apis/tag')
                await createTag({ name })
            }

            closeDialog()
            onUpdate()
        } catch (error) {
            console.error('作成エラー:', error)
        } finally {
            setLoading(false)
        }
    }

    const handleEdit = async () => {
        if (!selectedItem || !validateName(name)) return

        setLoading(true)
        try {
            if (type === 'category') {
                const { updateCategory } = await import('@/apis/category')
                await updateCategory(selectedItem.uuid, { name })
            } else if (type === 'target') {
                const { updateTarget } = await import('@/apis/target')
                await updateTarget(selectedItem.uuid, { name })
            } else if (type === 'tag') {
                const { updateTag } = await import('@/apis/tag')
                await updateTag(selectedItem.uuid, { name })
            }

            closeDialog()
            onUpdate()
        } catch (error) {
            console.error('編集エラー:', error)
        } finally {
            setLoading(false)
        }
    }

    const handleDelete = async () => {
        if (!selectedItem) return

        setLoading(true)
        try {
            if (type === 'category') {
                const { deleteCategory } = await import('@/apis/category')
                await deleteCategory(selectedItem.uuid)
            } else if (type === 'target') {
                const { deleteTarget } = await import('@/apis/target')
                await deleteTarget(selectedItem.uuid)
            } else if (type === 'tag') {
                const { deleteTag } = await import('@/apis/tag')
                await deleteTag(selectedItem.uuid)
            }

            closeDialog()
            onUpdate()
        } catch (error) {
            console.error('削除エラー:', error)
        } finally {
            setLoading(false)
        }
    }

    const getDialogTitle = () => {
        const typeName = type === 'category' ? 'カテゴリー' : type === 'target' ? 'ターゲット' : 'タグ'
        switch (executionType) {
            case 'create':
                return `${typeName}を追加`
            case 'edit':
                return `${typeName}を編集`
            case 'delete':
                return `${typeName}を削除`
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
                    <div className={styles['item-list']}>
                        {items.map((item) => (
                            <div className={styles['list-item']} key={item.uuid} onClick={() => openDialog('edit', item)}>
                                <div className={styles['item-content']}>
                                    <span className={styles['item-name']}>{item.name}</span>
                                    <div className={styles['item-actions']}>
                                        <Delete
                                            className={styles['icon-button']}
                                            onClick={(e) => {
                                                e.stopPropagation()
                                                openDialog('delete', item)
                                            }}
                                        />
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
            <div className={styles['add-button-container']}>
                <Button onClick={() => openDialog('create')}>
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
                    label: loading ? '処理中...' : executionType === 'delete' ? '削除' : executionType === 'edit' ? '更新' : '追加',
                    onClick: loading || (executionType !== 'delete' && !name.trim()) ? () => {} : handleSubmit,
                }}
                isOpen={dialogOpen}
                onClose={loading ? () => {} : closeDialog}
                title={getDialogTitle()}
            >
                {executionType === 'delete' ? (
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
