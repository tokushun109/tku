import { MultipleImageInput } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof MultipleImageInput> = {
    component: MultipleImageInput,
    args: {
        label: '商品画像',
        helperText: '画像ファイルを選択してください（複数選択可能）',
        value: [],
    },
    argTypes: {
        onChange: { action: 'changed' },
        required: {
            control: { type: 'boolean' },
        },
        disabled: {
            control: { type: 'boolean' },
        },
    },
    parameters: {
        docs: {
            description: {
                component: '複数の画像ファイルをアップロードするためのファイル入力コンポーネントです。',
            },
        },
    },
}

export default meta
type Story = StoryObj<typeof MultipleImageInput>

export const Default: Story = {}

export const Required: Story = {
    args: {
        required: true,
    },
}

export const WithError: Story = {
    args: {
        error: '画像ファイルを選択してください',
        required: true,
    },
}

export const Disabled: Story = {
    args: {
        disabled: true,
    },
}

export const WithSelectedFiles: Story = {
    args: {
        value: [
            new File([''], 'sample1.jpg', { type: 'image/jpeg' }),
            new File([''], 'sample2.png', { type: 'image/png' }),
            new File([''], 'sample3.jpg', { type: 'image/jpeg' }),
        ],
    },
}
