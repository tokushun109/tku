import { getContacts } from '@/apis/contact'

import { AdminContactTemplate } from './template'

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
