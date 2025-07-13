import { IProduct, IProductsByCategory, IThumbnail } from '@/features/product/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetProductsByCategoryParams {
    category: 'all' | string
    mode: 'all' | 'active'
    target: 'all' | string
}

/** カテゴリーごとの商品リストを取得 */
export const getProductsByCategory = async (params: IGetProductsByCategoryParams): Promise<IProductsByCategory[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${process.env.API_URL}/category/product?${query}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('カテゴリーごとの商品リストの取得に失敗しました')
    }
}

/** 商品詳細を取得 */
export const getProduct = async (uuid: string): Promise<IProduct> => {
    try {
        const res = await fetch(`${process.env.API_URL}/product/${uuid}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('商品詳細の取得に失敗しました')
    }
}

export const getCarouselImages = async (): Promise<IThumbnail[]> => {
    try {
        const res = await fetch(`${process.env.API_URL}/carousel_image/`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('カルーセル画像の取得に失敗しました')
    }
}
