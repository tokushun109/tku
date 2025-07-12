import { IProduct } from '@/features/product/type'

import { ProductDetail } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const mockProduct: IProduct = {
    uuid: '1',
    name: 'ハンドメイドピアス',
    description:
        'シンプルで上品なデザインのハンドメイドピアスです。\n日常使いにも特別な日にもお使いいただけます。\n\n素材: 真鍮、コットンパール\nサイズ: 約2cm',
    price: 1500,
    isActive: true,
    isRecommend: false,
    productImages: [
        { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', order: 1 },
        { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', order: 2 },
        { uuid: '3', apiPath: '/image/about/concept2.jpg', name: 'image3', order: 3 },
    ],
    target: { uuid: '1', name: 'レディース' },
    category: { uuid: '1', name: 'ピアス' },
    tags: [
        { uuid: '1', name: 'ハンドメイド' },
        { uuid: '2', name: 'シンプル' },
        { uuid: '3', name: '上品' },
    ],
    siteDetails: [
        {
            uuid: '1',
            detailUrl: 'https://creema.jp/item/1',
            salesSite: { uuid: '1', name: 'Creema' },
        },
        {
            uuid: '2',
            detailUrl: 'https://minne.com/item/2',
            salesSite: { uuid: '2', name: 'minne' },
        },
    ],
}

const meta: Meta<typeof ProductDetail> = {
    component: ProductDetail,
    args: {
        product: mockProduct,
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof ProductDetail>

export const Default: Story = {}

export const WithoutTags: Story = {
    args: {
        product: {
            ...mockProduct,
            tags: [],
        },
    },
}

export const WithoutSalesLinks: Story = {
    args: {
        product: {
            ...mockProduct,
            siteDetails: [],
        },
    },
}

export const ExpensiveItem: Story = {
    args: {
        product: {
            ...mockProduct,
            name: 'プレミアムゴールドネックレス',
            price: 25000,
            description:
                '18金を使用した高級ネックレスです。\n特別な日にふさわしい上品な輝きを放ちます。\n\n素材: 18金、天然石\nサイズ: チェーン長さ 40cm\n付属品: 専用ケース、品質保証書',
            category: { uuid: '2', name: 'ネックレス' },
            tags: [
                { uuid: '1', name: 'プレミアム' },
                { uuid: '2', name: '18金' },
                { uuid: '3', name: '天然石' },
                { uuid: '4', name: 'ギフト' },
            ],
        },
    },
}

export const MensAccessory: Story = {
    args: {
        product: {
            ...mockProduct,
            name: 'メンズレザーブレスレット',
            price: 3500,
            description:
                '本革を使用したメンズ向けブレスレットです。\nシンプルなデザインで様々なスタイルに合わせやすく、\n長く愛用していただけます。\n\n素材: 本革、真鍮\nサイズ: 約18cm（調整可能）',
            target: { uuid: '2', name: 'メンズ' },
            category: { uuid: '3', name: 'ブレスレット' },
            tags: [
                { uuid: '1', name: 'レザー' },
                { uuid: '2', name: 'シンプル' },
                { uuid: '3', name: 'カジュアル' },
            ],
        },
    },
}

export const SingleImage: Story = {
    args: {
        product: {
            ...mockProduct,
            productImages: [{ uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', order: 1 }],
        },
    },
}
