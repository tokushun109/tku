<template>
    <v-container class="page-product">
        <v-container class="page-title-container">
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
        <c-breadcrumbs :items="breadCrumbs" />
    </v-container>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { newProduct } from '~/methods'
import { IBreadCrumb, IGetProductsParams, IProduct } from '~/types'
@Component({})
export default class PageProductIndex extends Vue {
    products: Array<IProduct> = []
    // form用のproductModel
    productModel: IProduct = newProduct()

    breadCrumbs: Array<IBreadCrumb> = [
        { text: 'トップページ', href: '/' },
        { text: '商品一覧', disabled: true },
    ]

    async asyncData({ app }: Context) {
        try {
            const params: IGetProductsParams = {
                mode: 'active',
            }
            const products: Array<IProduct> = await app.$axios.$get(`/product`, { params })
            return { products }
        } catch (e) {
            return { products: [] }
        }
    }

    head() {
        if (!this.products) {
            return
        }
        const title = '商品一覧 | とこりり'
        const description = 'とこりりの商品一覧ページです。'
        const image = this.products[0].productImages[0].apiPath
        return {
            title,
            meta: [
                {
                    hid: 'description',
                    name: 'description',
                    content: description,
                },
                {
                    hid: 'og:title',
                    property: 'og:title',
                    content: title,
                },
                {
                    hid: 'og:description',
                    property: 'og:description',
                    content: description,
                },
                {
                    hid: 'og:type',
                    property: 'og:type',
                    content: 'article',
                },
                {
                    hid: 'og:image',
                    property: 'og:image',
                    content: image,
                },
            ],
        }
    }

    clickHandler(item: IProduct) {
        this.$router.push(`/product/${item.uuid}`)
    }
}
</script>

<style lang="stylus" scoped>
.page-title-container
    +sm()
        display none
    .page-title
        margin-bottom 20px
        color $site-title-text-color
        text-align center
</style>
