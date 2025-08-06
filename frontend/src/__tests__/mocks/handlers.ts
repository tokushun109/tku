import { http, HttpResponse } from 'msw'

import { ICreator } from '@/features/creator/type'
import { IProduct, IProductsByCategory, IThumbnail } from '@/features/product/type'
import { ISite } from '@/features/site/type'

// モックデータの定義
const mockCategories = {
    earrings: {
        name: 'イヤリング',
        uuid: 'earrings-uuid',
    },
    rings: {
        name: 'リング',
        uuid: 'rings-uuid',
    },
    necklaces: {
        name: 'ネックレス',
        uuid: 'necklaces-uuid',
    },
}

const mockTargets = {
    women: {
        name: '女性',
        uuid: 'women-uuid',
    },
    men: {
        name: '男性',
        uuid: 'men-uuid',
    },
    unisex: {
        name: 'ユニセックス',
        uuid: 'unisex-uuid',
    },
}

const mockProducts: IProduct[] = [
    {
        uuid: 'earrings-women-1',
        name: '女性向けイヤリング1',
        price: 1500,
        description: '女性向けイヤリング1の詳細',
        isActive: true,
        isRecommend: true,
        category: mockCategories.earrings,
        target: mockTargets.women,
        tags: [mockCategories.earrings],
        productImages: [
            {
                apiPath: '/image/earrings-women-1.jpg',
                name: 'earrings-women-1.jpg',
                order: 1,
                uuid: 'earrings-women-1-image-uuid',
            },
        ],
        siteDetails: [],
    },
    {
        uuid: 'earrings-women-2',
        name: '女性向けイヤリング2',
        price: 2000,
        description: '女性向けイヤリング2の詳細',
        isActive: false,
        isRecommend: false,
        category: mockCategories.earrings,
        target: mockTargets.women,
        tags: [mockCategories.earrings],
        productImages: [
            {
                apiPath: '/image/earrings-women-2.jpg',
                name: 'earrings-women-2.jpg',
                order: 1,
                uuid: 'earrings-women-2-image-uuid',
            },
        ],
        siteDetails: [],
    },
    {
        uuid: 'rings-men-1',
        name: '男性向けリング1',
        price: 3000,
        description: '男性向けリング1の詳細',
        isActive: true,
        isRecommend: false,
        category: mockCategories.rings,
        target: mockTargets.men,
        tags: [mockCategories.rings],
        productImages: [
            {
                apiPath: '/image/rings-men-1.jpg',
                name: 'rings-men-1.jpg',
                order: 1,
                uuid: 'rings-men-1-image-uuid',
            },
        ],
        siteDetails: [],
    },
    {
        uuid: 'necklaces-unisex-1',
        name: 'ユニセックスネックレス1',
        price: 2500,
        description: 'ユニセックスネックレス1の詳細',
        isActive: true,
        isRecommend: true,
        category: mockCategories.necklaces,
        target: mockTargets.unisex,
        tags: [mockCategories.necklaces],
        productImages: [
            {
                apiPath: '/image/necklaces-unisex-1.jpg',
                name: 'necklaces-unisex-1.jpg',
                order: 1,
                uuid: 'necklaces-unisex-1-image-uuid',
            },
        ],
        siteDetails: [],
    },
]

// 固定のテストデータ（パラメータに関わらず同じデータを返す）
const mockProductsByCategory: IProductsByCategory[] = [
    {
        category: mockCategories.earrings,
        products: [mockProducts[0], mockProducts[1]],
    },
    {
        category: mockCategories.rings,
        products: [mockProducts[2]],
    },
    {
        category: mockCategories.necklaces,
        products: [mockProducts[3]],
    },
]

const mockThumbnails: IThumbnail[] = [
    {
        apiPath: '/image/carousel1.jpg',
        product: mockProducts[0],
    },
    {
        apiPath: '/image/carousel2.jpg',
        product: mockProducts[1],
    },
    {
        apiPath: '/image/carousel3.jpg',
        product: mockProducts[2],
    },
]

const mockCreator: ICreator = {
    name: 'テスト作家',
    introduction: 'テスト作家のプロフィール',
    logo: '/image/creator.jpg',
    apiPath: '/image/creator.jpg',
}

const mockTags = [
    { uuid: 'tag-1', name: 'シンプル' },
    { uuid: 'tag-2', name: 'カジュアル' },
    { uuid: 'tag-3', name: 'エレガント' },
]

const mockSalesSites = [
    { uuid: 'site-1', name: 'Creema' },
    { uuid: 'site-2', name: 'minne' },
]

const mockSalesTarget: ISite = {
    name: 'テスト対象',
    uuid: 'test-target-uuid',
}

// APIハンドラーの定義
export const handlers = [
    // 商品関連のAPI
    http.get('http://localhost:8080/category/product', () => {
        return HttpResponse.json(mockProductsByCategory)
    }),

    // 全商品一覧取得API（管理画面用）
    http.get('http://localhost:8080/product', ({ request }) => {
        const url = new URL(request.url)
        const mode = url.searchParams.get('mode')
        const category = url.searchParams.get('category')
        const target = url.searchParams.get('target')

        // 管理画面用（mode=all）の場合、全商品を返す
        if (mode === 'all') {
            let filteredProducts = [...mockProducts]

            // カテゴリフィルタ
            if (category && category !== 'all') {
                filteredProducts = filteredProducts.filter((p) => p.category.uuid === category)
            }

            // ターゲットフィルタ
            if (target && target !== 'all') {
                filteredProducts = filteredProducts.filter((p) => p.target.uuid === target)
            }

            return HttpResponse.json(filteredProducts)
        }

        // 通常の商品一覧（公開中のみ）
        const activeProducts = mockProducts.filter((p) => p.isActive)
        return HttpResponse.json(activeProducts)
    }),

    http.get('http://localhost:8080/product/:uuid', ({ params }) => {
        const { uuid } = params
        const product = mockProducts.find((p) => p.uuid === uuid)

        if (!product) {
            return new HttpResponse(null, { status: 404 })
        }

        return HttpResponse.json(product)
    }),

    http.get('http://localhost:8080/carousel_image/', () => {
        return HttpResponse.json(mockThumbnails)
    }),

    // カテゴリ関連のAPI
    http.get('http://localhost:8080/category', () => {
        return HttpResponse.json(Object.values(mockCategories))
    }),

    // ターゲット関連のAPI
    http.get('http://localhost:8080/target', () => {
        return HttpResponse.json(Object.values(mockTargets))
    }),

    // タグ関連のAPI
    http.get('http://localhost:8080/tag', () => {
        return HttpResponse.json(mockTags)
    }),

    // 販売サイト関連のAPI
    http.get('http://localhost:8080/sales_site/', () => {
        return HttpResponse.json(mockSalesSites)
    }),

    // その他のAPI
    http.get('http://localhost:8080/creator', () => {
        return HttpResponse.json(mockCreator)
    }),

    http.get('http://localhost:8080/sales_target', () => {
        return HttpResponse.json([mockSalesTarget])
    }),

    http.get('http://localhost:8080/health', () => {
        return HttpResponse.json({ status: 'ok' })
    }),

    // お問い合わせAPI
    http.post('http://localhost:8080/contact', async ({ request }) => {
        const body = (await request.json()) as any

        // バリデーション（簡単なチェック）
        if (!body?.name || !body?.email || !body?.content) {
            return new HttpResponse(null, { status: 400 })
        }

        return HttpResponse.json({ message: 'お問い合わせを受け付けました' })
    }),

    // エラーケースのテスト用
    http.get('http://localhost:8080/error/500', () => {
        return new HttpResponse(null, { status: 500 })
    }),

    http.get('http://localhost:8080/error/404', () => {
        return new HttpResponse(null, { status: 404 })
    }),

    // ネットワークエラーのテスト用
    http.get('http://localhost:8080/error/network', () => {
        return HttpResponse.error()
    }),

    // CSV関連のAPI
    http.get('http://localhost:8080/csv/product', () => {
        const csvContent = `name,price,description,isActive,isRecommend
女性向けイヤリング1,1500,女性向けイヤリング1の詳細,true,true
女性向けイヤリング2,2000,女性向けイヤリング2の詳細,false,false
男性向けリング1,3000,男性向けリング1の詳細,true,false
ユニセックスネックレス1,2500,ユニセックスネックレス1の詳細,true,true`

        return new HttpResponse(csvContent, {
            headers: {
                'Content-Type': 'text/csv',
                'Content-Disposition': 'attachment; filename="商品レコード.csv"',
            },
        })
    }),

    http.post('http://localhost:8080/csv/product', async ({ request }) => {
        const formData = await request.formData()
        const csvFile = formData.get('csv') as File

        // バリデーション（ファイルが存在するかチェック）
        if (!csvFile) {
            return new HttpResponse(null, { status: 400 })
        }

        // ファイルの内容を読み取り（簡単なバリデーション）
        const csvContent = await csvFile.text()
        if (!csvContent.includes('name') || !csvContent.includes('price')) {
            return new HttpResponse(null, { status: 400 })
        }

        return HttpResponse.json({ message: 'CSVアップロードが完了しました' })
    }),
]
