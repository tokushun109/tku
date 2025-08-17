import { IProduct, IProductsByCategory, IThumbnail } from '@/features/product/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetProductsParams {
    category: 'all' | string
    mode: 'all' | 'active'
    target: 'all' | string
}

/** カテゴリーごとの商品リストを取得 */
export const getProductsByCategory = async (params: IGetProductsParams): Promise<IProductsByCategory[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${process.env.API_BASE_URL}/category/product?${query}`, {
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
        const res = await fetch(`${process.env.API_BASE_URL}/product/${uuid}`, {
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
        const res = await fetch(`${process.env.API_BASE_URL}/carousel_image/`, {
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

export interface IProductImageParams {
    isChanged: boolean
    order: { [key: number]: number }
}

/** 商品リストを取得（管理画面用） */
export const getProducts = async (params: IGetProductsParams): Promise<IProduct[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${process.env.API_BASE_URL}/product?${query}`, {
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
export const createProduct = async (product: Omit<IProduct, 'uuid'>): Promise<IProduct> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/product`, {
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
export const updateProduct = async (uuid: string, product: IProduct): Promise<IProduct> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/product/${uuid}`, {
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
        const res = await fetch(`${process.env.API_BASE_URL}/product/${uuid}`, {
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

/** 商品画像をアップロード（順序指定付き） */
export const uploadProductImages = async (productUuid: string, files: File[], orderParams: IProductImageParams): Promise<void> => {
    try {
        const formData = new FormData()
        formData.append('order', JSON.stringify(orderParams))

        files.forEach((file, index) => {
            formData.append(`file${index}`, file)
        })

        const res = await fetch(`${process.env.API_BASE_URL}/product/${productUuid}/product_image`, {
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

/** 商品画像をアップロード */
export const uploadProductImage = async (productUuid: string, files: File[], orderParams: IProductImageParams): Promise<void> => {
    try {
        const formData = new FormData()
        formData.append('order', JSON.stringify(orderParams))

        files.forEach((file, index) => {
            formData.append(`file${index}`, file)
        })

        const res = await fetch(`${process.env.API_BASE_URL}/product/${productUuid}/product_image`, {
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
        const res = await fetch(`${process.env.API_BASE_URL}/product/duplicate`, {
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
