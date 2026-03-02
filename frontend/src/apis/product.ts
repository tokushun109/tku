import { getApiBaseUrl } from '@/apis/baseUrl'
import { IProduct, IProductsByCategory, IThumbnail } from '@/features/product/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetProductsByCategoryParams {
    category: 'all' | string
    target: 'all' | string
}

export interface IGetProductsParams {
    category: 'all' | string
    mode: 'all' | 'active'
    target: 'all' | string
}

/** カテゴリーごとの商品リストを取得 */
export const getProductsByCategory = async (params: IGetProductsByCategoryParams): Promise<IProductsByCategory[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${getApiBaseUrl()}/category/product?${query}`, {
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
        const res = await fetch(`${getApiBaseUrl()}/product/${uuid}`, {
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
        const res = await fetch(`${getApiBaseUrl()}/carousel_image/`, {
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

export interface IProductImageDisplayOrderParams {
    displayOrder: { [key: number]: number }
    isChanged: boolean
}

/** 商品リストを取得（管理画面用） */
export const getProducts = async (params: IGetProductsParams): Promise<IProduct[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${getApiBaseUrl()}/product?${query}`, {
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
        throw new Error('商品リストの取得に失敗しました')
    }
}

/** 商品を作成 */
export const createProduct = async (product: Omit<IProduct, 'uuid'>): Promise<{ uuid: string }> => {
    try {
        const res = await fetch(`${getApiBaseUrl()}/product`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
            body: JSON.stringify(product),
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('商品の作成に失敗しました')
    }
}

/** 商品を更新 */
export const updateProduct = async (uuid: string, product: IProduct): Promise<void> => {
    try {
        const res = await fetch(`${getApiBaseUrl()}/product/${uuid}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'PUT',
            body: JSON.stringify(product),
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('商品の更新に失敗しました')
    }
}

/** 商品を削除 */
export const deleteProduct = async (uuid: string): Promise<void> => {
    try {
        const res = await fetch(`${getApiBaseUrl()}/product/${uuid}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'DELETE',
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('商品の削除に失敗しました')
    }
}

/** 商品画像をアップロード */
export const uploadProductImages = async (productUuid: string, files: File[], displayOrderParams: IProductImageDisplayOrderParams): Promise<void> => {
    try {
        const formData = new FormData()
        formData.append('displayOrder', JSON.stringify(displayOrderParams))

        files.forEach((file, index) => {
            formData.append(`file${index}`, file)
        })

        const res = await fetch(`${getApiBaseUrl()}/product/${productUuid}/product_image`, {
            method: 'POST',
            body: formData,
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('商品画像のアップロードに失敗しました')
    }
}

export interface ICreemaDuplicateParams {
    url: string
}

/** Creemaから商品を複製 */
export const duplicateProductFromCreema = async (params: ICreemaDuplicateParams): Promise<void> => {
    try {
        const res = await fetch(`${getApiBaseUrl()}/product/duplicate`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
            body: JSON.stringify(params),
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('Creemaからの商品複製に失敗しました')
    }
}
