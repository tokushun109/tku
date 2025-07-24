import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCategories } from '@/apis/category'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'
import ClassificationPage from '@/app/admin/classification/page'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/category')
vi.mock('@/apis/tag')
vi.mock('@/apis/target')

const mockGetCategories = vi.mocked(getCategories)
const mockGetTags = vi.mocked(getTags)
const mockGetTargets = vi.mocked(getTargets)

describe('Admin Classification Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('管理画面分類ページが正常に表示される', async () => {
        // モックデータの設定
        mockGetCategories.mockResolvedValue([
            {
                uuid: 'category-1',
                name: 'カテゴリー1',
            },
            {
                uuid: 'category-2',
                name: 'カテゴリー2',
            },
        ])

        mockGetTags.mockResolvedValue([
            {
                uuid: 'tag-1',
                name: 'タグ1',
            },
            {
                uuid: 'tag-2',
                name: 'タグ2',
            },
        ])

        mockGetTargets.mockResolvedValue([
            {
                uuid: 'target-1',
                name: 'ターゲット1',
            },
            {
                uuid: 'target-2',
                name: 'ターゲット2',
            },
        ])

        // コンポーネントをレンダリング
        render(await ClassificationPage())

        // 分類管理画面の表示を確認
        await waitFor(() => {
            expect(screen.getByText('カテゴリー')).toBeInTheDocument()
            expect(screen.getByText('ターゲット')).toBeInTheDocument()
            expect(screen.getByText('タグ')).toBeInTheDocument()
        })

        // API呼び出しの確認
        expect(mockGetCategories).toHaveBeenCalledWith({ mode: 'all' })
        expect(mockGetTags).toHaveBeenCalledWith()
        expect(mockGetTargets).toHaveBeenCalledWith({ mode: 'all' })
    })

    it('APIエラー時のフォールバック処理', async () => {
        // コンソールエラーをモック
        const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        // モックでエラーを発生させる
        mockGetCategories.mockRejectedValue(new Error('Category API Error'))
        mockGetTags.mockRejectedValue(new Error('Tag API Error'))
        mockGetTargets.mockRejectedValue(new Error('Target API Error'))

        // コンポーネントをレンダリング
        render(await ClassificationPage())

        // エラーログが出力されることを確認
        await waitFor(() => {
            expect(consoleSpy).toHaveBeenCalledWith('データの取得に失敗しました:', expect.any(Error))
        })

        // 空配列でテンプレートが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('カテゴリー')).toBeInTheDocument()
        })

        consoleSpy.mockRestore()
    })

    it('一部のAPIが失敗した場合の処理', async () => {
        // コンソールエラーをモック
        const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})

        // カテゴリーAPIのみエラー、他は正常
        mockGetCategories.mockRejectedValue(new Error('Category API Error'))
        mockGetTags.mockResolvedValue([
            {
                uuid: 'tag-1',
                name: 'タグ1',
            },
        ])
        mockGetTargets.mockResolvedValue([
            {
                uuid: 'target-1',
                name: 'ターゲット1',
            },
        ])

        // コンポーネントをレンダリング
        render(await ClassificationPage())

        // エラーログが出力されることを確認
        await waitFor(() => {
            expect(consoleSpy).toHaveBeenCalledWith('データの取得に失敗しました:', expect.any(Error))
        })

        // フォールバック処理で空配列が渡されることを確認
        await waitFor(() => {
            expect(screen.getByText('カテゴリー')).toBeInTheDocument()
        })

        consoleSpy.mockRestore()
    })

    it('すべてのAPIが正常に呼び出される', async () => {
        // モックデータの設定
        mockGetCategories.mockResolvedValue([])
        mockGetTags.mockResolvedValue([])
        mockGetTargets.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await ClassificationPage())

        // 並列でAPIが呼び出されることを確認
        expect(mockGetCategories).toHaveBeenCalledWith({ mode: 'all' })
        expect(mockGetTags).toHaveBeenCalledWith()
        expect(mockGetTargets).toHaveBeenCalledWith({ mode: 'all' })

        // すべてのAPIが一度だけ呼ばれることを確認
        expect(mockGetCategories).toHaveBeenCalledTimes(1)
        expect(mockGetTags).toHaveBeenCalledTimes(1)
        expect(mockGetTargets).toHaveBeenCalledTimes(1)
    })

    it('空データの場合でも正常に表示される', async () => {
        // モックデータの設定（空配列）
        mockGetCategories.mockResolvedValue([])
        mockGetTags.mockResolvedValue([])
        mockGetTargets.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await ClassificationPage())

        // テンプレートが正常に表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('カテゴリー')).toBeInTheDocument()
        })
    })
})
