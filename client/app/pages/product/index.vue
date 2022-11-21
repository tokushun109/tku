<template>
    <v-container class="page-product">
        <v-container>
            <h2 class="page-title text-sm-h3 text-h4">PRODUCTS</h2>
        </v-container>
        <v-sheet>
            <v-container>
                <c-message v-if="products.length === 0" class="mt-4"> 登録されていません </c-message>
                <v-row>
                    <v-col v-for="listItem in products" :key="listItem.uuid" cols="12" sm="6" md="4">
                        <c-product-card :list-item="listItem" @c-click="clickHandler(listItem)" />
                    </v-col>
                </v-row>
            </v-container>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { newProduct } from '~/methods'
import { IGetProductsParams, IProduct } from '~/types'
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
            const params: IGetProductsParams = {
                mode: 'active',
            }
            const products = await app.$axios.$get(`/product`, { params })
            return { products }
        } catch (e) {
            return { products: [] }
        }
    }

    clickHandler(item: IProduct) {
        this.$router.push(`/product/${item.uuid}`)
    }
}
</script>

<style lang="stylus" scoped>
.page-title
    margin-bottom 20px
    color $site-title-text-color
    text-align center
    +sm()
        margin-bottom auto
</style>