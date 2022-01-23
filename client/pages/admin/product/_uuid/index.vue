<template>
    <c-product-detail :product="product" />
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { newProduct } from '~/methods'
import { IProduct, ColorType } from '~/types'
@Component({
    head: {
        title: '商品詳細',
    },
})
export default class PageAdminProductDetail extends Vue {
    ColorType: typeof ColorType = ColorType
    product: IProduct = newProduct()
    async asyncData({ app, route }: Context) {
        try {
            const product = await app.$axios.$get(`/product/${route.params.uuid}`)
            return { product }
        } catch (e) {
            return { products: [] }
        }
    }
}
</script>

<style lang="stylus" scoped></style>
