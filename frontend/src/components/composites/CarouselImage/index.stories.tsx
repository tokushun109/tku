import { IThumbnail } from '@/features/product/type'

import { CarouselImage } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const mockThumbnail: IThumbnail = {
    apiPath: '/image/about/concept1.jpg',
    product: {
        uuid: '1',
        name: 'ハンドメイドピアス',
        description: 'シンプルで上品なデザインのハンドメイドピアスです。',
        price: 1500,
        isActive: true,
        isRecommend: false,
        productImages: [],
        target: { uuid: '1', name: 'レディース' },
        category: { uuid: '1', name: 'ピアス' },
        tags: [],
        siteDetails: [],
    },
}

const meta: Meta<typeof CarouselImage> = {
    component: CarouselImage,
    args: {
        item: mockThumbnail,
        shadow: true,
    },
    argTypes: {
        item: {
            control: false,
        },
        shadow: {
            control: { type: 'boolean' },
        },
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
    decorators: [
        (Story) => (
            <div style={{ width: '480px', height: '480px' }}>
                <Story />
            </div>
        ),
    ],
}

export default meta
type Story = StoryObj<typeof CarouselImage>

export const Default: Story = {}

export const WithoutShadow: Story = {
    args: {
        shadow: false,
    },
}

export const MensAccessory: Story = {
    args: {
        item: {
            ...mockThumbnail,
            product: {
                ...mockThumbnail.product,
                uuid: '2',
                name: 'メンズネックレス',
                price: 2500,
                target: { uuid: '2', name: 'メンズ' },
                category: { uuid: '2', name: 'ネックレス' },
            },
        },
    },
}

export const ExpensiveItem: Story = {
    args: {
        item: {
            ...mockThumbnail,
            product: {
                ...mockThumbnail.product,
                uuid: '3',
                name: 'プレミアムリング',
                price: 15000,
                target: { uuid: '3', name: 'ユニセックス' },
                category: { uuid: '3', name: 'リング' },
            },
        },
    },
}

export const LongProductName: Story = {
    args: {
        item: {
            ...mockThumbnail,
            product: {
                ...mockThumbnail.product,
                uuid: '4',
                name: 'とても長い商品名のハンドメイドアクセサリー',
                price: 3000,
            },
        },
    },
}
