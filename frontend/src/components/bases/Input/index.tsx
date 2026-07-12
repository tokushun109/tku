import React, { useState } from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

export const InputVariant = {
    Default: 'default',
    Bare: 'bare',
} as const
export type InputVariant = (typeof InputVariant)[keyof typeof InputVariant]

interface Props extends React.InputHTMLAttributes<HTMLInputElement> {
    error?: string
    helperText?: string
    label?: string
    required?: boolean
    variant?: InputVariant
}

export const Input = React.forwardRef<HTMLInputElement, Props>(
    ({ label, error, required, helperText, variant = InputVariant.Default, className, onChange, ...props }, ref) => {
        const [internalValue, setInternalValue] = useState<string>(String(props.value || ''))
        const inputId = props.id || props.name

        const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            setInternalValue(e.target.value)
            onChange?.(e)
        }

        const isEmpty = !internalValue || internalValue.trim() === ''

        return (
            <div className={`${styles['form-field']} ${required ? styles['require-form'] : ''}`}>
                {label && (
                    <label className={styles.label} htmlFor={inputId}>
                        {label}
                    </label>
                )}

                <div className={styles['input-container']}>
                    {required && isEmpty && (
                        <div className={styles['chip-container']}>
                            <Chip color={ColorType.Danger} fontColor="#b84150" fontSize={FontSizeType.SmMd} size={ChipSize.Small}>
                                必須
                            </Chip>
                        </div>
                    )}
                    <input
                        className={`${styles.input} ${variant === InputVariant.Bare ? styles['input-bare'] : ''} ${error ? styles.error : ''} ${className || ''}`}
                        id={inputId}
                        onChange={handleChange}
                        ref={ref}
                        {...props}
                    />
                </div>

                {error && <span className={styles['field-error']}>{error}</span>}
                {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
            </div>
        )
    },
)

Input.displayName = 'Input'
