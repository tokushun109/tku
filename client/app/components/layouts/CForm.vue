<template>
    <section class="form-layout">
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
}
</script>

<style lang="stylus">
.form-layout
    position relative
    padding $form-padding
    .form-title
        position relative
        text-align $form-title-text-align
        font-weight 400
        font-size $font-large
        if $form-title-underline
            padding-bottom 8px
        margin-bottom $form-title-margin-bottom
        if $form-title-underline
            &:after
                position absolute
                bottom 0
                left 50%
                display block
                width $form-title-underline-width
                border-bottom 3px solid $form-title-underline-color
                content ''
                transform translateX(-50%)
    .form-content
        > *
            margin-bottom $form-item-margin
            &:last-child
                margin-bottom 0
        .form-actions
            margin-top 32px
            margin-bottom 16px
            text-align center
            > *
                margin-left 8px
                &.first-child
                    margin-left 0
</style>
