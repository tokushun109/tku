<template>
    <div normal class="page-product">
        <div class="page-title-container">
            <h2 class="page-title">Product</h2>
        </div>
        <v-sheet class="page-product__container">
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
            <template v-if="isExistProducts">
                <div v-for="categoryProducts in categoryProductsList" :key="categoryProducts.category.uuid" class="fade-up">
                    <c-category-products :category-products="categoryProducts" />
                </div>
            </template>
            <c-message v-else wide>該当する商品が<br class="sm" />見つかりませんでした</c-message>
        </v-sheet>
        <c-breadcrumbs :items="breadCrumbs" />
    </div>
</template>

<script lang="ts">
import { mdiMagnify } from '@mdi/js'
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { newProduct } from '~/methods'
import { ColorType, IBreadCrumb, ICategoryProducts, IClassification, IGetClassificationParams, IGetProductsParams, IProduct } from '~/types'
@Component({})
export default class PageProductIndex extends Vue {
    ColorType: typeof ColorType = ColorType

    categoryProductsList: Array<ICategoryProducts> = []
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
            const categoryProductsList: Array<ICategoryProducts> = await app.$axios.$get(`/category/product`, { params: productParams })
            const classificationParams: IGetClassificationParams = {
                mode: 'used',
            }
            const categories: Array<IProduct> = await app.$axios.$get(`/category`, { params: classificationParams })
            const targets: Array<IProduct> = await app.$axios.$get(`/target`, { params: classificationParams })
            return { categoryProductsList, categories, targets }
        } catch (e) {
            return { categoryProductsList: [], categories: [], targets: [] }
        }
    }

    async loadingProduct() {
        this.categoryProductsList = await this.$axios.$get(`/category/product`, { params: this.productParams })
    }

    get isExistProducts(): boolean {
        let result = false
        for (const categoryProducts of this.categoryProductsList) {
            if (categoryProducts.products.length > 0) {
                result = true
                break
            }
        }
        return result
    }

    head() {
        const title = '商品一覧 | とこりり'
        const description = 'とこりりの商品一覧ページです。'
        const image = '/img/about/story.jpg'
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
    &__container
        min-height 56vh

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
