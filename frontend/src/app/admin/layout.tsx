'use client'

import { usePathname } from 'next/navigation'

import AdminHeader from '@/components/layouts/AdminHeader'
import { NavigationType } from '@/types/enum/navigation'

import styles from './layout.module.scss'

const AdminLayout = ({ children }: { children: React.ReactNode }) => {
    const pathname = usePathname()
    const isLoggedIn = pathname !== NavigationType.AdminLogin

    return (
        <div className={styles['container']}>
            <AdminHeader isLoggedIn={isLoggedIn} />
            <main>{children}</main>
        </div>
    )
}

export default AdminLayout
