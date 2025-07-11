import { MobileMenu } from '.'

import type { Meta, StoryObj } from '@storybook/react'

const meta: Meta<typeof MobileMenu> = {
    title: 'Components/Layouts/Footer/Components/MobileMenu',
    component: MobileMenu,
    parameters: {
        nextjs: {
            appDirectory: true,
            navigation: {
                pathname: '/',
                push: () => {},
            },
        },
        layout: 'fullscreen',
    },
    decorators: [
        (Story) => (
            <div style={{ width: '100%', height: '60px' }}>
                <Story />
            </div>
        ),
    ],
}

export default meta
type Story = StoryObj<typeof MobileMenu>

export const Default: Story = {}