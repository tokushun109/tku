import { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'メンテナンス中です | とこりり',
    robots: {
        index: false,
        follow: false,
    },
}

interface MaintenanceLayoutProps {
    children: React.ReactNode
}

const MaintenanceLayout: React.FC<MaintenanceLayoutProps> = ({ children }) => {
    return children
}

export default MaintenanceLayout
