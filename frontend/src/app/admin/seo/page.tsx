import { Metadata } from 'next'

import { getCreator } from '@/apis/creator'

import { AdminSeoTemplate } from './template'

export const metadata: Metadata = {
    title: 'SEO設定 | admin',
    robots: {
        index: false,
        follow: false,
    },
}

const AdminSeoPage = async () => {
    try {
        const creator = await getCreator()

        return <AdminSeoTemplate creator={creator} />
    } catch (error) {
        console.error('作者情報の取得に失敗しました:', error)
        return <AdminSeoTemplate creator={null} />
    }
}

export default AdminSeoPage
