import React from 'react'

import styles from './styles.module.scss'

export const MessageType = {
    Error: 'error',
    Warning: 'warning',
    Info: 'info',
    Success: 'success',
} as const

export type MessageType = (typeof MessageType)[keyof typeof MessageType]

export interface MessageProps {
    children: React.ReactNode
    className?: string
    type?: MessageType
}

export const Message: React.FC<MessageProps> = ({ children, type = MessageType.Error, className = '' }) => {
    const messageClass = `${styles['message']} ${styles[`message--${type}`]} ${className}`.trim()

    return (
        <div aria-live="polite" className={messageClass}>
            {children}
        </div>
    )
}
