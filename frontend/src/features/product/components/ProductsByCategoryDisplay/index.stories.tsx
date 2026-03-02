import { IProductsByCategory } from '@/features/product/type'

import { ProductsByCategoryDisplay } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const mockProductsByCategory: IProductsByCategory = {
    category: { uuid: '1', name: 'ピアス' },
    pageInfo: {
        hasMore: true,
        nextCursor: 'next-cursor',
    },
    products: [
        {
            uuid: '1',
            name: 'ハンドメイドピアス 1',
            description: 'シンプルで上品なデザインのハンドメイドピアスです。',
            price: 1500,
            isActive: true,
            isRecommend: false,
            productImages: [{ uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', displayOrder: 1 }],
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
            productImages: [{ uuid: '2', apiPath: '/image/about/story.jpg', name: 'image2', displayOrder: 1 }],
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
            productImages: [{ uuid: '3', apiPath: '/image/about/story.jpg', name: 'image3', displayOrder: 1 }],
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
            productImages: [{ uuid: '4', apiPath: '/image/about/story.jpg', name: 'image4', displayOrder: 1 }],
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
            productImages: [{ uuid: '5', apiPath: '/image/about/story.jpg', name: 'image5', displayOrder: 1 }],
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
            productImages: [{ uuid: '6', apiPath: '/image/about/story.jpg', name: 'image6', displayOrder: 1 }],
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
        hasMore: true,
        isLoadingMore: false,
        onClickMoreButton: () => {},
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
        hasMore: false,
        productsByCategory: {
            ...mockProductsByCategory,
            products: mockProductsByCategory.products.slice(0, 3),
        },
    },
}

export const ExactlyFourProducts: Story = {
    args: {
        hasMore: false,
        productsByCategory: {
            ...mockProductsByCategory,
            products: mockProductsByCategory.products.slice(0, 4),
        },
    },
}

export const ManyProducts: Story = {
    args: {
        hasMore: true,
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
                    productImages: [{ uuid: '7', apiPath: '/image/about/story.jpg', name: 'image7', displayOrder: 1 }],
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
                    productImages: [{ uuid: '8', apiPath: '/image/about/story.jpg', name: 'image8', displayOrder: 1 }],
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
        hasMore: true,
        productsByCategory: {
            category: { uuid: '2', name: 'ネックレス' },
            pageInfo: {
                hasMore: true,
                nextCursor: 'next-cursor',
            },
            products: [
                {
                    uuid: '1',
                    name: 'シンプルネックレス',
                    description: 'シンプルで上品なデザインのネックレスです。',
                    price: 3500,
                    isActive: true,
                    isRecommend: false,
                    productImages: [{ uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', displayOrder: 1 }],
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
                    productImages: [{ uuid: '2', apiPath: '/image/about/story.jpg', name: 'image2', displayOrder: 1 }],
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
                    productImages: [{ uuid: '3', apiPath: '/image/about/story.jpg', name: 'image3', displayOrder: 1 }],
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
                    productImages: [{ uuid: '4', apiPath: '/image/about/story.jpg', name: 'image4', displayOrder: 1 }],
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
                    productImages: [{ uuid: '5', apiPath: '/image/about/story.jpg', name: 'image5', displayOrder: 1 }],
                    target: { uuid: '2', name: 'メンズ' },
                    category: { uuid: '2', name: 'ネックレス' },
                    tags: [],
                    siteDetails: [],
                },
            ],
        },
    },
}
