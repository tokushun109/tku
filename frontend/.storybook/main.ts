import path from 'path'

import type { StorybookConfig } from '@storybook/nextjs-vite'

const config: StorybookConfig = {
    stories: ['../src/**/*.stories.@(js|jsx|mjs|ts|tsx)'],
    addons: ['@chromatic-com/storybook', '@storybook/addon-docs', '@storybook/addon-a11y', '@storybook/addon-vitest'],
    framework: {
        name: '@storybook/nextjs-vite',
        options: {},
    },
    build: {
        test: {
            disableMDXEntries: true,
            disableAutoDocs: true,
            disableSourcemaps: true,
        },
    },
    staticDirs: ['../public'],
    viteFinal: async (config) => {
        // @を../srcのエイリアスとして設定
        if (config.resolve) {
            config.resolve.alias = {
                ...config.resolve.alias,
                '@': path.resolve(__dirname, '../src'),
            }
        }

        // 依存関係最適化の設定
        config.optimizeDeps = {
            ...config.optimizeDeps,
            entries: ['src/**/*.stories.@(ts|tsx|js|jsx)'],
            exclude: ['@mdx-js/react', '@storybook/blocks', '@storybook/addon-docs', '@storybook/addon-vitest'],
            include: [
                'react',
                'react-dom',
                'react/jsx-runtime',
                'react/jsx-dev-runtime',
                '@storybook/react',
                'storybook',
                'markdown-to-jsx',
                '@emotion/react',
                '@emotion/styled',
                '@mui/material',
                '@mui/icons-material',
                'framer-motion',
                'classnames',
                'sonner',
                'react-virtuoso',
            ],
        }

        // Storybook配信ルートに合わせてbaseを固定
        config.base = '/'

        // グローバルなscssファイルの読み込み
        if (config.css) {
            config.css.preprocessorOptions = {
                ...config.css.preprocessorOptions,
                scss: {
                    additionalData:
                        '@use "sass:color"; @use "@/styles/variables.scss" as *; @use "@/styles/mixins.scss" as *; @use "@/styles/layouts.scss" as *;',
                },
            }
        } else {
            config.css = {
                preprocessorOptions: {
                    scss: {
                        additionalData:
                            '@use "sass:color"; @use "@/styles/variables.scss" as *; @use "@/styles/mixins.scss" as *; @use "@/styles/layouts.scss" as *;',
                    },
                },
            }
        }

        return config
    },
}
export default config
