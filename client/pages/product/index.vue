<template>
    <div>{{ products }}</div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProduct } from '~/types'

@Component({
    head: {
        title: '商品一覧',
    },
})
export default class ProductIndex extends Vue {
    // TODO 型を指定する
    products: Array<IProduct> | null = []
    async asyncData({ app }: Context) {
        try {
            const products = await app.$axios.$get(`/product/`)
            return { products }
        } catch (e) {
            return { products: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
