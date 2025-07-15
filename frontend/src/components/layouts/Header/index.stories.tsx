import { Header } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Header> = {
    component: Header,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
        layout: 'fullscreen',
    },
}

export default meta
type Story = StoryObj<typeof Header>

export const Default: Story = {}
