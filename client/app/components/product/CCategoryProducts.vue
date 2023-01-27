<template>
    <v-sheet v-if="categoryProducts.products.length > 0" class="c-category-product">
        <div class="category-name-container">
            <h2 class="category-name">{{ categoryProducts.category.name }}</h2>
        </div>
        <v-list class="product-list">
            <v-row>
                <v-col
                    v-for="(product, index) in categoryProducts.products"
                    :key="product.uuid"
                    :class="{ 'no-display': index >= 4 && !isExpand }"
                    cols="6"
                    md="3"
                >
                    <transition name="fadeDown">
                        <div v-if="index < 4 || isExpand">
                            <div class="product-list__container">
                                <nuxt-link :to="`/product/${product.uuid}`">
                                    <v-img
                                        v-if="product.productImages.length > 0"
                                        class="product-list__container__image"
                                        lazy-src="/img/product/gray-image.png"
                                        :src="product.productImages[0].apiPath"
                                        :alt="product.productImages[0].name"
                                    />
                                    <v-img
                                        v-else
                                        class="product-list__container__image"
                                        lazy-src="/img/product/gray-image.png"
                                        src="/img/product/no-image.png"
                                        alt="no-image"
                                    />
                                </nuxt-link>
                                <div class="default">
                                    <v-chip
                                        v-if="product.target.name"
                                        :color="ColorType.Secondary"
                                        :text-color="ColorType.White"
                                        class="product-list__container__target"
                                    >
                                        {{ product.target.name }}
                                    </v-chip>
                                </div>
                                <div class="sm">
                                    <v-chip
                                        v-if="product.target.name"
                                        :color="ColorType.Secondary"
                                        :text-color="ColorType.White"
                                        class="product-list__container__target"
                                        x-small
                                    >
                                        {{ product.target.name }}
                                    </v-chip>
                                </div>
                            </div>
                            <div class="product-list__footer">
                                <div class="product-list__footer__name">{{ product.name }}</div>
                                <div v-if="product.price" :color="ColorType.Accent" :text-color="ColorType.White" class="product-list__footer__price">
                                    ￥{{ product.price | priceFormat }}<span class="text-caption">(税込)</span>
                                </div>
                            </div>
                        </div>
                    </transition>
                </v-col>
            </v-row>
            <div v-if="categoryProducts.products.length > 3" class="product-list__button" @click="expandHandler">
                <c-detail-button v-if="!isExpand" class="product-list__button__content" to="/product" fall-down content="もっと見る" :link="false" />
            </div>
        </v-list>
    </v-sheet>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType, ICategoryProducts } from '~/types'

@Component
export default class CCategoryProduct extends Vue {
    @Prop({ type: Object, default: () => {} }) categoryProducts!: ICategoryProducts

    ColorType: typeof ColorType = ColorType
    isExpand: boolean = false

    expandHandler() {
        this.isExpand = !this.isExpand
    }
}
</script>

<style lang="stylus" scoped>
::v-deep .v-chip__content
    display inline-block !important
    overflow hidden
    height auto !important
    text-overflow ellipsis
    white-space nowrap

.c-category-product
    padding 16px 0
    .category-name-container
        margin 16px 0
        +sm()
            margin 8px 0
        .category-name
            color $text-color
            text-align left
            font-size 30px
            +sm()
                font-size 20px
    .product-list
        padding-bottom 24px
        +sm()
            16px
        &__container
            position relative
            &__image
                width 100%
                border-radius $image-border-radius
                aspect-ratio 1 / 1
                object-fit cover
            &__target
                position absolute
                top 8px
                left 8px
                max-width 70%
        &__footer
            text-align center
            &__name
                overflow hidden
                padding 4px 0
                color $text-color
                text-overflow ellipsis
                white-space nowrap
                font-size $font-large
                +sm()
                    font-size $font-medium
            &__price
                color $primary-text-color
                font-size $font-xlarge
                +sm()
                    font-size $font-large
        &__button
            margin-top 16px
            text-align center

.default
    +sm()
        display none

.sm
    display none
    +sm()
        display block

.no-display
    display none

.fadeDown-enter-active
    animation fadeDown 1s
</style>
