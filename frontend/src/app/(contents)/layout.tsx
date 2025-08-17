'use client'

import { GoogleAnalytics } from '@next/third-parties/google'
import classNames from 'classnames'
import { usePathname } from 'next/navigation'

import { Footer } from '@/components/layouts/Footer'
import { Header } from '@/components/layouts/Header'
import { PageFadeTransition } from '@/components/layouts/PageFadeTransition'
import { NavigationTitleEnum, NavigationType } from '@/types/enum/navigation'
import { labelFontFace } from '@/utils/font'

import styles from './layout.module.scss'

const DetailsLayout = ({ children }: { children: React.ReactNode }) => {
    const pathname = usePathname() as NavigationType

    return (
        <div className={styles['container']}>
            <Header />
            <main className={styles['main']}>
                <PageFadeTransition>
                    {NavigationTitleEnum[pathname] ? (
                        <div className={classNames(styles['title'], styles['default'], labelFontFace.className)}>{NavigationTitleEnum[pathname]}</div>
                    ) : (
                        <div className={styles['no-title-space']} />
                    )}
                    {children}
                </PageFadeTransition>
            </main>
            <Footer />
            {process.env.GOOGLE_TAG && <GoogleAnalytics gaId={process.env.GOOGLE_TAG} />}
        </div>
    )
}

export default DetailsLayout
