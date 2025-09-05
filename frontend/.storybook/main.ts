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
                // TODO 個別に設定しなくても済む方法を調査する
                // MUI Material個別コンポーネント
                '@mui/material/useMediaQuery',
                // MUI Icons個別コンポーネント
                '@mui/icons-material/KeyboardArrowDown',
                '@mui/icons-material/Delete',
                '@mui/icons-material/Edit',
                '@mui/icons-material/Category',
                '@mui/icons-material/Label',
                '@mui/icons-material/People',
                '@mui/icons-material/Email',
                '@mui/icons-material/Person',
                '@mui/icons-material/Add',
                '@mui/icons-material/FileUpload',
                '@mui/icons-material/Close',
                '@mui/icons-material/Menu',
                '@mui/icons-material/ExpandMore',
                '@mui/icons-material/Facebook',
                '@mui/icons-material/Reply',
                '@mui/icons-material/X',
                '@mui/icons-material/Link',
                '@mui/icons-material/Store',
                'framer-motion',
                'classnames',
                'sonner',
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
