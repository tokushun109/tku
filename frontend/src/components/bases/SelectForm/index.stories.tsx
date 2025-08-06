import { useState } from 'react'

import { SelectForm, SelectFormOption } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

// 型推論のためのオプション定義
const stringOptions: SelectFormOption<string>[] = [
    { label: 'ピアス', value: 'earrings' },
    { label: 'ネックレス', value: 'necklace' },
    { label: 'ブレスレット', value: 'bracelet' },
    { label: 'リング', value: 'ring' },
    { label: 'アンクレット', value: 'anklet' },
    { label: 'ヘアアクセサリー', value: 'hair-accessory' },
]

const meta: Meta<typeof SelectForm> = {
    component: SelectForm,
    args: {
        label: 'カテゴリー',
        placeholder: '選択してください',
        required: false,
        options: stringOptions,
    },
    argTypes: {
        required: {
            control: { type: 'boolean' },
        },
        value: {
            control: { type: 'text' },
        },
    },
}

export default meta
type Story = StoryObj<typeof SelectForm>

export const Default: Story = {
    render: () => {
        const [value, setValue] = useState<string>('')
        return <SelectForm label="カテゴリー" onChange={setValue} options={stringOptions} placeholder="選択してください" value={value} />
    },
}

export const Required: Story = {
    render: () => {
        const [value, setValue] = useState<string>('')
        return <SelectForm label="カテゴリー" onChange={setValue} options={stringOptions} placeholder="選択してください" required value={value} />
    },
}

export const WithValue: Story = {
    render: () => {
        const [value, setValue] = useState<string>('earrings')
        return <SelectForm label="カテゴリー" onChange={setValue} options={stringOptions} placeholder="選択してください" value={value} />
    },
}

export const WithClearButton: Story = {
    render: () => {
        const [value, setValue] = useState<string>('necklace')
        return (
            <SelectForm
                helperText="選択済みの値の右側にクリアボタン（×）が表示されます"
                label="カテゴリー"
                onChange={setValue}
                options={stringOptions}
                placeholder="選択してください"
                value={value}
            />
        )
    },
}

export const WithError: Story = {
    render: () => {
        const [value, setValue] = useState<string>('')
        return (
            <SelectForm
                error="カテゴリーを選択してください"
                label="カテゴリー"
                onChange={setValue}
                options={stringOptions}
                placeholder="選択してください"
                required
                value={value}
            />
        )
    },
}

export const WithHelperText: Story = {
    render: () => {
        const [value, setValue] = useState<string>('')
        return (
            <SelectForm
                helperText="商品に最も適したカテゴリーを選択してください"
                label="カテゴリー"
                onChange={setValue}
                options={stringOptions}
                placeholder="選択してください"
                value={value}
            />
        )
    },
}

const manyOptions: SelectFormOption<string>[] = [
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
        const [value, setValue] = useState<string>('')
        return <SelectForm label="カテゴリー" onChange={setValue} options={manyOptions} placeholder="選択してください" value={value} />
    },
}

export const NumberValue: Story = {
    render: () => {
        const [value, setValue] = useState<number>()
        const options: SelectFormOption<number>[] = [
            { label: '1位', value: 1 },
            { label: '2位', value: 2 },
            { label: '3位', value: 3 },
            { label: '4位', value: 4 },
            { label: '5位', value: 5 },
        ]

        return <SelectForm<number> label="順位" onChange={setValue} options={options} placeholder="順位を選択してください" value={value} />
    },
}
