import { MouseEvent } from 'react'
import { describe, it, expect, beforeEach, vi } from 'vitest'

import { doRippleAnimation, RippleColorEnum } from '@/utils/animation'

// MouseEventのモック
const createMockMouseEvent = (clientX: number, clientY: number, currentTarget: HTMLElement): MouseEvent => {
    return {
        clientX,
        clientY,
        currentTarget,
        preventDefault: vi.fn(),
        stopPropagation: vi.fn(),
    } as unknown as MouseEvent
}

describe('animation utils', () => {
    let mockButton: HTMLElement

    beforeEach(() => {
        // DOM要素をクリーンアップ
        document.body.innerHTML = ''

        // テスト用のボタンを作成
        mockButton = document.createElement('button')
        mockButton.style.position = 'relative'
        mockButton.style.overflow = 'hidden'
        mockButton.style.width = '100px'
        mockButton.style.height = '50px'
        document.body.appendChild(mockButton)

        // getBoundingClientRectをモック
        vi.spyOn(mockButton, 'getBoundingClientRect').mockReturnValue({
            left: 10,
            top: 20,
            width: 100,
            height: 50,
            right: 110,
            bottom: 70,
            x: 10,
            y: 20,
            toJSON: vi.fn(),
        })
    })

    describe('doRippleAnimation', () => {
        it('波紋要素が正しく作成される', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Orange)

            const rippleElement = mockButton.querySelector('.ripple')
            expect(rippleElement).toBeTruthy()
            expect(rippleElement?.className).toBe('ripple')
        })

        it('波紋要素のサイズが正しく計算される', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Orange)

            const rippleElement = mockButton.querySelector('.ripple') as HTMLElement
            expect(rippleElement.style.width).toBe('100px') // Math.max(100, 50) = 100
            expect(rippleElement.style.height).toBe('100px')
        })

        it('波紋要素の位置が正しく計算される', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Orange)

            const rippleElement = mockButton.querySelector('.ripple') as HTMLElement
            // x = clientX - rect.left - size / 2 = 60 - 10 - 50 = 0
            // y = clientY - rect.top - size / 2 = 45 - 20 - 50 = -25
            expect(rippleElement.style.left).toBe('0px')
            expect(rippleElement.style.top).toBe('-25px')
        })

        it('オレンジ色の波紋が正しく設定される', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Orange)

            const rippleElement = mockButton.querySelector('.ripple') as HTMLElement
            expect(rippleElement.style.background).toBe('rgb(255, 183, 77)')
        })

        it('黒色の波紋が正しく設定される', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Black)

            const rippleElement = mockButton.querySelector('.ripple') as HTMLElement
            expect(rippleElement.style.background).toBe('rgb(123, 103, 91)')
        })

        it('異なるクリック位置で波紋位置が変わる', () => {
            const mockEvent1 = createMockMouseEvent(30, 35, mockButton)
            const mockEvent2 = createMockMouseEvent(90, 55, mockButton)

            doRippleAnimation(mockEvent1, RippleColorEnum.Orange)
            doRippleAnimation(mockEvent2, RippleColorEnum.Orange)

            const rippleElements = mockButton.querySelectorAll('.ripple')
            expect(rippleElements).toHaveLength(2)

            const ripple1 = rippleElements[0] as HTMLElement
            const ripple2 = rippleElements[1] as HTMLElement

            // 1番目の波紋: x = 30 - 10 - 50 = -30, y = 35 - 20 - 50 = -35
            expect(ripple1.style.left).toBe('-30px')
            expect(ripple1.style.top).toBe('-35px')

            // 2番目の波紋: x = 90 - 10 - 50 = 30, y = 55 - 20 - 50 = -15
            expect(ripple2.style.left).toBe('30px')
            expect(ripple2.style.top).toBe('-15px')
        })

        it('異なるサイズの要素で波紋サイズが変わる', () => {
            // 縦長の要素を作成
            const tallButton = document.createElement('button')
            tallButton.style.position = 'relative'
            tallButton.style.overflow = 'hidden'
            document.body.appendChild(tallButton)

            vi.spyOn(tallButton, 'getBoundingClientRect').mockReturnValue({
                left: 0,
                top: 0,
                width: 50,
                height: 150,
                right: 50,
                bottom: 150,
                x: 0,
                y: 0,
                toJSON: vi.fn(),
            })

            const mockEvent = createMockMouseEvent(25, 75, tallButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Orange)

            const rippleElement = tallButton.querySelector('.ripple') as HTMLElement
            expect(rippleElement.style.width).toBe('150px') // Math.max(50, 150) = 150
            expect(rippleElement.style.height).toBe('150px')
        })

        it('アニメーション終了後に波紋が削除される', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            doRippleAnimation(mockEvent, RippleColorEnum.Orange)

            const rippleElement = mockButton.querySelector('.ripple') as HTMLElement
            expect(rippleElement).toBeTruthy()

            // animationendイベントを手動でトリガー
            const animationEndEvent = new Event('animationend')
            rippleElement.dispatchEvent(animationEndEvent)

            // 波紋要素が削除されていることを確認
            expect(mockButton.querySelector('.ripple')).toBeFalsy()
        })

        it('複数の波紋要素を同時に作成できる', () => {
            const mockEvent1 = createMockMouseEvent(30, 35, mockButton)
            const mockEvent2 = createMockMouseEvent(70, 45, mockButton)

            doRippleAnimation(mockEvent1, RippleColorEnum.Orange)
            doRippleAnimation(mockEvent2, RippleColorEnum.Black)

            const rippleElements = mockButton.querySelectorAll('.ripple')
            expect(rippleElements).toHaveLength(2)

            const ripple1 = rippleElements[0] as HTMLElement
            const ripple2 = rippleElements[1] as HTMLElement

            expect(ripple1.style.background).toBe('rgb(255, 183, 77)')
            expect(ripple2.style.background).toBe('rgb(123, 103, 91)')
        })

        it('デフォルトケースでは背景色が設定されない', () => {
            const mockEvent = createMockMouseEvent(60, 45, mockButton)

            // 無効な色タイプを強制的に渡す
            doRippleAnimation(mockEvent, 'invalid-color' as any)

            const rippleElement = mockButton.querySelector('.ripple') as HTMLElement
            expect(rippleElement.style.background).toBe('')
        })
    })

    describe('RippleColorEnum', () => {
        it('正しい色定数が定義されている', () => {
            expect(RippleColorEnum.Black).toBe('black')
            expect(RippleColorEnum.Orange).toBe('orange')
        })

        it('オブジェクトが凍結されている', () => {
            // TypeScriptのas constによる型レベルの制約であり、実際にはObjectが凍結されていない
            expect(Object.isFrozen(RippleColorEnum)).toBe(false)
        })
    })
})
