import { ColorType } from '@/types'

import { Button } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Button> = {
    component: Button,
    args: {
        children: 'button',
        colorType: ColorType.Primary,
        contrast: false,
        disabled: false,
        noBoxShadow: false,
        outlined: false,
        onClick: () => {
            console.log('clickしました')
        },
    },
    argTypes: {
        colorType: {
            control: { type: 'select' },
            options: Object.values(ColorType),
        },
        contrast: {
            control: { type: 'boolean' },
        },
        disabled: {
            control: { type: 'boolean' },
        },
        noBoxShadow: {
            control: { type: 'boolean' },
        },
        outlined: {
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

export const Contrast: Story = {
    args: {
        children: 'Contrast Button',
        colorType: ColorType.Primary,
        contrast: true,
    },
}

export const Outlined: Story = {
    args: {
        children: 'Outlined Button',
        colorType: ColorType.Primary,
        outlined: true,
    },
}

export const OutlinedContrast: Story = {
    args: {
        children: 'Outlined Contrast',
        colorType: ColorType.Primary,
        contrast: true,
        outlined: true,
    },
}

export const NoBoxShadow: Story = {
    args: {
        children: 'No Shadow Button',
        colorType: ColorType.Primary,
        noBoxShadow: true,
    },
}

export const DialogButtons: Story = {
    render: () => (
        <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
            <Button colorType={ColorType.Primary} contrast outlined>
                いいえ
            </Button>
            <Button colorType={ColorType.Primary}>はい</Button>
        </div>
    ),
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

export const DisabledOutlined: Story = {
    args: {
        children: '無効化されたアウトライン',
        colorType: ColorType.Primary,
        disabled: true,
        outlined: true,
    },
}

export const FormExample: Story = {
    render: () => (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '12px', maxWidth: '300px' }}>
            <Button colorType={ColorType.Primary} type="submit">
                送信する
            </Button>
            <Button colorType={ColorType.Secondary} contrast outlined>
                キャンセル
            </Button>
        </div>
    ),
}

export const ActionButtons: Story = {
    render: () => (
        <div style={{ display: 'flex', gap: '12px', flexWrap: 'wrap' }}>
            <Button colorType={ColorType.Primary}>管理者ログイン</Button>
            <Button colorType={ColorType.Accent}>商品をカートに追加</Button>
            <Button colorType={ColorType.Secondary} outlined>
                お気に入りに追加
            </Button>
            <Button colorType={ColorType.Danger} contrast>
                削除
            </Button>
        </div>
    ),
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
        outlined: true,
        onClick: () => {
            console.log('お気に入りに追加しました')
        },
    },
}
