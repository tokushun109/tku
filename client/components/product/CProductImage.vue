<template>
    <v-sheet>
        <v-carousel v-if="product.productImages.length > 0" :show-arrows="product.productImages.length > 1" height="auto" hide-delimiters>
            <v-carousel-item v-for="image in product.productImages" :key="image.uuid">
                <v-img :src="image.apiPath" :alt="image.uuid" />
            </v-carousel-item>
        </v-carousel>
        <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters>
            <v-carousel-item>
                <v-img src="/img/product/no-image.png" />
            </v-carousel-item>
        </v-carousel>
    </v-sheet>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IProduct } from '~/types'

@Component({})
export default class CProductImage extends Vue {
    // 商品画像のリスト
    @Prop({ type: Object }) product!: IProduct
    // 商品名を画像上に表示する
    @Prop({ type: Boolean, default: false }) title!: boolean
    // カテゴリーを画像上に表示する
    @Prop({ type: Boolean, default: false }) category!: boolean
}
</script>

<style lang="stylus">
.v-image
    aspect-ratio 16 / 9
</style>
