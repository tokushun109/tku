import { http, HttpResponse } from 'msw'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { getApiBaseUrl } from '@/apis/baseUrl'
import { getCarouselImages, getProduct, getProducts, getProductsByCategory } from '@/apis/product'
import { ApiError } from '@/utils/error'

import { server } from '../mocks/server'

const apiBaseUrl = getApiBaseUrl()

describe('product API', () => {
    beforeEach(() => {
        server.resetHandlers()
    })

    afterEach(() => {
        vi.restoreAllMocks()
    })

    describe('getProductsByCategory', () => {
        it('正常にカテゴリごとの商品リストを取得する', async () => {
            const result = await getProductsByCategory({
                category: 'all',
                limit: 4,
                target: 'all',
            })

            expect(result).toHaveLength(3) // 3つのカテゴリが存在
            expect(result.find((cat) => cat.category.name === 'イヤリング')).toBeDefined()
            expect(result.find((cat) => cat.category.name === 'リング')).toBeDefined()
            expect(result.find((cat) => cat.category.name === 'ネックレス')).toBeDefined()

            // 公開画面ではアクティブ商品のみ返る
            const earringsCategory = result.find((cat) => cat.category.name === 'イヤリング')
            expect(earringsCategory?.products).toHaveLength(1)
            expect(earringsCategory?.pageInfo).toEqual({
                hasMore: false,
                nextCursor: '',
            })
        })

        it('カテゴリとターゲットで絞り込んだ結果を取得できる', async () => {
            const result = await getProductsByCategory({
                category: 'rings-uuid',
                limit: 4,
                target: 'men-uuid',
            })

            expect(result).toEqual([
                {
                    category: {
                        name: 'リング',
                        uuid: 'rings-uuid',
                    },
                    pageInfo: {
                        hasMore: false,
                        nextCursor: '',
                    },
                    products: [
                        expect.objectContaining({
                            name: '男性向けリング1',
                            uuid: 'rings-men-1',
                        }),
                    ],
                },
            ])
        })

        it('APIエラーの場合、ApiErrorが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/category/product`, () => {
                    return new HttpResponse(null, { status: 500 })
                }),
            )

            await expect(
                getProductsByCategory({
                    category: 'all',
                    limit: 4,
                    target: 'all',
                }),
            ).rejects.toThrow(ApiError)
        })

        it('ネットワークエラーの場合、一般的なエラーが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/category/product`, () => {
                    return HttpResponse.error()
                }),
            )

            await expect(
                getProductsByCategory({
                    category: 'all',
                    limit: 4,
                    target: 'all',
                }),
            ).rejects.toThrow('カテゴリーごとの商品リストの取得に失敗しました')
        })

        it('正しいクエリパラメータでリクエストされる', async () => {
            let capturedRequest: Request | undefined

            server.use(
                http.get(`${apiBaseUrl}/category/product`, ({ request }) => {
                    capturedRequest = request
                    return HttpResponse.json([])
                }),
            )

            await getProductsByCategory({
                category: 'earrings-uuid',
                cursor: 'next-cursor',
                limit: 8,
                target: 'women-uuid',
            })

            expect(capturedRequest).toBeDefined()
            const url = new URL(capturedRequest!.url)
            expect(url.searchParams.get('category')).toBe('earrings-uuid')
            expect(url.searchParams.get('cursor')).toBe('next-cursor')
            expect(url.searchParams.get('limit')).toBe('8')
            expect(url.searchParams.get('mode')).toBeNull()
            expect(url.searchParams.get('target')).toBe('women-uuid')
        })

        it('cursor未指定時はクエリパラメータに含めない', async () => {
            let capturedRequest: Request | undefined

            server.use(
                http.get(`${apiBaseUrl}/category/product`, ({ request }) => {
                    capturedRequest = request
                    return HttpResponse.json([])
                }),
            )

            await getProductsByCategory({
                category: 'all',
                limit: 4,
                target: 'all',
            })

            expect(capturedRequest).toBeDefined()
            const url = new URL(capturedRequest!.url)
            expect(url.searchParams.has('cursor')).toBe(false)
        })
    })

    describe('getProduct', () => {
        it('正常に商品詳細を取得する（女性向けイヤリング1）', async () => {
            const result = await getProduct('earrings-women-1')

            expect(result).toEqual({
                uuid: 'earrings-women-1',
                name: '女性向けイヤリング1',
                price: 1500,
                description: '女性向けイヤリング1の詳細',
                isActive: true,
                isRecommend: true,
                category: {
                    name: 'イヤリング',
                    uuid: 'earrings-uuid',
                },
                target: {
                    name: '女性',
                    uuid: 'women-uuid',
                },
                tags: [
                    {
                        name: 'イヤリング',
                        uuid: 'earrings-uuid',
                    },
                ],
                productImages: [
                    {
                        apiPath: '/image/earrings-women-1.jpg',
                        name: 'earrings-women-1.jpg',
                        displayOrder: 1,
                        uuid: 'earrings-women-1-image-uuid',
                    },
                ],
                siteDetails: [],
            })
        })

        it('正常に商品詳細を取得する（非アクティブ商品）', async () => {
            const result = await getProduct('earrings-women-2')

            expect(result.uuid).toBe('earrings-women-2')
            expect(result.name).toBe('女性向けイヤリング2')
            expect(result.isActive).toBe(false)
            expect(result.isRecommend).toBe(false)
        })

        it('正常に商品詳細を取得する（男性向けリング）', async () => {
            const result = await getProduct('rings-men-1')

            expect(result.uuid).toBe('rings-men-1')
            expect(result.name).toBe('男性向けリング1')
            expect(result.category.name).toBe('リング')
            expect(result.target.name).toBe('男性')
        })

        it('正常に商品詳細を取得する（ユニセックスネックレス）', async () => {
            const result = await getProduct('necklaces-unisex-1')

            expect(result.uuid).toBe('necklaces-unisex-1')
            expect(result.name).toBe('ユニセックスネックレス1')
            expect(result.category.name).toBe('ネックレス')
            expect(result.target.name).toBe('ユニセックス')
        })

        it('存在しないUUIDの場合、ApiErrorが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/product/:uuid`, () => {
                    return new HttpResponse(null, { status: 404 })
                }),
            )

            await expect(getProduct('non-existent-uuid')).rejects.toThrow(ApiError)
        })

        it('APIエラーの場合、ApiErrorが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/product/:uuid`, () => {
                    return new HttpResponse(null, { status: 500 })
                }),
            )

            await expect(getProduct('test-uuid-1')).rejects.toThrow(ApiError)
        })

        it('ネットワークエラーの場合、一般的なエラーが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/product/:uuid`, () => {
                    return HttpResponse.error()
                }),
            )

            await expect(getProduct('test-uuid-1')).rejects.toThrow('商品詳細の取得に失敗しました')
        })

        it('正しいパスパラメータでリクエストされる', async () => {
            let capturedUuid: string | undefined

            server.use(
                http.get(`${apiBaseUrl}/product/:uuid`, ({ params }) => {
                    capturedUuid = params.uuid as string
                    return HttpResponse.json({
                        uuid: params.uuid,
                        name: 'テスト商品',
                        price: 1000,
                        description: 'テスト詳細',
                        isActive: true,
                        isRecommend: false,
                        category: {
                            name: 'テストカテゴリ',
                            uuid: 'test-category-uuid',
                        },
                        target: {
                            name: 'テスト対象',
                            uuid: 'test-target-uuid',
                        },
                        tags: [],
                        productImages: [],
                        siteDetails: [],
                    })
                }),
            )

            await getProduct('specific-uuid-123')

            expect(capturedUuid).toBe('specific-uuid-123')
        })
    })

    describe('getProducts', () => {
        it('管理画面用一覧取得はCookie付きでリクエストする', async () => {
            const fetchSpy = vi.spyOn(globalThis, 'fetch').mockResolvedValue(
                new Response(JSON.stringify([]), {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    status: 200,
                }),
            )

            const result = await getProducts({
                category: 'all',
                mode: 'all',
                target: 'all',
            })

            const [requestUrl, requestInit] = fetchSpy.mock.calls[0]
            const url = new URL(requestUrl as string)

            expect(url.searchParams.get('category')).toBe('all')
            expect(url.searchParams.get('mode')).toBe('all')
            expect(url.searchParams.get('target')).toBe('all')
            expect(requestInit?.credentials).toBe('include')
            expect(result).toEqual([])
        })
    })

    describe('getCarouselImages', () => {
        it('正常にカルーセル画像を取得する', async () => {
            const result = await getCarouselImages()

            expect(result).toHaveLength(3)

            // 1番目のカルーセル画像
            expect(result[0].apiPath).toBe('/image/carousel1.jpg')
            expect(result[0].product.uuid).toBe('earrings-women-1')
            expect(result[0].product.name).toBe('女性向けイヤリング1')

            // 2番目のカルーセル画像
            expect(result[1].apiPath).toBe('/image/carousel2.jpg')
            expect(result[1].product.uuid).toBe('earrings-women-2')
            expect(result[1].product.name).toBe('女性向けイヤリング2')

            // 3番目のカルーセル画像
            expect(result[2].apiPath).toBe('/image/carousel3.jpg')
            expect(result[2].product.uuid).toBe('rings-men-1')
            expect(result[2].product.name).toBe('男性向けリング1')
        })

        it('各カルーセル画像に商品情報が含まれる', async () => {
            const result = await getCarouselImages()

            result.forEach((thumbnail) => {
                expect(thumbnail.product).toBeDefined()
                expect(thumbnail.product.uuid).toBeDefined()
                expect(thumbnail.product.name).toBeDefined()
                expect(thumbnail.product.category).toBeDefined()
                expect(thumbnail.product.target).toBeDefined()
                expect(thumbnail.product.productImages).toBeDefined()
            })
        })

        it('異なるカテゴリーの商品が含まれる', async () => {
            const result = await getCarouselImages()

            const categories = result.map((thumbnail) => thumbnail.product.category.name)
            expect(categories).toContain('イヤリング')
            expect(categories).toContain('リング')
        })

        it('APIエラーの場合、ApiErrorが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/carousel_image/`, () => {
                    return new HttpResponse(null, { status: 500 })
                }),
            )

            await expect(getCarouselImages()).rejects.toThrow(ApiError)
        })

        it('ネットワークエラーの場合、一般的なエラーが投げられる', async () => {
            server.use(
                http.get(`${apiBaseUrl}/carousel_image/`, () => {
                    return HttpResponse.error()
                }),
            )

            await expect(getCarouselImages()).rejects.toThrow('カルーセル画像の取得に失敗しました')
        })

        it('正しいエンドポイントでリクエストされる', async () => {
            let capturedPath: string | undefined

            server.use(
                http.get(`${apiBaseUrl}/carousel_image/`, ({ request }) => {
                    capturedPath = new URL(request.url).pathname
                    return HttpResponse.json([])
                }),
            )

            await getCarouselImages()

            expect(capturedPath).toBe(new URL(`${apiBaseUrl}/carousel_image/`).pathname)
        })
    })
})
