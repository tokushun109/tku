import { IClassification, IClassificationForm } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetCategoriesParams {
    mode: 'all' | 'used'
}

export interface IPostCategoryParams {
    form: IClassificationForm
}

export interface IPutCategoryParams {
    form: IClassificationForm
    uuid: string
}

export interface IDeleteCategoryParams {
    uuid: string
}

export interface ICategoryResponse {
    message: string
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

/** カテゴリを追加 */
export const postCategory = async (params: IPostCategoryParams): Promise<ICategoryResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/category`, {
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify(params.form),
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('カテゴリの追加に失敗しました')
    }
}

/** カテゴリを更新 */
export const putCategory = async (params: IPutCategoryParams): Promise<ICategoryResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/category/${params.uuid}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'PUT',
            body: JSON.stringify(params.form),
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

/** カテゴリを削除 */
export const deleteCategory = async (params: IDeleteCategoryParams): Promise<ICategoryResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/category/${params.uuid}`, {
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'DELETE',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('カテゴリの削除に失敗しました')
    }
}
