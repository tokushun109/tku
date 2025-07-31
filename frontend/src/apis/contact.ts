import { IContact, IContactListItem } from '@/features/contact/type'
import { ApiError } from '@/utils/error'

export interface IPostContactParams {
    form: IContact
}

export interface IContactResponse {
    message: string
}

/** お問い合わせ一覧を取得 */
export const getContacts = async (): Promise<IContactListItem[]> => {
    try {
        // Server Componentsでは手動でCookieヘッダーを設定する必要がある
        const { cookies } = await import('next/headers')
        const cookieStore = await cookies()
        const cookieHeader = cookieStore.toString()

        const res = await fetch(`${process.env.API_URL}/contact`, {
            headers: {
                'Content-Type': 'application/json',
                ...(cookieHeader && { Cookie: cookieHeader }),
            },
            method: 'GET',
        })

        if (!res.ok) throw new ApiError(res)

        return await res.json()
    } catch (error) {
        if (error instanceof ApiError) {
            throw error
        }
        throw new Error('お問い合わせ一覧の取得に失敗しました')
    }
}

/** お問い合わせを送信 */
export const postContact = async (params: IPostContactParams): Promise<IContactResponse> => {
    try {
        const res = await fetch(`${process.env.API_URL}/contact`, {
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
        throw new Error('お問い合わせの送信に失敗しました')
    }
}
