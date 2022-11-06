<template>
    <section class="page-layout">
        <h1 v-if="title && title.length > 0" class="page-title">{{ title }}</h1>
        <div class="page-content">
            <slot />
        </div>
    </section>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
@Component
export default class CPage extends Vue {
    // ページタイトル
    @Prop({ type: String, default: null }) title?: string
}
</script>

<style lang="stylus">
.page-layout
    position relative
    margin 0 auto
    padding $page-padding
    min-height calc(100vh - 103px - 140px)
    max-width $wide-width
    +sm()
        padding $page-sp-padding
    .page-title
        position relative
        text-align $page-title-text-align
        font-weight 400
        font-size $font-xxlarge
        if $page-title-underline
            padding-bottom 8px
        margin-bottom $page-title-margin-bottom
        if $page-title-underline
            &after
                position absolute
                left 50%
                display block
                content ''
                if $page-title-text-align == 'left'
                    left 0
                else if $page-title-text-align == 'center'
                    left 50%
                    transform translateX(-50%)
                else if $page-title-text-align == 'right'
                    right 0
                bottom 0
                width $page-title-underline-width
                border-bottom 3px solid $page-title-underline-color
    .page-content
        .section-layout
            margin-bottom $page-content-section-margin-bottom
        .form-layout
            margin-bottom $page-content-section-margin-bottom
        .panel-layout
            margin-bottom $page-content-section-margin-bottom
</style>
