export default {
    // Global page headers: https://go.nuxtjs.dev/config-head
    head: {
        title: 'client',
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
                content: '',
            },
        ],
        link: [
            {
                rel: 'icon',
                type: 'image/x-icon',
                href: '/favicon.ico',
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
    components: true,

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
    ],
    webfontloader: {
        google: {
            families: ['Lobster:400,700'],
        },
    },

    // Axios module configuration: https://go.nuxtjs.dev/config-axios
    axios: {
        baseURL: `${process.env.AXIOS_API_URL}/api` || 'http://localhost:8080/api',
    },
    // Build Configuration: https://go.nuxtjs.dev/config-build
    build: {},
    router: {
        middleware: 'auth',
    },
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
    },
}
