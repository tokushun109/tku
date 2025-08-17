import { IClassification, IClassificationForm } from '@/features/classification/type'
import { ApiError } from '@/utils/error'

export interface IPostTagParams {
    form: IClassificationForm
}

export interface IPutTagParams {
    form: IClassificationForm
    uuid: string
}

export interface IDeleteTagParams {
    uuid: string
}

export interface ITagResponse {
    message: string
}

export const getTags = async (): Promise<IClassification[]> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/tag`, {
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

/** タグを追加 */
export const postTag = async (params: IPostTagParams): Promise<ITagResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/tag`, {
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
        throw new Error('タグの追加に失敗しました')
    }
}

/** タグを更新 */
export const putTag = async (params: IPutTagParams): Promise<ITagResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/tag/${params.uuid}`, {
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
        throw new Error('タグの更新に失敗しました')
    }
}

/** タグを削除 */
export const deleteTag = async (params: IDeleteTagParams): Promise<ITagResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/tag/${params.uuid}`, {
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
        throw new Error('タグの削除に失敗しました')
    }
}
