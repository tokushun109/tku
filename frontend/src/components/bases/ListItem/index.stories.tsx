import { Delete, Edit } from '@mui/icons-material'

import { ListItem } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ListItem> = {
    component: ListItem,
    args: {
        children: 'リストアイテム',
    },
    argTypes: {
        className: {
            control: { type: 'text' },
        },
    },
}

export default meta
type Story = StoryObj<typeof ListItem>

export const Default: Story = {}

export const WithSubItem: Story = {
    args: {
        children: 'メインコンテンツ',
        subItem: 'サブコンテンツ',
    },
}

export const WithActions: Story = {
    args: {
        children: 'アクション付きアイテム',
        actions: (
            <>
                <Edit style={{ cursor: 'pointer', color: '#666' }} />
                <Delete style={{ cursor: 'pointer', color: '#f44336' }} />
            </>
        ),
    },
}

export const WithSubItemAndActions: Story = {
    args: {
        children: 'メインコンテンツ',
        subItem: 'https://example.com',
        actions: (
            <>
                <Edit style={{ cursor: 'pointer', color: '#666' }} />
                <Delete style={{ cursor: 'pointer', color: '#f44336' }} />
            </>
        ),
    },
}

export const Clickable: Story = {
    args: {
        children: 'クリック可能なアイテム',
        subItem: 'クリックしてください',
        onClick: () => alert('クリックされました！'),
    },
}

export const LongContent: Story = {
    args: {
        children: 'とても長いメインコンテンツのサンプルテキストです。このテキストは複数行にわたって表示される可能性があります。',
        subItem: 'https://very-long-url-example.com/path/to/some/resource/with/a/very/long/name',
        actions: (
            <>
                <Edit style={{ cursor: 'pointer', color: '#666' }} />
                <Delete style={{ cursor: 'pointer', color: '#f44336' }} />
            </>
        ),
    },
}

export const CustomClassName: Story = {
    args: {
        children: 'カスタムクラス付きアイテム',
        className: 'custom-item',
        subItem: 'カスタムスタイルが適用されています',
    },
}

export const MultipleItems: Story = {
    render: () => (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
            <ListItem actions={<Delete style={{ cursor: 'pointer', color: '#f44336' }} />}>アイテム 1</ListItem>
            <ListItem
                actions={
                    <>
                        <Edit style={{ cursor: 'pointer', color: '#666' }} />
                        <Delete style={{ cursor: 'pointer', color: '#f44336' }} />
                    </>
                }
                subItem="https://example1.com"
            >
                アイテム 2
            </ListItem>
            <ListItem onClick={() => alert('アイテム 3 がクリックされました')} subItem="サブコンテンツ">
                アイテム 3（クリック可能）
            </ListItem>
        </div>
    ),
}
