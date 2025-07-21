import React from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

interface Props extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
    error?: string
    helperText?: string
    label?: string
    required?: boolean
}

export const TextArea = React.forwardRef<HTMLTextAreaElement, Props>(({ label, error, required, helperText, className, ...props }, ref) => {
    const textareaId = props.id || props.name

    return (
        <div className={`${styles['form-field']} ${required ? styles['require-form'] : ''}`}>
            {label && (
                <label className={styles.label} htmlFor={textareaId}>
                    {label}
                </label>
            )}

            {required && (
                <div className={styles['chip-container']}>
                    <Chip color={ColorType.Danger} fontColor="#b84150" fontSize={FontSizeType.SmMd} size={ChipSize.Small}>
                        必須
                    </Chip>
                </div>
            )}

            <textarea className={`${styles.textarea} ${error ? styles.error : ''} ${className || ''}`} id={textareaId} ref={ref} {...props} />

            {error && <span className={styles['field-error']}>{error}</span>}
            {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
        </div>
    )
})

TextArea.displayName = 'TextArea'
