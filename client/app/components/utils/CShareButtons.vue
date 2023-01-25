<template>
    <v-sheet :color="ColorType.Primary" class="c-share-buttons">
        <div class="c-share-buttons__message">
            <v-icon :color="ColorType.White">{{ mdiShare }}</v-icon>
            <span class="content">Share This Page♪</span>
        </div>
        <div class="c-share-buttons__area">
            <v-btn
                class="share-button"
                :color="backgroundColor"
                :href="`https://twitter.com/share?url=${shareUrl}`"
                rel="nofollow"
                target="_blank"
                fab
                depressed
                small
                :large="!small"
            >
                <v-icon :color="contentColor" :x-large="!small">{{ mdiTwitter }}</v-icon>
            </v-btn>
            <v-btn
                class="share-button"
                :color="backgroundColor"
                :href="`https://www.facebook.com/share.php?u=${shareUrl}`"
                rel="nofollow noopener"
                target="_blank"
                fab
                depressed
                small
                :large="!small"
            >
                <v-icon :color="contentColor" :x-large="!small">{{ mdiFacebook }}</v-icon>
            </v-btn>
        </div>
    </v-sheet>
</template>

<script lang="ts">
import { mdiFacebook, mdiShare, mdiTwitter } from '@mdi/js'
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType } from '~/types'

@Component
export default class CShareButtons extends Vue {
    @Prop({ type: String, default: '' }) url!: string
    @Prop({ type: Boolean, default: true }) light!: boolean
    @Prop({ type: Boolean, default: true }) small!: boolean

    ColorType: typeof ColorType = ColorType
    mdiTwitter = mdiTwitter
    mdiFacebook = mdiFacebook
    mdiShare = mdiShare

    get shareUrl(): string {
        let shareUrl = this.url
        if (!shareUrl) {
            shareUrl = process.env.DOMAIN_URL + this.$route.path
        }
        return shareUrl
    }

    get backgroundColor(): string {
        return this.light ? ColorType.White : ColorType.Primary
    }

    get contentColor(): string {
        return this.light ? ColorType.Primary : ColorType.White
    }
}
</script>

<style lang="stylus" scoped>
.c-share-buttons
    text-align center
    &__message
        .content
            color $white-color
            font-size 12px
    &__area
        padding 10px
        .share-button
            margin 0 10px
</style>
