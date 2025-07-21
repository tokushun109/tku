import { IClassification } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetTargetsParams {
    mode: 'all' | 'used'
}

export interface ICreateTargetParams {
    name: string
}

export interface IUpdateTargetParams {
    name: string
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

export const createTarget = async (params: ICreateTargetParams): Promise<IClassification> => {
    try {
        const res = await fetch(`${process.env.API_URL}/target`, {
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
        throw new Error('ターゲットの作成に失敗しました')
    }
}

export const updateTarget = async (uuid: string, params: IUpdateTargetParams): Promise<IClassification> => {
    try {
        const res = await fetch(`${process.env.API_URL}/target/${uuid}`, {
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
        throw new Error('ターゲットの更新に失敗しました')
    }
}

export const deleteTarget = async (uuid: string): Promise<void> => {
    try {
        const res = await fetch(`${process.env.API_URL}/target/${uuid}`, {
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
        throw new Error('ターゲットの削除に失敗しました')
    }
}
