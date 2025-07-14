import { IThumbnail } from '@/features/product/type'

import { Carousel } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

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
    {
        apiPath: '/image/about/concept1.jpg',
        product: {
            uuid: '5',
            name: 'ハンドメイドアクセサリー 5',
            description: 'シンプルで上品なデザインのハンドメイドアクセサリーです。',
            price: 3000,
            isActive: true,
            isRecommend: false,
            productImages: [],
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '5', name: 'ヘアアクセサリー' },
            tags: [],
            siteDetails: [],
        },
    },
]

const meta: Meta<typeof Carousel> = {
    component: Carousel,
    args: {
        items: mockThumbnails,
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof Carousel>

export const Default: Story = {}

export const FewItems: Story = {
    args: {
        items: mockThumbnails.slice(0, 2),
    },
}

export const ManyItems: Story = {
    args: {
        items: [...mockThumbnails, ...mockThumbnails, ...mockThumbnails],
    },
}
