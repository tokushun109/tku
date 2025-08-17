import { TopLoading } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof TopLoading> = {
    component: TopLoading,
    parameters: {
        layout: 'fullscreen',
        docs: {
            description: {
                component: 'トップページ表示時に画面全体にロゴを表示するローディングコンポーネント。1秒後に自動でフェードアウトします。',
            },
        },
    },
}

export default meta
type Story = StoryObj<typeof TopLoading>

export const Default: Story = {
    parameters: {
        docs: {
            description: {
                story: 'デフォルトの状態。1秒後にフェードアウトが開始されます。',
            },
        },
    },
}
