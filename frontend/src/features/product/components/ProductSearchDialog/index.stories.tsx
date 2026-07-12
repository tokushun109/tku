import { useState } from 'react'

import { ProductSearchDialog } from '.'

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

const InteractiveDialog = ({
    initialFilters = defaultFilters,
    isClearDisabled = false,
}: {
    initialFilters?: ToolbarFilters
    isClearDisabled?: boolean
}) => {
    const [isOpen, setIsOpen] = useState<boolean>(true)
    const [searchText, setSearchText] = useState<string>('')
    const [filters, setFilters] = useState<ToolbarFilters>(initialFilters)

    return (
        <ProductSearchDialog
            activeStatusOptions={activeStatusOptions}
            activeStatusValue={filters.activeStatus}
            categoryOptions={categoryOptions}
            categoryValue={filters.category}
            isClearDisabled={isClearDisabled}
            isOpen={isOpen}
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
            onClose={() => {
                setIsOpen(false)
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

const meta: Meta<typeof ProductSearchDialog> = {
    component: ProductSearchDialog,
}

export default meta
type Story = StoryObj<typeof ProductSearchDialog>

export const Default: Story = {
    render: () => <InteractiveDialog />,
}

export const WithActiveFilters: Story = {
    render: () => (
        <InteractiveDialog
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
    render: () => <InteractiveDialog isClearDisabled />,
}
