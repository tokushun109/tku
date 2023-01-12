<template>
    <v-container class="c-layout-container" :class="getClass">
        <slot />
    </v-container>
</template>

<script lang="ts">
import { Prop } from 'nuxt-property-decorator'
import { Component, Vue } from 'vue-property-decorator'
@Component
export default class CLayoutContainer extends Vue {
    // サイズ
    @Prop({ type: Boolean, default: false }) narrow!: boolean
    @Prop({ type: Boolean, default: false }) normal!: boolean

    get getClass() {
        return {
            narrow: this.narrow,
            normal: this.normal,
        }
    }
}
</script>

<style lang="stylus">
.c-layout-container
    padding $page-padding
    +sm()
        padding $page-sp-padding
    &.normalとい
        max-width 1200px
    &.narrow
        max-width 800px
</style>
