<template>
    <c-page>
        <div class="admin-product-list">
            <c-button primary @c-click="toggle">新規追加</c-button>
            <c-product-edit :visible.sync="dialogVisible" :model.sync="productModel" @close="toggle" @create="loadingProduct()" />
            <ul v-for="product in products" :key="product.uuid">
                <c-column>
                    <li>{{ product.name }}</li>
                    <div class="button-wrapper">
                        <c-button class="delete-button" label="削除" @c-click="productDeleteHandler(product)" />
                    </div>
                </c-column>
            </ul>
        </div>
    </c-page>
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
export default class PageAdminProductIndex extends Vue {
    products: Array<IProduct> = []
    // modalの表示切り替え
    dialogVisible: boolean = false
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

    async loadingProduct() {
        this.products = await this.$axios.$get(`/product`)
        this.productModel = newProduct()
    }

    // ボタンの切り替え
    toggle() {
        this.dialogVisible = !this.dialogVisible
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

<style lang="stylus"></style>
