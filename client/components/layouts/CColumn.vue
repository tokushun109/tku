<template>
    <section class="column-layout" :class="getClass">
        <slot />
    </section>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
@Component({})
export default class CColumn extends Vue {
    // カラム数
    @Prop({ type: Number, default: 2 }) column!: number
    // 横位置
    @Prop(Boolean) left!: boolean
    @Prop(Boolean) right!: boolean
    // 縦位置
    @Prop(Boolean) center!: boolean
    @Prop(Boolean) baseline!: boolean

    get getClass() {
        const c = {
            left: this.left,
            right: this.right,
            center: this.center,
            baseline: this.baseline,
            [`column-${this.column}`]: true,
        }
        return c
    }
}
</script>

<style lang="stylus">
.column-layout
    +tablet()
        display: block
        width: 100%
        margin: auto
    flex-box(flex-start, row nowrap)
    > *
        margin-right: $column-spacing
        +tablet()
            margin-right: 0
        &:last-child
            margin-right: 0
    for num in (1..12)
        &.column-{num} > *
            flex: '1 1 calc((100% / %s) - (%s))' % (num $tile-spacing)
            &:nth-child({num}n)
                margin-right: 0
        *::-ms-backdrop, &.column-{num} > *
            width: 'calc((100% / %s) - (%s))' % (num $tile-spacing)
    &.left > *
        flex: 0 1 auto
    &.right
        flex-box(flex-end, row nowrap)
        > *
            flex: 0 1 auto
    &.center
        align-items: center
    &.baseline
        align-items: baseline
</style>
