import { ProductImageGallery } from '.'
import { IProduct } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockProduct: IProduct = {
    uuid: '1',
    name: 'ハンドメイドピアス',
    description: 'シンプルで上品なデザインのハンドメイドピアスです。',
    price: 1500,
    images: [
        { uuid: '1', url: 'https://via.placeholder.com/400x400' },
        { uuid: '2', url: 'https://via.placeholder.com/400x400' },
        { uuid: '3', url: 'https://via.placeholder.com/400x400' },
        { uuid: '4', url: 'https://via.placeholder.com/400x400' },
        { uuid: '5', url: 'https://via.placeholder.com/400x400' },
        { uuid: '6', url: 'https://via.placeholder.com/400x400' },
    ],
    target: { uuid: '1', name: 'レディース' },
    category: { uuid: '1', name: 'ピアス' },
    tags: [],
    siteDetails: [],
}

const meta: Meta<typeof ProductImageGallery> = {
    component: ProductImageGallery,
    args: {
        product: mockProduct,
    },
}

export default meta
type Story = StoryObj<typeof ProductImageGallery>

export const Default: Story = {}

export const SingleImage: Story = {
    args: {
        product: {
            ...mockProduct,
            images: [
                { uuid: '1', url: 'https://via.placeholder.com/400x400' },
            ],
        },
    },
}

export const TwoImages: Story = {
    args: {
        product: {
            ...mockProduct,
            images: [
                { uuid: '1', url: 'https://via.placeholder.com/400x400' },
                { uuid: '2', url: 'https://via.placeholder.com/400x400' },
            ],
        },
    },
}

export const ManyImages: Story = {
    args: {
        product: {
            ...mockProduct,
            images: [
                { uuid: '1', url: 'https://via.placeholder.com/400x400' },
                { uuid: '2', url: 'https://via.placeholder.com/400x400' },
                { uuid: '3', url: 'https://via.placeholder.com/400x400' },
                { uuid: '4', url: 'https://via.placeholder.com/400x400' },
                { uuid: '5', url: 'https://via.placeholder.com/400x400' },
                { uuid: '6', url: 'https://via.placeholder.com/400x400' },
                { uuid: '7', url: 'https://via.placeholder.com/400x400' },
                { uuid: '8', url: 'https://via.placeholder.com/400x400' },
                { uuid: '9', url: 'https://via.placeholder.com/400x400' },
                { uuid: '10', url: 'https://via.placeholder.com/400x400' },
            ],
        },
    },
}