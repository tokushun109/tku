import { ILoginForm, ISession } from '@/features/auth/type'
import { ApiError } from '@/utils/error'

// ログインAPIのレスポンス型
export interface ILoginResponse {
    session: ISession
}

// ログイン
export const login = async (loginData: ILoginForm): Promise<ILoginResponse> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/user/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                Email: loginData.email,
                Password: loginData.password,
            }),
        })

        if (!res.ok) {
            throw new ApiError(res)
        }

        const session: ISession = await res.json()
        return { session }
    } catch {
        throw new Error('ログイン処理に失敗しました')
    }
}

// ログアウトAPI関数
export const logout = async (sessionToken: string): Promise<void> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/user/logout`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Cookie: `__sess__=${sessionToken}`,
            },
            credentials: 'include',
        })

        if (!res.ok) {
            throw new ApiError(res)
        }
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('ログアウト処理に失敗しました')
    }
}

// セッション検証API関数
export const validateSession = async (sessionToken: string): Promise<boolean> => {
    try {
        const res = await fetch(`${process.env.API_BASE_URL}/user/login`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Cookie: `__sess__=${sessionToken}`,
            },
            credentials: 'include',
        })
        return res.ok
    } catch (error) {
        console.error('Session validation error:', error)
        return false
    }
}
