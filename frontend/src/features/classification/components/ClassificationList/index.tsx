'use client'

import { Add, Delete } from '@mui/icons-material'
import { useState } from 'react'

import { Button } from '@/components/bases/Button'
import { Pagination } from '@/components/bases/Pagination'
import { IClassification } from '@/features/classification/type'

import styles from './styles.module.scss'

interface Props {
    items: IClassification[]
}

export const ClassificationList = ({ items }: Props) => {
    const [currentPage, setCurrentPage] = useState<number>(1)
    const itemsPerPage = 10

    const totalPages = Math.ceil(items.length / itemsPerPage)
    const startIndex = (currentPage - 1) * itemsPerPage
    const currentItems = items.slice(startIndex, startIndex + itemsPerPage)

    const handlePageChange = (page: number) => {
        setCurrentPage(page)
    }

    return (
        <div className={styles['classification-list']}>
            <div className={styles['list-content']}>
                {currentItems.length === 0 ? (
                    <div className={styles['empty-message']}>登録されていません</div>
                ) : (
                    <div className={styles['item-list']}>
                        {currentItems.map((item) => (
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
            <div className={styles['pagination-container']}>
                <Pagination currentPage={currentPage} onPageChange={handlePageChange} totalPages={totalPages} />
                <div className={styles['add-button-container']}>
                    <Button onClick={() => {}}>
                        <div className={styles['add-button-content']}>
                            <Add className={styles['add-icon']} />
                            追加
                        </div>
                    </Button>
                </div>
            </div>
        </div>
    )
}
