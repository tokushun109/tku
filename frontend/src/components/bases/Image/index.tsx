import NextImage from 'next/image'
import { useEffect, useRef, useState } from 'react'

import styles from './styles.module.scss'

type Props = {
    alt: string
    priority?: boolean
    src: string
}

export const Image = ({ src, alt, priority = false }: Props) => {
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [hasError, setHasError] = useState<boolean>(false)
    const containerRef = useRef<HTMLDivElement>(null)
    const normalizedSrc = src.trim()

    // srcが変わったら状態を初期化
    useEffect(() => {
        setHasError(false)
        setIsLoading(normalizedSrc !== '')
    }, [normalizedSrc])

    const shouldRenderImage = normalizedSrc !== '' && !hasError

    useEffect(() => {
        if (!shouldRenderImage) {
            return
        }

        const imageElement = containerRef.current?.querySelector('img')

        if (!imageElement || !imageElement.complete) {
            return
        }

        if (imageElement.naturalWidth > 0) {
            setIsLoading(false)
            return
        }

        setHasError(true)
        setIsLoading(false)
    }, [shouldRenderImage, normalizedSrc])

    return (
        <div className={styles['container']} ref={containerRef}>
            {shouldRenderImage && (
                <NextImage
                    alt={alt}
                    className={styles['image']}
                    fill
                    loading={priority ? undefined : 'lazy'}
                    onError={() => {
                        setHasError(true)
                        setIsLoading(false)
                    }}
                    onLoad={() => {
                        setIsLoading(false)
                    }}
                    priority={priority}
                    sizes="100%"
                    src={normalizedSrc}
                    style={{ opacity: isLoading ? 0 : 1 }}
                />
            )}
        </div>
    )
}
