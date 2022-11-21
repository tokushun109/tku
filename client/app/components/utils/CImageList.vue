<template>
    <v-container v-if="registeredList.length > 0 || previewList.length > 0" class="c-image-list mb-4">
        <v-row>
            <v-col v-for="(registeredPath, index) in registeredList" :key="registeredPath" class="registered-item" cols="6" sm="4" md="3">
                <div class="image-wrapper">
                    <c-icon
                        class="close-icon"
                        :type="IconType.Close.name"
                        :color="ColorType.Red"
                        @c-click="$emit('c-delete-image-handler', index, ImageType.Registered)"
                    />
                    <v-img :src="registeredPath" :alt="`registered${index}`" class="registered-item-image" />
                </div>
            </v-col>
            <v-col v-for="(previewPath, index) in previewList" :key="previewPath" class="preview-item lighten-3" cols="6" sm="4" md="3">
                <div class="image-wrapper">
                    <c-icon
                        class="close-icon"
                        :type="IconType.Close.name"
                        :color="ColorType.Red"
                        @c-click="$emit('c-delete-image-handler', index, ImageType.Preview)"
                    />
                    <v-img :src="previewPath" :alt="`preview${index}`" class="preview-item-image" />
                </div>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType, IconType, ImageType } from '~/types'

@Component
export default class CImageList extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType
    ImageType: typeof ImageType = ImageType

    @Prop({ type: Array, default: [] }) registeredList!: Array<string>
    @Prop({ type: Array, default: [] }) previewList!: Array<string>
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
            right 5px
            z-index 10
</style>