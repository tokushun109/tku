import { useState } from 'react'

import { Button } from '@/components/bases/Button'
import { ColorType } from '@/types/enum/color'

import { Dialog } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Dialog> = {
    component: Dialog,
    args: {
        isOpen: true,
        title: 'ダイアログ',
        children: <p>これはダイアログの内容です。</p>,
        onClose: () => {
            console.log('ダイアログを閉じます')
        },
    },
    argTypes: {
        isOpen: {
            control: { type: 'boolean' },
        },
        title: {
            control: { type: 'text' },
        },
    },
}

export default meta
type Story = StoryObj<typeof Dialog>

export const Default: Story = {}

export const WithConfirmCancel: Story = {
    args: {
        title: 'ログアウト',
        children: <p>ログアウトします。よろしいですか？</p>,
        confirmOption: {
            label: 'はい',
            onClick: () => console.log('確認ボタンがクリックされました'),
        },
        cancelOption: {
            label: 'いいえ',
            onClick: () => console.log('キャンセルボタンがクリックされました'),
        },
    },
}

export const OnlyConfirm: Story = {
    args: {
        title: '通知',
        children: <p>操作が完了しました。</p>,
        confirmOption: {
            label: 'OK',
            onClick: () => console.log('確認ボタンがクリックされました'),
        },
    },
}

export const OnlyCancel: Story = {
    args: {
        title: '警告',
        children: <p>この操作は取り消すことができません。</p>,
        cancelOption: {
            label: '戻る',
            onClick: () => console.log('戻るボタンがクリックされました'),
        },
    },
}

export const NoButtons: Story = {
    args: {
        title: '情報',
        children: <p>このダイアログにはボタンがありません。背景をクリックして閉じてください。</p>,
    },
}

export const Interactive: Story = {
    render: () => {
        const [isOpen, setIsOpen] = useState<boolean>(false)
        const [dialogType, setDialogType] = useState<string>('confirm-cancel')

        const openDialog = (type: string) => {
            setDialogType(type)
            setIsOpen(true)
        }

        const closeDialog = () => {
            setIsOpen(false)
        }

        const getDialogProps = () => {
            switch (dialogType) {
                case 'confirm-cancel':
                    return {
                        title: 'ログアウト確認',
                        children: <p>ログアウトします。よろしいですか？</p>,
                        confirmOption: { label: 'はい', onClick: closeDialog },
                        cancelOption: { label: 'いいえ', onClick: closeDialog },
                    }
                case 'only-confirm':
                    return {
                        title: '操作完了',
                        children: <p>データの保存が完了しました。</p>,
                        confirmOption: { label: 'OK', onClick: closeDialog },
                    }
                case 'no-buttons':
                    return {
                        title: '読み込み中',
                        children: <p>データを読み込んでいます。しばらくお待ちください。</p>,
                    }
                default:
                    return {
                        title: 'デフォルト',
                        children: <p>デフォルトのダイアログです。</p>,
                    }
            }
        }

        return (
            <div style={{ display: 'flex', gap: '12px', flexWrap: 'wrap' }}>
                <Button colorType={ColorType.Primary} onClick={() => openDialog('confirm-cancel')}>
                    確認・キャンセル
                </Button>
                <Button colorType={ColorType.Accent} onClick={() => openDialog('only-confirm')}>
                    確認のみ
                </Button>
                <Button colorType={ColorType.Secondary} onClick={() => openDialog('no-buttons')}>
                    ボタンなし
                </Button>

                <Dialog isOpen={isOpen} onClose={closeDialog} {...getDialogProps()} />
            </div>
        )
    },
}

export const LongContent: Story = {
    args: {
        title: '利用規約',
        children: (
            <div>
                <p>この利用規約は、サービスの利用に関して定めるものです。</p>
                <p>
                    長いテキストの例：
                    ユーザーは本サービスを利用するにあたり、以下の事項を遵守するものとします。本サービスに関する一切の権利は当社に帰属します。ユーザーは本サービスを適切に利用し、他のユーザーに迷惑をかけないよう注意する必要があります。
                </p>
                <p>違反した場合は、アカウントの停止等の措置を行う場合があります。</p>
            </div>
        ),
        confirmOption: {
            label: '同意する',
            onClick: () => console.log('利用規約に同意しました'),
        },
        cancelOption: {
            label: '拒否する',
            onClick: () => console.log('利用規約を拒否しました'),
        },
    },
}
