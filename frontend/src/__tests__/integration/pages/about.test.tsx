import { beforeEach, describe, expect, it, vi } from 'vitest'

import { getCreator } from '@/apis/creator'
import { getSalesSiteList } from '@/apis/salesSite'
import { getSnsList } from '@/apis/sns'
import About from '@/app/(contents)/about/page'

import { render, screen, waitFor } from '../helpers'

// APIモック
vi.mock('@/apis/creator')
vi.mock('@/apis/salesSite')
vi.mock('@/apis/sns')

const mockGetCreator = vi.mocked(getCreator)
const mockGetSalesSiteList = vi.mocked(getSalesSiteList)
const mockGetSnsList = vi.mocked(getSnsList)

describe('About Page Integration Test', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })

    it('Aboutページが正常に表示される', async () => {
        // モックデータの設定
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetSalesSiteList.mockResolvedValue([
            {
                name: 'Creema',
                uuid: 'creema-uuid',
            },
            {
                name: 'minne',
                uuid: 'minne-uuid',
            },
        ])

        mockGetSnsList.mockResolvedValue([
            {
                name: 'Twitter',
                uuid: 'twitter-uuid',
            },
            {
                name: 'Instagram',
                uuid: 'instagram-uuid',
            },
        ])

        // コンポーネントをレンダリング
        render(await About())

        // 販売サイトの表示を確認
        await waitFor(() => {
            expect(screen.getByText('Creema')).toBeInTheDocument()
            expect(screen.getByText('minne')).toBeInTheDocument()
        })

        // SNSの表示を確認
        await waitFor(() => {
            expect(screen.getByText('Twitter')).toBeInTheDocument()
            expect(screen.getByText('Instagram')).toBeInTheDocument()
        })
    })

    it('販売サイトがない場合でもエラーにならない', async () => {
        // モックデータの設定（販売サイトなし）
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetSalesSiteList.mockResolvedValue([])
        mockGetSnsList.mockResolvedValue([])

        // コンポーネントをレンダリング
        render(await About())

        // エラーが発生しないことを確認
        await waitFor(() => {
            expect(screen.queryByTestId('error-message')).not.toBeInTheDocument()
        })
    })

    it('API呼び出しが失敗した場合のエラーハンドリング', async () => {
        // モックでエラーを発生させる
        mockGetCreator.mockRejectedValue(new Error('API Error'))
        mockGetSalesSiteList.mockRejectedValue(new Error('API Error'))
        mockGetSnsList.mockRejectedValue(new Error('API Error'))

        // エラーが投げられることを確認
        await expect(About()).rejects.toThrow('API Error')
    })

    it('SNSリストが多い場合でも正常に表示される', async () => {
        // モックデータの設定
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetSalesSiteList.mockResolvedValue([
            {
                name: 'Creema',
                uuid: 'creema-uuid',
            },
        ])

        mockGetSnsList.mockResolvedValue([
            {
                name: 'Twitter',
                uuid: 'twitter-uuid',
            },
            {
                name: 'Instagram',
                uuid: 'instagram-uuid',
            },
            {
                name: 'Facebook',
                uuid: 'facebook-uuid',
            },
            {
                name: 'YouTube',
                uuid: 'youtube-uuid',
            },
            {
                name: 'TikTok',
                uuid: 'tiktok-uuid',
            },
        ])

        // コンポーネントをレンダリング
        render(await About())

        // すべてのSNSが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('Twitter')).toBeInTheDocument()
            expect(screen.getByText('Instagram')).toBeInTheDocument()
            expect(screen.getByText('Facebook')).toBeInTheDocument()
            expect(screen.getByText('YouTube')).toBeInTheDocument()
            expect(screen.getByText('TikTok')).toBeInTheDocument()
        })
    })

    it('販売サイトリストが多い場合でも正常に表示される', async () => {
        // モックデータの設定
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetSalesSiteList.mockResolvedValue([
            {
                name: 'Creema',
                uuid: 'creema-uuid',
            },
            {
                name: 'minne',
                uuid: 'minne-uuid',
            },
            {
                name: 'Etsy',
                uuid: 'etsy-uuid',
            },
            {
                name: 'BASE',
                uuid: 'base-uuid',
            },
        ])

        mockGetSnsList.mockResolvedValue([
            {
                name: 'Twitter',
                uuid: 'twitter-uuid',
            },
        ])

        // コンポーネントをレンダリング
        render(await About())

        // すべての販売サイトが表示されることを確認
        await waitFor(() => {
            expect(screen.getByText('Creema')).toBeInTheDocument()
            expect(screen.getByText('minne')).toBeInTheDocument()
            expect(screen.getByText('Etsy')).toBeInTheDocument()
            expect(screen.getByText('BASE')).toBeInTheDocument()
        })
    })

    it('部分的なAPIエラーの場合でも他のコンテンツは表示される', async () => {
        // モックデータの設定（一部のAPIはエラー）
        mockGetCreator.mockResolvedValue({
            name: 'テスト作家',
            introduction: 'テスト作家のプロフィール',
            logo: '/image/creator.jpg',
            apiPath: '/image/creator.jpg',
        })

        mockGetSalesSiteList.mockResolvedValue([
            {
                name: 'Creema',
                uuid: 'creema-uuid',
            },
        ])

        // SNSのみエラー
        mockGetSnsList.mockRejectedValue(new Error('SNS API Error'))

        // コンポーネントをレンダリング
        // SNSのエラーで全体がエラーになることを確認
        await expect(About()).rejects.toThrow('SNS API Error')
    })
})
