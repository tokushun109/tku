<template>
    <v-icon :color="color" :large="large" :small="small" :x-large="xLarge" :x-small="xSmall" :disabled="disabled" @click.stop="$emit('c-click')">
        {{ getIcon }}
    </v-icon>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { IconType, TColorType } from '~/types'

@Component
export default class CIcon extends Vue {
    // アイコンの種類
    @Prop({ type: String, default: IconType.New.name }) type!: string
    // 色の種類
    @Prop({ type: String }) color?: TColorType
    // サイズの種類
    @Prop({ type: Boolean, default: false }) large!: boolean
    @Prop({ type: Boolean, default: false }) small!: boolean
    @Prop({ type: Boolean, default: false }) xLarge!: boolean
    @Prop({ type: Boolean, default: false }) xSmall!: boolean
    // 非活性な状態か
    @Prop({ type: Boolean, default: false }) disabled!: boolean
    // アイコンの取得
    get getIcon(): string {
        let icon = ''
        for (const type in IconType) {
            if (this.type === IconType[type].name) {
                icon = IconType[type].icon
            }
        }
        return icon
    }
}
</script>

<style lang="stylus" scoped></style>
