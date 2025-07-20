import { Metadata } from 'next'

export const metadata: Metadata = {
    title: 'メンテナンス中です | とこりり',
    robots: {
        index: false,
        follow: false,
    },
}

interface Props {
    children: React.ReactNode
}

const MaintenanceLayout = ({ children }: Props) => {
    return children
}

export default MaintenanceLayout
