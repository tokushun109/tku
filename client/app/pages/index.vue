<template>
    <v-sheet :color="ColorType.Transparent" class="page-top">
        <c-top-image class="top-image" title category :carousel-items="carouselItems" />
        <div class="more">
            <v-btn rounded outlined x-large :color="ColorType.Grey" nuxt to="/product">
                <div class="text-h6">MORE</div>
                <v-icon large>{{ mdiArrowRightThick }}</v-icon>
            </v-btn>
        </div>
    </v-sheet>
</template>

<script lang="ts">
import { mdiArrowRightThick } from '@mdi/js'
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ColorType, ICarouselItem } from '~/types'
@Component({})
export default class PageTop extends Vue {
    mdiArrowRightThick = mdiArrowRightThick
    ColorType: typeof ColorType = ColorType

    carouselItems: Array<ICarouselItem> = []
    async asyncData({ app }: Context) {
        try {
            const carouselItems: Array<ICarouselItem> = await app.$axios.$get(`/carousel_image`)

            return { carouselItems }
        } catch (e) {
            return { carouselItems: [] }
        }
    }

    head() {
        if (!this.carouselItems) {
            return
        }
        const title = 'トップページ | tocoriri'
        const description = 'マクラメ編みのアクセサリーショップtocoriri(とこりり)の紹介サイトです。'
        const image = this.carouselItems[0].apiPath
        return {
            title,
            meta: [
                {
                    hid: 'description',
                    name: 'description',
                    content: description,
                },
                {
                    hid: 'og:title',
                    property: 'og:title',
                    content: title,
                },
                {
                    hid: 'og:description',
                    property: 'og:description',
                    content: description,
                },
                {
                    hid: 'og:type',
                    property: 'og:type',
                    content: 'website',
                },
                {
                    hid: 'og:image',
                    property: 'og:image',
                    content: image,
                },
            ],
        }
    }
}
</script>

<style lang="stylus" scoped>
.page-top
    margin 0 auto
    text-align center
    +sm()
        padding-top 35px
    .site-sub-title
        display none
        padding 30px 40px
        color $text-color
        &.sm
            +sm()
                display block
    .top-image
        margin-bottom 16px
    .more
        margin-right 16px
        text-align right
        +sm()
            display none
</style>
