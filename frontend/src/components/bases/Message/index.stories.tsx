import { Message, MessageType } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Message> = {
    component: Message,
    args: {
        children: 'メッセージの内容です',
        type: MessageType.Error,
    },
    argTypes: {
        type: {
            control: { type: 'select' },
            options: Object.values(MessageType),
        },
    },
}

export default meta
type Story = StoryObj<typeof Message>

export const Default: Story = {}

export const Error: Story = {
    args: {
        type: MessageType.Error,
        children: 'エラーが発生しました。もう一度お試しください。',
    },
}

export const Warning: Story = {
    args: {
        type: MessageType.Warning,
        children: '警告：この操作は元に戻すことができません。',
    },
}

export const Info: Story = {
    args: {
        type: MessageType.Info,
        children: '情報：データの更新が完了しました。',
    },
}

export const Success: Story = {
    args: {
        type: MessageType.Success,
        children: '成功：アカウントが正常に作成されました。',
    },
}

export const LongMessage: Story = {
    args: {
        type: MessageType.Error,
        children:
            'これは非常に長いエラーメッセージの例です。複数行にわたってメッセージが表示される場合の見た目を確認するために、意図的に長い文章を用意しています。',
    },
}
