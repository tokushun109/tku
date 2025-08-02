import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCreator } from '@/apis/creator'
import AdminSeoPage, { metadata } from '@/app/admin/seo/page'
import { ICreator } from '@/features/creator/type'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/creator')

const mockGetCreator = vi.mocked(getCreator)

// テスト用のモックデータ
const createMockCreator = (overrides: Partial<ICreator> = {}): ICreator => ({
    apiPath: '/api/images/creator-logo.jpg',
    introduction: 'ハンドメイドアクセサリーの制作を行っています。\n一つ一つ丁寧に手作りしています。',
    logo: 'creator-logo.jpg',
    name: 'tocoriri',
    ...overrides,
})

describe('Admin SEO Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('管理画面SEO設定ページが正常に表示される', async () => {
        // モックデータの設定
        const mockCreator = createMockCreator()

        mockGetCreator.mockResolvedValue(mockCreator)

        // コンポーネントをレンダリング
        render(await AdminSeoPage())

        // ページタイトルの表示を確認
        await waitFor(() => {
            expect(screen.getByText('SEO設定')).toBeInTheDocument()
        })

        // SEO編集コンポーネントが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('サイトロゴ')).toBeInTheDocument()
            expect(screen.getByText('サイト説明')).toBeInTheDocument()
        })

        // 作者情報が表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('tocoriri')).toBeInTheDocument()
            expect(screen.getByText('ハンドメイドアクセサリーの制作を行っています。\n一つ一つ丁寧に手作りしています。')).toBeInTheDocument()
        })

        // 編集ボタンが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('編集')).toBeInTheDocument()
        })

        // API呼び出しの確認
        expect(mockGetCreator).toHaveBeenCalledTimes(1)
    })

    it('作者情報がない場合でもエラーメッセージが表示される', async () => {
        // モックでnullを返す
        mockGetCreator.mockResolvedValue(null as unknown as ICreator)

        // コンポーネントをレンダリング
        render(await AdminSeoPage())

        // ページタイトルの表示を確認
        await waitFor(() => {
            expect(screen.getByText('SEO設定')).toBeInTheDocument()
        })

        // エラーメッセージの表示を確認
        await waitFor(() => {
            expect(screen.getByText('作者情報の読み込みに失敗しました')).toBeInTheDocument()
        })
    })

    it('API呼び出しが失敗した場合のエラーハンドリング', async () => {
        // コンソールエラーをモック
        const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        // モックでエラーを発生させる
        mockGetCreator.mockRejectedValue(new Error('Creator API Error'))

        // コンポーネントをレンダリング
        render(await AdminSeoPage())

        // エラーログが出力されることを確認
        await waitFor(() => {
            expect(consoleSpy).toHaveBeenCalledWith('作者情報の取得に失敗しました:', expect.any(Error))
        })

        // エラーメッセージでテンプレートが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('SEO設定')).toBeInTheDocument()
            expect(screen.getByText('作者情報の読み込みに失敗しました')).toBeInTheDocument()
        })

        consoleSpy.mockRestore()
    })

    it('ロゴ画像がない場合でも正常に表示される', async () => {
        // ロゴ画像がないモックデータ
        const mockCreator = createMockCreator({
            apiPath: '',
            logo: '',
        })

        mockGetCreator.mockResolvedValue(mockCreator)

        // コンポーネントをレンダリング
        render(await AdminSeoPage())

        // ページタイトルの表示を確認
        await waitFor(() => {
            expect(screen.getByText('SEO設定')).toBeInTheDocument()
        })

        // セクションタイトルが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('サイトロゴ')).toBeInTheDocument()
            expect(screen.getByText('サイト説明')).toBeInTheDocument()
        })

        // ロゴ変更ボタンが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('ロゴを変更')).toBeInTheDocument()
        })
    })

    it('長い紹介文でも正常に表示される', async () => {
        // 長い紹介文を持つモックデータ
        const longIntroduction = Array.from({ length: 10 }, (_, i) => `これは${i + 1}行目の長い紹介文です。`).join('\n')
        const mockCreator = createMockCreator({
            introduction: longIntroduction,
        })

        mockGetCreator.mockResolvedValue(mockCreator)

        // コンポーネントをレンダリング
        render(await AdminSeoPage())

        // ページタイトルの表示を確認
        await waitFor(() => {
            expect(screen.getByText('SEO設定')).toBeInTheDocument()
        })

        // 長い紹介文の一部が表示されることを確認
        await waitFor(() => {
            expect(screen.getByText(longIntroduction)).toBeInTheDocument()
        })
    })

    it('API呼び出しが正常に実行される', async () => {
        // モックデータの設定
        const mockCreator = createMockCreator()

        mockGetCreator.mockResolvedValue(mockCreator)

        // コンポーネントをレンダリング
        render(await AdminSeoPage())

        // API呼び出しの確認
        expect(mockGetCreator).toHaveBeenCalledTimes(1)

        // 基本的な表示確認
        await waitFor(() => {
            expect(screen.getByText('SEO設定')).toBeInTheDocument()
            expect(screen.getByText('サイトロゴ')).toBeInTheDocument()
            expect(screen.getByText('サイト説明')).toBeInTheDocument()
        })
    })

    it('ページメタデータでnoindex, nofollowが設定されている', () => {
        // メタデータの確認
        expect(metadata.title).toBe('SEO設定 | admin')
        expect(metadata.robots).toEqual({
            index: false,
            follow: false,
        })
    })
})
