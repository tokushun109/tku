import { IClassification } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetCategoriesParams {
    mode: 'all' | 'used'
}

export interface ICreateCategoryParams {
    name: string
}

export interface IUpdateCategoryParams {
    name: string
}

export const getCategories = async (params: IGetCategoriesParams): Promise<IClassification[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${process.env.API_URL}/category?${query}`, {
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
        throw new Error('カテゴリ一覧の取得に失敗しました')
    }
}

export const createCategory = async (params: ICreateCategoryParams): Promise<IClassification> => {
    try {
        const res = await fetch(`${process.env.API_URL}/category`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
            body: JSON.stringify(params),
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('カテゴリの作成に失敗しました')
    }
}

export const updateCategory = async (uuid: string, params: IUpdateCategoryParams): Promise<IClassification> => {
    try {
        const res = await fetch(`${process.env.API_URL}/category/${uuid}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'PUT',
            body: JSON.stringify(params),
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('カテゴリの更新に失敗しました')
    }
}

export const deleteCategory = async (uuid: string): Promise<void> => {
    try {
        const res = await fetch(`${process.env.API_URL}/category/${uuid}`, {
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
        throw new Error('カテゴリの削除に失敗しました')
    }
}
