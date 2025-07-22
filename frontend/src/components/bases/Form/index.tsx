import React from 'react'

import styles from './styles.module.scss'

interface Props extends React.FormHTMLAttributes<HTMLFormElement> {
    children: React.ReactNode
}

export const Form = ({ children, className, ...props }: Props) => {
    return (
        <form className={`${styles['form']} ${className || ''}`} {...props}>
            {children}
        </form>
    )
}
