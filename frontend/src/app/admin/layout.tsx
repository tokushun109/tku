'use client'

import AdminFooter from '@/components/layouts/AdminFooter'
import AdminHeader from '@/components/layouts/AdminHeader'

import styles from './layout.module.scss'

const AdminLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <div className={styles['container']}>
            <AdminHeader />
            <main>{children}</main>
            <AdminFooter />
        </div>
    )
}

export default AdminLayout
