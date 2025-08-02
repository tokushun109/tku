import { Metadata } from 'next'

import { AdminLoginTemplate } from './template'

export const metadata: Metadata = {
    title: 'ログイン | admin',
    robots: {
        index: false,
        follow: false,
    },
}

const AdminLoginPage = () => {
    return <AdminLoginTemplate />
}

export default AdminLoginPage
