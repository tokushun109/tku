import path from 'path'

import type { StorybookConfig } from '@storybook/nextjs'

const config: StorybookConfig = {
    stories: ['./Configure.mdx', '../src/**/*.stories.@(js|jsx|mjs|ts|tsx)'],
    staticDirs: ['../public'],
    addons: [
        '@storybook/addon-links',
        '@storybook/addon-essentials',
        '@storybook/addon-interactions',
        '@storybook/addon-styling-webpack',
        '@storybook/addon-themes',
        '@storybook/addon-controls',
    ],
    framework: {
        name: '@storybook/nextjs',
        options: {},
    },
    docs: {
        autodocs: 'tag',
    },
    webpackFinal: async (config) => {
        // @を../srcのエイリアスとして設定
        if (config.resolve) {
            config.resolve.alias = {
                ...config.resolve.alias,
                '@': path.resolve(__dirname, '../src'),
            }
        }
        // グローバルなscssファイルの読み込み
        if (config.module && config.module.rules) {
            config.module.rules.push({
                test: /\.scss$/i,
                use: [
                    {
                        loader: 'sass-loader',
                        options: {
                            additionalData:
                                '@use "@/styles/variables.scss" as *; @use "@/styles/mixins.scss" as *; @use "@/styles/layouts.scss" as *;',
                        },
                    },
                ],
            })
        }
        return config
    },
}
export default config
