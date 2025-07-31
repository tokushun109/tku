import { useState } from 'react'

import { IContactListItem } from '@/features/contact/type'

import { ContactList } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

// テスト用のモックデータ生成関数
const createMockContact = (id: number, overrides: Partial<IContactListItem> = {}): IContactListItem => ({
    id,
    name: `テストユーザー${id}`,
    company: `テスト会社${id}`,
    phoneNumber: `090-0000-000${id}`,
    email: `test${id}@example.com`,
    content: `これは${id}番目のお問い合わせ内容です。詳細な内容が含まれています。商品についての質問やご要望など、様々な内容のお問い合わせがあります。`,
    createdAt: new Date(2024, 0, id).toISOString(),
    ...overrides,
})

const meta: Meta<typeof ContactList> = {
    component: ContactList,
    decorators: [
        (Story) => (
            <div style={{ height: '600px', background: '#f5f5f5', padding: '20px' }}>
                <Story />
            </div>
        ),
    ],
    args: {
        contacts: [createMockContact(1), createMockContact(2), createMockContact(3), createMockContact(4), createMockContact(5)],
    },
    argTypes: {
        contacts: {
            description: 'お問い合わせ一覧データ',
        },
    },
}

export default meta
type Story = StoryObj<typeof ContactList>

export const Default: Story = {}

export const EmptyList: Story = {
    args: {
        contacts: [],
    },
}

export const SingleContact: Story = {
    args: {
        contacts: [createMockContact(1)],
    },
}

export const ManyContacts: Story = {
    args: {
        contacts: Array.from({ length: 50 }, (_, i) => createMockContact(i + 1)),
    },
}

export const LongContent: Story = {
    args: {
        contacts: [
            createMockContact(1, {
                name: 'とても長い名前のお客様田中花子さん',
                company: '非常に長い会社名の株式会社サンプルエンタープライズコーポレーション',
                content: 'これは非常に長いお問い合わせ内容です。'.repeat(20),
            }),
            createMockContact(2, {
                content:
                    'このお問い合わせには、商品についての詳細な質問が含まれています。材質、サイズ、価格、配送方法、返品条件など、様々な項目について詳しく教えてほしいという内容です。また、カスタマイズの可能性についても質問されています。',
            }),
            createMockContact(3),
        ],
    },
}

export const NoCompanyContacts: Story = {
    args: {
        contacts: [createMockContact(1, { company: undefined }), createMockContact(2, { company: '' }), createMockContact(3, { company: undefined })],
    },
}

export const NoPhoneContacts: Story = {
    args: {
        contacts: [
            createMockContact(1, { phoneNumber: undefined }),
            createMockContact(2, { phoneNumber: '' }),
            createMockContact(3, { phoneNumber: undefined }),
        ],
    },
}

export const MixedData: Story = {
    args: {
        contacts: [
            createMockContact(1, {
                name: '田中太郎',
                company: '株式会社テスト',
                phoneNumber: '090-1234-5678',
                email: 'tanaka@test.co.jp',
                content: '商品についてお伺いしたいことがあります。',
            }),
            createMockContact(2, {
                name: '佐藤花子',
                company: undefined,
                phoneNumber: undefined,
                email: 'hanako@example.com',
                content: 'オーダーメイドは可能でしょうか？',
            }),
            createMockContact(3, {
                name: '山田次郎',
                company: 'フリーランス',
                phoneNumber: '080-9876-5432',
                email: 'yamada.jiro@freelance.jp',
                content: 'とても長いお問い合わせ内容です。'.repeat(10),
            }),
        ],
    },
}

export const Interactive: Story = {
    render: () => {
        const [selectedContact, setSelectedContact] = useState<IContactListItem | null>(null)
        const contacts = Array.from({ length: 10 }, (_, i) => createMockContact(i + 1))

        return (
            <div>
                <ContactList contacts={contacts} />
                {selectedContact && (
                    <div
                        style={{
                            position: 'fixed',
                            top: '50%',
                            left: '50%',
                            transform: 'translate(-50%, -50%)',
                            background: 'white',
                            padding: '20px',
                            border: '1px solid #ccc',
                            borderRadius: '8px',
                            boxShadow: '0 4px 8px rgba(0,0,0,0.2)',
                            maxWidth: '500px',
                            width: '90%',
                        }}
                    >
                        <h3>選択されたお問い合わせ</h3>
                        <p>
                            <strong>名前:</strong> {selectedContact.name}
                        </p>
                        <p>
                            <strong>会社:</strong> {selectedContact.company || '未入力'}
                        </p>
                        <p>
                            <strong>メール:</strong> {selectedContact.email}
                        </p>
                        <p>
                            <strong>内容:</strong> {selectedContact.content}
                        </p>
                        <button onClick={() => setSelectedContact(null)}>閉じる</button>
                    </div>
                )}
            </div>
        )
    },
}

export const VeryLargeDataset: Story = {
    args: {
        contacts: Array.from({ length: 1000 }, (_, i) =>
            createMockContact(i + 1, {
                name: `お客様${i + 1}`,
                company: i % 5 === 0 ? undefined : `会社${i + 1}`,
                phoneNumber: i % 3 === 0 ? undefined : `090-${String(i).padStart(4, '0')}-${String(i).padStart(4, '0')}`,
                content: `お問い合わせ内容${i + 1}: ${i % 10 === 0 ? 'とても長い内容です。'.repeat(5) : '通常の長さの内容です。'}`,
            }),
        ),
    },
}

export const RecentContacts: Story = {
    args: {
        contacts: [
            createMockContact(1, {
                name: '田中太郎',
                createdAt: new Date().toISOString(),
            }),
            createMockContact(2, {
                name: '佐藤花子',
                createdAt: new Date(Date.now() - 24 * 60 * 60 * 1000).toISOString(), // 1日前
            }),
            createMockContact(3, {
                name: '山田次郎',
                createdAt: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString(), // 1週間前
            }),
        ],
    },
}

export const LongAlphabetContent: Story = {
    args: {
        contacts: [
            createMockContact(1, {
                name: 'VeryVeryVeryVeryLongAlphabeticalNameWithoutSpaces',
                company: 'SuperLongCompanyNameWithoutAnySpacesOrHyphensToTestWordBreaking',
                email: 'verylongemailaddresswithoutspaces@extremelylongdomainnameforwordbreaktesting.com',
                phoneNumber: '090-1234-5678-9012-3456',
                content:
                    'ThisIsAVeryLongAlphabeticalTextWithoutAnySpacesOrBreaksToTestTheWordBreakingFunctionalityInTheDialogContentAreaAndMakeSureThatLongWordsDoNotOverflowBeyondTheDialogBoundariesButInsteadWrapCorrectlyToTheNextLineWithoutCausingHorizontalScrollingOrLayoutIssues. https://www.verylongurltotestwordbreaking.com/with/very/long/path/segments/that/should/break/correctly/when/displayed/in/the/dialog/component. AnotherLongWordFollowedByMoreTextToEnsureProperWrappingBehavior.',
            }),
            createMockContact(2, {
                name: 'John Smith',
                email: 'john.smith@company.com',
                content:
                    'Mixed content with normal text and VeryLongAlphabeticalWordsWithoutSpacesToTestWrapping. URL: https://example.com/very/long/url/path/that/should/wrap/properly/in/the/dialog/component/without/causing/overflow/issues.',
            }),
            createMockContact(3, {
                name: '田中太郎',
                email: 'tanaka@example.com',
                content:
                    '日本語とEnglishMixedContentWithVeryLongAlphabeticalWordsが混在しているテストケースです。URLやメールアドレスなど：verylongemailaddress@extremelylongdomainname.co.jp のような長い文字列も適切に折り返されることを確認します。',
            }),
        ],
    },
}
