<template>
    <v-sheet :color="ColorType.Transparent" class="page-top">
        <div class="sm site-sub-title text-h5">Cotton lace × Macrame</div>
        <c-top-image class="top-image" title category :carousel-items="carouselItems" />
        <div class="more">
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
.page-top
    margin 0 auto
    text-align center
    .site-sub-title
        display none
        padding 0 40px
        color $text-color
        &.sm
            +sm()
                display block
    .top-image
        margin-bottom 16px
    .more
        margin-right 16px
        text-align right
</style>
