'use client'

import { AccountCircle, Close, Description, Email, Menu, ShoppingCart, Tag, Web } from '@mui/icons-material'
import { useRouter } from 'next/navigation'
import { useState } from 'react'


import styles from './styles.module.scss'

interface AdminHeaderProps {
    isLoggedIn?: boolean
    onLogout?: () => void
}

interface NavigationItem {
    icon: React.ComponentType
    link: string
    name: string
}

const AdminHeader: React.FC<AdminHeaderProps> = ({ isLoggedIn = true, onLogout }) => {
    const router = useRouter()
    const [sidebarVisible, setSidebarVisible] = useState<boolean>(false)
    const [dialogVisible, setDialogVisible] = useState<boolean>(false)

    const navigationItems: NavigationItem[] = [
        { icon: AccountCircle, link: '/admin/seo', name: 'SEO' },
        { icon: ShoppingCart, link: '/admin/product', name: '商品' },
        { icon: Tag, link: '/admin/classification', name: '分類' },
        { icon: Web, link: '/admin/site', name: 'サイト' },
        { icon: Email, link: '/admin/contact', name: 'お問い合わせ' },
        { icon: Description, link: '/admin/csv', name: 'CSV' },
    ]

    const handleLogout = () => {
        setDialogVisible(false)
        onLogout?.()
        router.push('/admin/user/login')
    }

    const handleNavigationClick = (link: string) => {
        router.push(link)
        setSidebarVisible(false)
    }

    return (
        <div className={styles['admin-header']}>
            <header className={styles['app-bar']}>
                <div className={styles['toolbar']}>
                    {isLoggedIn && (
                        <button
                            aria-label="open drawer"
                            className={styles['menu-button']}
                            onClick={() => setSidebarVisible(!sidebarVisible)}
                            type="button"
                        >
                            <Menu />
                        </button>
                    )}
                    <h1 className={styles['title']}>tku</h1>
                    <div className={styles['spacer']} />
                    {isLoggedIn && (
                        <button
                            className={styles['logout-button']}
                            onClick={() => setDialogVisible(true)}
                            type="button"
                        >
                            ログアウト
                        </button>
                    )}
                </div>
            </header>

            {isLoggedIn && (
                <div className={`${styles['drawer']} ${sidebarVisible ? styles['drawer-open'] : ''}`}>
                    <div className={styles['drawer-backdrop']} onClick={() => setSidebarVisible(false)} />
                    <nav className={styles['drawer-content']}>
                        <div className={styles['drawer-header']}>
                            <h2 className={styles['drawer-title']}>設定</h2>
                            <button
                                className={styles['close-button']}
                                onClick={() => setSidebarVisible(false)}
                                type="button"
                            >
                                <Close />
                            </button>
                        </div>
                        <div className={styles['divider']} />
                        <ul className={styles['navigation-list']}>
                            {navigationItems.map((item) => {
                                const IconComponent = item.icon
                                return (
                                    <li key={item.name}>
                                        <button
                                            className={styles['nav-item']}
                                            onClick={() => handleNavigationClick(item.link)}
                                            type="button"
                                        >
                                            <span className={styles['nav-icon']}>
                                                <IconComponent />
                                            </span>
                                            <span className={styles['nav-text']}>{item.name}</span>
                                        </button>
                                    </li>
                                )
                            })}
                        </ul>
                    </nav>
                </div>
            )}

            {dialogVisible && (
                <div className={styles['dialog-overlay']}>
                    <div className={styles['dialog-backdrop']} onClick={() => setDialogVisible(false)} />
                    <div className={styles['dialog']}>
                        <div className={styles['dialog-header']}>
                            <h3 className={styles['dialog-title']}>ログアウト</h3>
                        </div>
                        <div className={styles['dialog-content']}>
                            <p>ログアウトします。よろしいですか？</p>
                        </div>
                        <div className={styles['dialog-actions']}>
                            <button
                                className={styles['dialog-button-secondary']}
                                onClick={() => setDialogVisible(false)}
                                type="button"
                            >
                                いいえ
                            </button>
                            <button
                                className={styles['dialog-button-primary']}
                                onClick={handleLogout}
                                type="button"
                            >
                                はい
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    )
}

export default AdminHeader