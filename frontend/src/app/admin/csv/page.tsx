import { Metadata } from 'next'

import { AdminCsvTemplate } from './template'

export const metadata: Metadata = {
    title: 'CSV操作 | tocoriri',
    description: '商品レコードのCSVダウンロード・アップロード機能',
}

const AdminCsvPage = () => {
    return <AdminCsvTemplate />
}

export default AdminCsvPage
