import { NextConfig } from 'next'

const nextConfig: NextConfig = {
    output: 'standalone',
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
        API_BASE_URL: process.env.API_BASE_URL ? process.env.API_BASE_URL : 'http://localhost:8080/api',
        DOMAIN_URL: process.env.DOMAIN_URL ? process.env.DOMAIN_URL : `http://localhost:${process.env.PORT ? process.env.PORT : '3000'}`,
        COOKIE_DOMAIN_URL: process.env.COOKIE_DOMAIN_URL,
        GOOGLE_TAG: process.env.GOOGLE_TAG,
    },
}

module.exports = nextConfig
