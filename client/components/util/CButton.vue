<template>
    <button class="c-button c-form-item" :class="getClass" :disabled="disabled" @click="click">
        {{ label }}
        <slot />
    </button>
</template>

<script lang="ts">
import { Component, Vue, Prop, Emit } from 'nuxt-property-decorator'

@Component({})
export default class CButton extends Vue {
    // カラー
    @Prop({ type: Boolean, default: false }) primary!: boolean
    // サイズ
    @Prop({ type: Boolean, default: false }) tiny!: boolean
    @Prop({ type: Boolean, default: false }) small!: boolean
    @Prop({ type: Boolean, default: false }) large!: boolean
    // 位置
    @Prop({ type: Boolean, default: false }) leftTop!: boolean
    @Prop({ type: Boolean, default: false }) leftBottom!: boolean
    @Prop({ type: Boolean, default: false }) rightTop!: boolean
    @Prop({ type: Boolean, default: false }) rightBottom!: boolean

    @Prop({ type: Boolean, default: false }) block!: boolean

    @Prop({ type: String }) label?: string
    @Prop({ type: Boolean, default: false }) disabled!: boolean

    @Emit('c-click')
    private click() {}

    get getClass() {
        return {
            primary: this.primary,
            tiny: this.tiny,
            small: this.small,
            large: this.large,
            block: this.block,
            'left-top': this.leftTop,
            'left-bottom': this.leftBottom,
            'right-top': this.rightTop,
            'right-bottom': this.rightBottom,
        }
    }
}
</script>

<style lang="stylus" scoped>
.c-button
    padding 4px 24px
    border none
    border 1px solid $primary-color
    border-radius 20px
    background-color $white-color
    color $primary-color
    text-decoration none
    white-space nowrap
    font-weight 400
    font-size 16px
    cursor pointer
    .inner
        display flex
        align-items center
        margin 0
        padding 0
        max-width auto
    &.primary
        border-color $primary-color
        background $primary-color
        color $white-color
    &.tiny
        padding 4px 12px
        border-radius 4px
        font-size 11px
    &.small
        padding 6px 16px
        border-radius 4px
        font-size 13px
    &.large
        padding 8px 28px
        border-radius 8px
        font-size 24px
    &.block
        display block
        width 100%
    &:disabled
        opacity 0.6
        cursor not-allowed
    &.left-top
        position absolute
        top 0
        left 0
    &.left-bottom
        position absolute
        bottom 0
        left 0
    &.right-top
        position absolute
        top 0
        right 0
    &.right-bottom
        position absolute
        right 0
        bottom 0
    &:hover
        filter brightness(0.95)
</style>
