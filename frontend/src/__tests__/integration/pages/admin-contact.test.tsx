import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getContacts } from '@/apis/contact'
import AdminContactPage, { metadata } from '@/app/admin/contact/page'
import { IContactListItem } from '@/features/contact/type'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/contact')

const mockGetContacts = vi.mocked(getContacts)

// テスト用のモックデータ
const createMockContact = (id: number, overrides: Partial<IContactListItem> = {}): IContactListItem => ({
    id,
    name: `テストユーザー${id}`,
    company: `テスト会社${id}`,
    phoneNumber: `090-0000-000${id}`,
    email: `test${id}@example.com`,
    content: `これは${id}番目のお問い合わせ内容です。詳細な内容が含まれています。`,
    createdAt: new Date(2024, 0, id).toISOString(),
    ...overrides,
})

describe('Admin Contact Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('管理画面お問い合わせ一覧ページが正常に表示される', async () => {
        // モックデータの設定
        const mockContacts: IContactListItem[] = [createMockContact(1), createMockContact(2), createMockContact(3)]

        mockGetContacts.mockResolvedValue(mockContacts)

        // コンポーネントをレンダリング
        render(await AdminContactPage())

        // ページタイトルの表示を確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせ一覧')).toBeInTheDocument()
        })

        // 件数表示の確認
        await waitFor(() => {
            expect(screen.getByText('3件のお問い合わせ')).toBeInTheDocument()
        })

        // ContactListコンポーネントが表示されることを確認
        await waitFor(() => {
            expect(screen.getByTestId('virtuoso-scroller')).toBeInTheDocument()
        })

        // API呼び出しの確認
        expect(mockGetContacts).toHaveBeenCalledTimes(1)
    })

    it('お問い合わせがない場合でもエラーにならない', async () => {
        // 空配列を返すモック
        mockGetContacts.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await AdminContactPage())

        // ページタイトルの表示を確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせ一覧')).toBeInTheDocument()
        })

        // 件数表示の確認
        await waitFor(() => {
            expect(screen.getByText('0件のお問い合わせ')).toBeInTheDocument()
        })

        // 空メッセージの表示を確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせがありません')).toBeInTheDocument()
        })
    })

    it('API呼び出しが失敗した場合のエラーハンドリング', async () => {
        // コンソールエラーをモック
        const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        // モックでエラーを発生させる
        mockGetContacts.mockRejectedValue(new Error('Contact API Error'))

        // コンポーネントをレンダリング
        render(await AdminContactPage())

        // エラーログが出力されることを確認
        await waitFor(() => {
            expect(consoleSpy).toHaveBeenCalledWith('お問い合わせ一覧の取得に失敗しました:', expect.any(Error))
        })

        // 空配列でテンプレートが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせ一覧')).toBeInTheDocument()
            expect(screen.getByText('0件のお問い合わせ')).toBeInTheDocument()
        })

        consoleSpy.mockRestore()
    })

    it('大量のお問い合わせがある場合でも正常に表示される', async () => {
        // 大量のモックデータを作成
        const mockContacts: IContactListItem[] = Array.from({ length: 50 }, (_, index) => createMockContact(index + 1))

        mockGetContacts.mockResolvedValue(mockContacts)

        // コンポーネントをレンダリング
        render(await AdminContactPage())

        // ページタイトルと件数の表示を確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせ一覧')).toBeInTheDocument()
            expect(screen.getByText('50件のお問い合わせ')).toBeInTheDocument()
        })

        // 仮想スクロールコンポーネントが表示されることを確認
        await waitFor(() => {
            expect(screen.getByTestId('virtuoso-scroller')).toBeInTheDocument()
        })
    })

    it('API呼び出しが正常に実行される', async () => {
        // モックデータの設定
        const mockContacts: IContactListItem[] = [createMockContact(1)]

        mockGetContacts.mockResolvedValue(mockContacts)

        // コンポーネントをレンダリング
        render(await AdminContactPage())

        // API呼び出しの確認
        expect(mockGetContacts).toHaveBeenCalledTimes(1)

        // 基本的な表示確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせ一覧')).toBeInTheDocument()
            expect(screen.getByText('1件のお問い合わせ')).toBeInTheDocument()
        })
    })

    it('ページメタデータでnoindex, nofollowが設定されている', () => {
        // メタデータの確認
        expect(metadata.title).toBe('お問い合わせ管理 | admin')
        expect(metadata.robots).toEqual({
            index: false,
            follow: false,
        })
    })
})
