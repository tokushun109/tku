<template>
    <v-sheet :color="ColorType.Transparent" class="mx-auto page-top text-center">
        <v-chip :color="ColorType.White" class="font-weight-bold text-subtitle-2 text-sm-h6 grey--text lighten-1 my-8">
            コットンレースのマクラメアクセサリー
        </v-chip>
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

<style lang="stylus" scoped></style>
