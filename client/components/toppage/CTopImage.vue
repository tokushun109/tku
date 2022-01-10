<template>
    <v-card rounded elevation="20">
        <v-container class="carousel-wrapper">
            <v-carousel v-if="displayImages.length > 0" :show-arrows="carouselItems.length > 1" cycle hide-delimiters height="auto">
                <v-carousel-item
                    v-for="(image, index) in displayImages"
                    :key="index"
                    class="carousel-item-wrapper"
                    nuxt
                    :to="`/product/${displayProduct(index).uuid}`"
                >
                    <v-chip
                        v-if="category && displayProduct(index).category.uuid"
                        class="category-label"
                        small
                        :color="ColorType.LightGreen"
                        :text-color="ColorType.White"
                        >{{ displayProduct(index).category.name }}</v-chip
                    >
                    <v-img :src="image" :alt="`image-${index}`" />
                </v-carousel-item>
            </v-carousel>
            <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters>
                <v-carousel-item>
                    <v-img src="/img/product/no-image.png" />
                </v-carousel-item>
            </v-carousel>
        </v-container>
    </v-card>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { ColorType, ICarouselItem, IProduct } from '~/types'

@Component({})
export default class CTopImage extends Vue {
    ColorType: typeof ColorType = ColorType
    // 商品のリスト
    @Prop({ type: Object }) carouselItems!: Array<ICarouselItem>
    // 商品名を画像上に表示する
    @Prop({ type: Boolean, default: false }) title!: boolean
    // カテゴリーを画像上に表示する
    @Prop({ type: Boolean, default: false }) category!: boolean

    get displayImages(): Array<string> {
        return this.carouselItems.map((i) => {
            return i.apiPath
        })
    }

    displayProduct(index: number): IProduct {
        return this.carouselItems[index].product
    }
}
</script>

<style lang="stylus">
.carousel-wrapper
    .carousel-item-wrapper
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
