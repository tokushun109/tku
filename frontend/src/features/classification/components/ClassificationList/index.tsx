'use client'

import { Add, Delete } from '@mui/icons-material'
import { Virtuoso } from 'react-virtuoso'

import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { IClassification } from '@/features/classification/type'
import { ClassificationLabel, ClassificationType } from '@/types'

import { ClassificationFormDialog } from '../ClassificationFormDialog'
import { useClassificationList } from './hooks'
import styles from './styles.module.scss'

interface Props {
    classificationType: ClassificationType
    initialItems: IClassification[]
}

export const ClassificationList = ({ initialItems, classificationType }: Props) => {
    const {
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
    } = useClassificationList({
        initialItems,
        classificationType,
    })

    const classificationName = ClassificationLabel[classificationType]

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
                            <div
                                className={styles['list-item']}
                                onClick={() => {
                                    handleOpenDialog(item)
                                }}
                            >
                                <div className={styles['item-content']}>
                                    <span className={styles['item-name']}>{item.name}</span>
                                    <div className={styles['item-actions']}>
                                        <Delete
                                            className={styles['icon-button']}
                                            onClick={(e) => {
                                                e.stopPropagation()
                                                handleOpenDeleteDialog(item)
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
                <Button
                    onClick={() => {
                        handleOpenDialog(null)
                    }}
                >
                    <div className={styles['add-button-content']}>
                        <Add className={styles['add-icon']} />
                        追加
                    </div>
                </Button>
            </div>
            <ClassificationFormDialog
                classificationType={classificationType}
                isOpen={isOpen}
                isSubmitting={isSubmitting}
                onClose={handleCloseDialog}
                onSubmit={handleFormSubmit}
                submitError={submitError}
                updateItem={targetItem}
            />
            <Dialog
                cancelOption={{
                    label: 'キャンセル',
                    onClick: handleCloseDeleteDialog,
                }}
                confirmOption={{
                    label: '削除',
                    onClick: handleConfirmDelete,
                }}
                isOpen={isDeleteDialogOpen}
                onClose={handleCloseDeleteDialog}
                title="削除確認"
            >
                {targetItem && (
                    <>
                        <p>
                            {classificationName}「{targetItem.name}」を削除しますか？
                        </p>
                        <p>この操作は取り消せません。</p>
                    </>
                )}
            </Dialog>
        </div>
    )
}
