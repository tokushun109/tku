import 'ress'

import { Footer } from '@/components/layouts/Footer'
import { Header } from '@/components/layouts/Header'

const HomeLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <div>
            <Header />
            <main>{children}</main>
            <Footer />
        </div>
    )
}

export default HomeLayout
