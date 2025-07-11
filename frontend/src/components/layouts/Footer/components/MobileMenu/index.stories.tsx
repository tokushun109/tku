import { MobileMenu } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof MobileMenu> = {
    component: MobileMenu,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
        layout: 'fullscreen',
    },
}

export default meta
type Story = StoryObj<typeof MobileMenu>

export const Default: Story = {}