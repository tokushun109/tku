import { Metadata } from 'next'

import { AdminProductTemplate } from './template'

export const metadata: Metadata = {
    title: '商品管理 | admin',
    robots: 'noindex, nofollow',
}

const AdminProductPage = () => {
    return <AdminProductTemplate />
}

export default AdminProductPage
