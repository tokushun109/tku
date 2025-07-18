'use client'

import { Close, Menu } from '@mui/icons-material'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

import { Button } from '@/components/bases/Button'
import { ColorCode, ColorType } from '@/types/enum/color'
import { NavigationItems } from '@/types/enum/navigation'

import styles from './styles.module.scss'

interface AdminHeaderProps {
    isLoggedIn?: boolean
    onLogout?: () => void
}

const AdminHeader: React.FC<AdminHeaderProps> = ({ isLoggedIn = true, onLogout }) => {
    const router = useRouter()
    const [sidebarVisible, setSidebarVisible] = useState<boolean>(false)
    const [dialogVisible, setDialogVisible] = useState<boolean>(false)

    const navigationItems = NavigationItems

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
                        <Button
                            className={styles['logout-button']}
                            colorType={ColorType.Primary}
                            noBoxShadow
                            onClick={() => setDialogVisible(true)}
                            outlined
                        >
                            ログアウト
                        </Button>
                    )}
                </div>
            </header>

            {isLoggedIn && (
                <div className={`${styles['drawer']} ${sidebarVisible ? styles['drawer-open'] : ''}`}>
                    <div className={styles['drawer-backdrop']} onClick={() => setSidebarVisible(false)} />
                    <nav className={styles['drawer-content']}>
                        <div className={styles['drawer-header']}>
                            <h2 className={styles['drawer-title']}>設定</h2>
                            <button className={styles['close-button']} onClick={() => setSidebarVisible(false)} type="button">
                                <Close sx={{ color: ColorCode[ColorType.Primary] }} />
                            </button>
                        </div>
                        <div className={styles['divider']} />
                        <ul className={styles['navigation-list']}>
                            {navigationItems.map((item) => {
                                const IconComponent = item.icon
                                return (
                                    <li key={item.name}>
                                        <button className={styles['nav-item']} onClick={() => handleNavigationClick(item.link)} type="button">
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
                            <Button colorType={ColorType.Primary} contrast onClick={() => setDialogVisible(false)} outlined>
                                いいえ
                            </Button>
                            <Button colorType={ColorType.Primary} onClick={handleLogout}>
                                はい
                            </Button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    )
}

export default AdminHeader
