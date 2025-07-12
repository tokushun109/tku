import { IThumbnail } from '@/features/product/type'

import { SlideShow } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const mockThumbnails: IThumbnail[] = [
    {
        apiPath: '/image/about/concept1.jpg',
        product: {
            uuid: '1',
            name: 'ハンドメイドアクセサリー 1',
            description: 'シンプルで上品なデザインのハンドメイドアクセサリーです。',
            price: 1500,
            isActive: true,
            isRecommend: false,
            productImages: [],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
            tags: [],
            siteDetails: [],
        },
    },
    {
        apiPath: '/image/about/concept2.jpg',
        product: {
            uuid: '2',
            name: 'ハンドメイドアクセサリー 2',
            description: 'シンプルで上品なデザインのハンドメイドアクセサリーです。',
            price: 2000,
            isActive: true,
            isRecommend: false,
            productImages: [],
            target: { uuid: '2', name: 'メンズ' },
            category: { uuid: '2', name: 'ネックレス' },
            tags: [],
            siteDetails: [],
        },
    },
    {
        apiPath: '/image/about/concept1.jpg',
        product: {
            uuid: '3',
            name: 'ハンドメイドアクセサリー 3',
            description: 'シンプルで上品なデザインのハンドメイドアクセサリーです。',
            price: 1800,
            isActive: true,
            isRecommend: false,
            productImages: [],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '3', name: 'ブレスレット' },
            tags: [],
            siteDetails: [],
        },
    },
    {
        apiPath: '/image/about/concept2.jpg',
        product: {
            uuid: '4',
            name: 'ハンドメイドアクセサリー 4',
            description: 'シンプルで上品なデザインのハンドメイドアクセサリーです。',
            price: 2500,
            isActive: true,
            isRecommend: false,
            productImages: [],
            target: { uuid: '3', name: 'ユニセックス' },
            category: { uuid: '4', name: 'リング' },
            tags: [],
            siteDetails: [],
        },
    },
]

const meta: Meta<typeof SlideShow> = {
    component: SlideShow,
    args: {
        items: mockThumbnails,
        size: '300px',
        autoPlay: true,
        innerPadding: 16,
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof SlideShow>

export const Default: Story = {}

export const WithoutAutoPlay: Story = {
    args: {
        autoPlay: false,
    },
}

export const LargeSize: Story = {
    args: {
        size: '400px',
    },
}

export const SmallSize: Story = {
    args: {
        size: '200px',
    },
}

export const CustomPadding: Story = {
    args: {
        innerPadding: 32,
    },
}

export const SingleItem: Story = {
    args: {
        items: [mockThumbnails[0]],
    },
}
