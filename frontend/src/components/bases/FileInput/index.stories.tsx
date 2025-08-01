import { useState } from 'react'

import { FileInput } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof FileInput> = {
    component: FileInput,
    args: {
        label: 'ファイル選択',
        accept: '.csv',
    },
    argTypes: {
        required: {
            control: { type: 'boolean' },
        },
        error: {
            control: { type: 'text' },
        },
        helperText: {
            control: { type: 'text' },
        },
        accept: {
            control: { type: 'text' },
        },
    },
}

export default meta
type Story = StoryObj<typeof FileInput>

export const Default: Story = {}

export const Required: Story = {
    args: {
        required: true,
    },
}

export const WithError: Story = {
    args: {
        error: 'ファイルが選択されていません',
    },
}

export const WithHelperText: Story = {
    args: {
        helperText: 'CSVファイルを選択してください',
    },
}

export const CSVOnly: Story = {
    args: {
        accept: '.csv',
        helperText: 'CSV形式のファイルのみ選択可能です',
    },
}

export const ImageOnly: Story = {
    args: {
        accept: 'image/*',
        label: '画像ファイル選択',
        helperText: '画像ファイルを選択してください',
    },
}

export const FormExample: Story = {
    render: (args) => {
        const [file, setFile] = useState<File | null>(null)
        const [error, setError] = useState<string>('')

        const handleFileChange = (selectedFile: File | null) => {
            setFile(selectedFile)
            if (selectedFile) {
                setError('')
            } else {
                setError('ファイルを選択してください')
            }
        }

        return (
            <div style={{ maxWidth: '400px' }}>
                <FileInput {...args} error={error} onChange={handleFileChange} required value={file} />
                <div style={{ marginTop: '16px', fontSize: '14px', color: '#666' }}>
                    選択されたファイル: {file ? file.name : 'なし'}
                    <br />
                    ファイルサイズ: {file ? `${(file.size / 1024).toFixed(2)} KB` : '-'}
                </div>
            </div>
        )
    },
    args: {
        label: 'CSVファイル',
        accept: '.csv',
        helperText: 'CSVファイルを選択してください',
    },
}
