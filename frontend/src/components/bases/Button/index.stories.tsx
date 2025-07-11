import { Button } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof Button> = {
    component: Button,
    args: {
        children: 'button',
        onClick: () => {
            console.log('clickしました')
        },
    },
}

export default meta
type Story = StoryObj<typeof Button>

export const Default: Story = {}
