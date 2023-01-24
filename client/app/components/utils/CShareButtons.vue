<template>
    <v-sheet :color="ColorType.Transparent" class="c-share-buttons">
        <v-btn
            class="share-button"
            :color="backgroundColor"
            :href="`https://twitter.com/share?url=${url}&text=${text}`"
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
            :href="`https://www.facebook.com/share.php?u=${url}`"
            rel="nofollow noopener"
            target="_blank"
            fab
            depressed
            small
            :large="!small"
        >
            <v-icon :color="contentColor" :x-large="!small">{{ mdiFacebook }}</v-icon>
        </v-btn>
    </v-sheet>
</template>

<script lang="ts">
import { mdiFacebook, mdiTwitter } from '@mdi/js'
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType } from '~/types'

@Component
export default class CShareButtons extends Vue {
    @Prop({ type: String, default: 'https://tocoriri.com/' }) url!: string
    @Prop({ type: String, default: '' }) text!: string
    @Prop({ type: Boolean, default: true }) light!: boolean
    @Prop({ type: Boolean, default: true }) small!: boolean

    ColorType: typeof ColorType = ColorType
    mdiTwitter = mdiTwitter
    mdiFacebook = mdiFacebook

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
    .share-button
        margin 0 10px
</style>
