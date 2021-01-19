module.exports = {
    root: true,
    env: {
        browser: true,
        node: true,
    },
    extends: ['@nuxtjs/eslint-config-typescript', 'prettier', 'prettier/vue', 'plugin:prettier/recommended', 'plugin:nuxt/recommended'],
    plugins: ['prettier'],
    // add your custom rules here
    rules: {
        // おまじない
        'nuxt/no-cjs-in-config': 0,
        // 不要なカッコは消す
        'no-extra-parens': 1,
        // 無駄なスペースは削除
        'no-multi-spaces': 2,
        // 不要な改行は削除
        'no-multiple-empty-lines': [
            2,
            {
                max: 1,
            },
        ],
        // 関数とカッコはあけない
        'space-before-function-paren': [0, 'never'],
        // true/falseを無駄に使うな
        'no-unneeded-ternary': 2,
        // varは禁止
        'no-var': 2,
        // コンソールはwarning
        'no-console': 1,
        // 配列のindexには空白入れるな(hogehoge[ x ])
        'computed-property-spacing': 2,
        // キー
        'key-spacing': 2,
        // キーワードの前後には適切なスペースを
        'keyword-spacing': 2,
        // 使ってない変数は警告
        'no-unused-vars': 0,
        '@typescript-eslint/no-unused-vars': 1,
        // prettier
        'prettier/prettier': [
            'error',
            {
                printWidth: 150,
                tabWidth: 4,
                semi: false,
                singleQuote: true,
                arrowParens: 'always',
                jsxBracketSameLine: true,
            },
        ],
    },
}
