<template>
    <v-sheet class="carousel-wrapper">
        <v-chip v-if="category && product.category.uuid" class="category-label" small :color="ColorType.LightGreen" :text-color="ColorType.White">{{
            product.category.name
        }}</v-chip>
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
import { ColorType, IProduct } from '~/types'

@Component({})
export default class CProductImage extends Vue {
    ColorType: typeof ColorType = ColorType
    // 商品画像のリスト
    @Prop({ type: Object }) product!: IProduct
    // 商品名を画像上に表示する
    @Prop({ type: Boolean, default: false }) title!: boolean
    // カテゴリーを画像上に表示する
    @Prop({ type: Boolean, default: false }) category!: boolean
}
</script>

<style lang="stylus">
.carousel-wrapper
    position relative
    .category-label
        position absolute
        top 10px
        right 5px
        z-index 5
        opacity 0.8
    .v-image
        aspect-ratio 16 / 9
</style>
