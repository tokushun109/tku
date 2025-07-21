import { IClassification, IClassificationForm } from '@/features/classification/type'
import { ApiError } from '@/utils/error'

export interface IPostTagParams {
    form: IClassificationForm
}

export interface ITagResponse {
    message: string
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

/** タグを追加 */
export const postTag = async (params: IPostTagParams): Promise<ITagResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/tag`, {
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
        throw new Error('タグの追加に失敗しました')
    }
}
