import { ISite, ISiteForm } from '@/features/site/type'
import { ApiError } from '@/utils/error'

export interface IPostSnsParams {
    form: ISiteForm
}

export interface IPutSnsParams {
    form: ISiteForm
    uuid: string
}

export interface IDeleteSnsParams {
    uuid: string
}

export interface ISnsResponse {
    message: string
}

export const getSnsList = async (): Promise<ISite[]> => {
    try {
        const res = await fetch(`${process.env.API_URL}/sns/`, {
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
        throw new Error('SNS一覧の取得に失敗しました')
    }
}

/** SNSを追加 */
export const postSns = async (params: IPostSnsParams): Promise<ISnsResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/sns`, {
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
        throw new Error('SNSの追加に失敗しました')
    }
}

/** SNSを更新 */
export const putSns = async (params: IPutSnsParams): Promise<ISnsResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/sns/${params.uuid}`, {
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
        throw new Error('SNSの更新に失敗しました')
    }
}

/** SNSを削除 */
export const deleteSns = async (params: IDeleteSnsParams): Promise<ISnsResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/sns/${params.uuid}`, {
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
        throw new Error('SNSの削除に失敗しました')
    }
}
