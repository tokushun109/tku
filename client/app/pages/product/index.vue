<template>
    <div normal class="page-product">
        <div class="page-title-container">
            <h2 class="page-title">Product</h2>
        </div>
        <v-sheet>
            <div class="fade-up">
                <v-btn text :color="ColorType.Grey" class="search-button sm" @click="toggleSearchArea">
                    <v-icon>{{ mdiMagnify }}</v-icon>
                    <span class="icon-text">SEARCH</span>
                </v-btn>
                <v-expand-transition>
                    <div class="search-area">
                        <div v-if="isSearchAreaDisplay" class="search-area__content">
                            <c-select-search
                                group-name="Category"
                                :items="categories"
                                :target-content.sync="productParams.category"
                                @c-select-search="selectSearchHandler"
                            />
                            <c-select-search
                                group-name="Target"
                                :items="targets"
                                :target-content.sync="productParams.target"
                                @c-select-search="selectSearchHandler"
                            />
                        </div>
                    </div>
                </v-expand-transition>
            </div>
            <c-message v-if="products.length === 0" wide>該当する商品が<br class="sm" />見つかりませんでした</c-message>
            <v-row class="fade-up">
                <v-col v-for="listItem in products" :key="listItem.uuid" cols="12" sm="6" md="4" lg="3">
                    <c-product-card :list-item="listItem" @c-click="clickHandler(listItem)" />
                </v-col>
            </v-row>
        </v-sheet>
        <c-breadcrumbs :items="breadCrumbs" />
    </div>
</template>

<script lang="ts">
import { mdiMagnify } from '@mdi/js'
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { newProduct } from '~/methods'
import { ColorType, IBreadCrumb, IClassification, IGetClassificationParams, IGetProductsParams, IProduct } from '~/types'
@Component({})
export default class PageProductIndex extends Vue {
    ColorType: typeof ColorType = ColorType

    products: Array<IProduct> = []
    categories: Array<IClassification> = []
    targets: Array<IClassification> = []
    searchCategory: string = 'all'
    // form用のproductModel
    productModel: IProduct = newProduct()
    mdiMagnify = mdiMagnify

    isSearchAreaDisplay: boolean = false

    breadCrumbs: Array<IBreadCrumb> = [
        { text: 'トップページ', href: '/' },
        { text: '商品一覧', disabled: true },
    ]

    productParams: IGetProductsParams = {
        mode: 'active',
        category: 'all',
        target: 'all',
    }

    async asyncData({ app }: Context) {
        try {
            const productParams: IGetProductsParams = {
                mode: 'active',
                category: 'all',
                target: 'all',
            }
            const products: Array<IProduct> = await app.$axios.$get(`/product`, { params: productParams })
            const classificationParams: IGetClassificationParams = {
                mode: 'used',
            }
            const categories: Array<IProduct> = await app.$axios.$get(`/category`, { params: classificationParams })
            const targets: Array<IProduct> = await app.$axios.$get(`/target`, { params: classificationParams })
            return { products, categories, targets }
        } catch (e) {
            return { products: [], categories: [], targets: [] }
        }
    }

    async loadingProduct() {
        this.products = await this.$axios.$get(`/product`, { params: this.productParams })
    }

    head() {
        if (this.products.length === 0 || this.products[0].productImages.length === 0) {
            return
        }
        const title = '商品一覧 | とこりり'
        const description = 'とこりりの商品一覧ページです。'
        const image = this.products[0].productImages[0].apiPath ? this.products[0].productImages[0].apiPath : ''
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
                {
                    hid: 'twitter:image',
                    property: 'twitter:image',
                    content: image,
                },
                {
                    hid: 'twitter:title',
                    property: 'twitter:title',
                    content: title,
                },
                {
                    hid: 'twitter:description',
                    property: 'twitter:description',
                    content: description,
                },
            ],
        }
    }

    mounted() {
        this.isSearchAreaDisplay = window.innerWidth > 600
    }

    toggleSearchArea() {
        this.isSearchAreaDisplay = !this.isSearchAreaDisplay
    }

    clickHandler(item: IProduct) {
        this.$router.push(`/product/${item.uuid}`)
    }

    async selectSearchHandler() {
        await this.loadingProduct()
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

.search-button
    justify-content start
    margin 8px 0
    width 100%
    .icon-text
        height 24px
        line-height 24px

.search-area
    margin-bottom 8px
    min-height 56px
    +sm()
        min-height inherit
    &__content
        display flex
        flex-wrap wrap

.sm
    display none
    +sm()
        display block
</style>
