'use client'

import { FilterList, Search } from '@mui/icons-material'
import { type FormEvent, useState } from 'react'

import { Button } from '@/components/bases/Button'
import { Input, InputVariant } from '@/components/bases/Input'
import { MultiSelectForm, type MultiSelectFormOption } from '@/components/bases/MultiSelectForm'
import { SelectForm, type SelectFormOption } from '@/components/bases/SelectForm'

import styles from './styles.module.scss'

interface Props {
    activeFilterCount: number
    activeStatusOptions: SelectFormOption<string>[]
    activeStatusValue: string
    categoryOptions: SelectFormOption<string>[]
    categoryValue: string
    isClearDisabled?: boolean
    isSearchDisabled?: boolean
    maxPriceValue: string
    minPriceValue: string
    onActiveStatusChange: (_value: string) => void
    onCategoryChange: (_value: string) => void
    onClear: () => void
    onMaxPriceChange: (_value: string) => void
    onMinPriceChange: (_value: string) => void
    onRecommendStatusChange: (_value: string) => void
    onSearchTextChange: (_value: string) => void
    onSubmit: (_event: FormEvent<HTMLFormElement>) => void
    onTagsChange: (_value: string[]) => void
    recommendStatusOptions: SelectFormOption<string>[]
    recommendStatusValue: string
    searchText: string
    tagOptions: MultiSelectFormOption<string>[]
    tagValue: string[]
}

export const ProductSearchToolbar = ({
    activeFilterCount,
    activeStatusOptions,
    activeStatusValue,
    categoryOptions,
    categoryValue,
    isClearDisabled = false,
    isSearchDisabled = false,
    maxPriceValue,
    minPriceValue,
    onActiveStatusChange,
    onCategoryChange,
    onClear,
    onMaxPriceChange,
    onMinPriceChange,
    onRecommendStatusChange,
    onSearchTextChange,
    onSubmit,
    onTagsChange,
    recommendStatusOptions,
    recommendStatusValue,
    searchText,
    tagOptions,
    tagValue,
}: Props) => {
    const [isDetailOpen, setIsDetailOpen] = useState<boolean>(false)

    return (
        <form className={styles['search-form']} onSubmit={onSubmit}>
            {isDetailOpen && (
                <div
                    className={styles['detail-backdrop']}
                    onClick={() => {
                        setIsDetailOpen(false)
                    }}
                />
            )}
            <div className={styles['search-bar']}>
                <div className={styles['search-bar-field']}>
                    <Search className={styles['search-bar-icon']} fontSize="small" />
                    <Input
                        aria-label="商品名で検索"
                        onChange={(event) => {
                            onSearchTextChange(event.target.value)
                        }}
                        placeholder="商品名で検索"
                        value={searchText}
                        variant={InputVariant.Bare}
                    />
                </div>
                <Button contrast disabled={isClearDisabled} onClick={onClear} pill type="button">
                    クリア
                </Button>
                <Button
                    aria-expanded={isDetailOpen}
                    contrast={!isDetailOpen}
                    onClick={() => {
                        setIsDetailOpen((current) => !current)
                    }}
                    pill
                    type="button"
                >
                    <div className={styles['detail-toggle-content']}>
                        <FilterList fontSize="small" />
                        詳細
                        {activeFilterCount > 0 && <span className={styles['detail-badge']}>{activeFilterCount}</span>}
                    </div>
                </Button>
                <Button disabled={isSearchDisabled} pill type="submit">
                    検索
                </Button>
            </div>
            <div className={`${styles['detail-panel']} ${isDetailOpen ? styles['detail-panel-open'] : ''}`}>
                <div className={styles['detail-grid']}>
                    <SelectForm
                        id="product-search-toolbar-category"
                        onChange={(value) => {
                            onCategoryChange(value ?? '')
                        }}
                        options={categoryOptions}
                        placeholder="カテゴリ"
                        value={categoryValue}
                    />
                    <MultiSelectForm
                        id="product-search-toolbar-tag"
                        onChange={onTagsChange}
                        options={tagOptions}
                        placeholder="タグ"
                        value={tagValue}
                    />
                    <div className={styles['price-range']}>
                        <Input
                            aria-label="最低価格"
                            inputMode="numeric"
                            min={0}
                            onChange={(event) => {
                                onMinPriceChange(event.target.value)
                            }}
                            placeholder="最低価格"
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
                            placeholder="最高価格"
                            type="number"
                            value={maxPriceValue}
                        />
                    </div>
                    <SelectForm
                        id="product-search-toolbar-active-status"
                        onChange={(value) => {
                            onActiveStatusChange(value ?? '')
                        }}
                        options={activeStatusOptions}
                        placeholder="公開状態"
                        value={activeStatusValue}
                    />
                    <SelectForm
                        id="product-search-toolbar-recommend-status"
                        onChange={(value) => {
                            onRecommendStatusChange(value ?? '')
                        }}
                        options={recommendStatusOptions}
                        placeholder="おすすめ状態"
                        value={recommendStatusValue}
                    />
                </div>
            </div>
        </form>
    )
}
