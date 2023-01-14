<template>
    <v-card
        width="100%"
        :class="{ 'is-active': listItem.isActive }"
        hover
        class="c-product-card"
        @click.native="$emit('c-click', ExecutionType.Detail, listItem)"
    >
        <v-card-text class="product-card-wrapper">
            <div class="product-card-header">
                <div class="product-name">
                    {{ listItem.name }}
                    <v-chip v-if="!listItem.isActive" x-small :color="ColorType.Grey" :text-color="ColorType.White">展示</v-chip>
                </div>
            </div>
            <v-img
                class="product-card-image"
                lazy-src="/img/product/gray-image.png"
                :src="listItem.productImages[0].apiPath"
                :alt="listItem.productImages[0].name"
            />
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
    .product-card-wrapper
        .product-card-header
            margin 16px 0
            .product-name
                overflow hidden
                padding-bottom 10px
                color $title-primary-color
                text-overflow ellipsis
                white-space nowrap
                font-weight $title-font-weight
                font-size $font-xlarge
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
