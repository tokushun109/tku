import { Image } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Image> = {
    component: Image,
    args: {
        src: '/image/about/story.jpg',
        alt: 'Sample image',
    },
    decorators: [
        (Story) => (
            <div style={{ width: '300px', height: '200px' }}>
                <Story />
            </div>
        ),
    ],
}

export default meta
type Story = StoryObj<typeof Image>

export const Default: Story = {}

export const ProductImage: Story = {
    args: {
        src: '/image/about/story.jpg',
        alt: 'Product image',
    },
    decorators: [
        (Story) => (
            <div style={{ width: '400px', height: '400px' }}>
                <Story />
            </div>
        ),
    ],
}

export const LargeImage: Story = {
    args: {
        src: '/image/about/story.jpg',
        alt: 'Large image',
    },
    decorators: [
        (Story) => (
            <div style={{ width: '400px', height: '300px' }}>
                <Story />
            </div>
        ),
    ],
}

export const BrokenImage: Story = {
    args: {
        src: '/image/gray-image.png',
        alt: 'Broken image - should show fallback',
    },
}

export const JapaneseAlt: Story = {
    args: {
        src: '/image/about/story.jpg',
        alt: 'ハンドメイドアクセサリーの商品画像',
    },
}
