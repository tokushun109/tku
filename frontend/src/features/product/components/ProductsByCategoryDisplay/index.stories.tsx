import { ProductsByCategoryDisplay } from '.'
import { IProductsByCategory } from '@/features/product/type'

import type { Meta, StoryObj } from '@storybook/react'

const mockProductsByCategory: IProductsByCategory = {
    uuid: '1',
    name: 'ピアス',
    products: [
        {
            uuid: '1',
            name: 'ハンドメイドピアス 1',
            price: 1500,
            imageUrl: 'https://via.placeholder.com/300x300',
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
        },
        {
            uuid: '2',
            name: 'ハンドメイドピアス 2',
            price: 2000,
            imageUrl: 'https://via.placeholder.com/300x300',
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
        },
        {
            uuid: '3',
            name: 'ハンドメイドピアス 3',
            price: 1800,
            imageUrl: 'https://via.placeholder.com/300x300',
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
        },
        {
            uuid: '4',
            name: 'ハンドメイドピアス 4',
            price: 2500,
            imageUrl: 'https://via.placeholder.com/300x300',
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
        },
        {
            uuid: '5',
            name: 'ハンドメイドピアス 5',
            price: 3000,
            imageUrl: 'https://via.placeholder.com/300x300',
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
        },
        {
            uuid: '6',
            name: 'ハンドメイドピアス 6',
            price: 2200,
            imageUrl: 'https://via.placeholder.com/300x300',
            target: { uuid: '1', name: 'レディース' },
            category: { uuid: '1', name: 'ピアス' },
        },
    ],
}

const meta: Meta<typeof ProductsByCategoryDisplay> = {
    component: ProductsByCategoryDisplay,
    args: {
        productsByCategory: mockProductsByCategory,
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof ProductsByCategoryDisplay>

export const Default: Story = {}

export const FewProducts: Story = {
    args: {
        productsByCategory: {
            ...mockProductsByCategory,
            products: mockProductsByCategory.products.slice(0, 3),
        },
    },
}

export const ExactlyFourProducts: Story = {
    args: {
        productsByCategory: {
            ...mockProductsByCategory,
            products: mockProductsByCategory.products.slice(0, 4),
        },
    },
}

export const ManyProducts: Story = {
    args: {
        productsByCategory: {
            ...mockProductsByCategory,
            products: [
                ...mockProductsByCategory.products,
                {
                    uuid: '7',
                    name: 'ハンドメイドピアス 7',
                    price: 2800,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '1', name: 'ピアス' },
                },
                {
                    uuid: '8',
                    name: 'ハンドメイドピアス 8',
                    price: 3500,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '1', name: 'ピアス' },
                },
            ],
        },
    },
}

export const NecklaceCategory: Story = {
    args: {
        productsByCategory: {
            uuid: '2',
            name: 'ネックレス',
            products: [
                {
                    uuid: '1',
                    name: 'シンプルネックレス',
                    price: 3500,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '2', name: 'ネックレス' },
                },
                {
                    uuid: '2',
                    name: 'パールネックレス',
                    price: 4500,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '1', name: 'レディース' },
                    category: { uuid: '2', name: 'ネックレス' },
                },
                {
                    uuid: '3',
                    name: 'ゴールドネックレス',
                    price: 8000,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '3', name: 'ユニセックス' },
                    category: { uuid: '2', name: 'ネックレス' },
                },
                {
                    uuid: '4',
                    name: 'メンズネックレス',
                    price: 5500,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '2', name: 'メンズ' },
                    category: { uuid: '2', name: 'ネックレス' },
                },
                {
                    uuid: '5',
                    name: 'チェーンネックレス',
                    price: 6000,
                    imageUrl: 'https://via.placeholder.com/300x300',
                    target: { uuid: '2', name: 'メンズ' },
                    category: { uuid: '2', name: 'ネックレス' },
                },
            ],
        },
    },
}