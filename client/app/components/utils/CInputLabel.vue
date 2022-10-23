<template>
    <div class="c-input-label" :class="{ inline: inline, center: center }">
        <div v-if="label" class="c-input-label-name">
            <c-label :label="label" :required="required" />
        </div>
        <div class="c-input-label-content">
            <slot :disabled="disabled" />
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
@Component
export default class CInputLabel extends Vue {
    // label名
    @Prop(String) label!: string
    // 必須項目であるか
    @Prop(Boolean) required?: boolean
    // インラインレイアウトか否か
    @Prop(Boolean) inline?: boolean
    // 無効状態か否か
    @Prop(Boolean) disabled?: boolean
    // 中央揃え
    @Prop(Boolean) center?: boolean
}
</script>

<style lang="stylus">
.c-input-label
    .c-input-label-name
        margin-bottom 10px
        .c-label-content
            border-bottom 2px solid $form-bordered-border-color
            font-size $form-item-font-size
            padding 5px
    &.inline
        flex-box()
        +tablet()
            display block
        .c-input-label-name
            flex 0 0 $form-item-inline-label-width
            +tablet()
                width 100%
            margin-bottom 0
            padding-top 4px
        .c-input-label-content
            flex 1 1 auto
            +tablet()
                width 100%
    &.center
        text-align center
</style>
