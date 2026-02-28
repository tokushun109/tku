import { IProduct } from '@/features/product/type'

import { ProductImageGallery } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const mockProduct: IProduct = {
    uuid: '1',
    name: 'ハンドメイドピアス',
    description: 'シンプルで上品なデザインのハンドメイドピアスです。',
    price: 1500,
    isActive: true,
    isRecommend: false,
    productImages: [
        { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', displayOrder: 1 },
        { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', displayOrder: 2 },
        { uuid: '3', apiPath: '/image/about/concept2.jpg', name: 'image3', displayOrder: 3 },
        { uuid: '4', apiPath: '/image/about/concept1.jpg', name: 'image4', displayOrder: 4 },
        { uuid: '5', apiPath: '/image/about/concept2.jpg', name: 'image5', displayOrder: 5 },
        { uuid: '6', apiPath: '/image/about/concept1.jpg', name: 'image6', displayOrder: 6 },
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
            productImages: [{ uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', displayOrder: 1 }],
        },
    },
}

export const TwoImages: Story = {
    args: {
        product: {
            ...mockProduct,
            productImages: [
                { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', displayOrder: 1 },
                { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', displayOrder: 2 },
            ],
        },
    },
}

export const ManyImages: Story = {
    args: {
        product: {
            ...mockProduct,
            productImages: [
                { uuid: '1', apiPath: '/image/about/story.jpg', name: 'image1', displayOrder: 1 },
                { uuid: '2', apiPath: '/image/about/concept1.jpg', name: 'image2', displayOrder: 2 },
                { uuid: '3', apiPath: '/image/about/concept2.jpg', name: 'image3', displayOrder: 3 },
                { uuid: '4', apiPath: '/image/about/concept1.jpg', name: 'image4', displayOrder: 4 },
                { uuid: '5', apiPath: '/image/about/concept2.jpg', name: 'image5', displayOrder: 5 },
                { uuid: '6', apiPath: '/image/about/concept1.jpg', name: 'image6', displayOrder: 6 },
                { uuid: '7', apiPath: '/image/about/concept2.jpg', name: 'image7', displayOrder: 7 },
                { uuid: '8', apiPath: '/image/about/concept1.jpg', name: 'image8', displayOrder: 8 },
                { uuid: '9', apiPath: '/image/about/concept2.jpg', name: 'image9', displayOrder: 9 },
                { uuid: '10', apiPath: '/image/about/concept1.jpg', name: 'image10', displayOrder: 10 },
            ],
        },
    },
}
