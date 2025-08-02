import { FileUpload } from '@mui/icons-material'
import React, { useRef } from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

interface Props extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'onChange' | 'type' | 'value'> {
    error?: string
    helperText?: string
    label?: string
    onChange?: (_file: File | null) => void
    required?: boolean
    value?: File | null
}

export const FileInput = React.forwardRef<HTMLInputElement, Props>(
    ({ label, error, required, helperText, className, onChange, value, accept, ...props }, _ref) => {
        const fileInputRef = useRef<HTMLInputElement>(null)
        const inputId = props.id || props.name

        const handleContainerClick = () => {
            fileInputRef.current?.click()
        }

        const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
            const selectedFile = event.target.files?.[0] || null
            onChange?.(selectedFile)
        }

        const fileName = value?.name || 'ファイルが選択されていません'

        return (
            <div className={`${styles['form-field']} ${required ? styles['require-form'] : ''}`}>
                {label && (
                    <label className={styles.label} htmlFor={inputId}>
                        {label}
                    </label>
                )}

                <div className={`${styles['file-input-container']} ${error ? styles.error : ''} ${className || ''}`} onClick={handleContainerClick}>
                    <input
                        accept={accept}
                        className={styles['hidden-input']}
                        id={inputId}
                        onChange={handleFileChange}
                        ref={fileInputRef}
                        type="file"
                        {...props}
                    />

                    {required && (
                        <div className={styles['chip-container']}>
                            <Chip color={ColorType.Danger} fontColor="#b84150" fontSize={FontSizeType.SmMd} size={ChipSize.Small}>
                                必須
                            </Chip>
                        </div>
                    )}

                    <div className={styles['file-display']}>
                        <div className={styles['file-info']}>
                            <FileUpload className={styles['file-icon']} />
                            <span className={`${styles['file-name']} ${!value ? styles['placeholder'] : ''}`}>{fileName}</span>
                        </div>
                    </div>
                </div>

                {error && <span className={styles['field-error']}>{error}</span>}
                {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
            </div>
        )
    },
)

FileInput.displayName = 'FileInput'
