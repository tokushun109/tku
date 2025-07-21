import { IClassification, IClassificationForm } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetCategoriesParams {
    mode: 'all' | 'used'
}

export interface IPostCategoryParams {
    form: IClassificationForm
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
