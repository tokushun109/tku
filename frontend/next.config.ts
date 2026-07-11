import { NextConfig } from 'next'

const defaultApiBaseUrl = 'http://localhost:8080/api'
const apiBaseUrl = process.env.API_BASE_URL ? process.env.API_BASE_URL : defaultApiBaseUrl
const browserBaseUrl = process.env.BROWSER_BASE_URL ? process.env.BROWSER_BASE_URL : defaultApiBaseUrl
const domainUrl = process.env.DOMAIN_URL ? process.env.DOMAIN_URL : `http://localhost:${process.env.PORT ? process.env.PORT : '3000'}`
const isProduction = process.env.NODE_ENV === 'production'

const getOrigin = (url: string): string | undefined => {
    try {
        return new URL(url).origin
    } catch {
        return undefined
    }
}

const unique = (values: string[]): string[] => Array.from(new Set(values))

const apiOrigins = unique([apiBaseUrl, browserBaseUrl].map((url) => getOrigin(url)).filter((origin): origin is string => Boolean(origin)))

const buildContentSecurityPolicy = (): string => {
    const scriptSources = ["'self'", "'unsafe-inline'", 'https://www.googletagmanager.com', 'https://www.google-analytics.com']
    const connectSources = ["'self'", ...apiOrigins, 'https://www.google-analytics.com', 'https://region1.google-analytics.com']

    if (!isProduction) {
        scriptSources.push("'unsafe-eval'")
        connectSources.push('http://localhost:*', 'ws://localhost:*')
    }

    const directives = [
        ['default-src', "'self'"],
        ['base-uri', "'self'"],
        ['object-src', "'none'"],
        ['frame-ancestors', "'none'"],
        ['form-action', "'self'"],
        ['script-src', ...scriptSources],
        ['style-src', "'self'", "'unsafe-inline'"],
        [
            'img-src',
            "'self'",
            'data:',
            'blob:',
            'http://localhost:*',
            'http://minio:*',
            'https://tocoriri.com',
            'https://tku-api-ck57lb-prod.s3.ap-northeast-1.amazonaws.com',
            'https://www.google-analytics.com',
            'https://www.googletagmanager.com',
        ],
        ['font-src', "'self'", 'data:'],
        ['connect-src', ...connectSources],
        ['media-src', "'self'", 'blob:'],
        ['worker-src', "'self'", 'blob:'],
        ['manifest-src', "'self'"],
        ...(isProduction ? [['upgrade-insecure-requests']] : []),
    ]

    return directives.map((directive) => directive.join(' ')).join('; ')
}

const nextConfig: NextConfig = {
    output: 'standalone',
    async headers() {
        return [
            {
                source: '/:path*',
                headers: [
                    {
                        key: 'Content-Security-Policy',
                        value: buildContentSecurityPolicy(),
                    },
                ],
            },
        ]
    },
    sassOptions: {
        prependData: '@use "sass:color"; @use "@/styles/variables.scss" as *; @use "@/styles/mixins.scss" as *; @use "@/styles/layouts.scss" as *;',
    },
    images: {
        remotePatterns: [
            {
                protocol: 'http',
                hostname: 'localhost',
                pathname: '**',
            },
            {
                protocol: 'http',
                hostname: 'minio',
                pathname: '**',
            },
            {
                protocol: 'https',
                hostname: 'tocoriri.com',
                pathname: '**',
            },
            {
                protocol: 'https',
                hostname: 'tku-api-ck57lb-prod.s3.ap-northeast-1.amazonaws.com',
                pathname: '**',
            },
        ],
    },
    env: {
        API_BASE_URL: apiBaseUrl,
        BROWSER_BASE_URL: browserBaseUrl,
        DOMAIN_URL: domainUrl,
        COOKIE_DOMAIN_URL: process.env.COOKIE_DOMAIN_URL,
        GOOGLE_TAG: process.env.GOOGLE_TAG,
    },
}

module.exports = nextConfig
