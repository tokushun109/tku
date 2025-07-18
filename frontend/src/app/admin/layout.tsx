'use client'

import AdminFooter from '@/components/layouts/AdminFooter'
import AdminHeader from '@/components/layouts/AdminHeader'

const AdminLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <div>
            <AdminHeader />
            <main>{children}</main>
            <AdminFooter />
        </div>
    )
}

export default AdminLayout
