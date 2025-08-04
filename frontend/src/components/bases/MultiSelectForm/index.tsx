'use client'

import { Close, ExpandMore } from '@mui/icons-material'
import { useEffect, useRef, useState } from 'react'

import { Checkbox } from '@/components/bases/Checkbox'
import { Chip, ChipSize } from '@/components/bases/Chip'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'

export type MultiSelectFormOption<T = string> = {
    label: string
    value: T
}

interface Props<T = string> {
    error?: string
    helperText?: string
    id?: string
    label?: string
    onChange?: (_value: T[]) => void
    options: MultiSelectFormOption<T>[]
    placeholder?: string
    required?: boolean
    value?: T[]
}

export const MultiSelectForm = <T,>({ label, options, value = [], onChange, placeholder, required, error, helperText, id }: Props<T>) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [searchText, setSearchText] = useState<string>('')
    const [filteredOptions, setFilteredOptions] = useState<MultiSelectFormOption<T>[]>(options)
    const [isFocused, setIsFocused] = useState<boolean>(false)
    const dropdownRef = useRef<HTMLDivElement>(null)
    const inputRef = useRef<HTMLInputElement>(null)

    // 選択された項目を取得
    const selectedOptions = options.filter((option) => value.includes(option.value))

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
                setIsFocused(false)
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
        if (!isFocused) {
            setIsOpen(!isOpen)
            setIsFocused(true)
            if (!isOpen) {
                setSearchText('')
            }
        }
    }

    const handleInputFocus = () => {
        setIsFocused(true)
        setIsOpen(true)
    }

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchText(e.target.value)
        if (!isOpen) {
            setIsOpen(true)
        }
    }

    const handleOptionClick = (option: MultiSelectFormOption<T>) => {
        const newValues = value.includes(option.value) ? value.filter((v) => v !== option.value) : [...value, option.value]
        onChange?.(newValues)
    }

    const handleRemoveSelectedOption = (optionValue: T) => {
        const newValues = value.filter((v) => v !== optionValue)
        onChange?.(newValues)
    }

    const inputId = id || `multi-select-form-${Math.random().toString(36).substr(2, 9)}`

    return (
        <div className={`${styles['multi-select-form']} ${required ? styles['require-form'] : ''}`}>
            {label && (
                <label className={styles['label']} htmlFor={inputId}>
                    {label}
                </label>
            )}

            <div className={styles['select-container']} ref={dropdownRef}>
                {required && value.length === 0 && (
                    <div className={styles['chip-container']}>
                        <Chip color={ColorType.Danger} fontColor="#b84150" fontSize={FontSizeType.SmMd} size={ChipSize.Small}>
                            必須
                        </Chip>
                    </div>
                )}
                <div className={`${styles['select-trigger']} ${error ? styles['error'] : ''} ${isOpen ? styles['open'] : ''}`} onClick={handleToggle}>
                    <div className={styles['trigger-content']}>
                        {selectedOptions.map((option) => (
                            <div className={styles['selected-chip-container']} key={String(option.value)}>
                                <Chip color={ColorType.Secondary} fontColor="#ffffff" fontSize={FontSizeType.SmMd} size={ChipSize.Small}>
                                    {option.label}
                                    <button
                                        className={styles['chip-close-button']}
                                        onClick={(e) => {
                                            e.stopPropagation()
                                            handleRemoveSelectedOption(option.value)
                                        }}
                                        type="button"
                                    >
                                        <Close />
                                    </button>
                                </Chip>
                            </div>
                        ))}
                        {isFocused || isOpen ? (
                            <input
                                className={`${styles['search-trigger-input']} ${selectedOptions.length > 0 ? styles['with-chip'] : ''}`}
                                onChange={handleInputChange}
                                onFocus={handleInputFocus}
                                placeholder="検索..."
                                ref={inputRef}
                                type="text"
                                value={searchText}
                            />
                        ) : (
                            selectedOptions.length === 0 && <span className={styles['placeholder']}>{placeholder || '選択してください'}</span>
                        )}
                    </div>
                    <ExpandMore className={`${styles['chevron']} ${isOpen ? styles['chevron-open'] : ''}`} />
                </div>

                {isOpen && (
                    <div className={styles['dropdown']}>
                        <div className={styles['options-container']}>
                            {filteredOptions.length === 0 ? (
                                <div className={styles['no-options']}>該当する項目が見つかりません</div>
                            ) : (
                                filteredOptions.map((option) => {
                                    const isSelected = value.includes(option.value)
                                    return (
                                        <div
                                            className={`${styles['option']} ${isSelected ? styles['option-selected'] : ''}`}
                                            key={String(option.value)}
                                            onClick={() => handleOptionClick(option)}
                                        >
                                            <Checkbox checked={isSelected} className={styles['option-checkbox']} onChange={() => {}} />
                                            <span className={styles['option-label']}>{option.label}</span>
                                        </div>
                                    )
                                })
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
