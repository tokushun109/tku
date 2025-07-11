import { ProductImageGallery } from '.'
import { IProduct } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockProduct: IProduct = {
    uuid: '1',
    name: 'ハンドメイドピアス',
    description: 'シンプルで上品なデザインのハンドメイドピアスです。',
    price: 1500,
    isActive: true,
    isRecommend: false,
    productImages: [
        { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', order: 1 },
        { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', order: 2 },
        { uuid: '3', apiPath: '/image/about/concept2.jpg', name: 'image3', order: 3 },
        { uuid: '4', apiPath: '/image/about/concept1.jpg', name: 'image4', order: 4 },
        { uuid: '5', apiPath: '/image/about/concept2.jpg', name: 'image5', order: 5 },
        { uuid: '6', apiPath: '/image/about/concept1.jpg', name: 'image6', order: 6 },
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
            productImages: [
                { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', order: 1 },
            ],
        },
    },
}

export const TwoImages: Story = {
    args: {
        product: {
            ...mockProduct,
            productImages: [
                { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', order: 1 },
                { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', order: 2 },
            ],
        },
    },
}

export const ManyImages: Story = {
    args: {
        product: {
            ...mockProduct,
            productImages: [
                { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', order: 1 },
                { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', order: 2 },
                { uuid: '3', apiPath: '/image/about/concept2.jpg', name: 'image3', order: 3 },
                { uuid: '4', apiPath: '/image/about/concept1.jpg', name: 'image4', order: 4 },
                { uuid: '5', apiPath: '/image/about/concept2.jpg', name: 'image5', order: 5 },
                { uuid: '6', apiPath: '/image/about/concept1.jpg', name: 'image6', order: 6 },
                { uuid: '7', apiPath: '/image/about/concept2.jpg', name: 'image7', order: 7 },
                { uuid: '8', apiPath: '/image/about/concept1.jpg', name: 'image8', order: 8 },
                { uuid: '9', apiPath: '/image/about/concept2.jpg', name: 'image9', order: 9 },
                { uuid: '10', apiPath: '/image/about/concept1.jpg', name: 'image10', order: 10 },
            ],
        },
    },
}