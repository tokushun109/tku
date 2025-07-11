import { ProductsByCategoryDisplay } from '.'
import { IProductsByCategory } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockProductsByCategory: IProductsByCategory = {
    category: { uuid: '1', name: 'ピアス' },
    products: [
        {
            uuid: '1',
            name: 'ハンドメイドピアス 1',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 1500,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '1', apiPath: '/image/about/concept1.jpg', name: 'image1', order: 1 }],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
        {
            uuid: '2',
            name: 'ハンドメイドピアス 2',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 2000,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '2', apiPath: '/image/about/concept2.jpg', name: 'image2', order: 1 }],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
        {
            uuid: '3',
            name: 'ハンドメイドピアス 3',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 1800,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '3', apiPath: '/image/about/concept1.jpg', name: 'image3', order: 1 }],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
        {
            uuid: '4',
            name: 'ハンドメイドピアス 4',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 2500,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '4', apiPath: '/image/about/concept2.jpg', name: 'image4', order: 1 }],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
        {
            uuid: '5',
            name: 'ハンドメイドピアス 5',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 3000,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '5', apiPath: '/image/about/concept1.jpg', name: 'image5', order: 1 }],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
        {
            uuid: '6',
            name: 'ハンドメイドピアス 6',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 2200,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '6', apiPath: '/image/about/concept2.jpg', name: 'image6', order: 1 }],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
    ],
}

const meta: Meta<typeof ProductsByCategoryDisplay> = {
    component: ProductsByCategoryDisplay,
    args: {
        productsByCategory: mockProductsByCategory,
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof ProductsByCategoryDisplay>

export const Default: Story = {}

export const FewProducts: Story = {
    args: {
        productsByCategory: {
            ...mockProductsByCategory,
            products: mockProductsByCategory.products.slice(0, 3),
        },
    },
}

export const ExactlyFourProducts: Story = {
    args: {
        productsByCategory: {
            ...mockProductsByCategory,
            products: mockProductsByCategory.products.slice(0, 4),
        },
    },
}

export const ManyProducts: Story = {
    args: {
        productsByCategory: {
            ...mockProductsByCategory,
            products: [
                ...mockProductsByCategory.products,
                {
                    uuid: '7',
                    name: 'ハンドメイドピアス 7',
                    description: 'シンプルで上品なデザインのハンドメイドピアスです。',
                    price: 2800,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '7', apiPath: '/image/about/concept1.jpg', name: 'image7', order: 1 }],
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '1', name: 'ピアス' },
                    tags: [],
                    siteDetails: [],
                },
                {
                    uuid: '8',
                    name: 'ハンドメイドピアス 8',
                    description: 'シンプルで上品なデザインのハンドメイドピアスです。',
                    price: 3500,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '8', apiPath: '/image/about/concept2.jpg', name: 'image8', order: 1 }],
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '1', name: 'ピアス' },
                    tags: [],
                    siteDetails: [],
                },
            ],
        },
    },
}

export const NecklaceCategory: Story = {
    args: {
        productsByCategory: {
            category: { uuid: '2', name: 'ネックレス' },
            products: [
                {
                    uuid: '1',
                    name: 'シンプルネックレス',
                    description: 'シンプルで上品なデザインのネックレスです。',
                    price: 3500,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '1', apiPath: '/image/about/concept1.jpg', name: 'image1', order: 1 }],
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '2', name: 'ネックレス' },
                    tags: [],
                    siteDetails: [],
                },
                {
                    uuid: '2',
                    name: 'パールネックレス',
                    description: 'エレガントなパールネックレスです。',
                    price: 4500,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '2', apiPath: '/image/about/concept2.jpg', name: 'image2', order: 1 }],
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '2', name: 'ネックレス' },
                    tags: [],
                    siteDetails: [],
                },
                {
                    uuid: '3',
                    name: 'ゴールドネックレス',
                    description: '高級感のあるゴールドネックレスです。',
                    price: 8000,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '3', apiPath: '/image/about/concept1.jpg', name: 'image3', order: 1 }],
                    target: { uuid: '3', name: 'ユニセックス' },
                    category: { uuid: '2', name: 'ネックレス' },
                    tags: [],
                    siteDetails: [],
                },
                {
                    uuid: '4',
                    name: 'メンズネックレス',
                    description: 'スタイリッシュなメンズネックレスです。',
                    price: 5500,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '4', apiPath: '/image/about/concept2.jpg', name: 'image4', order: 1 }],
                    target: { uuid: '2', name: 'メンズ' },
                    category: { uuid: '2', name: 'ネックレス' },
                    tags: [],
                    siteDetails: [],
                },
                {
                    uuid: '5',
                    name: 'チェーンネックレス',
                    description: 'クールなチェーンネックレスです。',
                    price: 6000,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '5', apiPath: '/image/about/concept2.jpg', name: 'image5', order: 1 }],
                    target: { uuid: '2', name: 'メンズ' },
                    category: { uuid: '2', name: 'ネックレス' },
                    tags: [],
                    siteDetails: [],
                },
            ],
        },
    },
}