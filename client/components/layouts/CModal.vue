<template>
    <div v-if="syncVisible" class="c-modal">
        <div class="c-modal-bg" @click="cancelButton()" />
        <div class="c-modal-container" :style="{ width: width, height: height }">
            <div v-if="isHeader" class="c-modal-container-title">
                {{ title }}
            </div>
            <div
                class="c-modal-container-content"
                :class="{ 'c-modal-container-content-none-button': !isButton, 'c-modal-container-content-button': isButton }"
            >
                <slot />
            </div>
            <div v-if="isButton && !isOnlyClose" class="c-modal-container-bottom">
                <c-button class="c-modal-container-bottom-button" label="キャンセル" @c-click="cancelButton()" />
                <c-button
                    class="c-modal-container-bottom-button"
                    :disabled="confirmButtonDisabled"
                    :label="confirmButtonTitle"
                    primary
                    @c-click="confirmButton()"
                />
            </div>
            <div v-else-if="isOnlyClose" class="c-modal-container-bottom">
                <c-button class="c-modal-container-bottom-button" label="閉じる" @c-click="cancelButton()" />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, PropSync, Emit } from 'nuxt-property-decorator'
@Component
export default class CModal extends Vue {
    // モーダルの横幅(px)
    @Prop({ type: String, default: '512px' }) width!: string
    // モーダルの縦幅
    @Prop({ type: String, default: '90%' }) height!: string
    // 選択ボタンの有無
    @Prop({ default: true }) isButton!: boolean
    // 閉じるボタンの有無
    @Prop({ default: false }) isOnlyClose!: boolean
    // ヘッダーの有無
    @Prop({ default: true }) isHeader!: boolean
    // サブボタンのボタン名
    @Prop({ type: String, default: '' }) subConfirmButtonTitle?: string

    @PropSync('visible', { type: Boolean }) syncVisible!: boolean
    // モーダルタイトル
    @Prop(String) title!: string
    // 確定ボタン
    @Prop({ type: String, default: '確定' }) confirmButtonTitle?: string
    // 非活性の確定ボタン
    @Prop({ type: Boolean, default: false }) confirmButtonDisabled?: boolean
    // キャンセルイベント
    @Emit('close')
    private cancelButton() {}

    // 確定イベント
    @Emit('confirm')
    private confirmButton() {}
}
</script>

<style lang="stylus">
.c-modal
    position fixed
    top 0
    left 0
    z-index 999
    width 100%
    height 100%
    &-bg
        position fixed
        top 0
        left 0
        z-index 100
        width 100vw
        height 100vh
        background rgba(0, 0, 0, 0.7)
    &-container
        position absolute
        top 50%
        left 50%
        z-index 1000
        overflow hidden
        box-sizing border-box
        // overflow-y scroll
        padding 0
        max-width 95% !important
        max-height 95% !important
        border-radius 3px
        background #fff
        color #333
        text-align left
        transform translate(-50%, -50%)
        // 下部ボタン
        &-bottom
            display flex
            justify-content center
            padding-top 16px
            padding-bottom 16px
            &-button
                display block
                margin 0 8px
                padding 8px !important
                min-width 110px !important
        &-title
            margin 0 auto
            height 56px
            background $white-color
            color $primary-color
            text-align center
            font-weight bold
            font-size 16px
            line-height 56px
        &-content
            overflow-y scroll
            padding 24px
            height calc(100% - 56px - 48px)
            color #333
            &-none-button
                overflow auto
            &-button
                height calc(100% - 56px - 48px - 74px)
</style>
