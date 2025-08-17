'use client'

import Image from 'next/image'
import { useEffect, useState } from 'react'

import styles from './styles.module.scss'

export const TopLoading = () => {
    const [isDisplay, setIsDisplay] = useState<boolean>(true)
    const [shouldRemove, setShouldRemove] = useState<boolean>(false)

    useEffect(() => {
        const fadeTimer = setTimeout(() => {
            setIsDisplay(false)
        }, 1000)

        const removeTimer = setTimeout(() => {
            setShouldRemove(true)
        }, 1700) // フェードアウト完了後に削除

        return () => {
            clearTimeout(fadeTimer)
            clearTimeout(removeTimer)
        }
    }, [])

    if (shouldRemove) return null

    return (
        <div className={`${styles['top-loading']} ${!isDisplay ? styles['fade-out'] : ''}`}>
            <div className={styles['top-loading__logo']}>
                <Image alt="tocoriri ロゴ" className={styles['fade-up']} height={400} priority src="/logo/tocoriri_logo_white.png" width={400} />
            </div>
        </div>
    )
}
