import { TextArea } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof TextArea> = {
    component: TextArea,
    args: {
        label: 'ラベル',
        placeholder: 'プレースホルダー',
        required: false,
        disabled: false,
        rows: 4,
    },
    argTypes: {
        required: {
            control: { type: 'boolean' },
        },
        disabled: {
            control: { type: 'boolean' },
        },
        rows: {
            control: { type: 'number', min: 1, max: 10 },
        },
    },
}

export default meta
type Story = StoryObj<typeof TextArea>

export const Default: Story = {}

export const WithLabel: Story = {
    args: {
        label: 'お問い合わせ内容',
        placeholder: 'ご質問やご要望をお聞かせください',
    },
}

export const Required: Story = {
    args: {
        label: 'メッセージ',
        placeholder: 'メッセージを入力してください',
        required: true,
        rows: 5,
    },
}

export const WithError: Story = {
    args: {
        label: 'コメント',
        error: 'コメントは必須項目です',
        required: true,
        rows: 3,
    },
}

export const WithDefaultValue: Story = {
    args: {
        label: '自己紹介',
        defaultValue: 'こんにちは。よろしくお願いします。',
        placeholder: '自己紹介を入力してください',
        rows: 4,
    },
}

export const WithHelperText: Story = {
    args: {
        label: '詳細説明',
        helperText: '200文字以内で入力してください',
        placeholder: '詳細を入力してください',
        rows: 6,
    },
}

export const Disabled: Story = {
    args: {
        label: '無効化されたテキストエリア',
        defaultValue: 'この内容は編集できません。',
        disabled: true,
        rows: 3,
    },
}

export const LargeTextArea: Story = {
    args: {
        label: '長文入力',
        placeholder: '長文を入力してください',
        rows: 8,
        required: false,
    },
}

export const FormExample: Story = {
    args: {
        label: 'お問い合わせ詳細',
        defaultValue: '',
        placeholder: '具体的なお問い合わせ内容をご記入ください。',
        required: true,
        rows: 5,
    },
}
