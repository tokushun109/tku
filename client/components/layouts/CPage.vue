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
@Component({})
export default class CPage extends Vue {
    // ページタイトル
    @Prop({ type: String, default: null }) title?: string
}
</script>

<style lang="stylus">
.page-layout
    position relative
    max-width $wide-width
    min-height calc(100vh - 103px - 140px)
    margin 0 auto
    padding $page-padding
    +sp()
        padding $page-sp-padding
    .page-title
        position relative
        font-size $font-xxlarge
        font-weight 400
        text-align $page-title-text-align
        if $page-title-underline
            padding-bottom 8px
        margin-bottom $page-title-margin-bottom
        if $page-title-underline
            &after
                content ''
                position absolute
                display block
                left 50%
                if $page-title-text-align == 'left'
                    left 0
                else if $page-title-text-align == 'center'
                    left 50%
                    transform translateX(-50%)
                else if $page-title-text-align == 'right'
                    right 0
                width $page-title-underline-width
                bottom 0
                border-bottom 3px solid $page-title-underline-color
    .page-content
        .section-layout
            margin-bottom $page-content-section-margin-bottom
        .form-layout
            margin-bottom $page-content-section-margin-bottom
        .panel-layout
            margin-bottom $page-content-section-margin-bottom
</style>
