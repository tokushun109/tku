'use client'

import { ExpandMore } from '@mui/icons-material'
import { useEffect, useRef, useState } from 'react'

import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

export type SelectFormOption<T = string> = {
    label: string
    value: T
}

interface Props<T = string> {
    error?: string
    helperText?: string
    id?: string
    label?: string
    onChange?: (_value: T) => void
    options: SelectFormOption<T>[]
    placeholder?: string
    required?: boolean
    value?: T
}

export const SelectForm = <T,>({ label, options, value, onChange, placeholder, required, error, helperText, id }: Props<T>) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [searchText, setSearchText] = useState<string>('')
    const [filteredOptions, setFilteredOptions] = useState<SelectFormOption<T>[]>(options)
    const dropdownRef = useRef<HTMLDivElement>(null)
    const inputRef = useRef<HTMLInputElement>(null)

    // 選択された項目を取得
    const selectedOption = options.find((option) => option.value === value)

    // 検索機能
    useEffect(() => {
        if (searchText) {
            const filtered = options.filter((option) => option.label.toLowerCase().includes(searchText.toLowerCase()))
            setFilteredOptions(filtered)
        } else {
            setFilteredOptions(options)
        }
    }, [searchText, options])

    // 外部クリックでドロップダウンを閉じる
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
                setIsOpen(false)
                setSearchText('')
            }
        }

        document.addEventListener('mousedown', handleClickOutside)
        return () => {
            document.removeEventListener('mousedown', handleClickOutside)
        }
    }, [])

    // ドロップダウンが開いたときに入力フィールドにフォーカス
    useEffect(() => {
        if (isOpen && inputRef.current) {
            inputRef.current.focus()
        }
    }, [isOpen])

    const handleToggle = () => {
        setIsOpen(!isOpen)
        if (!isOpen) {
            setSearchText('')
        }
    }

    const handleOptionClick = (option: SelectFormOption<T>) => {
        onChange?.(option.value)
        setIsOpen(false)
        setSearchText('')
    }

    const handleClear = () => {
        onChange?.(undefined as T)
        setIsOpen(false)
        setSearchText('')
    }

    const inputId = id || `select-form-${Math.random().toString(36).substr(2, 9)}`

    return (
        <div className={`${styles['select-form']} ${required ? styles['require-form'] : ''}`}>
            {label && (
                <label className={styles['label']} htmlFor={inputId}>
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

            <div className={styles['select-container']} ref={dropdownRef}>
                <div className={`${styles['select-trigger']} ${error ? styles['error'] : ''} ${isOpen ? styles['open'] : ''}`} onClick={handleToggle}>
                    {selectedOption ? (
                        <div className={styles['selected-item']}>
                            <span className={styles['selected-text']}>{selectedOption.label}</span>
                            <button
                                className={styles['clear-button']}
                                onClick={(e) => {
                                    e.stopPropagation()
                                    handleClear()
                                }}
                                type="button"
                            >
                                ×
                            </button>
                        </div>
                    ) : (
                        <span className={styles['placeholder']}>{placeholder || '選択してください'}</span>
                    )}
                    <ExpandMore className={`${styles['chevron']} ${isOpen ? styles['chevron-open'] : ''}`} />
                </div>

                {isOpen && (
                    <div className={styles['dropdown']}>
                        <div className={styles['search-container']}>
                            <input
                                className={styles['search-input']}
                                onChange={(e) => setSearchText(e.target.value)}
                                placeholder="検索..."
                                ref={inputRef}
                                type="text"
                                value={searchText}
                            />
                        </div>

                        <div className={styles['options-container']}>
                            {filteredOptions.length === 0 ? (
                                <div className={styles['no-options']}>該当する項目が見つかりません</div>
                            ) : (
                                filteredOptions.map((option) => (
                                    <div
                                        className={`${styles['option']} ${option.value === value ? styles['option-selected'] : ''}`}
                                        key={String(option.value)}
                                        onClick={() => handleOptionClick(option)}
                                    >
                                        {option.label}
                                    </div>
                                ))
                            )}
                        </div>
                    </div>
                )}
            </div>

            {error && <span className={styles['field-error']}>{error}</span>}
            {helperText && !error && <span className={styles['helper-text']}>{helperText}</span>}
        </div>
    )
}
