'use client'

import { Add, Delete } from '@mui/icons-material'
import { Virtuoso } from 'react-virtuoso'

import { Button } from '@/components/bases/Button'
import { IClassification } from '@/features/classification/type'
import { ClassificationType } from '@/types'

import { ClassificationFormDialog } from '../ClassificationFormDialog'
import { useClassificationList } from './hooks'
import styles from './styles.module.scss'

interface Props {
    classificationType: ClassificationType
    initialItems: IClassification[]
}

export const ClassificationList = ({ initialItems, classificationType }: Props) => {
    const { items, isOpen, isSubmitting, submitError, handleOpenDialog, handleCloseDialog, handleFormSubmit } = useClassificationList({
        initialItems,
        classificationType,
    })

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
            />
        </div>
    )
}
