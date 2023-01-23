<template>
    <v-container class="c-message">
        <v-sheet :color="color" :class="getClass" rounded class="c-message__container">
            <slot />
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ColorType } from '~/types'
@Component
export default class CMessage extends Vue {
    @Prop({ type: String, default: ColorType.White }) color!: typeof ColorType
    @Prop({ type: Boolean, default: false }) wide!: boolean

    ColorType: typeof ColorType = ColorType

    get getClass() {
        return {
            wide: this.wide,
        }
    }
}
</script>

<style lang="stylus" scoped>
.c-message
    &__container
        padding 32px
        color $primary
        text-align center
        font-weight $title-font-weight
        font-size $font-xlarge
        +sm()
            font-size $font-normal
        &.wide
            margin 32px 0
            padding 64px
</style>
