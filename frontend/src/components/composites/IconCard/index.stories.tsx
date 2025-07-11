import { MaterialIconEnum } from '@/types'

import { IconCard } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof IconCard> = {
    component: IconCard,
    decorators: [(story) => <div style={{ width: 480 }}>{story()}</div>],
}

export default meta
type Story = StoryObj<typeof IconCard>

export const About: Story = {
    args: {
        Icon: MaterialIconEnum.Face,
        label: 'about',
    },
}
export const Product: Story = {
    args: {
        Icon: MaterialIconEnum.Diamond,
        label: 'product',
    },
}

export const Contact: Story = {
    args: {
        Icon: MaterialIconEnum.Email,
        label: 'contact',
    },
}
