<template>
    <v-sheet :color="ColorType.Transparent" class="page-top">
        <c-top-image class="top-image" title category :carousel-items="carouselItems" />
        <c-layout-container normal>
            <div class="page-top-container">
                <div class="about-section">
                    <h2 class="page-top-title">About</h2>
                    <v-container class="about-section__message default">
                        <p>仕事も育児も頑張る女性の日常に寄り添うアクセサリーを創りたい。</p>
                        <p>
                            そんな想いで
                            <span class="emphasis">オシャレなだけじゃない、あったらちょっと嬉しい作品</span> を制作・販売しております。
                        </p>
                    </v-container>
                    <v-container class="about-section__message sm">
                        <p>仕事も育児も頑張る女性の日常に寄り添う<br />アクセサリーを創りたい。</p>
                        <p>そんな想いで</p>
                        <p><span class="emphasis">オシャレなだけじゃない</span><br /><span class="emphasis">あったらちょっと嬉しい作品</span></p>
                        <p>を制作・販売しております。</p>
                    </v-container>
                </div>
            </div>
        </c-layout-container>
    </v-sheet>
</template>

<script lang="ts">
import { mdiArrowRightThick } from '@mdi/js'
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ColorType, ICarouselItem, ICreator } from '~/types'
@Component({})
export default class PageTop extends Vue {
    mdiArrowRightThick = mdiArrowRightThick
    ColorType: typeof ColorType = ColorType

    carouselItems: Array<ICarouselItem> = []
    creator: ICreator | null = null
    async asyncData({ app }: Context) {
        try {
            const creator: ICreator = await app.$axios.$get(`/creator`)
            const carouselItems: Array<ICarouselItem> = await app.$axios.$get(`/carousel_image`)

            return { creator, carouselItems }
        } catch (e) {
            return { creator: null, carouselItems: [] }
        }
    }

    head() {
        if (!this.carouselItems) {
            return
        }
        const title = 'アクセサリーショップ とこりり'
        const description = 'マクラメ編みのアクセサリーショップ【とこりり】の紹介サイトです。'
        const image = this.creator && this.creator.apiPath ? this.creator.apiPath : ''
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
        +sm()
            margin-bottom 40px
    .page-top-title
        margin-bottom 10px
        color $primary
        text-align center
        font-size 60px !important
        font-family $title-font-face !important
        +sm()
            font-size 40px !important
    .about-section
        &__message
            color $text-color
            font-weight 800
            &.default
                +sm()
                    display none
            &.sm
                display none
                +sm()
                    display block
            .emphasis
                color $accent
</style>
