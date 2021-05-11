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
    // カラーリング
    @Prop({ type: Boolean, default: false }) primary!: boolean
    @Prop({ type: Boolean, default: false }) secondary!: boolean
    @Prop({ type: Boolean, default: false }) light!: boolean
    @Prop({ type: Boolean, default: false }) dark!: boolean
    @Prop({ type: Boolean, default: false }) info!: boolean
    @Prop({ type: Boolean, default: false }) success!: boolean
    @Prop({ type: Boolean, default: false }) warning!: boolean
    @Prop({ type: Boolean, default: false }) danger!: boolean
    @Prop({ type: Boolean, default: false }) orange!: boolean
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
            secondary: this.secondary,
            light: this.light,
            dark: this.dark,
            info: this.info,
            success: this.success,
            warning: this.warning,
            danger: this.danger,
            orange: this.orange,
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
    border-radius 4px
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
    &.secondary
        border-color $primary-color
        background-color $white-color
        color $primary-color
    &.light
        border-color $light-bg-colorr
        background-color $light-bg-colorr
        color $light-color
    &.dark
        border-color $dark-bg-color
        background-color $dark-bg-color
        color $dark-color
    &.info
        border-color $info-bg-color
        background-color $info-bg-color
        color $info-color
    &.success
        border-color $primary-color
        background $primary-color
        color $white-color
    &.warning
        border-color $warning-color
        background-color $warning-color
        color $white-color
    &.danger
        border-color $danger-color
        background-color $danger-color
        color $white-color
    &.orange
        border-color $orange-color
        background-color $orange-color
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
