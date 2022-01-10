<template>
    <v-container class="page-top text-center">
        <v-sheet :color="ColorType.Transparent" max-width="900" class="mx-auto">
            <v-chip :color="ColorType.White" class="font-weight-bold text-subtitle-2 text-sm-h6 grey--text lighten-1 mt-4 mb-10">
                コットンレースのマクラメアクセサリー
            </v-chip>
            <v-sheet class="light-green lighten-4">
                <c-top-image class="mb-4" category :carousel-items="carouselItems" />
            </v-sheet>
            <div class="more text-right">
                <v-btn rounded outlined x-large :color="ColorType.Grey" nuxt to="/product">
                    <div class="text-h6">MORE</div>
                    <v-icon large>mdi-arrow-right-thick</v-icon>
                </v-btn>
            </div>
        </v-sheet>
    </v-container>
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
