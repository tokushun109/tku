import 'ress'

import '@/styles/globals.scss'
import { Metadata } from 'next'

import { PageFadeTransition } from '@/components/layouts/PageFadeTransition'
import { mainFontFace } from '@/utils/font'

import Favicon from '/public/favicon/favicon.ico'

export const metadata: Metadata = {
    icons: [{ rel: 'icon', url: Favicon.src }],
}

const RootLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <html lang="ja">
            <body className={mainFontFace.className}>
                <PageFadeTransition>{children}</PageFadeTransition>
            </body>
        </html>
    )
}

export default RootLayout
