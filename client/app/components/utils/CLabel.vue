<template>
    <div class="c-label" :class="{ required: required }">
        <label class="c-label-content">
            {{ label }}
            <slot />
        </label>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'

@Component
export default class CLabel extends Vue {
    @Prop({ default: '' }) label!: String
    @Prop({ default: false }) required!: boolean
}
</script>

<style lang="stylus" scoped>
.c-label
    position relative
    text-align center
    font-size $font-normal
    &.required
        .c-label-content:after
            margin-left 2px
            color $c-label-required-color
            content $c-label-required-content
            font-weight $c-label-required-weight
            font-size $c-label-required-size
</style>
