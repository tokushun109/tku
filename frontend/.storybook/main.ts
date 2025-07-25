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
        disableMDXEntries: true,
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
            exclude: ['@mdx-js/react', 'markdown-to-jsx', '@storybook/blocks', '@storybook/addon-docs', '@storybook/addon-vitest'],
            include: ['react', 'react-dom', '@storybook/react', 'storybook'],
        }

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
