import path from 'node:path'
import { fileURLToPath } from 'node:url'

import { FlatCompat } from '@eslint/eslintrc'
import js from '@eslint/js'
import prettierConfig from 'eslint-config-prettier'
import jsxA11y from 'eslint-plugin-jsx-a11y'
import perfectionistPlugin from 'eslint-plugin-perfectionist'
import unusedImports from 'eslint-plugin-unused-imports'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const compat = new FlatCompat({
    baseDirectory: __dirname,
    recommendedConfig: js.configs.recommended,
    allConfig: js.configs.all,
})

const configs = [
    ...compat.extends('next', 'next/core-web-vitals', 'eslint:recommended', 'plugin:storybook/recommended'),
    {
        plugins: {
            'unused-imports': unusedImports,
            'jsx-a11y': jsxA11y,
            perfectionist: perfectionistPlugin,
        },

        languageOptions: {
            globals: {
                React: 'readonly',
                JSX: 'readonly',
            },
        },

        rules: {
            'no-unused-vars': [
                'error',
                {
                    args: 'after-used',
                    argsIgnorePattern: '^_',
                    varsIgnorePattern: '^_$',
                },
            ],

            'no-console': [
                'error',
                {
                    allow: ['error'],
                },
            ],

            'import/order': [
                'error',
                {
                    groups: ['builtin', 'external', 'internal', ['parent', 'sibling'], 'index', 'object', 'type'],

                    'newlines-between': 'always',

                    alphabetize: {
                        order: 'asc',
                    },
                },
            ],

            // interface のプロパティの並び順をアルファベット順に統一
            'perfectionist/sort-interfaces': 'warn',
            // Object 型のプロパティの並び順をアルファベット順に統一
            'perfectionist/sort-object-types': 'warn',

            'unused-imports/no-unused-imports': 'error',
            // Props などの分割代入を強制
            'react/destructuring-assignment': 'error',
            // コンポーネントの定義方法をアロー関数に統一
            'react/function-component-definition': [
                'error',
                {
                    namedComponents: 'arrow-function',
                    unnamedComponents: 'arrow-function',
                },
            ],
            // useState の返り値の命名を [value, setValue] に統一
            'react/hook-use-state': 'error',
            // boolean 型の Props の渡し方を統一
            'react/jsx-boolean-value': 'error',
            // React Fragment の書き方を統一
            'react/jsx-fragments': 'error',
            // Props と children で不要な中括弧を削除
            'react/jsx-curly-brace-presence': 'error',
            // 不要な React Fragment を削除
            'react/jsx-no-useless-fragment': 'error',
            // Props の並び順をアルファベット順に統一
            'react/jsx-sort-props': 'error',
            // 子要素がない場合は自己終了タグを使う
            'react/self-closing-comp': 'error',
            // コンポーネント名をパスカルケースに統一
            'react/jsx-pascal-case': 'error',
            // Props の型チェックは TS で行う
            'react/prop-types': 'off',
            // exhaustive-depsに関してもエラーを出力する
            'react-hooks/exhaustive-deps': 'error',
        },
    },

    {
        files: ['**/*.stories.*'],
        rules: {
            'no-console': 'off',
        },
    },
    prettierConfig, // フォーマット は Prettier で行うため、フォーマット関連のルールを無効化
]

export default configs
