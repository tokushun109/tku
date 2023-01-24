<template>
    <v-sheet :color="ColorType.Transparent" class="page-top">
        <c-top-image class="top-image" title category :carousel-items="carouselItems" />
        <div class="page-top-container">
            <div class="about-section page-top-section">
                <c-layout-container normal no-vertical-padding>
                    <h2 class="page-top-title">About</h2>
                    <v-container class="about-section__message default">
                        <p>仕事や出産、育児、家事...</p>
                        <p>頑張る女性の味方になりたい、</p>
                        <p>そんな想いでマクラメ編みのアクセサリーを作っています。</p>
                    </v-container>
                    <v-container class="about-section__message sm">
                        <p>仕事や出産、育児、家事...</p>
                        <p>頑張る女性の味方になりたい、</p>
                        <p>そんな想いで</p>
                        <p>マクラメ編みのアクセサリーを作っています。</p>
                    </v-container>
                    <c-detail-button to="about" />
                </c-layout-container>
            </div>
            <div class="contact-section page-top-section reverse">
                <c-layout-container normal no-vertical-padding>
                    <h2 class="page-top-title reverse">Contact</h2>
                    <v-container class="contact-section__message">
                        <p>お問い合わせ・ご意見・ご相談はこちらから</p>
                    </v-container>
                    <c-detail-button content="お問い合わせフォーム" to="contact" />
                </c-layout-container>
            </div>
            <v-divider />
            <div class="share-section">
                <c-layout-container normal no-vertical-padding>
                    <c-share-buttons></c-share-buttons>
                </c-layout-container>
            </div>
        </div>
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
        const description = this.creator && this.creator.introduction ? this.creator.introduction : ''
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
                {
                    hid: 'twitter:image',
                    property: 'twitter:image',
                    content: image,
                },
                {
                    hid: 'twitter:title',
                    property: 'twitter:title',
                    content: title,
                },
                {
                    hid: 'twitter:description',
                    property: 'twitter:description',
                    content: description,
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
        padding-top 16px
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
            margin-bottom 0
    .page-top-container
        .page-top-section
            padding 48px 0 80px
            background-color $white-color
            &.reverse
                background-color $primary
            &.sub
                padding 8px 0
            .page-top-title
                margin-bottom 10px
                color $primary
                text-align center
                font-size 60px
                font-family $title-font-face !important
                +sm()
                    font-size 40px
                &.reverse
                    color $white-color
                &.sub
                    font-size 30px
                    +sm()
                        font-size 20px
        .about-section
            &__message
                padding 48px 0
                color $text-color
                &.default
                    +sm()
                        display none
                &.sm
                    display none
                    +sm()
                        display block
                .emphasis
                    color $accent
        .contact-section
            &__message
                padding 48px 0
                color $white-color !important
        .share-section
            padding-top 16px
            background-color $primary
</style>
