<template>
    <v-sheet class="light-green lighten-4"> </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProduct, newProduct } from '~/types'
@Component({
    head: {
        title: '商品一覧',
    },
})
export default class PageProductIndex extends Vue {
    products: Array<IProduct> = []
    // form用のproductModel
    productModel: IProduct = newProduct()

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
