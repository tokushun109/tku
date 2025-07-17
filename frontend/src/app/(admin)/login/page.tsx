import React from 'react'

import { LoginTemplate } from './template'

const AdminLoginPage = () => {
    return <LoginTemplate />
}

export default AdminLoginPage

// adminページはindexさせない
export const metadata = {
    robots: 'noindex, nofollow',
}
