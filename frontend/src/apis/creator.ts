import { ICreator } from '@/features/creator/type'
import { ApiError } from '@/utils/error'

export const getCreator = async (): Promise<ICreator> => {
    try {
        const res = await fetch(`${process.env.API_URL}/creator/`, {
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
        throw new Error('作者情報の取得に失敗しました')
    }
}
