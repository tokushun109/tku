import { NextResponse } from 'next/server'

import { getCurrentUser } from '@/apis/auth'
import { healthCheck } from '@/apis/healthCheck'
import { NavigationType } from '@/types/enum/navigation'

import type { NextRequest } from 'next/server'

const NOT_FOUND_PATH = '/not-found'

/**
 * セッション Cookie からログイン中ユーザーを取得し、管理者として扱えるかを判定する。
 */
async function checkAuth(request: NextRequest) {
    const sessionToken = request.cookies.get('__sess__')?.value

    if (!sessionToken) {
        return false
    }

    // ログイン中ユーザーを取得し、管理者のみ許可
    const currentUser = await getCurrentUser(sessionToken)
    return currentUser?.isAdmin ?? false
}

/**
 * リバースプロキシ経由で渡されるIP文字列を、環境変数と比較しやすい形に整える。
 *
 * IPv4 mapped IPv6（::ffff:192.0.2.1）やIPv4:port形式を許容し、
 * 空文字やunknownは比較対象外として空文字にする。
 */
function getNormalizedIP(ip: string) {
    const trimmedIP = ip.trim()
    if (!trimmedIP || trimmedIP.toLowerCase() === 'unknown') {
        return ''
    }

    if (trimmedIP.startsWith('::ffff:')) {
        return trimmedIP.replace('::ffff:', '')
    }

    if (trimmedIP.startsWith('[')) {
        const closingBracketIndex = trimmedIP.indexOf(']')
        return closingBracketIndex === -1 ? trimmedIP : trimmedIP.slice(1, closingBracketIndex)
    }

    const maybeIPv4WithPort = trimmedIP.match(/^(\d{1,3}(?:\.\d{1,3}){3}):\d+$/)
    return maybeIPv4WithPort?.[1] ?? trimmedIP
}

function getForwardedIPs(request: NextRequest) {
    return (
        request.headers
            .get('x-forwarded-for')
            ?.split(',')
            .map((ip) => getNormalizedIP(ip))
            .filter(Boolean) ?? []
    )
}

/**
 * クライアントIPをリクエストヘッダから取得する。
 *
 * x-forwarded-for は複数IPが入るため、最初に追加された先頭の値を優先する。
 * 取得できない場合は x-real-ip を参照する。
 */
function getClientIP(request: NextRequest) {
    const forwardedIPs = getForwardedIPs(request)
    const forwardedIP = forwardedIPs.at(0) ?? ''
    if (forwardedIP) {
        return forwardedIP
    }

    return getNormalizedIP(request.headers.get('x-real-ip') ?? '')
}

/**
 * admin配下へのアクセスを許可するIPかどうかを判定する。
 *
 * ENV=local の場合は、ローカル直アクセスでIPヘッダを取得できないため許可する。
 * local以外でMY_IP_ADDRESSが未設定、または空の値だけの場合は設定漏れとして拒否する。
 * カンマ区切りで複数IPを指定できる。
 */
function canAccessAdmin(request: NextRequest) {
    if (process.env.ENV === 'local') {
        return true
    }

    const allowedIP = process.env.MY_IP_ADDRESS
    if (!allowedIP) {
        return false
    }

    const allowedIPs = allowedIP
        .split(',')
        .map((ip) => getNormalizedIP(ip))
        .filter(Boolean)

    if (allowedIPs.length === 0) {
        return false
    }

    return allowedIPs.includes(getClientIP(request))
}

/**
 * adminのIP制限に失敗したリクエストを404ページへrewriteする。
 */
function rewriteNotFound(request: NextRequest) {
    return NextResponse.rewrite(new URL(NOT_FOUND_PATH, request.url), { status: 404 })
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

    const isAdminPath = request.nextUrl.pathname.startsWith('/admin')
    if (isAdminPath && !canAccessAdmin(request)) {
        return rewriteNotFound(request)
    }

    try {
        // ヘルスチェックAPIを呼び出し
        await healthCheck()

        // admin配下のページで認証が必要なルートをチェック
        if (isAdminPath) {
            const isAuthenticated = await checkAuth(request)
            const isAdminRoot = request.nextUrl.pathname === '/admin' || request.nextUrl.pathname === '/admin/'

            // /admin は認証状態に応じて遷移先を振り分ける
            if (isAdminRoot) {
                const destination = isAuthenticated ? NavigationType.AdminProduct : NavigationType.AdminLogin
                return NextResponse.redirect(new URL(destination, request.url))
            }

            // ログインページの場合
            if (request.nextUrl.pathname === NavigationType.AdminLogin) {
                // 既にログイン済みの場合は商品管理ページにリダイレクト
                if (isAuthenticated) {
                    return NextResponse.redirect(new URL(NavigationType.AdminProduct, request.url))
                }
                return NextResponse.next()
            }

            // その他のadminページは認証チェック
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
