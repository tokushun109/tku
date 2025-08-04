import { useState } from 'react'

import { Checkbox } from './index'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Checkbox> = {
    title: 'Components/Bases/Checkbox',
    component: Checkbox,
    parameters: {
        layout: 'centered',
    },
    tags: ['autodocs'],
    argTypes: {
        checked: {
            control: 'boolean',
        },
        disabled: {
            control: 'boolean',
        },
        label: {
            control: 'text',
        },
    },
}

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
    args: {
        label: 'チェックボックス',
    },
}

export const Checked: Story = {
    args: {
        label: 'チェック済み',
        checked: true,
    },
}

export const Disabled: Story = {
    args: {
        label: '無効状態',
        disabled: true,
    },
}

export const DisabledChecked: Story = {
    args: {
        label: '無効状態（チェック済み）',
        disabled: true,
        checked: true,
    },
}

export const WithoutLabel: Story = {
    args: {},
}

export const Interactive: Story = {
    render: (args) => {
        const [checked, setChecked] = useState<boolean>(false)
        return <Checkbox {...args} checked={checked} onChange={(e) => setChecked(e.target.checked)} />
    },
    args: {
        label: 'インタラクティブなチェックボックス',
    },
}
