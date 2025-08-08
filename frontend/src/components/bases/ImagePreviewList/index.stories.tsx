import { ImagePreviewList } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ImagePreviewList> = {
    component: ImagePreviewList,
    args: {
        title: '画像プレビュー',
        images: [
            {
                id: 'existing-1',
                src: 'https://images.unsplash.com/photo-1573408301185-9146fe634ad0?w=300&h=300&fit=crop',
                type: 'existing',
                order: 1,
            },
            {
                id: 'new-1',
                src: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=300&fit=crop',
                type: 'new',
                order: 2,
            },
            {
                id: 'existing-2',
                src: 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=300&h=300&fit=crop',
                type: 'existing',
                order: 3,
            },
        ],
    },
    argTypes: {
        onDelete: { action: 'deleted' },
        onReorder: { action: 'reordered' },
    },
    parameters: {
        docs: {
            description: {
                component: '画像のプレビュー、削除、並び替え機能を提供するコンポーネントです。',
            },
        },
    },
}

export default meta
type Story = StoryObj<typeof ImagePreviewList>

export const Default: Story = {}

export const Empty: Story = {
    args: {
        images: [],
        title: '画像なし',
    },
}

export const ExistingImagesOnly: Story = {
    args: {
        images: [
            {
                id: 'existing-1',
                src: 'https://images.unsplash.com/photo-1573408301185-9146fe634ad0?w=300&h=300&fit=crop',
                type: 'existing',
                order: 1,
            },
            {
                id: 'existing-2',
                src: 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=300&h=300&fit=crop',
                type: 'existing',
                order: 2,
            },
        ],
        title: '既存画像のみ',
    },
}

export const NewImagesOnly: Story = {
    args: {
        images: [
            {
                id: 'new-1',
                src: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=300&fit=crop',
                type: 'new',
            },
            {
                id: 'new-2',
                src: 'https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=300&h=300&fit=crop',
                type: 'new',
            },
        ],
        title: '新規画像のみ',
    },
}

export const ManyImages: Story = {
    args: {
        images: Array.from({ length: 8 }, (_, index) => ({
            id: `image-${index}`,
            src: `https://images.unsplash.com/photo-${1573408301185 + index}?w=300&h=300&fit=crop`,
            type: index % 2 === 0 ? ('existing' as const) : ('new' as const),
            order: index + 1,
        })),
        title: '多数の画像',
    },
}
