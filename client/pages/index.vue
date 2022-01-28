<template>
    <v-sheet :color="ColorType.Transparent" class="mx-auto page-top text-center">
        <div class="sm site-sub-title text-h5 grey--text text--darken-1 py-10">Cotton lace × Macrame</div>
        <c-top-image class="mb-4" title category :carousel-items="carouselItems" />
        <div class="more text-right mr-4">
            <v-btn rounded outlined x-large :color="ColorType.Grey" nuxt to="/product">
                <div class="text-h6">MORE</div>
                <v-icon large>mdi-arrow-right-thick</v-icon>
            </v-btn>
        </div>
    </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ColorType, ICarouselItem } from '~/types'
@Component({
    head: {
        title: 'tku',
    },
})
export default class PageTop extends Vue {
    ColorType: typeof ColorType = ColorType

    carouselItems: Array<ICarouselItem> = []
    async asyncData({ app }: Context) {
        try {
            const carouselItems = await app.$axios.$get(`/carousel_image`)
            return { carouselItems }
        } catch (e) {
            return { carouselItems: [] }
        }
    }
}
</script>

<style lang="stylus" scoped>
.site-sub-title
    display none
    &.sm
        +sm()
            display block
            font-family 'Lobster' !important
</style>
