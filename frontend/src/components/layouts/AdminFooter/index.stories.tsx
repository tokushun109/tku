import AdminFooter from './index'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof AdminFooter> = {
    component: AdminFooter,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
        layout: 'fullscreen',
    },
    args: {},
    argTypes: {
        className: {
            control: { type: 'text' },
        },
    },
}

export default meta
type Story = StoryObj<typeof AdminFooter>

export const Default: Story = {}

export const WithCustomClass: Story = {
    args: {
        className: 'custom-footer-class',
    },
}
