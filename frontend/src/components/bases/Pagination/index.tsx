import React from 'react'

import styles from './styles.module.scss'

interface Props {
    currentPage: number
    delta?: number
    onPageChange: (_page: number) => void
    totalPages: number
}

export const Pagination = ({ currentPage, totalPages, onPageChange, delta = 2 }: Props) => {
    if (totalPages <= 1) return null

    const getVisiblePages = () => {
        const range = []
        const rangeWithDots = []

        // 常に表示する最初と最後のページ
        for (let i = Math.max(2, currentPage - delta); i <= Math.min(totalPages - 1, currentPage + delta); i++) {
            range.push(i)
        }

        // 最初のページを追加
        if (currentPage - delta > 2) {
            rangeWithDots.push(1, '...')
        } else {
            rangeWithDots.push(1)
        }

        // 中間のページを追加
        rangeWithDots.push(...range)

        // 最後のページを追加
        if (currentPage + delta < totalPages - 1) {
            rangeWithDots.push('...', totalPages)
        } else if (totalPages > 1) {
            rangeWithDots.push(totalPages)
        }

        return rangeWithDots
    }

    const visiblePages = getVisiblePages()

    return (
        <div className={styles['pagination']}>
            <button className={styles['pagination-button']} disabled={currentPage === 1} onClick={() => onPageChange(currentPage - 1)}>
                前へ
            </button>
            {visiblePages.map((page, index) => {
                if (page === '...') {
                    return (
                        <span className={styles['pagination-dots']} key={`dots-${index}`}>
                            ...
                        </span>
                    )
                }
                return (
                    <button
                        className={`${styles['pagination-button']} ${page === currentPage ? styles['active'] : ''}`}
                        key={page}
                        onClick={() => onPageChange(page as number)}
                    >
                        {page}
                    </button>
                )
            })}
            <button className={styles['pagination-button']} disabled={currentPage === totalPages} onClick={() => onPageChange(currentPage + 1)}>
                次へ
            </button>
        </div>
    )
}
