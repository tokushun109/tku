'use client'

import Head from 'next/head'
import Image from 'next/image'
import Link from 'next/link'
import React from 'react'

import { Header } from '@/components/layouts/Header'

import styles from './error.module.scss'

interface Props {
    errorMessage: React.ReactNode
    showHomeButton?: boolean
    statusCode?: number
}

const ErrorPage = ({ errorMessage, statusCode, showHomeButton = true }: Props) => {
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
                <Header isDisabled />
                <div className={styles.container}>
                    <div className={styles['error-content']}>
                        <div className={styles['error-message']}>
                            {statusCode && <div className={styles['error-code']}>{statusCode}</div>}
                            <div>{errorMessage}</div>
                            {showHomeButton && (
                                <div className={styles['home-button-wrapper']}>
                                    <Link className={styles['home-link']} href="/">
                                        トップページに戻る
                                    </Link>
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
