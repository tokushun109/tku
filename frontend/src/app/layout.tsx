import 'ress'

import '@/styles/globals.scss'
import { Metadata } from 'next'
import { Toaster } from 'sonner'

import { PageFadeTransition } from '@/components/layouts/PageFadeTransition'
import { mainFontFace } from '@/utils/font'

export const metadata: Metadata = {
    icons: [{ rel: 'icon', url: '/favicon/favicon.ico' }],
}

const RootLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <html lang="ja">
            <body className={mainFontFace.className}>
                <PageFadeTransition>{children}</PageFadeTransition>
                <Toaster
                    duration={5000}
                    icons={{
                        error: null,
                        info: null,
                        loading: null,
                        success: null,
                        warning: null,
                    }}
                    position="bottom-right"
                    toastOptions={{
                        style: {
                            fontFamily: mainFontFace.style.fontFamily,
                        },
                    }}
                />
            </body>
        </html>
    )
}

export default RootLayout
