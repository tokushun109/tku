<template>
    <section class="form-layout" :class="{ bordered: bordered, inline: inline, slim: slim }">
        <h3 v-if="title && title.length > 0" class="form-title">{{ title }}</h3>
        <div class="form-content">
            <slot />
        </div>
    </section>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
@Component
export default class NsForm extends Vue {
    // フォームタイトル
    @Prop({ type: String, default: null }) title?: string
    // ボーダーをつけるか
    @Prop(Boolean) bordered!: boolean
    // スリムフォーム(タブレットサイズ)
    @Prop(Boolean) slim!: boolean
    // インラインレイアウトか否か
    @Prop(Boolean) inline?: boolean
}
</script>

<style lang="stylus">
.form-layout
    position: relative
    padding: $form-padding
    &.bordered
        border: 1px solid $form-bordered-border-color
        border-radius: $form-bordered-border-radius
        background-color: $white-color
        padding: $form-bordered-padding
    .form-title
        position: relative
        font-size: $font-large
        font-weight: 400
        text-align: $form-title-text-align
        if $form-title-underline
            padding-bottom: 8px
        margin-bottom: $form-title-margin-bottom
        if $form-title-underline
            &:after
                content: ''
                position: absolute
                display: block
                if $form-title-text-align == 'left'
                    left: 0
                else if $form-title-text-align == 'center'
                    left: 50%
                    transform: translateX(-50%)
                else if $form-title-text-align == 'right'
                    right: 0
                bottom: 0
                width: $form-title-underline-width
                border-bottom: 3px solid $form-title-underline-color
    .form-content
        > *
            margin-bottom: $form-item-margin
            &:last-child
                margin-bottom: 0
        .form-actions
            text-align: center
            margin-top: 32px
            margin-bottom: 16px
            > *
                margin-left: 8px
                &.first-child
                    margin-left: 0
    &.inline
        .form-content
            .c-labeled-item
                display: flex
                justify-content: space-between
                .c-labeled-item-label
                    flex: 0 0 $form-item-inline-label-width
                    padding-top: 4px
                    margin-bottom: 0
                .c-labeled-item-content
                    flex: 1 1 auto
    &.slim
        width: $tab-width
        margin: 0 auto
        +tablet()
            width: 80%
</style>
