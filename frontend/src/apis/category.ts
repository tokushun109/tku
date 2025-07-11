import { IClassification } from '@/features/classification/type'
import { ApiError } from '@/utils/error'
import { convertObjectToURLSearchParams } from '@/utils/request'

export interface IGetCategoriesParams {
    mode: 'all' | 'used'
}

export const getCategories = async (params: IGetCategoriesParams): Promise<IClassification[]> => {
    const query = convertObjectToURLSearchParams(params)
    const res = await fetch(`${process.env.API_URL}/category?${query}`, {
        headers: {
            'Content-Type': 'application/json',
        },
        method: 'GET',
    })

    if (!res.ok) throw new ApiError(res)

    return await res.json()
}
