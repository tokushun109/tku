<template>
    <div>
        <nuxt-link v-if="link" :to="`/${to}`" class="c-detail-button" :class="getClass">
            <span>{{ content }}</span>
        </nuxt-link>
        <div v-else class="c-detail-button" :class="getClass">
            <span>{{ content }}</span>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'

@Component
export default class CDetailButton extends Vue {
    @Prop({ type: String, default: '/' }) to?: string
    @Prop({ type: Boolean, default: true }) link?: boolean
    @Prop({ type: String, default: '詳しくはこちら' }) content?: string
    @Prop({ type: Boolean, default: false }) fallDown?: boolean

    get getClass() {
        return {
            'fall-down': this.fallDown,
        }
    }
}
</script>

<style lang="stylus" scoped>
.c-detail-button
    cursor pointer

.fall-down
    position relative
    display inline-block
    overflow hidden
    padding 10px 30px
    outline none
    border 1px solid $secondary
    background-color $white-color
    color $primary
    text-decoration none
    font-weight $title-font-weight
    &::before
        position absolute
        top 0
        left 0
        z-index -1
        width 100%
        height 0
        background-color $primary
        content ''
        transition all 0.3s
        +md()
            transition none
    span
        z-index 2
        display block
        &::before
        &::after
            position absolute
            top 0
            width 2px
            height 0
            background $accent
            content ''
            transition none
        &::before
            left 0
        &::after
            right 0
    &:hover
        border-color transparent
        background $accent
        color $white-color
        transition all 0.3s
        transition-delay 0.6s
        +md()
            transition none
            transition-delay 0
        span
            &::before
            &::after
                height 100%
        &::before
            height 100%
            transition-delay 0.4s
            +md()
                transition-delay 0
</style>
