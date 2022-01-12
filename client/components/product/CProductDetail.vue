<template>
    <v-sheet class="c-product-detail-page grey lighten-4">
        <v-container class="c-product-detail-page grey lighten-4">
            <v-sheet class="pa-4 lighten-4">
                <p class="title-weight green--text text--darken-3">{{ product.name }}</p>
                <v-row>
                    <v-col cols="12" sm="6">
                        <c-product-image class="mb-4" category :product="product" />
                    </v-col>
                    <v-col cols="12" sm="6">
                        <p class="title-weight grey--text text--light-1 text-h6">
                            ￥{{ product.price | priceFormat }}
                            <span class="text-body-2">(税込)</span>
                        </p>
                        <pre style="white-space: pre-wrap" class="mb-4 text-body-2">{{ product.description }}</pre>
                        <p v-if="product.tags.length > 0" class="title-weight">関連タグ<v-divider /></p>
                        <div class="d-flex">
                            <div v-for="tag in product.tags" :key="tag.uuid">
                                <v-chip small :color="ColorType.Lime" :text-color="ColorType.White" class="mx-1">{{ tag.name }}</v-chip>
                            </div>
                        </div>
                        <div class="site-buttons-area">
                            <div class="text-body-1 mt-4">
                                <p class="title-weight">販売サイト<v-divider /></p>
                            </div>
                            <v-row>
                                <v-col v-for="siteDetail in product.siteDetails" :key="siteDetail.uuid" md="6" sm="12">
                                    <v-btn
                                        :color="ColorType.LightGreen"
                                        :href="siteDetail.url"
                                        class="white--text text-h6"
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
        </v-container>
        <div class="purchase-button">
            <v-btn :color="ColorType.Orange" :disabled="product.siteDetails.length === 0" x-large fab @click="dialogVisible = true">
                <c-icon :type="IconType.Cart.name" :color="ColorType.White" @c-click="dialogVisible = true" />
            </v-btn>
        </div>
        <c-dialog :visible.sync="dialogVisible" :is-button="false">
            <template #content>
                <div class="text-body-1">
                    <p class="title-weight">以下のサイトで販売中です！<v-divider /></p>
                </div>
                <c-container>
                    <v-row>
                        <v-col v-for="siteDetail in product.siteDetails" :key="siteDetail.uuid" cols="12" md="6">
                            <v-btn
                                :color="ColorType.LightGreen"
                                :href="siteDetail.url"
                                class="white--text text-h6 site-modal"
                                target="_blank"
                                rel="noopener noreferrer"
                                block
                            >
                                {{ siteDetail.salesSite.name }}
                            </v-btn>
                        </v-col>
                    </v-row>
                </c-container>
            </template>
        </c-dialog>
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

    // ダイアログの表示
    dialogVisible: boolean = false
}
</script>

<style lang="stylus" scoped>
.c-product-detail-page
    position relative
    .site-buttons-area
        +sm()
            display none
    .purchase-button
        position fixed
        right 20px
        bottom 50px
        display none
        +sm()
            display block
</style>
