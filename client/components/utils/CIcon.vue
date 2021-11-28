<template>
    <v-icon :color="color" @click="$emit('c-click')">{{ getIcon }}</v-icon>
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
