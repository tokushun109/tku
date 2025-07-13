import { ISite } from '@/features/site/type'
import { ApiError } from '@/utils/error'

export const getSalesSiteList = async (): Promise<ISite[]> => {
    try {
        const res = await fetch(`${process.env.API_URL}/sales_site/`, {
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
