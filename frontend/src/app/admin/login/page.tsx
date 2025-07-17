import { Metadata } from 'next'

import { AdminLoginTemplate } from './template'

export const metadata: Metadata = {
    title: 'ログイン | admin',
    robots: 'noindex, nofollow',
}

const AdminLoginPage = () => {
    return <AdminLoginTemplate />
}

export default AdminLoginPage
