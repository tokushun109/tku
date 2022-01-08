<template>
    <c-product-list :items="products" :categories="categories" :tags="tags" :sales-sites="salesSites" @c-change="loadingProduct" />
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IClassification, IProduct, ISite } from '~/types'
@Component({
    head: {
        title: '商品一覧',
    },
})
export default class PageAdminProductIndex extends Vue {
    products: Array<IProduct> = []

    categories: Array<IClassification> = []
    tags: Array<IClassification> = []
    salesSites: Array<ISite> = []

    // 新規作成ダイアログの表示
    createDialogVisible: boolean = false
    async asyncData({ app }: Context) {
        try {
            const products = await app.$axios.$get(`/product`)
            const categories = await app.$axios.$get(`/category`)
            const tags = await app.$axios.$get(`/tag`)
            const salesSites = await app.$axios.$get(`/sales_site`)
            return { products, categories, tags, salesSites }
        } catch (e) {
            return { products: [], categories: [], tags: [], salesSites: [] }
        }
    }

    async loadingProduct() {
        this.products = await this.$axios.$get(`/product`)
    }

    // 商品の削除
    async productDeleteHandler(product: IProduct) {
        if (confirm(`${product.name}を削除します。よろしいですか？`)) {
            await this.$axios.$delete(`/product/${product.uuid}`)
            this.loadingProduct()
        }
    }
}
</script>

<style lang="stylus" scoped></style>
