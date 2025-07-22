import { ClassificationType } from '@/types'

import { ClassificationList } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ClassificationList> = {
    component: ClassificationList,
    args: {
        initialItems: [
            { uuid: '1', name: 'ネックレス' },
            { uuid: '2', name: 'ピアス' },
            { uuid: '3', name: 'リング' },
            { uuid: '4', name: 'ブレスレット' },
            { uuid: '5', name: 'イヤリング' },
        ],
        type: ClassificationType.Category,
    },
    argTypes: {
        initialItems: {
            description: '分類アイテムの配列',
        },
        type: {
            description: '分類の種類',
        },
    },
}

export default meta
type Story = StoryObj<typeof ClassificationList>

export const Default: Story = {}

export const EmptyList: Story = {
    args: {
        initialItems: [],
    },
}

export const SingleItem: Story = {
    args: {
        initialItems: [{ uuid: '1', name: 'ネックレス' }],
    },
}

export const ManyItems: Story = {
    args: {
        initialItems: Array.from({ length: 25 }, (_, i) => ({
            uuid: `${i + 1}`,
            name: `アイテム ${i + 1}`,
        })),
    },
}

export const LongNames: Story = {
    args: {
        initialItems: [
            { uuid: '1', name: 'とても長い名前のアクセサリーアイテム' },
            { uuid: '2', name: 'Super Long Accessory Item Name Example' },
            { uuid: '3', name: '超超超超超長いアイテム名前例' },
        ],
    },
}

export const ExactlyTenItems: Story = {
    args: {
        initialItems: Array.from({ length: 10 }, (_, i) => ({
            uuid: `${i + 1}`,
            name: `アイテム ${i + 1}`,
        })),
    },
}

export const ElevenItems: Story = {
    args: {
        initialItems: Array.from({ length: 11 }, (_, i) => ({
            uuid: `${i + 1}`,
            name: `アイテム ${i + 1}`,
        })),
    },
}
