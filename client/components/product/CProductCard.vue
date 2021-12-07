<template>
    <v-card width="100%" :color="Color">
        <v-card-text>
            <div class="my-4">
                <div class="d-flex">
                    <div class="green--text text--darken-3">
                        {{ listItem.name }}
                        <v-chip v-if="!listItem.isActive" x-small :color="ColorType.Grey" :text-color="ColorType.White">展示</v-chip>
                    </div>
                    <v-spacer />
                    <div>
                        <c-icon :type="IconType.Edit.name" @c-click="$emit('c-open', ExecutionType.Edit, listItem)" />
                        <c-icon :type="IconType.Delete.name" @c-click="$emit('c-open', ExecutionType.Delete, listItem)" />
                    </div>
                </div>
                <v-divider />
            </div>
            <v-carousel v-if="listItem.productImages.length > 0" :show-arrows="listItem.productImages.length > 1" height="auto" hide-delimiters>
                <v-carousel-item v-for="image in listItem.productImages" :key="image.uuid">
                    <v-img :src="image.apiPath" :alt="image.uuid" />
                </v-carousel-item>
            </v-carousel>
            <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters>
                <v-carousel-item>
                    <v-img src="/img/product/no-image.png" />
                </v-carousel-item>
            </v-carousel>
            <div class="text-right">
                <p>{{ listItem.price | priceFormat }}円(税込)</p>
            </div>
        </v-card-text>
    </v-card>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import _ from 'lodash'
import { ColorType, ExecutionType, IconType, IProduct, TExecutionType } from '~/types'
@Component({})
export default class CProductCard extends Vue {
    @Prop({ type: Object }) listItem!: IProduct
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType
    ExecutionType: typeof ExecutionType = ExecutionType
    executionType: TExecutionType = ExecutionType.Create

    get Color(): string {
        return this.listItem.isActive ? 'light-green lighten-5' : 'grey lighten-3'
    }
}
</script>

<style lang="stylus">
.v-image
    aspect-ratio 16 / 9
</style>
