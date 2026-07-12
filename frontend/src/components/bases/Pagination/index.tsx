import { KeyboardArrowLeft, KeyboardArrowRight } from '@mui/icons-material'

import styles from './styles.module.scss'

type PaginationItem = 'ellipsis' | number

interface Props {
    currentPage: number
    disabled?: boolean
    onPageChange: (_page: number) => void
    siblingCount?: number
    totalPages: number
}

/** 現在ページの周辺と先頭/末尾を残し、長いページ列を省略表示用に整形する。 */
const createPaginationItems = (currentPage: number, totalPages: number, siblingCount: number): PaginationItem[] => {
    const visiblePageCount = siblingCount * 2 + 5

    if (totalPages <= visiblePageCount) {
        return Array.from({ length: totalPages }, (_, index) => index + 1)
    }

    const firstPage = 1
    const lastPage = totalPages
    const leftSibling = Math.max(currentPage - siblingCount, firstPage)
    const rightSibling = Math.min(currentPage + siblingCount, lastPage)
    const shouldShowLeftEllipsis = leftSibling > firstPage + 1
    const shouldShowRightEllipsis = rightSibling < lastPage - 1

    if (!shouldShowLeftEllipsis && shouldShowRightEllipsis) {
        const leftRange = Array.from({ length: 3 + siblingCount * 2 }, (_, index) => index + 1)
        return [...leftRange, 'ellipsis', lastPage]
    }

    if (shouldShowLeftEllipsis && !shouldShowRightEllipsis) {
        const rightRangeStart = lastPage - (2 + siblingCount * 2)
        const rightRange = Array.from({ length: lastPage - rightRangeStart + 1 }, (_, index) => rightRangeStart + index)
        return [firstPage, 'ellipsis', ...rightRange]
    }

    const middleRange = Array.from({ length: rightSibling - leftSibling + 1 }, (_, index) => leftSibling + index)
    return [firstPage, 'ellipsis', ...middleRange, 'ellipsis', lastPage]
}

export const Pagination = ({ currentPage, disabled = false, onPageChange, siblingCount = 1, totalPages }: Props) => {
    if (totalPages <= 1) return null

    const paginationItems = createPaginationItems(currentPage, totalPages, siblingCount)
    const canGoPrevious = currentPage > 1
    const canGoNext = currentPage < totalPages

    const handlePageChange = (page: number) => {
        if (disabled || page === currentPage || page < 1 || page > totalPages) return

        onPageChange(page)
    }

    return (
        <nav aria-label="ページネーション" className={styles['container']}>
            <button
                aria-label="前のページへ"
                className={styles['nav-button']}
                disabled={disabled || !canGoPrevious}
                onClick={() => handlePageChange(currentPage - 1)}
                type="button"
            >
                <KeyboardArrowLeft fontSize="small" />
                <span>前へ</span>
            </button>
            <div className={styles['page-list']}>
                {paginationItems.map((item, index) =>
                    item === 'ellipsis' ? (
                        <span aria-hidden="true" className={styles['ellipsis']} key={`ellipsis-${index}`}>
                            ...
                        </span>
                    ) : (
                        <button
                            aria-current={item === currentPage ? 'page' : undefined}
                            aria-label={`${item}ページへ`}
                            className={`${styles['page-button']} ${item === currentPage ? styles['active'] : ''}`}
                            disabled={disabled || item === currentPage}
                            key={item}
                            onClick={() => handlePageChange(item)}
                            type="button"
                        >
                            {item}
                        </button>
                    ),
                )}
            </div>
            <button
                aria-label="次のページへ"
                className={styles['nav-button']}
                disabled={disabled || !canGoNext}
                onClick={() => handlePageChange(currentPage + 1)}
                type="button"
            >
                <span>次へ</span>
                <KeyboardArrowRight fontSize="small" />
            </button>
        </nav>
    )
}
