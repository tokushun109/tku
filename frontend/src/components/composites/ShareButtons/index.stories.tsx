import { ShareButtons } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ShareButtons> = {
    component: ShareButtons,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
    args: {
        url: '',
    },
    argTypes: {
        url: {
            control: { type: 'text' },
            description: '共有するURL（指定しない場合は現在のページのURLを使用）',
        },
    },
}

export default meta
type Story = StoryObj<typeof ShareButtons>

export const Default: Story = {}

export const WithCustomUrl: Story = {
    args: {
        url: 'https://tocoriri.com/product/123',
    },
}
