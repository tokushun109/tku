import AdminHeader from './index'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof AdminHeader> = {
    component: AdminHeader,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
        layout: 'fullscreen',
    },
    args: {
        isLoggedIn: true,
    },
    argTypes: {
        isLoggedIn: {
            control: { type: 'boolean' },
        },
    },
}

export default meta
type Story = StoryObj<typeof AdminHeader>

export const Default: Story = {}

export const LoggedIn: Story = {
    args: {
        isLoggedIn: true,
    },
}

export const LoggedOut: Story = {
    args: {
        isLoggedIn: false,
    },
}

export const WithLogout: Story = {
    args: {
        isLoggedIn: true,
    },
}
