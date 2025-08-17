import { ISite, ISiteForm } from '@/features/site/type'
import { ApiError } from '@/utils/error'

export interface IPostSalesSiteParams {
    form: ISiteForm
}

export interface IPutSalesSiteParams {
    form: ISiteForm
    uuid: string
}

export interface IDeleteSalesSiteParams {
    uuid: string
}

export interface ISalesSiteResponse {
    message: string
}

export const getSalesSiteList = async (): Promise<ISite[]> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/sales_site/`, {
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
        throw new Error('販売サイト一覧の取得に失敗しました')
    }
}

/** 販売サイトを追加 */
export const postSalesSite = async (params: IPostSalesSiteParams): Promise<ISalesSiteResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/sales_site`, {
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
        throw new Error('販売サイトの追加に失敗しました')
    }
}

/** 販売サイトを更新 */
export const putSalesSite = async (params: IPutSalesSiteParams): Promise<ISalesSiteResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/sales_site/${params.uuid}`, {
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
        throw new Error('販売サイトの更新に失敗しました')
    }
}

/** 販売サイトを削除 */
export const deleteSalesSite = async (params: IDeleteSalesSiteParams): Promise<ISalesSiteResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/sales_site/${params.uuid}`, {
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
        throw new Error('販売サイトの削除に失敗しました')
    }
}
