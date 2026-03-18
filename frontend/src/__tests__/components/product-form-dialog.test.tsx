import { fireEvent, render, screen } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'

import { ProductFormDialog } from '@/features/product/components/ProductFormDialog'

import type { ComponentProps } from 'react'

const createProps = (overrides: Partial<ComponentProps<typeof ProductFormDialog>> = {}): ComponentProps<typeof ProductFormDialog> => ({
    categories: [{ name: 'ピアス', uuid: 'category-1' }],
    isOpen: true,
    isSubmitting: false,
    onClose: vi.fn(),
    onCreemaDuplicate: vi.fn().mockResolvedValue(undefined),
    onSubmit: vi.fn().mockResolvedValue(undefined),
    salesSites: [{ name: 'Creema', uuid: 'site-1' }],
    submitError: null,
    tags: [{ name: 'シンプル', uuid: 'tag-1' }],
    targets: [{ name: '大人女性', uuid: 'target-1' }],
    updateItem: null,
    ...overrides,
})

describe('ProductFormDialog', () => {
    it('送信中でもキャンセルで閉じられる', () => {
        const onClose = vi.fn()

        render(<ProductFormDialog {...createProps({ isSubmitting: true, onClose })} />)

        const cancelButton = screen.getByRole('button', { name: 'キャンセル' })

        expect(cancelButton).toBeEnabled()

        fireEvent.click(cancelButton)

        expect(onClose).toHaveBeenCalledTimes(1)
    })
})
