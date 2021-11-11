<template>
    <v-icon @click="$emit('c-click')">{{ getIcon }}</v-icon>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

interface IIconType {
    [key: string]: { name: string; icon: string }
}

const IconType: IIconType = {
    New: { name: 'new', icon: 'mdi-note-plus' },
    Edit: { name: 'edit', icon: 'mdi-pencil' },
    Delete: { name: 'delete', icon: 'mdi-delete' },
} as const

@Component
export default class CMessage extends Vue {
    // アイコンの種類
    @Prop({ type: String, default: IconType.New.name }) type!: string
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
