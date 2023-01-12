<template>
    <v-sheet class="carousel-wrapper">
        <v-chip v-if="category && product.category.uuid" class="category-label" small :color="ColorType.LightGreen" :text-color="ColorType.White">{{
            product.category.name
        }}</v-chip>
        <template v-if="!listDisplay">
            <v-carousel
                v-if="product.productImages.length > 0"
                :show-arrows="showArrows && product.productImages.length > 1"
                height="auto"
                hide-delimiters
                class="carousel-area"
            >
                <template v-if="product.productImages.length > 1" #prev="{ on, attrs }">
                    <v-icon large v-bind="attrs" v-on="on">{{ mdiChevronLeft }}</v-icon>
                </template>
                <template v-if="product.productImages.length > 1" #next="{ on, attrs }">
                    <v-icon large v-bind="attrs" v-on="on">{{ mdiChevronRight }}</v-icon>
                </template>
                <v-carousel-item v-for="image in product.productImages" :key="image.uuid">
                    <v-img :src="image.apiPath" :alt="image.uuid" class="carousel-image" />
                </v-carousel-item>
            </v-carousel>
            <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters class="carousel-area">
                <v-carousel-item>
                    <v-img src="/img/product/no-image.png" alt="no-image" class="carousel-image" />
                </v-carousel-item>
            </v-carousel>
        </template>
        <template v-else>
            <div class="display-area">
                <v-img v-if="product.productImages.length > 1" :src="selectImage.apiPath" :alt="selectImage.uuid" class="carousel-image" />
                <v-img v-else src="/img/product/no-image.png" alt="no-image" class="carousel-image" />
            </div>
            <v-list>
                <v-row dense>
                    <v-col v-for="(productImage, index) in product.productImages" :key="productImage.uuid" cols="2">
                        <v-img
                            :src="productImage.apiPath"
                            :alt="productImage.uuid"
                            class="carousel-image list-display"
                            :class="{ 'is-select': isSelect(index) }"
                            @click="setSelectIndex(index)"
                        />
                    </v-col>
                </v-row>
            </v-list>
        </template>
    </v-sheet>
</template>

<script lang="ts">
import { mdiChevronLeft, mdiChevronRight } from '@mdi/js'
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { ColorType, IProduct, IProductImage } from '~/types'

@Component({})
export default class CProductImage extends Vue {
    mdiChevronLeft = mdiChevronLeft
    mdiChevronRight = mdiChevronRight
    ColorType: typeof ColorType = ColorType
    // 商品画像のリスト
    @Prop({ type: Object }) product!: IProduct
    // 商品名を画像上に表示する
    @Prop({ type: Boolean, default: false }) title!: boolean
    // カテゴリーを画像上に表示する
    @Prop({ type: Boolean, default: false }) category!: boolean
    // スライド用の矢印を表示する
    @Prop({ type: Boolean, default: true }) showArrows!: boolean
    // 登録画像をリスト形式で表示
    @Prop({ type: Boolean, default: false }) listDisplay!: boolean

    selectIndex: number = 0

    setSelectIndex(index: number) {
        this.selectIndex = index
    }

    isSelect(index: number): boolean {
        return this.selectIndex === index
    }

    get selectImage(): IProductImage {
        return this.product.productImages[this.selectIndex]
    }
}
</script>

<style lang="stylus" scoped>
.carousel-wrapper
    position relative
    border-radius $image-border-radius
    .category-label
        position absolute
        top 10px
        right 5px
        z-index 5
        opacity 0.8
    .carousel-area
        border-radius $image-border-radius
    .carousel-image
        width 100%
        aspect-ratio 1 / 1
        object-fit cover
        &.list-display
            border solid 0.5px
        &.is-select
            border solid red 2px
    .display-area
        margin-bottom 8px
        border-radius $image-border-radius
</style>
