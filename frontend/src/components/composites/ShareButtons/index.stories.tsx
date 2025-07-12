import { ShareButtons } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs'

const meta: Meta<typeof ShareButtons> = {
    component: ShareButtons,
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof ShareButtons>

export const Default: Story = {}
