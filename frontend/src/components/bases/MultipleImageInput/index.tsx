import { FileUpload } from '@mui/icons-material'
import React, { useRef } from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

interface Props extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'onChange' | 'type' | 'value' | 'accept'> {
    error?: string
    helperText?: string
    label?: string
    onChange?: (_files: File[]) => void
    required?: boolean
    value?: File[]
}

export const MultipleImageInput = React.forwardRef<HTMLInputElement, Props>(
    ({ label, error, required, helperText, className, onChange, value = [], ...props }, _ref) => {
        const fileInputRef = useRef<HTMLInputElement>(null)
        const inputId = props.id || props.name

        const handleContainerClick = () => {
            fileInputRef.current?.click()
        }

        const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
            const selectedFiles = Array.from(event.target.files || [])
            const imageFiles = selectedFiles.filter((file) => file.type.startsWith('image/'))
            onChange?.(imageFiles)
        }

        const fileCount = value.length
        const displayText = fileCount > 0 ? `${fileCount}個の画像が選択されています` : '画像を選択してください'

        console.log(value, 'アップロードファイル')

        return (
            <div className={`${styles['form-field']} ${required ? styles['require-form'] : ''}`}>
                {label && (
                    <label className={styles.label} htmlFor={inputId}>
                        {label}
                    </label>
                )}

                <div className={`${styles['file-input-container']} ${error ? styles.error : ''} ${className || ''}`} onClick={handleContainerClick}>
                    <input
                        key={value.map((file) => file.name).join(', ')}
                        accept="image/*"
                        className={styles['hidden-input']}
                        id={inputId}
                        multiple
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
                            <span className={`${styles['file-name']} ${fileCount === 0 ? styles['placeholder'] : ''}`}>{displayText}</span>
                        </div>
                    </div>
                </div>

                {error && <span className={styles['field-error']}>{error}</span>}
                {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
            </div>
        )
    },
)

MultipleImageInput.displayName = 'MultipleImageInput'
