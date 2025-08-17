import { ApiError } from '@/utils/error'

export interface IHealthCheckResponse {
    message?: string
}

export const healthCheck = async (): Promise<IHealthCheckResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/health_check`, {
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
        throw new Error('APIサーバーのヘルスチェックに失敗しました')
    }
}
