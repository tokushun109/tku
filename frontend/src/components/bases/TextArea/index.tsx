import React, { useState } from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

interface Props extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
    error?: string
    helperText?: string
    label?: string
    required?: boolean
}

export const TextArea = React.forwardRef<HTMLTextAreaElement, Props>(({ label, error, required, helperText, className, onChange, ...props }, ref) => {
    const [internalValue, setInternalValue] = useState<string>(String(props.value || ''))
    const textareaId = props.id || props.name

    const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        setInternalValue(e.target.value)
        onChange?.(e)
    }

    const isEmpty = !internalValue || internalValue.trim() === ''

    return (
        <div className={`${styles['form-field']} ${required ? styles['require-form'] : ''}`}>
            {label && (
                <label className={styles.label} htmlFor={textareaId}>
                    {label}
                </label>
            )}

            <div className={styles['textarea-container']}>
                {required && isEmpty && (
                    <div className={styles['chip-container']}>
                        <Chip color={ColorType.Danger} fontColor="#b84150" fontSize={FontSizeType.SmMd} size={ChipSize.Small}>
                            必須
                        </Chip>
                    </div>
                )}
                <textarea
                    className={`${styles.textarea} ${error ? styles.error : ''} ${className || ''}`}
                    id={textareaId}
                    onChange={handleChange}
                    ref={ref}
                    {...props}
                />
            </div>

            {error && <span className={styles['field-error']}>{error}</span>}
            {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
        </div>
    )
})

TextArea.displayName = 'TextArea'
