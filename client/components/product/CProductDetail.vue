<template>
    <v-main class="c-product-detail-page grey lighten-4">
        <v-container>
            <v-sheet class="pa-4 lighten-4">
                <p class="title-weight green--text text--darken-3">{{ product.name }}</p>
                <c-product-image class="mb-4" category :product="product" />
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
            </v-sheet>
        </v-container>
        <div class="purchase-button">
            <v-btn :color="ColorType.Orange" :disabled="product.salesSites.length === 0" x-large fab @click="dialogVisible = true">
                <c-icon :type="IconType.Cart.name" :color="ColorType.White" @c-click="dialogVisible = true" />
            </v-btn>
        </div>
        <c-dialog :visible.sync="dialogVisible" title="ご購入はこちらから" :is-button="false">
            <template #content>
                <div v-for="site in product.salesSites" :key="site.uuid">
                    <v-btn :color="ColorType.LightGreen" :href="site.url" class="my-4 white--text" block>
                        {{ site.name }}
                    </v-btn>
                </div>
            </template>
        </c-dialog>
    </v-main>
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
    .purchase-button
        position absolute
        right 20px
        bottom 20px
</style>
