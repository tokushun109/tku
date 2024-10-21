<template>
    <v-card
        width="100%"
        :class="{ 'is-active': listItem.isActive, 'is-recommend': listItem.isActive && listItem.isRecommend }"
        hover
        class="c-product-card"
        @click.native="$emit('c-click', ExecutionType.Detail, listItem)"
    >
        <v-card-text class="product-card-wrapper">
            <div class="product-card-header">
                <div class="product-name">
                    {{ listItem.name }}
                </div>
                <div class="product-status">
                    <v-chip v-if="admin && listItem.isRecommend" x-small :color="ColorType.Accent" :text-color="ColorType.White">おすすめ</v-chip>
                    <v-chip v-if="!listItem.isActive" x-small :color="ColorType.Grey" :text-color="ColorType.White">展示</v-chip>
                </div>
            </div>
            <div class="product-card-image-container">
                <v-img
                    v-if="listItem.productImages.length > 0"
                    class="product-card-image"
                    lazy-src="/img/product/gray-image.png"
                    :src="listItem.productImages[0].apiPath"
                    :alt="listItem.productImages[0].name"
                />
                <v-img v-else class="product-card-image" lazy-src="/img/product/gray-image.png" src="/img/product/no-image.png" alt="no-image" />
                <div v-if="listItem.category.uuid" class="product-category">
                    <v-chip :color="ColorType.Accent" :text-color="ColorType.White" x-small>
                        {{ listItem.category.name }}
                    </v-chip>
                </div>
                <div v-if="listItem.target.uuid" class="product-target">
                    <v-chip :color="ColorType.Accent" :text-color="ColorType.White" x-small>
                        {{ listItem.target.name }}
                    </v-chip>
                </div>
            </div>
            <div class="product-card-footer">
                <div class="product-card-footer-content">
                    <template v-if="admin">
                        <c-icon :type="IconType.Edit.name" @c-click="$emit('c-click', ExecutionType.Edit, listItem)" />
                        <c-icon :type="IconType.Delete.name" @c-click="$emit('c-click', ExecutionType.Delete, listItem)" />
                    </template>
                    <v-spacer />
                    <div class="price">￥{{ listItem.price | priceFormat }}<span class="text-caption">税込</span></div>
                </div>
            </div>
        </v-card-text>
    </v-card>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType, ExecutionType, IconType, IProduct, TExecutionType } from '~/types'
@Component({})
export default class CProductCard extends Vue {
    @Prop({ type: Object }) listItem!: IProduct
    @Prop({ type: Boolean, default: false }) admin!: boolean

    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType
    ExecutionType: typeof ExecutionType = ExecutionType
    executionType: TExecutionType = ExecutionType.Create
}
</script>

<style lang="stylus" scoped>
.c-product-card
    background-color $light-dark-color
    &.is-active
        background-color $accent-light-color
    &.is-recommend
        background-color $bright-color
    .product-card-wrapper
        .product-card-header
            margin 5px 0
            .product-name
                overflow hidden
                padding-bottom 10px
                color $title-primary-color
                text-overflow ellipsis
                white-space nowrap
                font-weight $title-font-weight
                font-size $font-xlarge
            .product-status
                height 24px
                display flex
                align-items center
                gap 8px
                justify-content flex-end
        .product-card-image-container
            position relative
            .product-category
                position absolute
                top 5px
                left 10px
            .product-target
                position absolute
                right 10px
                bottom 5px
            .product-card-image
                width 100%
                border-radius $image-border-radius
                aspect-ratio 1 / 1
                object-fit cover
        .product-card-footer
            margin-top 15px
            text-align right
            .product-card-footer-content
                display flex
                .price
                    color $title-primary-color
                    font-weight 600
                    font-size $font-xlarge
</style>
