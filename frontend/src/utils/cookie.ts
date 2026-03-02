/**
 * Server Components用のCookie関連ユーティリティ関数
 */

/**
 * Server ComponentsでAPI呼び出し時に使用するCookieヘッダーを取得する
 * Next.js App Router Server Componentsでは credentials: 'include' が機能しないため、
 * 手動でCookieヘッダーを設定する必要がある
 *
 * @returns Cookie header string or undefined if no cookies
 */
export const getCookieHeaderForAPI = async (): Promise<string | undefined> => {
    try {
        // dynamic importを使用してNext.jsのServer環境でのみ実行
        const { cookies } = await import('next/headers')
        const cookieStore = await cookies()
        const cookieHeader = cookieStore.toString()

        return cookieHeader || undefined
    } catch (error) {
        console.error('Failed to get cookies for API call:', error)
        return undefined
    }
}

/**
 * API fetchリクエスト用のheadersオブジェクトを作成する
 * Server ComponentsではCookieヘッダーを自動的に含める
 *
 * @param additionalHeaders - 追加のヘッダー
 * @returns fetch用のheadersオブジェクト
 */
export const createAPIHeaders = async (additionalHeaders: Record<string, string> = {}): Promise<Record<string, string>> => {
    const baseHeaders = {
        'Content-Type': 'application/json',
        ...additionalHeaders,
    }

    if (typeof window !== 'undefined') {
        return baseHeaders
    }

    const cookieHeader = await getCookieHeaderForAPI()

    return {
        ...baseHeaders,
        ...(cookieHeader && { Cookie: cookieHeader }),
    }
}
