import { fireEvent, waitFor } from '@testing-library/react'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { postContact } from '@/apis/contact'
import ContactPage from '@/app/(contents)/contact/page'

import { render, screen, mockUseRouter, mockRouter, resetRouterMock } from '../helpers'

// APIモック
vi.mock('@/apis/contact')
vi.mock('next/navigation', () => ({
    useRouter: vi.fn(),
}))

const mockPostContact = vi.mocked(postContact)

describe('Contact Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
        resetRouterMock()
        mockUseRouter()
    })

    it('お問い合わせページが正常に表示される', async () => {
        // コンポーネントをレンダリング
        render(<ContactPage />)

        // フォーム要素の表示を確認
        await waitFor(() => {
            expect(screen.getByText('お問い合わせ・ご意見・ご相談はこちらから')).toBeInTheDocument()
            expect(screen.getByLabelText('お名前')).toBeInTheDocument()
            expect(screen.getByLabelText('会社名')).toBeInTheDocument()
            expect(screen.getByLabelText('電話番号(-を入れずに入力)')).toBeInTheDocument()
            expect(screen.getByLabelText('メールアドレス')).toBeInTheDocument()
            expect(screen.getByLabelText('お問い合わせ内容')).toBeInTheDocument()
            expect(screen.getByRole('button', { name: '送信する' })).toBeInTheDocument()
        })
    })

    it('フォームの入力値が正常に処理される', async () => {
        // モックデータの設定
        mockPostContact.mockResolvedValue({ message: 'お問い合わせを受け付けました' })

        // コンポーネントをレンダリング
        render(<ContactPage />)

        // フォームに入力
        const nameInput = screen.getByLabelText('お名前')
        const companyInput = screen.getByLabelText('会社名')
        const phoneInput = screen.getByLabelText('電話番号(-を入れずに入力)')
        const emailInput = screen.getByLabelText('メールアドレス')
        const contentInput = screen.getByLabelText('お問い合わせ内容')

        fireEvent.change(nameInput, { target: { value: '山田太郎' } })
        fireEvent.change(companyInput, { target: { value: '株式会社テスト' } })
        fireEvent.change(phoneInput, { target: { value: '09012345678' } })
        fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
        fireEvent.change(contentInput, { target: { value: 'テストお問い合わせ内容' } })

        // 送信ボタンを有効化するまで待機
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).not.toBeDisabled()
        })

        // フォームを送信
        fireEvent.click(screen.getByRole('button', { name: '送信する' }))

        // APIが正しい値で呼び出されることを確認
        await waitFor(() => {
            expect(mockPostContact).toHaveBeenCalledWith({
                form: {
                    name: '山田太郎',
                    company: '株式会社テスト',
                    phoneNumber: '09012345678',
                    email: 'test@example.com',
                    content: 'テストお問い合わせ内容',
                },
            })
        })

        // 成功メッセージが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText(/お問い合わせを送信しました/)).toBeInTheDocument()
        })

        // 3秒後にリダイレクトされることを確認
        await waitFor(
            () => {
                expect(mockRouter.push).toHaveBeenCalledWith('/')
            },
            { timeout: 4000 },
        )
    })

    it('必須フィールドの入力がない場合はバリデーションエラーが表示される', async () => {
        // コンポーネントをレンダリング
        render(<ContactPage />)

        // 必須フィールドの要素を取得
        const nameInput = screen.getByLabelText('お名前')
        const emailInput = screen.getByLabelText('メールアドレス')
        const contentInput = screen.getByLabelText('お問い合わせ内容')

        // 必須フィールドにまず値を入力
        fireEvent.change(nameInput, { target: { value: '山田太郎' } })
        fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
        fireEvent.change(contentInput, { target: { value: 'テスト内容' } })

        // その後、空文字にしてバリデーションを実行
        fireEvent.change(nameInput, { target: { value: '' } })
        fireEvent.change(emailInput, { target: { value: '' } })
        fireEvent.change(contentInput, { target: { value: '' } })

        // フォーカスを外してバリデーションを実行
        fireEvent.blur(nameInput)
        fireEvent.blur(emailInput)
        fireEvent.blur(contentInput)

        // バリデーションエラーが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('お名前は必須項目です')).toBeInTheDocument()
            expect(screen.getByText('メールアドレスは必須項目です')).toBeInTheDocument()
            expect(screen.getByText('お問い合わせ内容は必須項目です')).toBeInTheDocument()
        })

        // 送信ボタンが無効化されることを確認
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).toBeDisabled()
        })
    })

    it('メールアドレスの形式が正しくない場合はバリデーションエラーが表示される', async () => {
        // コンポーネントをレンダリング
        render(<ContactPage />)

        // 不正なメールアドレスを入力
        const emailInput = screen.getByLabelText('メールアドレス')
        fireEvent.change(emailInput, { target: { value: 'invalid-email' } })
        fireEvent.blur(emailInput)

        // バリデーションエラーが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText(/正しいメールアドレスを入力してください/)).toBeInTheDocument()
        })

        // 送信ボタンが無効化されることを確認
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).toBeDisabled()
        })
    })

    it('API呼び出しが失敗した場合はエラーメッセージが表示される', async () => {
        // モックでエラーを発生させる
        mockPostContact.mockRejectedValue(new Error('API Error'))

        // コンポーネントをレンダリング
        render(<ContactPage />)

        // 有効なフォーム入力
        const nameInput = screen.getByLabelText('お名前')
        const emailInput = screen.getByLabelText('メールアドレス')
        const contentInput = screen.getByLabelText('お問い合わせ内容')

        fireEvent.change(nameInput, { target: { value: '山田太郎' } })
        fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
        fireEvent.change(contentInput, { target: { value: 'テストお問い合わせ内容' } })

        // 送信ボタンを有効化するまで待機
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).not.toBeDisabled()
        })

        // フォームを送信
        fireEvent.click(screen.getByRole('button', { name: '送信する' }))

        // エラーメッセージが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText(/送信中にエラーが発生しました。もう一度お試しください。/)).toBeInTheDocument()
        })

        // 送信ボタンが再度有効化されることを確認
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).not.toBeDisabled()
        })
    })

    it('送信中は送信ボタンが無効化される', async () => {
        // モックを遅延させる
        mockPostContact.mockImplementation(() => new Promise((resolve) => setTimeout(resolve, 1000)))

        // コンポーネントをレンダリング
        render(<ContactPage />)

        // 有効なフォーム入力
        const nameInput = screen.getByLabelText('お名前')
        const emailInput = screen.getByLabelText('メールアドレス')
        const contentInput = screen.getByLabelText('お問い合わせ内容')

        fireEvent.change(nameInput, { target: { value: '山田太郎' } })
        fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
        fireEvent.change(contentInput, { target: { value: 'テストお問い合わせ内容' } })

        // 送信ボタンを有効化するまで待機
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).not.toBeDisabled()
        })

        // フォームを送信
        fireEvent.click(screen.getByRole('button', { name: '送信する' }))

        // 送信中は送信ボタンが無効化されることを確認
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信中...' })).toBeDisabled()
        })
    })

    it('任意フィールドのみの入力でも送信が可能', async () => {
        // モックデータの設定
        mockPostContact.mockResolvedValue({ message: 'お問い合わせを受け付けました' })

        // コンポーネントをレンダリング
        render(<ContactPage />)

        // 必須フィールドのみ入力
        const nameInput = screen.getByLabelText('お名前')
        const emailInput = screen.getByLabelText('メールアドレス')
        const contentInput = screen.getByLabelText('お問い合わせ内容')

        fireEvent.change(nameInput, { target: { value: '山田太郎' } })
        fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
        fireEvent.change(contentInput, { target: { value: 'テストお問い合わせ内容' } })

        // 送信ボタンを有効化するまで待機
        await waitFor(() => {
            expect(screen.getByRole('button', { name: '送信する' })).not.toBeDisabled()
        })

        // フォームを送信
        fireEvent.click(screen.getByRole('button', { name: '送信する' }))

        // APIが正しい値で呼び出されることを確認（任意フィールドは空）
        await waitFor(() => {
            expect(mockPostContact).toHaveBeenCalledWith({
                form: {
                    name: '山田太郎',
                    company: '',
                    phoneNumber: '',
                    email: 'test@example.com',
                    content: 'テストお問い合わせ内容',
                },
            })
        })

        // 成功メッセージが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText(/お問い合わせを送信しました/)).toBeInTheDocument()
        })
    })
})
