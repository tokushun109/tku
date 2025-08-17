import 'ress'

import { GoogleAnalytics } from '@next/third-parties/google'

import { Footer } from '@/components/layouts/Footer'
import { Header } from '@/components/layouts/Header'

const HomeLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <div>
            <Header />
            <main>{children}</main>
            <Footer />
            {process.env.GOOGLE_TAG && <GoogleAnalytics gaId={process.env.GOOGLE_TAG} />}
        </div>
    )
}

export default HomeLayout
