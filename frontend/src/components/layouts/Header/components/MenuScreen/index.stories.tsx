import { MenuScreen } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof MenuScreen> = {
    component: MenuScreen,
    args: {
        onCloseClick: () => {
            console.log('Close button clicked')
        },
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
        layout: 'fullscreen',
    },
}

export default meta
type Story = StoryObj<typeof MenuScreen>

export const Default: Story = {}
