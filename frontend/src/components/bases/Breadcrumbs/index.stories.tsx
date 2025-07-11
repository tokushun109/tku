import { Breadcrumbs } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof Breadcrumbs> = {
    component: Breadcrumbs,
    args: {
        breadcrumbs: [{ label: 'label1', link: '/link1' }, { label: 'label2', link: '/link2' }, { label: 'label3' }],
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof Breadcrumbs>

export const Default: Story = {}

export const TwoItems: Story = {
    args: {
        breadcrumbs: [
            { label: 'ホーム', link: '/' },
            { label: '商品一覧' }
        ],
    },
}

export const LongPath: Story = {
    args: {
        breadcrumbs: [
            { label: 'ホーム', link: '/' },
            { label: 'カテゴリ', link: '/category' },
            { label: 'ネックレス', link: '/category/necklace' },
            { label: 'シルバーネックレス', link: '/category/necklace/silver' },
            { label: '商品詳細' }
        ],
    },
}

export const AllLinked: Story = {
    args: {
        breadcrumbs: [
            { label: 'ホーム', link: '/' },
            { label: '商品一覧', link: '/products' },
            { label: 'カテゴリ', link: '/category' }
        ],
    },
}

export const SingleItem: Story = {
    args: {
        breadcrumbs: [
            { label: 'ホーム' }
        ],
    },
}

export const JapaneseLongLabels: Story = {
    args: {
        breadcrumbs: [
            { label: 'ハンドメイドアクセサリー', link: '/' },
            { label: 'ネックレス・ペンダント', link: '/necklace' },
            { label: 'シルバー925チェーンネックレス' }
        ],
    },
}
