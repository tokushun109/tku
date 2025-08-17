import { ICreator } from '@/features/creator/type'
import { ApiError } from '@/utils/error'

export const getCreator = async (): Promise<ICreator> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/creator/`, {
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

export const updateCreator = async (creator: Omit<ICreator, 'apiPath' | 'logo'>) => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/creator/`, {
            body: JSON.stringify(creator),
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'PUT',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('作者情報の更新に失敗しました')
    }
}

export const updateCreatorLogo = async (file: File) => {
    try {
        const formData = new FormData()
        formData.append('logo', file)

        const res = await fetch(`${process.env.API_BASE_URL}/creator/logo/`, {
            body: formData,
            method: 'PUT',
            credentials: 'include',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('作者ロゴの更新に失敗しました')
    }
}
