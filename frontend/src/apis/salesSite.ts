import { ISite } from '@/features/site/type'
import { ApiError } from '@/utils/error'

export const getSalesSiteList = async (): Promise<ISite[]> => {
    const res = await fetch(`${process.env.API_URL}/sales_site/`, {
        headers: {
            'Content-Type': 'application/json',
        },
        method: 'GET',
    })

    if (!res.ok) throw new ApiError(res)

    return await res.json()
}
