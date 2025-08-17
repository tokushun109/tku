'use server'

import { cookies } from 'next/headers'
import { redirect } from 'next/navigation'

import { login, logout, validateSession } from '@/apis/auth'

import { LoginSchema } from './schema'

// サーバーアクションで返すレスポンスの型定義
interface LoginActionResult {
    error?: string
    success: boolean
}

export async function loginAction(prevState: LoginActionResult, formData: FormData): Promise<LoginActionResult> {
    try {
        // フォームデータの検証
        const result = LoginSchema.safeParse({
            email: formData.get('email'),
            password: formData.get('password'),
        })

        if (!result.success) {
            return {
                success: false,
                error: result.error.issues[0]?.message || 'フォームデータが不正です',
            }
        }

        // ログインAPI呼び出し
        const { session } = await login(result.data)

        // セッションCookieを設定
        const cookieStore = await cookies()
        cookieStore.set('__sess__', session.uuid, {
            httpOnly: true,
            secure: process.env.NODE_ENV === 'production',
            sameSite: 'lax',
            path: '/',
            // セッションの有効期限を設定（例：7日間）
            maxAge: 60 * 60 * 24 * 7,
            // サブドメイン間でのcookie共有のためにdomainを設定
            ...(process.env.COOKIE_DOMAIN_URL && { domain: process.env.COOKIE_DOMAIN_URL }),
        })

        return { success: true }
    } catch (error) {
        console.error('Login error:', error)
        return {
            success: false,
            error: error instanceof Error ? error.message : 'ログイン処理中にエラーが発生しました',
        }
    }
}

// ログアウト用のServer Action
export async function logoutAction(): Promise<void> {
    try {
        const cookieStore = await cookies()
        const sessionToken = cookieStore.get('__sess__')?.value

        // セッションがある場合はサーバーサイドでログアウト処理を実行
        if (sessionToken) {
            await logout(sessionToken)
        }

        // セッションCookieを削除
        cookieStore.delete('__sess__')
    } catch (error) {
        console.error('Logout error:', error)
        // エラーが発生してもCookieは削除する
        const cookieStore = await cookies()
        cookieStore.delete('__sess__')
    }

    // ログインページにリダイレクト
    redirect('/admin/login')
}

// セッションチェック用のServer Action
export async function checkSession(): Promise<boolean> {
    try {
        const cookieStore = await cookies()
        const sessionToken = cookieStore.get('__sess__')?.value

        if (!sessionToken) {
            return false
        }

        // セッション検証API呼び出し
        return await validateSession(sessionToken)
    } catch (error) {
        console.error('Session check error:', error)
        return false
    }
}
