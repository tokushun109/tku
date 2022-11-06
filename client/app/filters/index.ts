import Vue from 'vue'

// 金額フォーマット
Vue.filter('priceFormat', (value: number) => {
    const formatter = new Intl.NumberFormat('ja-JP')
    return formatter.format(value)
})
