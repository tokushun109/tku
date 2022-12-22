const axios = require('axios')

module.exports = {
    ssr: 'true',
    srcDir: 'app/',
    // Global page headers: https://go.nuxtjs.dev/config-head
    head: {
        htmlAttrs: {
            lang: 'ja',
        },
        meta: [
            // 基本設定
            {
                charset: 'utf-8',
            },
            {
                name: 'viewport',
                content: 'width=device-width, initial-scale=1',
            },
            {
                'http-equiv': 'X-UA-Compatible',
                content: 'IE=edge',
            },
            {
                name: 'format-detection',
                content: 'telephone=no',
            },

            // SEO設定
            {
                name: 'twitter:card',
                content: 'summary',
            },

            // OGP設定
            {
                hid: 'og:site_name',
                property: 'og:site_name',
                content: 'tocoriri',
            },
        ],

        link: [
            {
                rel: 'icon',
                type: 'image/x-icon',
                href: '/favicon/favicon.ico',
            },
            {
                rel: 'apple-touch-icon',
                sizes: '180x180',
                href: '/apple-touch-icon/apple-touch-icon.png',
            },
        ],
    },

    // Global CSS: https://go.nuxtjs.dev/config-css
    css: [
        'ress',
        {
            src: './assets/style/main.styl',
            lang: 'stylus',
        },
    ],
    styleResources: {
        stylus: ['./assets/style/variables.styl', './assets/style/mixins.styl'],
    },

    // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
    plugins: ['~/plugins/dompurify', '~/filters'],

    // Auto import components: https://go.nuxtjs.dev/config-components
    components: [
        {
            path: '@/components/',
            pathPrefix: false,
        },
    ],

    // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
    buildModules: [
        // https://go.nuxtjs.dev/typescript
        '@nuxt/typescript-build',
        '@nuxtjs/vuetify',
        '@nuxtjs/google-gtag',
    ],

    'google-gtag': {
        id: process.env.GOOGLE_TAG,
        debug: false,
    },

    // Modules: https://go.nuxtjs.dev/config-modules
    modules: [
        // https://go.nuxtjs.dev/axios
        '@nuxtjs/axios',
        '@nuxtjs/style-resources',
        ['cookie-universal-nuxt', { parseJSON: false }],
        'nuxt-webfontloader',
        '@nuxtjs/sitemap',
        '@nuxtjs/robots',
    ],

    robots: {
        UserAgent: '*',
        Disallow: ['/admin', '/contact'],
        Sitemap: 'https://tocoriri.com/sitemap.xml',
    },

    webfontloader: {
        google: {
            families: ['Lobster:400,700'],
        },
    },

    sitemap: {
        hostname: 'https://tocoriri.com',
        exclude: ['/admin', '/admin/**', '/order', '/contact'],
        routes: async () => {
            // 商品詳細ページ
            const baseURL = process.env.API_BASE_URL || 'http://localhost:8080/api'
            const { data } = await axios.get(`${baseURL}/product?mode=active`)
            return data.map((product) => `/product/${product.uuid}`)
        },
    },

    // Axios module configuration: https://go.nuxtjs.dev/config-axios
    publicRuntimeConfig: {
        axios: {
            browserBaseURL: process.env.BROWSER_BASE_URL || 'http://localhost:8080/api',
        },
    },
    privateRuntimeConfig: {
        axios: {
            baseURL: process.env.API_BASE_URL || 'http://localhost:8080/api',
        },
    },
    // Build Configuration: https://go.nuxtjs.dev/config-build
    build: {
        publicPath: '_nuxt/',
    },
    router: {
        base: `/`,
        middleware: 'auth',
    },
    performance: {
        gzip: false,
    },
    dev: false,
    vuetify: {
        theme: {
            themes: {
                light: {
                    primary: '#4caf50',
                    secondary: '#8bc34a',
                    accent: '#cddc39',
                    warning: '#fff9c4',
                    info: '#a7ffeb',
                    success: '#b3e5fc',
                    error: '#ff80ab',
                },
            },
        },
        defaultAssets: false,
    },
}
