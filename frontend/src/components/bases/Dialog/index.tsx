import classNames from 'classnames'
import React from 'react'

import { Button } from '@/components/bases/Button'
import { ColorType } from '@/types/enum/color'

import styles from './styles.module.scss'

interface DialogButtonOption {
    disabled?: boolean
    label: string
    onClick: () => void
}

interface Props {
    cancelOption?: DialogButtonOption
    children: React.ReactNode
    confirmOption?: DialogButtonOption
    isOpen: boolean
    onClose: () => void
    title?: string
    wide?: boolean
}

export const Dialog = ({ isOpen, onClose, title, children, confirmOption, cancelOption, wide = false }: Props) => {
    if (!isOpen) return null

    const handleBackdropClick = () => {
        onClose()
    }

    const showActions = confirmOption || cancelOption

    return (
        <div className={styles['dialog-overlay']} onClick={handleBackdropClick}>
            <div className={styles['dialog-backdrop']} />
            <div
                className={classNames(`${styles['dialog']} ${wide ? styles['wide'] : ''}`)}
                onClick={(e) => {
                    e.stopPropagation()
                }}
            >
                {title && (
                    <div className={styles['dialog-header']}>
                        <h3 className={styles['dialog-title']}>{title}</h3>
                    </div>
                )}
                <div className={styles['dialog-content']}>{children}</div>
                {showActions && (
                    <div className={styles['dialog-actions']}>
                        {cancelOption && (
                            <Button colorType={ColorType.Primary} contrast disabled={cancelOption.disabled} onClick={cancelOption.onClick} outlined>
                                {cancelOption.label}
                            </Button>
                        )}
                        {confirmOption && (
                            <Button colorType={ColorType.Primary} disabled={confirmOption.disabled} onClick={confirmOption.onClick} type="submit">
                                {confirmOption.label}
                            </Button>
                        )}
                    </div>
                )}
            </div>
        </div>
    )
}
