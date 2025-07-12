import NextImage from 'next/image'
import { useEffect, useState } from 'react'

import styles from './styles.module.scss'

type Props = {
    alt: string
    src: string
}

export const Image = ({ src, alt }: Props) => {
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [displaySrc, setDisplaySrc] = useState<string>(src)

    // srcが変わったらdisplaySrcも更新
    useEffect(() => {
        setDisplaySrc(src)
        setIsLoading(true)
    }, [src])

    return (
        <div className={styles['container']}>
            <NextImage
                alt={alt}
                className={styles['image']}
                fill
                loading="lazy"
                onError={() => {
                    setDisplaySrc('/image/gray-image.png')
                }}
                onLoad={() => {
                    setIsLoading(false)
                }}
                sizes="100%"
                src={isLoading ? '/image/gray-image.png' : displaySrc}
            />
        </div>
    )
}
