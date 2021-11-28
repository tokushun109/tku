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
                <input id="file-input" ref="preview" type="file" multiple @change="uploadFile($event)" />
                <c-button primary @c-click="selectFile">{{ 'ファイルを選択' }}</c-button>
            </div>
        </div>
        <c-input-label v-if="previewList.length > 0" label="プレビュー">
            <ul class="preview">
                <li v-for="(previewUrl, index) in previewList" :key="index" class="preview-item">
                    <img :src="previewUrl" :alt="`preview${index}`" class="preview-item-image" />
                    <img src="/icon/preview_close.png" alt="プレビューを削除" class="preview-item-close" @click="deletePreviewItem(index)" />
                </li>
            </ul>
        </c-input-label>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'

@Component
export default class CFileUpload extends Vue {
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
        this.addPreviewList()
        await this.$emit('c-file-uploaded', this.files)
    }

    previewList: Array<string> = []
    // プレビューリスト
    addPreviewList() {
        for (const file of this.files) {
            const url = URL.createObjectURL(file)
            this.previewList.push(url)
        }
    }

    // プレビュー画像を削除する
    deletePreviewItem(index: number) {
        this.previewList.splice(index, 1)
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
    .preview
        display flex
        border 1px dashed $light-dark-color
        border-radius 3px
        text-align center
        &-item
            position relative
            padding 15px
            width 100%
            &-image
                width 100%
                object-fit cover
                aspect-ratio 4 / 3
            &-close
                position absolute
                top 20px
                right 20px
    .drag-drop-area
        margin-bottom 16px
        padding 4px
        border 1px solid $light-dark-color
        border-radius 3px
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
