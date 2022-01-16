<template>
    <div class="c-top-image">
        <!-- md幅以上での表示 -->
        <div class="default carousel-wrapper">
            <v-btn fab x-large class="arrow-btn left" @click="leftSlide"><v-icon>mdi-arrow-left</v-icon></v-btn>
            <v-btn fab x-large class="arrow-btn right" @click="rightSlide"><v-icon>mdi-arrow-right</v-icon></v-btn>
            <v-row>
                <v-col v-if="leftCarouselItem" cols="6" md="4">
                    <v-img :src="leftCarouselItem.apiPath" class="carousel-image" />
                </v-col>
                <v-col v-if="centerCarouselItem" cols="6" md="4">
                    <v-img :src="centerCarouselItem.apiPath" class="carousel-image" />
                </v-col>
                <v-col v-if="rightCarouselItem" cols="6" md="4">
                    <v-img :src="rightCarouselItem.apiPath" class="carousel-image md" />
                </v-col>
            </v-row>
        </div>
        <!-- sm幅以下での表示 -->
        <v-card class="sm" rounded elevation="20">
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
                        <v-chip v-if="title" class="title-label" small :color="ColorType.Orange" :text-color="ColorType.White">
                            {{ displayProduct(index).name }}
                        </v-chip>
                        <v-img :src="image" :alt="`image-${index}`" class="carousel-image" />
                    </v-carousel-item>
                </v-carousel>
                <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters>
                    <v-carousel-item class="carousel-item-wrapper">
                        <v-img src="/img/product/no-image.png" alt="no-image" class="carousel-image" />
                    </v-carousel-item>
                </v-carousel>
            </v-container>
        </v-card>
    </div>
</template>

<script lang="ts">
import _ from 'lodash'
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { ColorType, ICarouselItem, IProduct } from '~/types'

@Component({})
export default class CTopImage extends Vue {
    ColorType: typeof ColorType = ColorType
    // 商品のリスト
    @Prop({ type: Array }) carouselItems!: Array<ICarouselItem>
    // 商品名を画像上に表示する
    @Prop({ type: Boolean, default: false }) title!: boolean
    // カテゴリーを画像上に表示する
    @Prop({ type: Boolean, default: false }) category!: boolean

    intervalId: number = 0
    mounted() {
        this.intervalId = window.setInterval(this.leftSlide, 6000)
    }

    // md幅以上での処理
    sortCarouselItems: Array<ICarouselItem> = _.cloneDeep(this.carouselItems)

    // 右の矢印を押した時の並び替え
    rightSlide() {
        // インターバルを解除
        clearInterval(this.intervalId)
        // 先頭の要素を除去
        const firstItem: ICarouselItem = this.sortCarouselItems.shift() as ICarouselItem
        // 除去した要素を末尾に追加
        this.sortCarouselItems.push(firstItem)
        // インターバルを開始
        this.intervalId = window.setInterval(this.rightSlide, 6000)
    }

    // 左の画像を押した時の並び替え
    leftSlide() {
        // インターバルを解除
        clearInterval(this.intervalId)
        // 末尾の要素を除去
        const lasItem: ICarouselItem = this.sortCarouselItems.pop() as ICarouselItem
        // 除去した要素を先頭に追加
        this.sortCarouselItems.unshift(lasItem)
        // インターバルを開始
        this.intervalId = window.setInterval(this.rightSlide, 6000)
    }

    // 中心のアイテム
    get centerCarouselItem(): ICarouselItem | null {
        return this.sortCarouselItems.length > 0 ? this.sortCarouselItems[0] : null
    }

    // 左のアイテム
    get leftCarouselItem(): ICarouselItem | null {
        return this.sortCarouselItems.length > 1 ? this.sortCarouselItems[this.sortCarouselItems.length - 1] : null
    }

    // 右のアイテム
    get rightCarouselItem(): ICarouselItem | null {
        return this.sortCarouselItems.length > 2 ? this.sortCarouselItems[1] : null
    }

    // sm幅以下での処理
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
    .carousel-wrapper
        position relative
        &.default
            .arrow-btn
                position absolute
                top 50%
                z-index 1
                opacity 0.8
                transform translateY(-50%)
                &.left
                    left 10px
                &.right
                    right 10px
            .carousel-image
                width 100%
                aspect-ratio 1 / 1
                object-fit cover
                &.md
                    +md()
                        display none
            +sm()
                display none
    .sm
        display none
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
</style>
