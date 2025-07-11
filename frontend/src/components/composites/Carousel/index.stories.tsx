import { Carousel } from '.'
import { IThumbnail } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockThumbnails: IThumbnail[] = [
    {
        uuid: '1',
        name: 'ハンドメイドアクセサリー 1',
        price: 1500,
        imageUrl: 'https://via.placeholder.com/300x300',
        target: { uuid: '1', name: 'レディース' },
        category: { uuid: '1', name: 'ピアス' },
    },
    {
        uuid: '2',
        name: 'ハンドメイドアクセサリー 2',
        price: 2000,
        imageUrl: 'https://via.placeholder.com/300x300',
        target: { uuid: '2', name: 'メンズ' },
        category: { uuid: '2', name: 'ネックレス' },
    },
    {
        uuid: '3',
        name: 'ハンドメイドアクセサリー 3',
        price: 1800,
        imageUrl: 'https://via.placeholder.com/300x300',
        target: { uuid: '1', name: 'レディース' },
        category: { uuid: '3', name: 'ブレスレット' },
    },
    {
        uuid: '4',
        name: 'ハンドメイドアクセサリー 4',
        price: 2500,
        imageUrl: 'https://via.placeholder.com/300x300',
        target: { uuid: '3', name: 'ユニセックス' },
        category: { uuid: '4', name: 'リング' },
    },
    {
        uuid: '5',
        name: 'ハンドメイドアクセサリー 5',
        price: 3000,
        imageUrl: 'https://via.placeholder.com/300x300',
        target: { uuid: '1', name: 'レディース' },
        category: { uuid: '5', name: 'ヘアアクセサリー' },
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