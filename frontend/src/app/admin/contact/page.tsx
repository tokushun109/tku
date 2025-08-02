import { Metadata } from 'next'

import { getContacts } from '@/apis/contact'

import { AdminContactTemplate } from './template'

export const metadata: Metadata = {
    title: 'お問い合わせ管理 | admin',
    robots: {
        index: false,
        follow: false,
    },
}

const AdminContactPage = async () => {
    try {
        const contacts = await getContacts()

        return <AdminContactTemplate contacts={contacts} />
    } catch (error) {
        console.error('お問い合わせ一覧の取得に失敗しました:', error)
        return <AdminContactTemplate contacts={[]} />
    }
}

export default AdminContactPage
