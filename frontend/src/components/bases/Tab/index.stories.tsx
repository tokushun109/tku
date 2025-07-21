import { Category, Label, People } from '@mui/icons-material'

import { Tab } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Tab> = {
    component: Tab,
    args: {
        items: [
            { key: 'tab1', label: 'タブ1', icon: <Category /> },
            { key: 'tab2', label: 'タブ2', icon: <People /> },
            { key: 'tab3', label: 'タブ3', icon: <Label /> },
        ],
        activeKey: 'tab1',
        onTabChange: (key) => console.log('Tab changed to:', key),
    },
    argTypes: {
        items: {
            description: 'タブアイテムの配列',
        },
        activeKey: {
            control: { type: 'text' },
            description: 'アクティブなタブのキー',
        },
        onTabChange: {
            action: 'tab changed',
            description: 'タブ変更時のコールバック関数',
        },
    },
}

export default meta
type Story = StoryObj<typeof Tab>

export const Default: Story = {}

export const ManyTabs: Story = {
    args: {
        items: Array.from({ length: 6 }, (_, i) => ({
            key: `tab${i + 1}`,
            label: `タブ${i + 1}`,
        })),
        activeKey: 'tab3',
    },
}

export const LongLabels: Story = {
    args: {
        items: [
            { key: 'very-long', label: 'とても長いタブ名前の例' },
            { key: 'short', label: '短い' },
            { key: 'medium', label: '中程度の長さ' },
        ],
        activeKey: 'very-long',
    },
}

export const SingleTab: Story = {
    args: {
        items: [{ key: 'only', label: '唯一のタブ' }],
        activeKey: 'only',
    },
}

export const WithoutIcons: Story = {
    args: {
        items: [
            { key: 'tab1', label: 'アイコンなし1' },
            { key: 'tab2', label: 'アイコンなし2' },
            { key: 'tab3', label: 'アイコンなし3' },
        ],
        activeKey: 'tab1',
    },
}

export const WithNumberKeys: Story = {
    args: {
        items: [
            { key: 1, label: 'タブ1' },
            { key: 2, label: 'タブ2' },
            { key: 3, label: 'タブ3' },
        ],
        activeKey: 1,
        onTabChange: (key) => console.log('Number key changed to:', key),
    },
}
