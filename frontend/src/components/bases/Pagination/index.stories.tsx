import { useState } from 'react'

import { Pagination } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Pagination> = {
    component: Pagination,
    args: {
        currentPage: 1,
        disabled: false,
        siblingCount: 1,
        totalPages: 10,
        onPageChange: (page) => console.log('page changed:', page),
    },
    argTypes: {
        currentPage: {
            control: { type: 'number', min: 1 },
        },
        disabled: {
            control: { type: 'boolean' },
        },
        siblingCount: {
            control: { type: 'number', min: 0, max: 3 },
        },
        totalPages: {
            control: { type: 'number', min: 1 },
        },
    },
}

export default meta
type Story = StoryObj<typeof Pagination>

export const Default: Story = {}

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
        totalPages: 4,
    },
}

export const Disabled: Story = {
    args: {
        currentPage: 3,
        disabled: true,
        totalPages: 10,
    },
}

export const Interactive: Story = {
    render: (args) => {
        const [currentPage, setCurrentPage] = useState(args.currentPage)

        return <Pagination {...args} currentPage={currentPage} onPageChange={setCurrentPage} />
    },
}
