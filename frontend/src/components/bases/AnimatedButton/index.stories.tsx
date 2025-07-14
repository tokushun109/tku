import { AnimatedButton } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof AnimatedButton> = {
    component: AnimatedButton,
    args: {
        children: 'animated button',
        onClick: () => {
            console.log('clickしました')
        },
        enabledAnimation: true,
    },
    argTypes: {
        enabledAnimation: {
            control: 'boolean',
            description: 'Enable button hover animation',
        },
    },
}

export default meta
type Story = StoryObj<typeof AnimatedButton>

export const Default: Story = {}

export const WithoutAnimation: Story = {
    args: {
        enabledAnimation: false,
        children: 'アニメーション無効',
    },
}

export const LongText: Story = {
    args: {
        children: '商品をカートに追加しました',
        onClick: () => {
            console.log('商品をカートに追加しました')
        },
    },
}

export const ShortText: Story = {
    args: {
        children: '購入',
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
        onClick: () => {
            console.log('お気に入りに追加しました')
        },
    },
}

export const ActionButtons: Story = {
    args: {
        children: '管理者ログイン',
        onClick: () => {
            console.log('管理者ログインボタンがクリックされました')
        },
    },
}

export const NoClickHandler: Story = {
    args: {
        children: 'クリック無効ボタン',
    },
}

export const EnglishText: Story = {
    args: {
        children: 'Add to Cart',
        onClick: () => {
            console.log('Add to Cart clicked')
        },
    },
}
