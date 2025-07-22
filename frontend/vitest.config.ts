import path from 'node:path'
import { fileURLToPath } from 'node:url'

import { storybookTest } from '@storybook/addon-vitest/vitest-plugin'
import { defineConfig } from 'vitest/config'

const dirname = typeof __dirname !== 'undefined' ? __dirname : path.dirname(fileURLToPath(import.meta.url))

// More info at: https://storybook.js.org/docs/next/writing-tests/integrations/vitest-addon
export default defineConfig({
    resolve: {
        alias: {
            '@': path.resolve(dirname, './src'),
        },
    },
    optimizeDeps: {
        exclude: ['@mdx-js/react', 'markdown-to-jsx'],
    },
    test: {
        coverage: {
            provider: 'v8',
            reporter: ['text', 'html', 'json'],
            include: ['src/**/*.{ts,tsx}'],
            exclude: [
                'node_modules/',
                'dist/',
                'build/',
                'storybook-static/',
                'coverage/',
                '**/*.stories.{js,ts,jsx,tsx}',
                '**/*.config.{js,ts}',
                '**/*.test.{js,ts,jsx,tsx}',
                '**/__tests__/**',
                '.storybook/**',
                'public/**',
            ],
            thresholds: {
                lines: 50,
                functions: 50,
                branches: 50,
                statements: 50,
            },
        },
        projects: [
            {
                extends: true,
                plugins: [
                    // The plugin will run tests for the stories defined in your Storybook config
                    // See options at: https://storybook.js.org/docs/next/writing-tests/integrations/vitest-addon#storybooktest
                    storybookTest({ configDir: path.join(dirname, '.storybook') }),
                ],
                optimizeDeps: {
                    exclude: ['@mdx-js/react', 'markdown-to-jsx'],
                },
                test: {
                    name: 'storybook',
                    browser: {
                        enabled: true,
                        headless: true,
                        provider: 'playwright',
                        instances: [{ browser: 'chromium' }],
                    },
                    setupFiles: ['.storybook/vitest.setup.ts'],
                },
            },
            {
                resolve: {
                    alias: {
                        '@': path.resolve(dirname, './src'),
                    },
                },
                test: {
                    name: 'unit',
                    environment: 'jsdom',
                    setupFiles: ['./src/__tests__/setup.ts'],
                    globals: true,
                },
            },
            {
                resolve: {
                    alias: {
                        '@': path.resolve(dirname, './src'),
                    },
                },
                test: {
                    name: 'integration',
                    environment: 'jsdom',
                    setupFiles: ['./src/__tests__/setup.ts'],
                    globals: true,
                    include: ['src/__tests__/integration/**/*.test.{ts,tsx}'],
                },
                esbuild: {
                    jsxFactory: 'React.createElement',
                    jsxFragment: 'React.Fragment',
                },
            },
        ],
    },
})
