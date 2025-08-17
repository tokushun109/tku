import { NextResponse } from 'next/server'

import { validateSession } from '@/apis/auth'
import { NavigationType } from '@/types/enum/navigation'

import type { NextRequest } from 'next/server'

// 認証チェック関数
async function checkAuth(request: NextRequest) {
    const sessionToken = request.cookies.get('__sess__')?.value

    if (!sessionToken) {
        return false
    }

    // セッション検証API呼び出し
    return await validateSession(sessionToken)
}

export async function middleware(request: NextRequest) {
    // メンテナンスページへのアクセスの場合はそのまま通す
    if (request.nextUrl.pathname === '/maintenance') {
        return NextResponse.next()
    }

    // 静的ファイルやAPI routes、Next.js内部ファイルはスキップ
    if (
        request.nextUrl.pathname.startsWith('/_next') ||
        request.nextUrl.pathname.startsWith('/api') ||
        request.nextUrl.pathname.startsWith('/static') ||
        request.nextUrl.pathname.includes('.')
    ) {
        return NextResponse.next()
    }

    try {
        // ヘルスチェックAPIを呼び出し
        const response = await fetch(`${process.env.API_BASE_URL}/health_check`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        })

        // ヘルスチェックが失敗した場合、メンテナンスページにリダイレクト
        if (!response.ok) {
            return NextResponse.redirect(new URL('/maintenance', request.url))
        }

        // admin配下のページで認証が必要なルートをチェック
        if (request.nextUrl.pathname.startsWith('/admin')) {
            // ログインページの場合
            if (request.nextUrl.pathname === NavigationType.AdminLogin) {
                // 既にログイン済みの場合は商品管理ページにリダイレクト
                const isAuthenticated = await checkAuth(request)
                if (isAuthenticated) {
                    return NextResponse.redirect(new URL(NavigationType.AdminProduct, request.url))
                }
                return NextResponse.next()
            }

            // その他のadminページは認証チェック
            const isAuthenticated = await checkAuth(request)
            if (!isAuthenticated) {
                return NextResponse.redirect(new URL(NavigationType.AdminLogin, request.url))
            }
        }

        // ヘルスチェックが成功した場合、通常通り処理を続行
        return NextResponse.next()
    } catch {
        // ネットワークエラーなどでヘルスチェックが失敗した場合もメンテナンスページにリダイレクト
        return NextResponse.redirect(new URL('/maintenance', request.url))
    }
}

export const config = {
    matcher: [
        /*
         * Match all request paths except for the ones starting with:
         * - api (API routes)
         * - _next/static (static files)
         * - _next/image (image optimization files)
         * - favicon.ico (favicon file)
         */
        '/((?!api|_next/static|_next/image|favicon.ico).*)',
    ],
}
