import { fireEvent, render, screen } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'

import { Dialog } from '@/components/bases/Dialog'

describe('Dialog', () => {
    it('デフォルトでは背景クリックで閉じる', () => {
        const onClose = vi.fn()
        const { container } = render(
            <Dialog isOpen onClose={onClose} title="テストダイアログ">
                <p>ダイアログ内容</p>
            </Dialog>,
        )

        fireEvent.click(container.firstChild as HTMLElement)

        expect(onClose).toHaveBeenCalledTimes(1)
    })

    it('closeOnBackdropClick が false のとき背景クリックで閉じない', () => {
        const onClose = vi.fn()
        const { container } = render(
            <Dialog closeOnBackdropClick={false} isOpen onClose={onClose} title="テストダイアログ">
                <p>ダイアログ内容</p>
            </Dialog>,
        )

        fireEvent.click(container.firstChild as HTMLElement)

        expect(onClose).not.toHaveBeenCalled()
        expect(screen.getByText('ダイアログ内容')).toBeInTheDocument()
    })
})
