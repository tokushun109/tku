'use client'

import { Close, Email, Home, Inventory, Menu, Settings, Store, Tag } from '@mui/icons-material'
import {
    AppBar,
    Box,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Divider,
    Drawer,
    IconButton,
    List,
    ListItem,
    ListItemIcon,
    ListItemText,
    Toolbar,
    Typography,
} from '@mui/material'
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
        { icon: Settings, link: '/admin/seo', name: 'SEO' },
        { icon: Inventory, link: '/admin/product', name: '商品' },
        { icon: Tag, link: '/admin/classification', name: '分類' },
        { icon: Store, link: '/admin/site', name: 'サイト' },
        { icon: Email, link: '/admin/contact', name: 'お問い合わせ' },
        { icon: Home, link: '/admin/csv', name: 'CSV' },
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
            <AppBar className={styles['app-bar']} position="fixed">
                <Toolbar>
                    {isLoggedIn && (
                        <IconButton
                            aria-label="open drawer"
                            className={styles['menu-button']}
                            color="inherit"
                            edge="start"
                            onClick={() => setSidebarVisible(!sidebarVisible)}
                        >
                            <Menu />
                        </IconButton>
                    )}
                    <Typography className={styles['title']} component="div" noWrap variant="h6">
                        tku
                    </Typography>
                    <Box sx={{ flexGrow: 1 }} />
                    {isLoggedIn && (
                        <Button color="inherit" onClick={() => setDialogVisible(true)} variant="outlined">
                            ログアウト
                        </Button>
                    )}
                </Toolbar>
            </AppBar>

            {isLoggedIn && (
                <Drawer
                    className={styles['drawer']}
                    classes={{
                        paper: styles['drawer-paper'],
                    }}
                    onClose={() => setSidebarVisible(false)}
                    open={sidebarVisible}
                    variant="temporary"
                >
                    <div className={styles['drawer-header']}>
                        <Typography className={styles['drawer-title']} variant="h6">
                            設定
                        </Typography>
                        <IconButton onClick={() => setSidebarVisible(false)}>
                            <Close />
                        </IconButton>
                    </div>
                    <Divider />
                    <List>
                        {navigationItems.map((item) => {
                            const IconComponent = item.icon
                            return (
                                <ListItem
                                    className={styles['list-item']}
                                    component="div"
                                    key={item.name}
                                    onClick={() => handleNavigationClick(item.link)}
                                    sx={{ cursor: 'pointer' }}
                                >
                                    <ListItemIcon>
                                        <IconComponent />
                                    </ListItemIcon>
                                    <ListItemText primary={item.name} />
                                </ListItem>
                            )
                        })}
                    </List>
                </Drawer>
            )}

            <Dialog fullWidth maxWidth="sm" onClose={() => setDialogVisible(false)} open={dialogVisible}>
                <DialogTitle>ログアウト</DialogTitle>
                <DialogContent>
                    <Typography>ログアウトします。よろしいですか？</Typography>
                </DialogContent>
                <DialogActions className={styles['dialog-actions']}>
                    <Button color="primary" onClick={() => setDialogVisible(false)} variant="outlined">
                        いいえ
                    </Button>
                    <Button color="primary" onClick={handleLogout} variant="contained">
                        はい
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    )
}

export default AdminHeader
