<template>
    <v-card width="100%" :color="Color" hover @click.native="$emit('c-click', ExecutionType.Detail, listItem)">
        <v-card-text>
            <div class="my-4">
                <div>
                    <div class="green--text text--darken-3">
                        {{ listItem.name }}
                        <v-chip v-if="!listItem.isActive" x-small :color="ColorType.Grey" :text-color="ColorType.White">展示</v-chip>
                    </div>
                </div>
                <v-divider />
            </div>
            <c-product-image :product="listItem" />
            <div class="text-right mt-2">
                <div class="d-flex">
                    <template v-if="admin">
                        <c-icon :type="IconType.Edit.name" @c-click="$emit('c-click', ExecutionType.Edit, listItem)" />
                        <c-icon :type="IconType.Delete.name" @c-click="$emit('c-click', ExecutionType.Delete, listItem)" />
                    </template>
                    <v-spacer />
                    <div class="text-h6">￥{{ listItem.price | priceFormat }}<span class="text-caption">税込</span></div>
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

    get Color(): string {
        return this.listItem.isActive ? 'light-green lighten-5' : 'grey lighten-3'
    }
}
</script>

<style lang="stylus" scoped></style>
