<template>
    <v-dialog v-model="dialogVisible" :width="width">
        <v-card>
            <v-card-title v-if="title" class="card-title text-h5 justify-center">{{ title }}</v-card-title>
            <v-card-text class="pt-5">
                <div class="title text-center"><slot name="content" /></div>
                <div v-if="isButton" class="text-center">
                    <v-btn color="primary" outlined @click="cancelButton">キャンセル</v-btn>
                    <v-btn v-if="subConfirmButtonTitle" color="secondary" @click="subConfirmButton">{{ subConfirmButtonTitle }}</v-btn>
                    <v-btn color="primary" :disabled="confirmButtonDisabled" @click="confirmButton">{{ confirmButtonTitle }}</v-btn>
                </div>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
import { Component, Vue, Prop, PropSync, Emit } from 'nuxt-property-decorator'
@Component
export default class CDialog extends Vue {
    // 表示フラグ
    @PropSync('visible', { type: Boolean }) dialogVisible!: boolean
    // モーダルの横幅(px)
    @Prop({ type: String, default: '800px' }) width!: string
    // 選択ボタンの有無
    @Prop({ default: true }) isButton!: boolean
    // タイトル
    @Prop({ type: String, default: '' }) title?: string
    // 確定ボタン
    @Prop({ type: String, default: '確定' }) confirmButtonTitle?: string
    // 非活性の確定ボタン
    @Prop({ type: Boolean, default: false }) confirmButtonDisabled?: boolean
    // 確定ボタン
    @Prop({ type: String, default: '' }) subConfirmButtonTitle!: string
    // 確定イベント
    @Emit('confirm')
    private confirmButton() {}

    // キャンセルイベント
    @Emit('close')
    private cancelButton() {}

    // サブボタンの確定イベント
    @Emit('sub-confirm')
    private subConfirmButton() {}
}
</script>

<style lang="stylus" scoped>
.card-title
    background-color $primary
    color $white-color
</style>
