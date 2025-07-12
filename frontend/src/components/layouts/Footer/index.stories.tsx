import { Footer } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs'

const meta: Meta<typeof Footer> = {
    component: Footer,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
        layout: 'fullscreen',
    },
}

export default meta
type Story = StoryObj<typeof Footer>

export const Default: Story = {}
