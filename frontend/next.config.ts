import { NextConfig } from 'next'

const nextConfig: NextConfig = {
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
                hostname: 'tocoriri.com',
                pathname: '**',
            },
        ],
    },
    env: {
        API_URL: process.env.API_URL ? process.env.API_URL : 'http://localhost:8080/api',
        DOMAIN_URL: process.env.DOMAIN_URL ? process.env.DOMAIN_URL : `http://localhost:${process.env.PORT ? process.env.PORT : '3000'}`,
        MY_IP_ADDRESS: process.env.MY_IP_ADDRESS,
    },
}

module.exports = nextConfig
