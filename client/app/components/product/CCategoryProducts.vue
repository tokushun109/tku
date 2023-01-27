<template>
    <v-sheet v-if="categoryProducts.products.length > 0" class="c-category-product">
        <div class="category-name-container">
            <h2 class="category-name">{{ categoryProducts.category.name }}</h2>
        </div>
        <v-list class="product-list">
            <v-row>
                <v-col v-for="(product, index) in categoryProducts.products" :key="product.uuid" cols="6" md="3">
                    <template v-if="index < 4 || isExpand">
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
                                <v-chip
                                    v-if="product.price"
                                    :color="ColorType.Accent"
                                    :text-color="ColorType.White"
                                    class="product-list__container__price"
                                >
                                    ￥{{ product.price | priceFormat }}<span class="text-caption">(税込)</span>
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
                                <v-chip
                                    v-if="product.price"
                                    :color="ColorType.Accent"
                                    :text-color="ColorType.White"
                                    class="product-list__container__price"
                                    x-small
                                >
                                    ￥{{ product.price | priceFormat }}<span class="text-caption">(税込)</span>
                                </v-chip>
                            </div>
                        </div>
                        <p class="product-list__name">{{ product.name }}</p>
                    </template>
                </v-col>
            </v-row>
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
    .category-name-container
        margin 16px 0
        .category-name
            color $secondary
            text-align left
            font-size 30px
            +sm()
                font-size 18px
    .product-list
        padding-bottom 24px
        +sm()
            16px
        &__name
            overflow hidden
            padding 4px 0
            color $text-color
            text-align center
            text-overflow ellipsis
            white-space nowrap
            font-size $font-large
            +sm()
                font-size $font-medium
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
            &__price
                position absolute
                right 8px
                bottom 8px

.default
    +sm()
        display none

.sm
    display none
    +sm()
        display block
</style>
