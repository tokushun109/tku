<template>
    <div class="c-top-image">
        <!-- md幅以上での表示 -->
        <v-sheet class="default">
            <div class="carousel-wrapper">
                <div v-for="(image, index) in displayImages" :key="index" class="slide-show">
                    <v-card class="carousel-item-wrapper" elevation="20" nuxt :to="`/product/${displayProduct(index).uuid}`">
                        <v-chip
                            v-if="category && displayProduct(index).category.uuid"
                            class="category-label"
                            small
                            :color="ColorType.Accent"
                            :text-color="ColorType.White"
                            >{{ displayProduct(index).category.name }}</v-chip
                        >
                        <v-chip v-if="title" class="title-label" small :color="ColorType.Accent" :text-color="ColorType.White">
                            {{ displayProduct(index).name }}
                        </v-chip>
                        <v-img eager lazy-src="/img/product/gray-image.png" :src="image" :alt="`image-${index}`" width="530" class="carousel-image" />
                    </v-card>
                </div>
                <div v-for="(image, index) in displayImages" :key="`${index}-sm`" class="slide-show">
                    <v-card class="carousel-item-wrapper" elevation="20" nuxt :to="`/product/${displayProduct(index).uuid}`">
                        <v-chip
                            v-if="category && displayProduct(index).category.uuid"
                            class="category-label"
                            small
                            :color="ColorType.Accent"
                            :text-color="ColorType.White"
                            >{{ displayProduct(index).category.name }}</v-chip
                        >
                        <v-chip v-if="title" class="title-label" small :color="ColorType.Accent" :text-color="ColorType.White">
                            {{ displayProduct(index).name }}
                        </v-chip>
                        <v-img eager lazy-src="/img/product/gray-image.png" :src="image" :alt="`image-${index}`" width="530" class="carousel-image" />
                    </v-card>
                </div>
            </div>
        </v-sheet>
        <!-- sm幅以下での表示 -->
        <v-sheet class="sm">
            <v-card rounded elevation="20">
                <v-container class="carousel-wrapper">
                    <v-carousel
                        v-if="displayImages && displayImages.length > 0"
                        :show-arrows="false"
                        cycle
                        :interval="3000"
                        hide-delimiters
                        height="auto"
                    >
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
                                :color="ColorType.Accent"
                                :text-color="ColorType.White"
                                >{{ displayProduct(index).category.name }}</v-chip
                            >
                            <v-chip v-if="title" class="title-label" small :color="ColorType.Accent" :text-color="ColorType.White">
                                {{ displayProduct(index).name }}
                            </v-chip>
                            <v-img :src="image" lazy-src="/img/product/gray-image.png" :alt="`image-${index}`" class="carousel-image" />
                        </v-carousel-item>
                    </v-carousel>
                    <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters>
                        <v-carousel-item class="carousel-item-wrapper">
                            <v-img src="/img/product/no-image.png" alt="no-image" class="carousel-image" />
                        </v-carousel-item>
                    </v-carousel>
                </v-container>
            </v-card>
        </v-sheet>
    </div>
</template>

<script lang="ts">
import { mdiChevronLeft, mdiChevronRight } from '@mdi/js'
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { ICarouselItem, IProduct, ColorType } from '~/types'

@Component({})
export default class CTopImage extends Vue {
    mdiChevronLeft = mdiChevronLeft
    mdiChevronRight = mdiChevronRight

    ColorType: typeof ColorType = ColorType

    // 商品のリスト
    @Prop({ type: Array }) carouselItems!: Array<ICarouselItem>
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

<style lang="stylus" scoped>
.c-top-image
    .default
        +sm()
            display none
        .carousel-wrapper
            display flex
            align-items center
            overflow hidden
            background-color $accent-light-color
            .slide-show
                display flex
                padding 20px
                animation loop-slide 40s infinite linear 1s both
                .carousel-item-wrapper
                    position relative
                    transition all 0.2s
                    &:hover
                        opacity 0.8
                        cursor pointer
                        transform translateY(-10px)
                    .title-label
                        position absolute
                        top 5px
                        left 5px
                        z-index 5
                        opacity 0.8
                    .category-label
                        position absolute
                        right 5px
                        bottom 5px
                        z-index 5
                        opacity 0.8
                    .carousel-image
                        width 100%
                        aspect-ratio 1 / 1
                        object-fit cover
                        &.sm
                            +sm()
                                display none
                    +sm()
                        display none
    .sm
        display none
        margin 0 16px
        +sm()
            display block
            .carousel-wrapper
                .carousel-item-wrapper
                    position relative
                    .title-label
                        position absolute
                        top 5px
                        left 5px
                        z-index 5
                        opacity 0.8
                    .category-label
                        position absolute
                        right 5px
                        bottom 5px
                        z-index 5
                        opacity 0.8
                    .carousel-image
                        width 100%
                        aspect-ratio 1 / 1
                        object-fit cover

@keyframes loop-slide
    from
        transform translateX(0)
    to
        transform translateX(calc(-1 * (530px + 40px) * 5))
</style>
