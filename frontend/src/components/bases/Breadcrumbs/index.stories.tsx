import { Breadcrumbs } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof Breadcrumbs> = {
    component: Breadcrumbs,
    args: {
        breadcrumbs: [{ label: 'label1', link: '/link1' }, { label: 'label2', link: '/link2' }, { label: 'label3' }],
    },
    parameters: {
        nextjs: {
            appDirectory: true,
        },
    },
}

export default meta
type Story = StoryObj<typeof Breadcrumbs>

export const Default: Story = {}
