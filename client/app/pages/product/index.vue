<template>
    <div normal class="page-product">
        <div class="page-title-container">
            <h2 class="page-title">Product</h2>
        </div>
        <v-sheet>
            <div>
                <c-select-search group-name="Category" :items="categories" />
                <c-message v-if="products.length === 0" class="mt-4"> 登録されていません </c-message>
                <v-row>
                    <v-col v-for="listItem in products" :key="listItem.uuid" cols="12" sm="6" md="4" lg="3">
                        <c-product-card :list-item="listItem" @c-click="clickHandler(listItem)" />
                    </v-col>
                </v-row>
            </div>
        </v-sheet>
        <c-breadcrumbs :items="breadCrumbs" />
    </div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { newProduct } from '~/methods'
import { IBreadCrumb, IClassification, IGetCategoriesParams, IGetProductsParams, IProduct } from '~/types'
@Component({})
export default class PageProductIndex extends Vue {
    products: Array<IProduct> = []
    categories: Array<IClassification> = []
    // form用のproductModel
    productModel: IProduct = newProduct()

    breadCrumbs: Array<IBreadCrumb> = [
        { text: 'トップページ', href: '/' },
        { text: '商品一覧', disabled: true },
    ]

    async asyncData({ app }: Context) {
        try {
            const productParams: IGetProductsParams = {
                mode: 'active',
            }
            const products: Array<IProduct> = await app.$axios.$get(`/product`, { params: productParams })
            const categoryParams: IGetCategoriesParams = {
                mode: 'used',
            }
            const categories: Array<IProduct> = await app.$axios.$get(`/category`, { params: categoryParams })
            return { products, categories }
        } catch (e) {
            return { products: [], categories: [] }
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
.page-product
    margin 0 auto
    padding 48px
    max-width $xl-width
    +sm()
        padding 16px

.page-title-container
    +sm()
        display none
</style>
