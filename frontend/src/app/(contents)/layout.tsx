'use client'

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
                    <div className={classNames(styles['title'], styles['default'], labelFontFace.className)}>{NavigationTitleEnum[pathname]}</div>
                    {children}
                </PageFadeTransition>
            </main>
            <Footer />
        </div>
    )
}

export default DetailsLayout
