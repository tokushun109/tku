import { IClassification } from '@/features/classification/type'
import { ApiError } from '@/utils/error'

export interface ICreateTagParams {
    name: string
}

export interface IUpdateTagParams {
    name: string
}

export const getTags = async (): Promise<IClassification[]> => {
    try {
        const res = await fetch(`${process.env.API_URL}/tag`, {
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
        throw new Error('タグ一覧の取得に失敗しました')
    }
}

export const createTag = async (params: ICreateTagParams): Promise<IClassification> => {
    try {
        const res = await fetch(`${process.env.API_URL}/tag`, {
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
        throw new Error('タグの作成に失敗しました')
    }
}

export const updateTag = async (uuid: string, params: IUpdateTagParams): Promise<IClassification> => {
    try {
        const res = await fetch(`${process.env.API_URL}/tag/${uuid}`, {
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
        throw new Error('タグの更新に失敗しました')
    }
}

export const deleteTag = async (uuid: string): Promise<void> => {
    try {
        const res = await fetch(`${process.env.API_URL}/tag/${uuid}`, {
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
        throw new Error('タグの削除に失敗しました')
    }
}
