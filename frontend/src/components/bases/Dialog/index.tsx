import React from 'react'

import { Button } from '@/components/bases/Button'
import { ColorType } from '@/types/enum/color'

import styles from './styles.module.scss'

interface DialogButtonOption {
    label: string
    onClick: () => void
}

interface DialogProps {
    cancelOption?: DialogButtonOption
    children: React.ReactNode
    confirmOption?: DialogButtonOption
    isOpen: boolean
    onClose: () => void
    title?: string
}

export const Dialog: React.FC<DialogProps> = ({ isOpen, onClose, title, children, confirmOption, cancelOption }) => {
    if (!isOpen) return null

    const handleBackdropClick = () => {
        onClose()
    }

    const handleDialogClick = (e: React.MouseEvent) => {
        e.stopPropagation()
    }

    const showActions = confirmOption || cancelOption

    return (
        <div className={styles['dialog-overlay']} onClick={handleBackdropClick}>
            <div className={styles['dialog-backdrop']} />
            <div className={styles['dialog']} onClick={handleDialogClick}>
                {title && (
                    <div className={styles['dialog-header']}>
                        <h3 className={styles['dialog-title']}>{title}</h3>
                    </div>
                )}
                <div className={styles['dialog-content']}>{children}</div>
                {showActions && (
                    <div className={styles['dialog-actions']}>
                        {cancelOption && (
                            <Button colorType={ColorType.Primary} contrast onClick={cancelOption.onClick} outlined>
                                {cancelOption.label}
                            </Button>
                        )}
                        {confirmOption && (
                            <Button colorType={ColorType.Primary} onClick={confirmOption.onClick}>
                                {confirmOption.label}
                            </Button>
                        )}
                    </div>
                )}
            </div>
        </div>
    )
}
