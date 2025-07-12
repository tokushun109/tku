import { Slide } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof Slide> = {
    component: Slide,
    args: {
        children: (
            <div style={{ padding: '20px', backgroundColor: '#f0f0f0', borderRadius: '8px' }}>
                <h2>Sample Content</h2>
                <p>This content will slide in with animation when it comes into view.</p>
            </div>
        ),
    },
}

export default meta
type Story = StoryObj<typeof Slide>

export const Default: Story = {}

export const WithText: Story = {
    args: {
        children: <p>Simple text that slides in</p>,
    },
}

export const WithCard: Story = {
    args: {
        children: (
            <div
                style={{
                    padding: '24px',
                    backgroundColor: '#ffffff',
                    borderRadius: '12px',
                    boxShadow: '0 2px 8px rgba(0,0,0,0.1)',
                    maxWidth: '300px',
                }}
            >
                <h3>Card Title</h3>
                <p>This card content will animate in smoothly.</p>
            </div>
        ),
    },
}
