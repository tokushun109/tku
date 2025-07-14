import { useEffect, useRef, useState } from 'react'

import { IThumbnail } from '@/features/product/type'

import { ImageIndex, SwipeDirection, SwipePosition } from '../types'

export const useSlideShow = (items: IThumbnail[], autoPlay: boolean) => {
    // 商品画像の数
    const maxLength = items.length

    const [imageIndex, setImageIndex] = useState<ImageIndex>({ previous: maxLength - 1, display: 0, next: 1 })
    const [swipePosition, setSwipePosition] = useState<SwipePosition>({ start: undefined, end: undefined })
    const [swipeDirection, setSwipeDirection] = useState<SwipeDirection | undefined>(undefined)

    // 画像をスライドする
    const slideImage = (swipeDirection: SwipeDirection) => {
        switch (swipeDirection) {
            case 'right':
                setImageIndex({
                    previous: imageIndex.previous !== 0 ? imageIndex.previous - 1 : maxLength - 1,
                    display: imageIndex.display !== 0 ? imageIndex.display - 1 : maxLength - 1,
                    next: imageIndex.next !== 0 ? imageIndex.next - 1 : maxLength - 1,
                })
                break
            case 'left':
                setImageIndex({
                    previous: imageIndex.previous !== maxLength - 1 ? imageIndex.previous + 1 : 0,
                    display: imageIndex.display !== maxLength - 1 ? imageIndex.display + 1 : 0,
                    next: imageIndex.next !== maxLength - 1 ? imageIndex.next + 1 : 0,
                })
                break
        }
        setSwipeDirection(swipeDirection)
        window.setTimeout(() => {
            // アニメーション終了後にスワイプ方向のリセットする
            setSwipeDirection(undefined)
        }, 300)
    }

    const useInterval = (callback: () => void) => {
        const callbackRef = useRef<() => void>(callback)
        useEffect(() => {
            callbackRef.current = callback
        }, [callback])

        useEffect(() => {
            if (!autoPlay) return
            const tick = () => {
                callbackRef.current()
            }
            const id = window.setInterval(tick, 5000)
            return () => {
                window.clearInterval(id)
            }
            // eslint-disable-next-line react-hooks/exhaustive-deps
        }, [imageIndex])
    }

    useInterval(() => {
        slideImage('left')
    })

    // 画像をクリックした時のハンドラ
    const clickHandler = () => {
        slideImage('left')
    }

    // 画像をスワイプした時のハンドラ
    const swipeHandler = () => {
        // start、endに値がない(イベントが発火していない)ときは処理をスルー
        if (!swipePosition.start || !swipePosition.end) {
            setSwipePosition({ start: undefined, end: undefined })
            return
        }
        if (swipePosition.end - swipePosition.start > 0) {
            slideImage('right')
        } else if (swipePosition.end - swipePosition.start < 0) {
            slideImage('left')
        }
        setSwipePosition({ start: undefined, end: undefined })
    }

    return {
        imageIndex,
        swipePosition,
        swipeDirection,
        setSwipePosition,
        clickHandler,
        swipeHandler,
    }
}
