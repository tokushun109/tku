import { Indicator } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Indicator> = {
    component: Indicator,
    args: {
        children: (
            <div style={{ padding: '10px' }}>
                <span>Sample indicator content</span>
            </div>
        ),
    },
    decorators: [
        (Story) => (
            <div style={{ height: 96, width: 80, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                <Story />
            </div>
        ),
    ],
}

export default meta
type Story = StoryObj<typeof Indicator>

export const Default: Story = {}

export const WithText: Story = {
    args: {
        children: <span>Loading...</span>,
    },
}

export const WithIcon: Story = {
    args: {
        children: <div>✓ Complete</div>,
    },
}

export const WithNumber: Story = {
    args: {
        children: <span>1</span>,
    },
}

export const WithCustomContent: Story = {
    args: {
        children: (
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                <div style={{ width: '8px', height: '8px', backgroundColor: '#ff0000', borderRadius: '50%' }} />
                <span>Status</span>
            </div>
        ),
    },
}
