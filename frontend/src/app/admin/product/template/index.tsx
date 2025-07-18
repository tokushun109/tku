'use client'

import { logoutAction } from '@/features/auth/action'

import styles from './styles.module.scss'

export const AdminProductTemplate = () => {
    const handleLogout = async () => {
        if (confirm('ログアウトしますか？')) {
            await logoutAction()
        }
    }

    return (
        <div className={styles['page-admin-product']}>
            <div className={styles['content-area']}>
                <h1 className={styles['page-title']}>商品管理</h1>
                <div className={styles['test-content']}>
                    <p>これは商品管理のテストページです。</p>
                    <p>ログイン時の挙動を確認するためのページです。</p>

                    <div className={styles['test-actions']}>
                        <button className={styles['logout-test-button']} onClick={handleLogout}>
                            テスト用ログアウト
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
}
