'use client'

import styles from './styles.module.scss'

interface AdminFooterProps {
    className?: string
}

const AdminFooter: React.FC<AdminFooterProps> = ({ className }) => {
    return (
        <footer className={`${styles['admin-footer']} ${className || ''}`}>
            <p className={styles['footer-text']}>tku</p>
        </footer>
    )
}

export default AdminFooter
