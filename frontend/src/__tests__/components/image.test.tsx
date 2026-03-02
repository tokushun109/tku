import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import { afterEach, describe, expect, it, vi } from 'vitest'

import { Image } from '@/components/bases/Image'

describe('Image', () => {
    afterEach(() => {
        vi.restoreAllMocks()
    })

    it('初回描画から実画像のURLを設定する', () => {
        render(<Image alt="sample" src="https://example.com/product.png" />)

        expect(screen.getByAltText('sample')).toHaveAttribute('src', 'https://example.com/product.png')
    })

    it('srcの前後の空白を取り除いて描画する', () => {
        render(<Image alt="trimmed" src="  https://example.com/trimmed.png  " />)

        expect(screen.getByAltText('trimmed')).toHaveAttribute('src', 'https://example.com/trimmed.png')
    })

    it('priority指定時はNextImageへpriorityを渡してlazy loadを外す', () => {
        render(<Image alt="priority" priority src="https://example.com/priority.png" />)

        expect(screen.getByAltText('priority')).toHaveAttribute('data-priority', 'true')
        expect(screen.getByAltText('priority')).not.toHaveAttribute('loading')
    })

    it('srcが空文字なら画像要素を描画しない', () => {
        render(<Image alt="empty" src="" />)

        expect(screen.queryByAltText('empty')).not.toBeInTheDocument()
    })

    it('hydration時点で読み込み済みの画像でも表示状態へ切り替える', async () => {
        vi.spyOn(HTMLImageElement.prototype, 'complete', 'get').mockReturnValue(true)
        vi.spyOn(HTMLImageElement.prototype, 'naturalWidth', 'get').mockReturnValue(320)

        render(<Image alt="loaded" src="https://example.com/already-loaded.png" />)

        await waitFor(() => {
            expect(screen.getByAltText('loaded')).toHaveStyle({ opacity: '1' })
        })
    })

    it('読み込みエラー時は画像要素を取り除く', async () => {
        render(<Image alt="error" src="https://example.com/error.png" />)

        fireEvent.error(screen.getByAltText('error'))

        await waitFor(() => {
            expect(screen.queryByAltText('error')).not.toBeInTheDocument()
        })
    })
})
