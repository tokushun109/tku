'use client'

import { Add, Delete } from '@mui/icons-material'

import { Button } from '@/components/bases/Button'
import { IClassification } from '@/features/classification/type'

import styles from './styles.module.scss'

interface Props {
    items: IClassification[]
}

export const ClassificationList = ({ items }: Props) => {
    return (
        <div className={styles['classification-list']}>
            <div className={styles['list-content']}>
                {items.length === 0 ? (
                    <div className={styles['empty-message']}>登録されていません</div>
                ) : (
                    <div className={styles['item-list']}>
                        {items.map((item) => (
                            <div className={styles['list-item']} key={item.uuid} onClick={() => {}}>
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
                        ))}
                    </div>
                )}
            </div>
            <div className={styles['add-button-container']}>
                <Button onClick={() => {}}>
                    <div className={styles['add-button-content']}>
                        <Add className={styles['add-icon']} />
                        追加
                    </div>
                </Button>
            </div>
        </div>
    )
}
