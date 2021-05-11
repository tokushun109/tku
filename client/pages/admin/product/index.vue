<template>
    <c-page>
        <div>
            <c-button primary @c-click="modalVisible = !modalVisible">新規追加</c-button>
            <c-product-edit :visible.sync="modalVisible" :model.sync="editProductModel" @close="modalVisible = !modalVisible" />
            <ul v-for="product in products" :key="product.uuid">
                <li>{{ product }}</li>
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
export default class PageAdminProductIndex extends Vue {
    products: Array<IProduct> = []
    // modalの表示切り替え
    modalVisible: boolean = false
    editProductModel: IProduct | null = null
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
