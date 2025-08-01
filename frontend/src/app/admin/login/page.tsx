import { Metadata } from 'next'

import { AdminLoginTemplate } from './template'

export const metadata: Metadata = {
    title: 'ログイン | admin',
}

const AdminLoginPage = () => {
    return <AdminLoginTemplate />
}

export default AdminLoginPage
