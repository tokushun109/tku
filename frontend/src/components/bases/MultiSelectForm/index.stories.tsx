import { useState } from 'react'

import { MultiSelectForm, MultiSelectFormOption } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

// 型推論のためのオプション定義
const stringOptions: MultiSelectFormOption<string>[] = [
    { label: 'ピアス', value: 'earrings' },
    { label: 'ネックレス', value: 'necklace' },
    { label: 'ブレスレット', value: 'bracelet' },
    { label: 'リング', value: 'ring' },
    { label: 'アンクレット', value: 'anklet' },
    { label: 'ヘアアクセサリー', value: 'hair-accessory' },
]

const meta: Meta<typeof MultiSelectForm> = {
    component: MultiSelectForm,
    args: {
        label: 'タグ',
        placeholder: '選択してください',
        required: false,
        options: stringOptions,
    },
    argTypes: {
        required: {
            control: { type: 'boolean' },
        },
        value: {
            control: { type: 'object' },
        },
    },
}

export default meta
type Story = StoryObj<typeof MultiSelectForm>

export const Default: Story = {
    render: () => {
        const [value, setValue] = useState<string[]>([])
        return <MultiSelectForm label="タグ" onChange={setValue} options={stringOptions} placeholder="タグを選択してください" value={value} />
    },
}

export const WithValue: Story = {
    render: () => {
        const [value, setValue] = useState<string[]>(['earrings', 'necklace'])
        return <MultiSelectForm label="タグ" onChange={setValue} options={stringOptions} placeholder="タグを選択してください" value={value} />
    },
}

export const Required: Story = {
    render: () => {
        const [value, setValue] = useState<string[]>([])
        return (
            <MultiSelectForm label="タグ" onChange={setValue} options={stringOptions} placeholder="タグを選択してください" required value={value} />
        )
    },
}

export const WithError: Story = {
    render: () => {
        const [value, setValue] = useState<string[]>([])
        return (
            <MultiSelectForm
                error="少なくとも1つのタグを選択してください"
                label="タグ"
                onChange={setValue}
                options={stringOptions}
                placeholder="タグを選択してください"
                required
                value={value}
            />
        )
    },
}

export const WithHelperText: Story = {
    render: () => {
        const [value, setValue] = useState<string[]>([])
        return (
            <MultiSelectForm
                helperText="商品に適用するタグを複数選択できます"
                label="タグ"
                onChange={setValue}
                options={stringOptions}
                placeholder="タグを選択してください"
                value={value}
            />
        )
    },
}

const manyOptions: MultiSelectFormOption<string>[] = [
    { label: 'ピアス', value: 'earrings' },
    { label: 'ネックレス', value: 'necklace' },
    { label: 'ブレスレット', value: 'bracelet' },
    { label: 'リング', value: 'ring' },
    { label: 'アンクレット', value: 'anklet' },
    { label: 'ヘアアクセサリー', value: 'hair-accessory' },
    { label: 'バッグチャーム', value: 'bag-charm' },
    { label: 'キーホルダー', value: 'key-holder' },
    { label: 'ブローチ', value: 'brooch' },
    { label: 'カフス', value: 'cufflinks' },
    { label: 'ウォッチ', value: 'watch' },
    { label: 'サングラス', value: 'sunglasses' },
]

export const ManyOptions: Story = {
    render: () => {
        const [value, setValue] = useState<string[]>(['earrings', 'necklace', 'bracelet'])
        return (
            <MultiSelectForm
                helperText="検索機能を使って必要なタグを見つけてください"
                label="タグ"
                onChange={setValue}
                options={manyOptions}
                placeholder="タグを選択してください"
                value={value}
            />
        )
    },
}

export const NumberValue: Story = {
    render: () => {
        const [value, setValue] = useState<number[]>([])
        const options: MultiSelectFormOption<number>[] = [
            { label: '高優先度', value: 1 },
            { label: '中優先度', value: 2 },
            { label: '低優先度', value: 3 },
            { label: '緊急', value: 4 },
            { label: '後回し', value: 5 },
        ]

        return <MultiSelectForm<number> label="優先度" onChange={setValue} options={options} placeholder="優先度を選択してください" value={value} />
    },
}
