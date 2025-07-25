import { IClassification, IClassificationForm } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetTargetsParams {
    mode: 'all' | 'used'
}

export interface IPostTargetParams {
    form: IClassificationForm
}

export interface IPutTargetParams {
    form: IClassificationForm
    uuid: string
}

export interface IDeleteTargetParams {
    uuid: string
}

export interface ITargetResponse {
    message: string
}

export const getTargets = async (params: IGetTargetsParams): Promise<IClassification[]> => {
    try {
        const query = convertObjectToURLSearchParams(params)
        const res = await fetch(`${process.env.API_URL}/target?${query}`, {
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
        throw new Error('ターゲット一覧の取得に失敗しました')
    }
}

/** ターゲットを追加 */
export const postTarget = async (params: IPostTargetParams): Promise<ITargetResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/target`, {
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
        throw new Error('ターゲットの追加に失敗しました')
    }
}

/** ターゲットを更新 */
export const putTarget = async (params: IPutTargetParams): Promise<ITargetResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/target/${params.uuid}`, {
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
        throw new Error('ターゲットの更新に失敗しました')
    }
}

/** ターゲットを削除 */
export const deleteTarget = async (params: IDeleteTargetParams): Promise<ITargetResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/target/${params.uuid}`, {
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
        throw new Error('ターゲットの削除に失敗しました')
    }
}
