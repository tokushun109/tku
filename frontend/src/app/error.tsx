'use client'

import Head from 'next/head'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import React from 'react'

import { Button } from '@/components/bases/Button'

import styles from './error.module.scss'

interface ErrorPageProps {
    errorMessage: React.ReactNode
    showHomeButton?: boolean
    statusCode?: number
}

const ErrorPage: React.FC<ErrorPageProps> = ({ errorMessage, statusCode, showHomeButton = true }) => {
    const router = useRouter()

    const handleHomeClick = () => {
        router.push('/')
    }
    return (
        <>
            <Head>
                <meta content="noindex, nofollow" name="robots" />
            </Head>
            <div className={styles['error-wrapper']}>
                <div className={styles['site-title-area']}>
                    <Link href="/">
                        <Image
                            alt="アクセサリーショップ とこりり"
                            className={styles['site-title']}
                            height={150}
                            src="/logo/tocoriri_logo.png"
                            width={370}
                        />
                    </Link>
                </div>
                <div className={styles.container}>
                    <div className={styles['error-content']}>
                        <div className={styles['error-message']}>
                            {statusCode && <div className={styles['error-code']}>{statusCode}</div>}
                            <div>{errorMessage}</div>
                            {showHomeButton && (
                                <div className={styles['home-button-wrapper']}>
                                    <Button onClick={handleHomeClick}>トップページに戻る</Button>
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default ErrorPage
