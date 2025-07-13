import { ColorType } from '@/types'

import { Button } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs'

const meta: Meta<typeof Button> = {
    component: Button,
    args: {
        children: 'button',
        colorType: ColorType.Primary,
        disabled: false,
        onClick: () => {
            console.log('clickしました')
        },
    },
    argTypes: {
        colorType: {
            control: { type: 'select' },
            options: Object.values(ColorType),
        },
        disabled: {
            control: { type: 'boolean' },
        },
    },
}

export default meta
type Story = StoryObj<typeof Button>

export const Default: Story = {}

export const Primary: Story = {
    args: {
        children: 'Primary Button',
        colorType: ColorType.Primary,
    },
}

export const Secondary: Story = {
    args: {
        children: 'Secondary Button',
        colorType: ColorType.Secondary,
    },
}

export const Accent: Story = {
    args: {
        children: 'Accent Button',
        colorType: ColorType.Accent,
    },
}

export const LongText: Story = {
    args: {
        children: '商品をカートに追加しました',
        colorType: ColorType.Primary,
        onClick: () => {
            console.log('商品をカートに追加しました')
        },
    },
}

export const ShortText: Story = {
    args: {
        children: '購入',
        colorType: ColorType.Accent,
        onClick: () => {
            console.log('購入ボタンがクリックされました')
        },
    },
}

export const WithIcon: Story = {
    args: {
        children: (
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                <span>♡</span>
                <span>お気に入りに追加</span>
            </div>
        ),
        colorType: ColorType.Secondary,
        onClick: () => {
            console.log('お気に入りに追加しました')
        },
    },
}

export const ActionButtons: Story = {
    args: {
        children: '管理者ログイン',
        colorType: ColorType.Primary,
        onClick: () => {
            console.log('管理者ログインボタンがクリックされました')
        },
    },
}

export const NoClickHandler: Story = {
    args: {
        children: 'クリック無効ボタン',
        colorType: ColorType.Secondary,
    },
}

export const Disabled: Story = {
    args: {
        children: '無効化されたボタン',
        colorType: ColorType.Primary,
        disabled: true,
        onClick: () => {
            console.log('このクリックは実行されません')
        },
    },
}

export const DisabledSecondary: Story = {
    args: {
        children: '送信中...',
        colorType: ColorType.Secondary,
        disabled: true,
    },
}

export const FormSubmit: Story = {
    args: {
        children: '送信する',
        colorType: ColorType.Primary,
        type: 'submit',
        disabled: false,
    },
}
