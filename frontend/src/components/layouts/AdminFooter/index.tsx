'use client'

import { Box, Typography } from '@mui/material'

import styles from './styles.module.scss'

interface AdminFooterProps {
    className?: string
}

const AdminFooter: React.FC<AdminFooterProps> = ({ className }) => {
    return (
        <Box className={`${styles['admin-footer']} ${className || ''}`} component="footer">
            <Typography className={styles['footer-text']} variant="body2">
                tku
            </Typography>
        </Box>
    )
}

export default AdminFooter
