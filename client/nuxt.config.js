module.exports = {
    ssr: 'true',
    srcDir: 'app/',
    // Global page headers: https://go.nuxtjs.dev/config-head
    head: {
        title: 'tocoriri',
        htmlAttrs: {
            lang: 'ja',
        },
        meta: [
            {
                charset: 'utf-8',
            },
            {
                name: 'viewport',
                content: 'width=device-width, initial-scale=1',
            },
            {
                hid: 'description',
                name: 'description',
                content: 'tocoriri web site',
            },
        ],
        link: [
            {
                rel: 'icon',
                type: 'image/x-icon',
                href: '/favicon/favicon.ico',
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
        pathPrefix: false
        }
    ],

    // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
    buildModules: [
        // https://go.nuxtjs.dev/typescript
        '@nuxt/typescript-build',
        '@nuxtjs/vuetify',
    ],

    // Modules: https://go.nuxtjs.dev/config-modules
    modules: [
        // https://go.nuxtjs.dev/axios
        '@nuxtjs/axios',
        '@nuxtjs/style-resources',
        ['cookie-universal-nuxt', { parseJSON: false }],
        'nuxt-webfontloader',
        '@nuxtjs/sitemap',
    ],

    webfontloader: {
        google: {
            families: ['Lobster:400,700'],
        },
    },

    sitemap: {
        hostname: 'https://tocoriri.com',
        exclude: [
            '/admin',
            '/admin/**',
            '/order',
        ],
        routes: async () => {
            const axios = require('axios')
            // 商品詳細ページ
            const baseURL = process.env.API_BASE_URL || 'http://localhost:8080/api'
            const { data } = await axios.get(`${baseURL}/product?mode=active`)
            return data.map((product) => `/product/${product.uuid}`)
        }
    },

    // Axios module configuration: https://go.nuxtjs.dev/config-axios
    publicRuntimeConfig: {
    axios: {
        browserBaseURL:
            process.env.BROWSER_BASE_URL || 'http://localhost:8080/api'
        }
    },
    privateRuntimeConfig: {
    axios: {
        baseURL:
            process.env.API_BASE_URL || 'http://localhost:8080/api'
        }
    },    
    // Build Configuration: https://go.nuxtjs.dev/config-build
    build: {
        publicPath: '_nuxt/'
    },
    router: {
        base: `/`,
        middleware: 'auth',
    },
    performance: {
        gzip: false
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
        defaultAssets: false
    },
}
