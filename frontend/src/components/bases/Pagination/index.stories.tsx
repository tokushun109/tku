import { Pagination } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Pagination> = {
    component: Pagination,
    args: {
        currentPage: 1,
        totalPages: 10,
        onPageChange: (_page: number) => {},
        delta: 2,
    },
    argTypes: {
        currentPage: {
            control: { type: 'number', min: 1 },
            description: '現在のページ番号',
        },
        totalPages: {
            control: { type: 'number', min: 1 },
            description: '総ページ数',
        },
        delta: {
            control: { type: 'number', min: 1, max: 5 },
            description: '現在のページの前後に表示するページ数',
        },
        onPageChange: {
            action: 'page changed',
            description: 'ページ変更時のコールバック関数',
        },
    },
}

export default meta
type Story = StoryObj<typeof Pagination>

export const Default: Story = {}

export const FirstPage: Story = {
    args: {
        currentPage: 1,
        totalPages: 10,
    },
}

export const MiddlePage: Story = {
    args: {
        currentPage: 5,
        totalPages: 10,
    },
}

export const LastPage: Story = {
    args: {
        currentPage: 10,
        totalPages: 10,
    },
}

export const FewPages: Story = {
    args: {
        currentPage: 2,
        totalPages: 3,
    },
}

export const ManyPages: Story = {
    args: {
        currentPage: 15,
        totalPages: 50,
    },
}

export const LargeDelta: Story = {
    args: {
        currentPage: 10,
        totalPages: 20,
        delta: 5,
    },
}

export const SmallDelta: Story = {
    args: {
        currentPage: 10,
        totalPages: 20,
        delta: 1,
    },
}

export const SinglePage: Story = {
    args: {
        currentPage: 1,
        totalPages: 1,
    },
}
