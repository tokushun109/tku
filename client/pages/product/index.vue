<template>
    <c-page>
        <div>
            <ul v-for="product in products" :key="product.uuid">
                <li>{{ product }}</li>
                <li>
                    <nuxt-link :to="`/product/${product.uuid}`">詳細へ</nuxt-link>
                </li>
            </ul>
        </div>
    </c-page>
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
export default class PageProductIndex extends Vue {
    products: Array<IProduct> = []
    async asyncData({ app }: Context) {
        try {
            const products = await app.$axios.$get(`/product`)
            return { products }
        } catch (e) {
            return { products: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
