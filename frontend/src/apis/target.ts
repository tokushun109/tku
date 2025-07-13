import { IClassification } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetTargetsParams {
    mode: 'all' | 'used'
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
