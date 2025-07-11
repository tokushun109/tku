import { ProductThumbnail } from '.'
import { IThumbnail } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockThumbnail: IThumbnail = {
    uuid: '1',
    name: 'ハンドメイドピアス',
    price: 1500,
    imageUrl: 'https://via.placeholder.com/300x300',
    target: { uuid: '1', name: 'レディース' },
    category: { uuid: '1', name: 'ピアス' },
}

const meta: Meta<typeof ProductThumbnail> = {
    component: ProductThumbnail,
    args: {
        item: mockThumbnail,
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof ProductThumbnail>

export const Default: Story = {}

export const MensAccessory: Story = {
    args: {
        item: {
            ...mockThumbnail,
            uuid: '2',
            name: 'メンズネックレス',
            price: 2500,
            target: { uuid: '2', name: 'メンズ' },
            category: { uuid: '2', name: 'ネックレス' },
        },
    },
}

export const ExpensiveItem: Story = {
    args: {
        item: {
            ...mockThumbnail,
            uuid: '3',
            name: 'プレミアムリング',
            price: 15000,
            target: { uuid: '3', name: 'ユニセックス' },
            category: { uuid: '3', name: 'リング' },
        },
    },
}

export const LongProductName: Story = {
    args: {
        item: {
            ...mockThumbnail,
            uuid: '4',
            name: 'とても長い商品名のハンドメイドアクセサリー',
            price: 3000,
            category: { uuid: '4', name: 'ブレスレット' },
        },
    },
}

export const AffordableItem: Story = {
    args: {
        item: {
            ...mockThumbnail,
            uuid: '5',
            name: 'シンプルピアス',
            price: 800,
            category: { uuid: '5', name: 'ピアス' },
        },
    },
}

export const HairAccessory: Story = {
    args: {
        item: {
            ...mockThumbnail,
            uuid: '6',
            name: 'ヘアクリップ',
            price: 1200,
            category: { uuid: '6', name: 'ヘアアクセサリー' },
        },
    },
}