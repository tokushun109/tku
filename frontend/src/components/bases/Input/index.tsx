import React from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

interface Props extends React.InputHTMLAttributes<HTMLInputElement> {
    error?: string
    helperText?: string
    label?: string
    required?: boolean
}

export const Input = React.forwardRef<HTMLInputElement, Props>(({ label, error, required, helperText, className, ...props }, ref) => {
    const inputId = props.id || props.name

    return (
        <div className={`${styles['form-field']} ${required ? styles['require-form'] : ''}`}>
            {label && (
                <label className={styles.label} htmlFor={inputId}>
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

            <input className={`${styles.input} ${error ? styles.error : ''} ${className || ''}`} id={inputId} ref={ref} {...props} />

            {error && <span className={styles['field-error']}>{error}</span>}
            {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
        </div>
    )
})

Input.displayName = 'Input'
