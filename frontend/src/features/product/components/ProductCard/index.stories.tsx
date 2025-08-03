import { ProductCard } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ProductCard> = {
    component: ProductCard,
    args: {
        product: {
            uuid: '1',
            name: 'サンプル商品',
            description: 'これはサンプル商品の説明です。',
            price: 2500,
            isActive: true,
            isRecommend: false,
            productImages: [
                {
                    uuid: '1',
                    apiPath: '/image/gray-image.png',
                    name: 'sample-image',
                    order: 1,
                },
            ],
            category: {
                uuid: 'category-1',
                name: 'アクセサリー',
            },
            target: {
                uuid: 'target-1',
                name: '女性',
            },
            tags: [],
            siteDetails: [],
        },
        admin: true,
        onEdit: () => console.log('Edit clicked'),
        onDelete: () => console.log('Delete clicked'),
    },
    argTypes: {
        admin: {
            control: { type: 'boolean' },
        },
        onEdit: { action: 'edit' },
        onDelete: { action: 'delete' },
    },
}

export default meta
type Story = StoryObj<typeof ProductCard>

export const Default: Story = {}

export const AdminMode: Story = {
    args: {
        admin: true,
    },
}

export const CustomerMode: Story = {
    args: {
        admin: false,
    },
}

export const RecommendProduct: Story = {
    args: {
        product: {
            uuid: '2',
            name: 'おすすめ商品',
            description: 'これはおすすめ商品です。',
            price: 3500,
            isActive: true,
            isRecommend: true,
            productImages: [
                {
                    uuid: '2',
                    apiPath: '/image/gray-image.png',
                    name: 'recommend-image',
                    order: 1,
                },
            ],
            category: {
                uuid: 'category-2',
                name: 'ネックレス',
            },
            target: {
                uuid: 'target-2',
                name: '女性',
            },
            tags: [],
            siteDetails: [],
        },
        admin: true,
    },
}

export const InactiveProduct: Story = {
    args: {
        product: {
            uuid: '3',
            name: '展示中商品',
            description: 'これは展示中の商品です。',
            price: 1800,
            isActive: false,
            isRecommend: false,
            productImages: [],
            category: {
                uuid: 'category-3',
                name: 'ピアス',
            },
            target: {
                uuid: 'target-3',
                name: '女性',
            },
            tags: [],
            siteDetails: [],
        },
        admin: true,
    },
}

export const NoImage: Story = {
    args: {
        product: {
            uuid: '4',
            name: '画像なし商品',
            description: '画像がない商品です。',
            price: 1200,
            isActive: true,
            isRecommend: false,
            productImages: [],
            category: {
                uuid: 'category-4',
                name: 'リング',
            },
            target: {
                uuid: 'target-4',
                name: '女性',
            },
            tags: [],
            siteDetails: [],
        },
        admin: true,
    },
}
