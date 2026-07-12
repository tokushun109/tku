'use client'

import { Search } from '@mui/icons-material'
import { type FormEvent } from 'react'

import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { Input, InputVariant } from '@/components/bases/Input'
import { MultiSelectForm, type MultiSelectFormOption } from '@/components/bases/MultiSelectForm'
import { SelectForm, type SelectFormOption } from '@/components/bases/SelectForm'
import { PRODUCT_KEYWORD_MAX_LENGTH } from '@/features/product/constants'

import styles from './styles.module.scss'

interface Props {
    activeStatusOptions: SelectFormOption<string>[]
    activeStatusValue: string
    categoryOptions: SelectFormOption<string>[]
    categoryValue: string
    isClearDisabled?: boolean
    isOpen: boolean
    isSearchDisabled?: boolean
    maxPriceValue: string
    minPriceValue: string
    onActiveStatusChange: (_value: string) => void
    onCategoryChange: (_value: string) => void
    onClear: () => void
    onClose: () => void
    onMaxPriceChange: (_value: string) => void
    onMinPriceChange: (_value: string) => void
    onRecommendStatusChange: (_value: string) => void
    onSearch: () => void
    onSearchTextChange: (_value: string) => void
    onTagsChange: (_value: string[]) => void
    recommendStatusOptions: SelectFormOption<string>[]
    recommendStatusValue: string
    searchText: string
    tagOptions: MultiSelectFormOption<string>[]
    tagValue: string[]
}

export const ProductSearchDialog = ({
    activeStatusOptions,
    activeStatusValue,
    categoryOptions,
    categoryValue,
    isClearDisabled = false,
    isOpen,
    isSearchDisabled = false,
    maxPriceValue,
    minPriceValue,
    onActiveStatusChange,
    onCategoryChange,
    onClear,
    onClose,
    onMaxPriceChange,
    onMinPriceChange,
    onRecommendStatusChange,
    onSearch,
    onSearchTextChange,
    onTagsChange,
    recommendStatusOptions,
    recommendStatusValue,
    searchText,
    tagOptions,
    tagValue,
}: Props) => {
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        onSearch()
    }

    return (
        <Dialog
            cancelOption={{ label: 'キャンセル', onClick: onClose }}
            confirmOption={{ disabled: isSearchDisabled, label: '検索', onClick: onSearch }}
            isOpen={isOpen}
            onClose={onClose}
            title="絞り込み"
            wide
        >
            <form className={styles['search-form']} onSubmit={handleSubmit}>
                <div className={styles['search-card']}>
                    <div className={styles['search-bar']}>
                        <div className={styles['search-bar-field']}>
                            <Search className={styles['search-bar-icon']} fontSize="small" />
                            <Input
                                aria-label="商品名で検索"
                                maxLength={PRODUCT_KEYWORD_MAX_LENGTH}
                                onChange={(event) => {
                                    onSearchTextChange(event.target.value)
                                }}
                                placeholder="商品名で検索"
                                value={searchText}
                                variant={InputVariant.Bare}
                            />
                        </div>
                    </div>
                    <div className={styles['search-divider']} />
                    <div className={styles['detail-grid']}>
                        <SelectForm
                            id="product-search-dialog-category"
                            label="カテゴリ"
                            onChange={(value) => {
                                onCategoryChange(value ?? '')
                            }}
                            options={categoryOptions}
                            placeholder="すべて"
                            value={categoryValue}
                        />
                        <MultiSelectForm
                            id="product-search-dialog-tag"
                            label="タグ"
                            onChange={onTagsChange}
                            options={tagOptions}
                            placeholder="すべて"
                            value={tagValue}
                        />
                        <div className={styles['price-field']}>
                            <span className={styles['field-label']}>価格帯（円）</span>
                            <div className={styles['price-range']}>
                                <Input
                                    aria-label="最低価格"
                                    inputMode="numeric"
                                    min={0}
                                    onChange={(event) => {
                                        onMinPriceChange(event.target.value)
                                    }}
                                    placeholder="最低"
                                    type="number"
                                    value={minPriceValue}
                                />
                                <span className={styles['price-separator']}>〜</span>
                                <Input
                                    aria-label="最高価格"
                                    inputMode="numeric"
                                    min={0}
                                    onChange={(event) => {
                                        onMaxPriceChange(event.target.value)
                                    }}
                                    placeholder="最高"
                                    type="number"
                                    value={maxPriceValue}
                                />
                            </div>
                        </div>
                        <SelectForm
                            id="product-search-dialog-active-status"
                            label="公開状態"
                            onChange={(value) => {
                                onActiveStatusChange(value ?? '')
                            }}
                            options={activeStatusOptions}
                            placeholder="すべて"
                            value={activeStatusValue}
                        />
                        <SelectForm
                            id="product-search-dialog-recommend-status"
                            label="おすすめ"
                            onChange={(value) => {
                                onRecommendStatusChange(value ?? '')
                            }}
                            options={recommendStatusOptions}
                            placeholder="すべて"
                            value={recommendStatusValue}
                        />
                    </div>
                    <div className={styles['card-actions']}>
                        <Button contrast disabled={isClearDisabled} onClick={onClear} outlined type="button">
                            クリア
                        </Button>
                    </div>
                </div>
            </form>
        </Dialog>
    )
}
