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
        children: <p style={{ textAlign: 'center' }}>ログアウトします。よろしいですか？</p>,
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

export const NoTitle: Story = {
    args: {
        children: <p>このダイアログにはタイトルがありません。ヘッダーは表示されません。</p>,
        title: undefined, // タイトルを省略
        confirmOption: {
            label: 'OK',
            onClick: () => console.log('確認ボタンがクリックされました'),
        },
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
                        children: <p style={{ textAlign: 'center' }}>ログアウトします。よろしいですか？</p>,
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
                case 'no-title':
                    return {
                        children: <p>タイトルなしのダイアログです。ヘッダーは表示されません。</p>,
                        confirmOption: { label: 'OK', onClick: closeDialog },
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
                <Button colorType={ColorType.Danger} onClick={() => openDialog('no-title')}>
                    タイトルなし
                </Button>

                <Dialog isOpen={isOpen} onClose={closeDialog} {...getDialogProps()} />
            </div>
        )
    },
}

export const Wide: Story = {
    args: {
        title: '横幅の広いダイアログ',
        wide: true,
        children: (
            <div>
                <p>このダイアログは横幅が広く設定されています。お問い合わせの詳細など、多くの情報を表示する際に使用します。</p>
                <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '16px', marginTop: '16px' }}>
                    <div>
                        <strong>名前:</strong> 田中太郎
                    </div>
                    <div>
                        <strong>会社:</strong> 株式会社サンプル
                    </div>
                    <div>
                        <strong>電話番号:</strong> 090-1234-5678
                    </div>
                    <div>
                        <strong>メール:</strong> tanaka@example.com
                    </div>
                </div>
                <div style={{ marginTop: '16px' }}>
                    <strong>内容:</strong>
                    <p>こちらのアクセサリーについて詳しく教えてください。</p>
                </div>
            </div>
        ),
        confirmOption: {
            label: '閉じる',
            onClick: () => console.log('ダイアログを閉じました'),
        },
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

export const ScrollableContent: Story = {
    args: {
        title: 'スクロール可能なコンテンツ',
        children: (
            <div>
                <p>このダイアログはコンテンツがスクロール可能です。内容が長い場合でもすべての情報を確認できます。</p>
                {Array.from({ length: 20 }, (_, i) => (
                    <p key={i} style={{ margin: '12px 0' }}>
                        {i + 1}. これは{i + 1}
                        番目の段落です。スクロールして全ての内容を確認してください。ダイアログの高さを超える場合、自動的にスクロールバーが表示されます。
                    </p>
                ))}
                <p style={{ marginTop: '20px', padding: '10px', backgroundColor: '#f0f0f0', borderRadius: '4px' }}>
                    最後の段落です。ここまでスクロールできていれば、スクロール機能が正常に動作しています。
                </p>
            </div>
        ),
        confirmOption: {
            label: 'OK',
            onClick: () => console.log('スクロール確認完了'),
        },
    },
}

export const WideScrollableContent: Story = {
    args: {
        title: '横幅広・スクロール可能',
        wide: true,
        children: (
            <div>
                <h4>お問い合わせ詳細</h4>
                <div style={{ display: 'grid', gridTemplateColumns: '150px 1fr', gap: '8px', marginBottom: '16px' }}>
                    <div>
                        <strong>受付日時:</strong>
                    </div>
                    <div>2024年1月15日 14:30</div>
                    <div>
                        <strong>お名前:</strong>
                    </div>
                    <div>田中花子</div>
                    <div>
                        <strong>会社名:</strong>
                    </div>
                    <div>株式会社サンプル企業</div>
                    <div>
                        <strong>電話番号:</strong>
                    </div>
                    <div>090-1234-5678</div>
                    <div>
                        <strong>メールアドレス:</strong>
                    </div>
                    <div>hanako.tanaka@sample-company.co.jp</div>
                </div>
                <div style={{ marginTop: '20px' }}>
                    <strong>お問い合わせ内容:</strong>
                    <div style={{ marginTop: '8px', padding: '16px', backgroundColor: '#f9f9f9', borderRadius: '4px' }}>
                        <p>この度は、素敵なアクセサリーを拝見させていただき、ありがとうございます。</p>
                        <p>商品について、いくつか質問があります：</p>
                        <ul>
                            <li>材質について詳しく教えてください</li>
                            <li>サイズ調整は可能でしょうか</li>
                            <li>オーダーメイドは承っていますか</li>
                            <li>配送方法と送料について</li>
                            <li>返品・交換の条件について</li>
                        </ul>
                        <p>また、今後の新作情報なども教えていただけると嬉しいです。</p>
                        <p>お忙しい中恐縮ですが、よろしくお願いいたします。</p>
                        {Array.from({ length: 10 }, (_, i) => (
                            <p key={i}>
                                追加の質問{i + 1}:
                                これは追加の質問内容です。長いお問い合わせの場合、このようにスクロールして全ての内容を確認できます。
                            </p>
                        ))}
                    </div>
                </div>
            </div>
        ),
        confirmOption: {
            label: '返信する',
            onClick: () => console.log('返信画面を開きます'),
        },
        cancelOption: {
            label: '閉じる',
            onClick: () => console.log('ダイアログを閉じます'),
        },
    },
}
