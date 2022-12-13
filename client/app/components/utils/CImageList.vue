<template>
    <v-container v-if="registeredList.length > 0 || previewList.length > 0" class="c-image-list mb-4">
        <p>※画像をクリックで表示順を設定</p>
        <v-row>
            <v-col
                v-for="(registeredPathOrder, index) in registeredList"
                :key="registeredPathOrder.path"
                class="registered-item"
                cols="6"
                sm="4"
                md="3"
            >
                <div class="image-wrapper">
                    <v-chip v-if="registeredPathOrder.order" :color="ColorType.Orange" :text-color="ColorType.White" class="order-number">
                        {{ displayOrder(registeredPathOrder.order) }}
                    </v-chip>
                    <c-icon
                        class="close-icon"
                        :type="IconType.Close.name"
                        :color="ColorType.Red"
                        @c-click="$emit('c-delete-image-handler', index, ImageType.Registered)"
                    />
                    <v-img
                        :src="registeredPathOrder.path"
                        :alt="`registered${index}`"
                        class="registered-item-image"
                        @click="$emit('c-order-image-handler', index, ImageType.Registered)"
                    />
                </div>
            </v-col>
            <v-col
                v-for="(previewPathOrder, index) in previewList"
                :key="previewPathOrder.path"
                class="preview-item lighten-3"
                cols="6"
                sm="4"
                md="3"
            >
                <div class="image-wrapper">
                    <v-chip v-if="previewPathOrder.order" :color="ColorType.Orange" :text-color="ColorType.White" class="order-number">
                        {{ displayOrder(previewPathOrder.order) }}
                    </v-chip>
                    <c-icon
                        class="close-icon"
                        :type="IconType.Close.name"
                        :color="ColorType.Red"
                        @c-click="$emit('c-delete-image-handler', index, ImageType.Preview)"
                    />
                    <v-img
                        :src="previewPathOrder.path"
                        :alt="`preview${index}`"
                        class="preview-item-image"
                        @click="$emit('c-order-image-handler', index, ImageType.Preview)"
                    />
                </div>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType, IconType, IImagePathOrder, ImageType } from '~/types'

@Component
export default class CImageList extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType
    ImageType: typeof ImageType = ImageType

    @Prop({ type: Array, default: [] }) registeredList!: Array<IImagePathOrder>
    @Prop({ type: Array, default: [] }) previewList!: Array<IImagePathOrder>
    @Prop({ type: Number, default: 100 }) maxOrder!: number

    displayOrder(imageOrder: number): number {
        return this.maxOrder + 1 - imageOrder
    }
}
</script>

<style lang="stylus" scoped>
.c-image-list
    border 1px dashed $light-dark-color
    border-radius $image-border-radius
    .image-wrapper
        position relative
        .registered-item-image
        .preview-item-image
            width 100%
            aspect-ratio 1 / 1
            object-fit cover
        .close-icon
            position absolute
            top 5px
            right 10px
            z-index 10
        .order-number
            position absolute
            top 5px
            left 10px
            z-index 10
</style>
