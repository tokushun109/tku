<template>
    <v-sheet class="c-product-detail-page-area">
        <p class="product-name">{{ product.name }}</p>
        <v-row class="detail-area">
            <v-col cols="12" sm="6">
                <c-product-image class="mb-4" :show-arrows="false" list-display :product="product" />
            </v-col>
            <v-col cols="12" sm="6">
                <div class="price-area">
                    <p class="price">
                        ￥{{ product.price | priceFormat }}
                        <span class="text-body-2">(税込)</span>
                    </p>
                </div>
                <div class="description-area">
                    <pre class="description text-body-2">{{ product.description }}</pre>
                </div>
                <div v-if="product.category.name" class="category-area">
                    <p class="category">
                        カテゴリー<client-only><v-divider /></client-only>
                    </p>

                    <div class="category-content">
                        <v-chip small :color="ColorType.Secondary" :text-color="ColorType.White" class="mx-1">{{ product.category.name }}</v-chip>
                    </div>
                </div>
                <div v-if="product.tags.length > 0" class="tag-area">
                    <p v class="tag">
                        関連タグ<client-only><v-divider /></client-only>
                    </p>

                    <div class="tag-content">
                        <div v-for="tag in product.tags" :key="tag.uuid">
                            <v-chip small :color="ColorType.Secondary" :text-color="ColorType.White" class="mx-1">{{ tag.name }}</v-chip>
                        </div>
                    </div>
                </div>
                <div v-if="product.siteDetails.length > 0" class="sales-site-area">
                    <div class="sales-site text-body-1">
                        <p>
                            販売サイト<client-only><v-divider /></client-only>
                        </p>
                    </div>
                    <v-row>
                        <v-col v-for="siteDetail in product.siteDetails" :key="siteDetail.uuid" md="6" sm="12" cols="12">
                            <v-btn
                                :color="ColorType.Accent"
                                :href="siteDetail.detailUrl"
                                class="site-buttons"
                                target="_blank"
                                rel="noopener noreferrer"
                                block
                            >
                                {{ siteDetail.salesSite.name }}
                            </v-btn>
                        </v-col>
                    </v-row>
                </div>
            </v-col>
        </v-row>
    </v-sheet>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType, IconType, IProduct } from '~/types'
@Component({})
export default class CProductDetail extends Vue {
    @Prop({ type: Object }) product!: IProduct
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType
}
</script>

<style lang="stylus" scoped>
.c-product-detail-page-area
    padding 16px
    .product-name
        color $title-primary-color
        font-weight $title-font-weight
        font-size 35px
        +sm()
            font-size $font-xlarge
    .detail-area
        .price-area
            text-align right
            .price
                margin-right auto
                color $title-primary-color
                font-weight $title-font-weight
                font-size 45px
                +sm()
                    font-size 30px
        .description-area
            .description
                margin-bottom 16px
                white-space pre-wrap
        .category-area
            .category
                font-weight $title-font-weight
            .category-content
                display flex
        .tag-area
            .tag
                margin-top 16px
                font-weight $title-font-weight
            .tag-content
                display flex
        .sales-site-area
            .sales-site
                margin-top 16px
                font-weight $title-font-weight
            .site-buttons
                color $white-color
                font-size $font-large
</style>
