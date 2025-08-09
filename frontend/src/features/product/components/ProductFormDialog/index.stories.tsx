import { ProductFormDialog } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ProductFormDialog> = {
    component: ProductFormDialog,
    args: {
        categories: [
            { uuid: 'category-1', name: 'ピアス' },
            { uuid: 'category-2', name: 'ネックレス' },
            { uuid: 'category-3', name: 'ブレスレット' },
        ],
        targets: [
            { uuid: 'target-1', name: '大人女性' },
            { uuid: 'target-2', name: '若い女性' },
            { uuid: 'target-3', name: '男性' },
        ],
        tags: [
            { uuid: 'tag-1', name: 'シンプル' },
            { uuid: 'tag-2', name: 'カジュアル' },
            { uuid: 'tag-3', name: 'エレガント' },
        ],
        salesSites: [
            { uuid: 'site-1', name: 'Creema' },
            { uuid: 'site-2', name: 'minne' },
        ],
        isOpen: true,
        isSubmitting: false,
        onClose: () => {
            console.log('ダイアログを閉じます')
        },
        onSubmit: async (data) => {
            console.log('フォーム送信:', data)
            await new Promise((resolve) => setTimeout(resolve, 1000))
        },
        onCreemaDuplicate: async (data) => {
            console.log('Creema複製:', data)
            await new Promise((resolve) => setTimeout(resolve, 1000))
        },
        submitError: null,
        updateItem: null,
    },
    argTypes: {
        isOpen: {
            control: { type: 'boolean' },
        },
        isSubmitting: {
            control: { type: 'boolean' },
        },
        updateItem: {
            control: { type: 'object' },
        },
    },
}

export default meta
type Story = StoryObj<typeof ProductFormDialog>

export const Create: Story = {
    args: {
        updateItem: null,
    },
}

export const Edit: Story = {
    args: {
        updateItem: {
            uuid: '1',
            name: 'テスト商品',
            description: 'これはテスト用の商品説明です。',
            price: 1500,
            isActive: true,
            isRecommend: false,
            category: {
                uuid: 'category-1',
                name: 'ピアス',
            },
            target: {
                uuid: 'target-1',
                name: '大人女性',
            },
            tags: [
                {
                    uuid: 'tag-1',
                    name: 'シンプル',
                },
                {
                    uuid: 'tag-2',
                    name: 'カジュアル',
                },
            ],
            productImages: [],
            siteDetails: [
                {
                    uuid: 'site-detail-1',
                    detailUrl: 'https://creema.jp/item/123456/detail',
                    salesSite: {
                        uuid: 'sales-site-1',
                        name: 'Creema',
                    },
                },
            ],
        },
    },
}

export const WithError: Story = {
    args: {
        submitError: '送信中にエラーが発生しました。もう一度お試しください。',
    },
}

export const Submitting: Story = {
    args: {
        isSubmitting: true,
    },
}

export const Closed: Story = {
    args: {
        isOpen: false,
    },
}
