import { fireEvent, render, screen } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'

import { Pagination } from '@/components/bases/Pagination'

describe('Pagination', () => {
    it('ページ番号と省略記号を表示する', () => {
        render(<Pagination currentPage={5} onPageChange={vi.fn()} totalPages={10} />)

        expect(screen.getByRole('button', { name: '前のページへ' })).toBeInTheDocument()
        expect(screen.getByRole('button', { name: '5ページへ' })).toHaveAttribute('aria-current', 'page')
        expect(screen.getAllByText('...')).toHaveLength(2)
        expect(screen.getByRole('button', { name: '次のページへ' })).toBeInTheDocument()
    })

    it('ページ番号クリックでページ変更を通知する', () => {
        const handlePageChange = vi.fn()
        render(<Pagination currentPage={1} onPageChange={handlePageChange} totalPages={3} />)

        fireEvent.click(screen.getByRole('button', { name: '2ページへ' }))

        expect(handlePageChange).toHaveBeenCalledWith(2)
    })

    it('1ページだけの場合は表示しない', () => {
        render(<Pagination currentPage={1} onPageChange={vi.fn()} totalPages={1} />)

        expect(screen.queryByRole('navigation', { name: 'ページネーション' })).not.toBeInTheDocument()
    })
})
