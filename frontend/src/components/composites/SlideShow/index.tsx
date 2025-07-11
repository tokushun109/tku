'use client'

import classNames from 'classnames'

import { IThumbnail } from '@/features/product/type'

import { useSlideShow } from './hooks'
import styles from './styles.module.scss'
import { ImageIndexEnum } from './types'
import { CarouselImage } from '../CarouselImage'

type Props = {
    autoPlay?: boolean
    innerPadding?: number
    items: IThumbnail[]
    size: string
}

const SlideShow = ({ items, size, innerPadding = 16, autoPlay = true }: Props) => {
    const { imageIndex, swipePosition, swipeDirection, setSwipePosition, clickHandler, swipeHandler } = useSlideShow(items, autoPlay)
    return (
        <div
            className={styles['container']}
            onClick={clickHandler}
            onTouchEnd={swipeHandler}
            onTouchMove={(e) => {
                setSwipePosition({ ...swipePosition, end: e.touches[0].pageX })
            }}
            onTouchStart={(e) => {
                setSwipePosition({ ...swipePosition, start: e.touches[0].pageX })
            }}
            style={{ width: `calc(${size})`, height: `calc(${size})` }}
        >
            <div className={styles['wrapper']}>
                {Object.values(ImageIndexEnum).map((v) => (
                    <div
                        className={classNames(styles['content'], styles[v], swipeDirection && styles[`${swipeDirection}-swipe`])}
                        key={v}
                        style={{ width: `calc(${size} - ${innerPadding}px)`, aspectRatio: '1 / 1' }}
                    >
                        <CarouselImage item={items[imageIndex[v]]} shadow={false} />
                    </div>
                ))}
            </div>
        </div>
    )
}

export default SlideShow
