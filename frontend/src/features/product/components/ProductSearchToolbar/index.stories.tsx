import { useState } from 'react'

import { ProductSearchToolbar } from '.'

import type { MultiSelectFormOption } from '@/components/bases/MultiSelectForm'
import type { SelectFormOption } from '@/components/bases/SelectForm'
import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const categoryOptions: SelectFormOption<string>[] = [
    { label: 'すべてのカテゴリ', value: 'all' },
    { label: 'ネックレス', value: 'necklace' },
    { label: 'ブレスレット', value: 'bracelet' },
    { label: 'ピアス', value: 'earrings' },
]

const tagOptions: MultiSelectFormOption<string>[] = [
    { label: '天然石', value: 'natural-stone' },
    { label: 'ベビー・キッズ', value: 'baby-kids' },
    { label: '静電気除去', value: 'anti-static' },
]

const activeStatusOptions: SelectFormOption<string>[] = [
    { label: '公開状態すべて', value: 'all' },
    { label: '公開中', value: 'active' },
    { label: '非公開', value: 'inactive' },
]

const recommendStatusOptions: SelectFormOption<string>[] = [
    { label: 'おすすめ状態すべて', value: 'all' },
    { label: 'おすすめ', value: 'recommended' },
    { label: 'おすすめ以外', value: 'not_recommended' },
]

type ToolbarFilters = {
    activeStatus: string
    category: string
    maxPrice: string
    minPrice: string
    recommendStatus: string
    tags: string[]
}

const defaultFilters: ToolbarFilters = {
    activeStatus: 'all',
    category: 'all',
    maxPrice: '',
    minPrice: '',
    recommendStatus: 'all',
    tags: [],
}

const countActiveFilters = (filters: ToolbarFilters) =>
    [
        filters.category !== 'all',
        filters.tags.length > 0,
        filters.minPrice !== '',
        filters.maxPrice !== '',
        filters.activeStatus !== 'all',
        filters.recommendStatus !== 'all',
    ].filter(Boolean).length

const InteractiveToolbar = ({
    initialFilters = defaultFilters,
    isClearDisabled = false,
}: {
    initialFilters?: ToolbarFilters
    isClearDisabled?: boolean
}) => {
    const [searchText, setSearchText] = useState<string>('')
    const [filters, setFilters] = useState<ToolbarFilters>(initialFilters)

    return (
        <ProductSearchToolbar
            activeFilterCount={countActiveFilters(filters)}
            activeStatusOptions={activeStatusOptions}
            activeStatusValue={filters.activeStatus}
            categoryOptions={categoryOptions}
            categoryValue={filters.category}
            isClearDisabled={isClearDisabled}
            maxPriceValue={filters.maxPrice}
            minPriceValue={filters.minPrice}
            onActiveStatusChange={(value) => {
                setFilters((current) => ({ ...current, activeStatus: value }))
            }}
            onCategoryChange={(value) => {
                setFilters((current) => ({ ...current, category: value }))
            }}
            onClear={() => {
                setSearchText('')
                setFilters(defaultFilters)
            }}
            onMaxPriceChange={(value) => {
                setFilters((current) => ({ ...current, maxPrice: value }))
            }}
            onMinPriceChange={(value) => {
                setFilters((current) => ({ ...current, minPrice: value }))
            }}
            onRecommendStatusChange={(value) => {
                setFilters((current) => ({ ...current, recommendStatus: value }))
            }}
            onSearchTextChange={setSearchText}
            onSubmit={(event) => {
                event.preventDefault()
                console.log('検索', { searchText, filters })
            }}
            onTagsChange={(value) => {
                setFilters((current) => ({ ...current, tags: value }))
            }}
            recommendStatusOptions={recommendStatusOptions}
            recommendStatusValue={filters.recommendStatus}
            searchText={searchText}
            tagOptions={tagOptions}
            tagValue={filters.tags}
        />
    )
}

const meta: Meta<typeof ProductSearchToolbar> = {
    component: ProductSearchToolbar,
}

export default meta
type Story = StoryObj<typeof ProductSearchToolbar>

export const Default: Story = {
    render: () => <InteractiveToolbar />,
}

export const WithActiveFilters: Story = {
    render: () => (
        <InteractiveToolbar
            initialFilters={{
                activeStatus: 'active',
                category: 'bracelet',
                maxPrice: '5000',
                minPrice: '1000',
                recommendStatus: 'all',
                tags: ['natural-stone'],
            }}
        />
    ),
}

export const ClearDisabled: Story = {
    render: () => <InteractiveToolbar isClearDisabled />,
}
