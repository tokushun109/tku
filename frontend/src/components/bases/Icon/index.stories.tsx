import { ColorType } from '@/types'

import { Icon } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Icon> = {
    component: Icon,
    args: {
        color: ColorType.Primary,
        size: 80,
        children: 'Icon',
        onClick: () => {
            console.log('clickしました')
        },
    },
}

export default meta
type Story = StoryObj<typeof Icon>

export const Default: Story = {}
