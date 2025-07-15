import { Input } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Input> = {
    component: Input,
    args: {
        label: 'ラベル',
        placeholder: 'プレースホルダー',
        required: false,
        disabled: false,
    },
    argTypes: {
        type: {
            control: { type: 'select' },
            options: ['text', 'email', 'password', 'tel', 'url', 'number'],
        },
        required: {
            control: { type: 'boolean' },
        },
        disabled: {
            control: { type: 'boolean' },
        },
    },
}

export default meta
type Story = StoryObj<typeof Input>

export const Default: Story = {}

export const WithLabel: Story = {
    args: {
        label: 'お名前',
        placeholder: '山田太郎',
    },
}

export const Required: Story = {
    args: {
        label: 'メールアドレス',
        placeholder: 'example@example.com',
        type: 'email',
        required: true,
    },
}

export const WithError: Story = {
    args: {
        label: 'パスワード',
        type: 'password',
        error: 'パスワードは8文字以上で入力してください',
        required: true,
    },
}

export const WithDefaultValue: Story = {
    args: {
        label: '会社名',
        defaultValue: '株式会社サンプル',
        placeholder: '会社名を入力してください',
    },
}

export const WithHelperText: Story = {
    args: {
        label: '電話番号',
        type: 'tel',
        helperText: 'ハイフンなしで入力してください',
        placeholder: '09012345678',
    },
}

export const Disabled: Story = {
    args: {
        label: '無効化されたフィールド',
        defaultValue: '編集できません',
        disabled: true,
    },
}

export const MaxLength: Story = {
    args: {
        label: 'ユーザー名',
        maxLength: 20,
        placeholder: '20文字以内で入力してください',
    },
}

export const FormExample: Story = {
    args: {
        label: 'お問い合わせ内容（件名）',
        defaultValue: '商品について',
        maxLength: 50,
        required: true,
        placeholder: '件名を入力してください',
    },
}
