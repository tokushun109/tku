import { ICreator } from '@/features/creator/type'
import { ApiError } from '@/utils/error'

export const getCreator = async (): Promise<ICreator> => {
    const res = await fetch(`${process.env.API_URL}/creator/`, {
        headers: {
            'Content-Type': 'application/json',
        },
        method: 'GET',
    })

    if (!res.ok) throw new ApiError(res)

    return await res.json()
}
