<template>
    <div class="c-file-upload">
        <div
            class="drag-drop-area"
            :class="{ hovering: hovering }"
            @dragover.prevent="dragoverHandler"
            @dragleave.prevent="dragleaveHandler"
            @drop.prevent="dropHandler"
        >
            <div class="drag-drop-inside">
                <p v-dompurify-html="'ファイルを選択、もしくは<br />ドロップ&ドラッグでアップロードする'" class="drag-drop-message"></p>
                <input id="file-input" type="file" multiple @change="uploadFile($event)" />
                <c-button primary @c-click="selectFile">{{ 'ファイルを選択' }}</c-button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'

@Component
export default class CTeamFileUpload extends Vue {
    // アップロードしたファイルのリスト
    files: Array<File> = []
    // 無効状態か否か
    @Prop(Boolean) disabled?: boolean
    // 許可しているファイル種別
    @Prop({ type: String, default: 'image/png,image/jpeg,image/gif' }) fileTypes?: string

    selectFile() {
        const input: HTMLInputElement | null = document.querySelector('#file-input')
        if (input) {
            input.click()
        }
    }

    // ファイルアップロード
    async uploadFile(event: any) {
        event.preventDefault()
        this.files = event.target.files ? event.target.files : event.dataTransfer.files
        if (this.files.length === 0) {
            return null
        }
        await this.$emit('c-file-uploaded', this.files)
    }

    hovering: boolean = false
    dragoverHandler(_event: DragEvent) {
        this.hovering = true
    }

    dragleaveHandler(_event: DragEvent) {
        this.hovering = false
    }

    dropHandler(event: DragEvent) {
        this.hovering = false
        this.uploadFile(event)
    }
}
</script>

<style lang="stylus">
.c-file-upload
    input[type=file]
        display none
    .drag-drop-area
        border 1px solid $light-dark-color
        padding 4px
        text-align center
        .drag-drop-inside
            padding 36px 18px
            .drag-drop-message
                margin-bottom 8px
        &.hovering
            background-color #ff9
    @media all and (-ms-high-contrast none)
        *::-ms-backdrop
        .drag-drop-area
            width 100%
            height 100%
</style>
