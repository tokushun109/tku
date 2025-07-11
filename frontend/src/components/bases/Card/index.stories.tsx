import { Card } from '.'
import { ColorType } from '@/types'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof Card> = {
    component: Card,
    args: {
        children: (
            <div style={{ padding: '20px' }}>
                <h3>Card Content</h3>
                <p>This is sample content inside a card component.</p>
            </div>
        ),
        color: ColorType.White,
        height: 'auto',
        shadow: true,
        width: 'auto',
    },
}

export default meta
type Story = StoryObj<typeof Card>

export const Default: Story = {}

export const WithoutShadow: Story = {
    args: {
        shadow: false,
    },
}

export const PrimaryColor: Story = {
    args: {
        color: ColorType.Primary,
    },
}

export const SecondaryColor: Story = {
    args: {
        color: ColorType.Secondary,
    },
}

export const CustomDimensions: Story = {
    args: {
        width: '300px',
        height: '200px',
    },
}

export const LargeCard: Story = {
    args: {
        width: '400px',
        height: '300px',
        children: (
            <div style={{ padding: '24px' }}>
                <h2>Large Card</h2>
                <p>This is a larger card with more content.</p>
                <ul>
                    <li>Feature 1</li>
                    <li>Feature 2</li>
                    <li>Feature 3</li>
                </ul>
            </div>
        ),
    },
}