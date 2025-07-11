import { SlideShow } from '.'
import { IThumbnail } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockThumbnails: IThumbnail[] = [
    {
        uuid: '1',
        name: 'ハンドメイドアクセサリー 1',
        price: 1500,
        imageUrl: 'https://via.placeholder.com/400x400',
        target: { uuid: '1', name: 'レディース' },
        category: { uuid: '1', name: 'ピアス' },
    },
    {
        uuid: '2',
        name: 'ハンドメイドアクセサリー 2',
        price: 2000,
        imageUrl: 'https://via.placeholder.com/400x400',
        target: { uuid: '2', name: 'メンズ' },
        category: { uuid: '2', name: 'ネックレス' },
    },
    {
        uuid: '3',
        name: 'ハンドメイドアクセサリー 3',
        price: 1800,
        imageUrl: 'https://via.placeholder.com/400x400',
        target: { uuid: '1', name: 'レディース' },
        category: { uuid: '3', name: 'ブレスレット' },
    },
    {
        uuid: '4',
        name: 'ハンドメイドアクセサリー 4',
        price: 2500,
        imageUrl: 'https://via.placeholder.com/400x400',
        target: { uuid: '3', name: 'ユニセックス' },
        category: { uuid: '4', name: 'リング' },
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