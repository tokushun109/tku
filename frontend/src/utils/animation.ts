import { MouseEvent } from 'react'

export const RippleColorEnum = {
    Black: 'black',
    Orange: 'orange',
} as const
type RippleColorType = (typeof RippleColorEnum)[keyof typeof RippleColorEnum]

/**
 * 波紋アニメーションを実行
 *
 * NOTE: 実行する要素には下記のスタイルが必要
 * - position: static以外を指定
 * - overflow: hiddenを指定
 */
export const doRippleAnimation = (e: MouseEvent, rippleColorType: RippleColorType) => {
    const target = e.currentTarget

    // 波紋要素を作成
    const ripple = document.createElement('span')
    ripple.className = 'ripple'

    // クリック位置を計算
    const rect = target.getBoundingClientRect()
    const size = Math.max(rect.width, rect.height)
    const x = e.clientX - rect.left - size / 2
    const y = e.clientY - rect.top - size / 2

    // スタイルを設定
    ripple.style.width = ripple.style.height = `${size}px`
    ripple.style.left = `${x}px`
    ripple.style.top = `${y}px`
    let rippleColor: string = ''
    switch (rippleColorType) {
        case RippleColorEnum.Orange:
            rippleColor = 'rgb(255, 183, 77)'
            break

        case RippleColorEnum.Black:
            rippleColor = 'rgb(123, 103, 91)'
            break
        default:
            break
    }
    ripple.style.background = rippleColor

    // ボタンに波紋を追加
    target.appendChild(ripple)

    // アニメーション終了後に波紋を削除
    ripple.addEventListener('animationend', () => {
        ripple.remove()
    })
}
