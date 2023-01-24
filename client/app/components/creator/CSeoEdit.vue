<template>
    <v-sheet class="c-creator-edit">
        <v-container class="logo-content">
            <h3 class="title product-list-title">ページサムネイル</h3>
            <v-divider class="divider"></v-divider>
            <v-img v-if="creator.apiPath" :src="creator.apiPath" max-width="400" aspect-ratio="1" class="logo-image" alt="ロゴ画像" />
        </v-container>
        <v-container class="description-content">
            <h3 class="title product-list-title">ページの説明</h3>
            <v-divider class="divider"></v-divider>
            <pre class="description-text">{{ creator.introduction }}</pre>
        </v-container>
        <div v-if="admin" class="edit-button">
            <v-btn color="primary" @click="openHandler"><c-icon :type="IconType.Edit.name" @c-click="openHandler" />編集</v-btn>
        </div>
        <c-dialog
            :visible.sync="dialogVisible"
            title="製作者の編集"
            width="800"
            :confirm-button-disabled="!valid"
            @confirm="saveHandler"
            @close="closeHandler"
        >
            <template #content>
                <c-error :errors.sync="errors" />
                <v-form ref="form" v-model="valid" lazy-validation>
                    <v-file-input v-model="uploadFile" label="ロゴ画像" :prepend-icon="mdiCamera" outlined />
                    <v-textarea v-model="creator.introduction" label="紹介文" outlined />
                </v-form>
            </template>
        </c-dialog>
        <c-notification :visible.sync="notificationVisible">製作者を更新しました</c-notification>
    </v-sheet>
</template>

<script lang="ts">
import { mdiCamera } from '@mdi/js'
import { Component, Vue, Watch, Prop, PropSync } from 'nuxt-property-decorator'
import _ from 'lodash'
import { ICreator, IError, ISite, IconType, ColorType } from '~/types'
import { newCreator } from '~/methods'
@Component({})
export default class CSeoEdit extends Vue {
    mdiCamera = mdiCamera
    IconType: typeof IconType = IconType
    ColorType: typeof ColorType = ColorType

    // 製作者
    @PropSync('item') creator!: ICreator
    // 管理画面での使用
    @Prop({ type: Boolean, default: false }) admin!: boolean

    // 最作者の初期情報
    initCreator: ICreator = newCreator()

    // アップロードするロゴファイル
    uploadFile: File | null = null

    // 販売サイトのリスト
    salesSites: Array<ISite> | null = []

    // 製作者編集ダイアログの表示
    dialogVisible: boolean = false

    // 通知の表示
    notificationVisible: boolean = false

    valid: boolean = true

    errors: Array<IError> = []

    async saveHandler() {
        try {
            this.errors = []
            await this.$axios.$put(`/creator`, this.creator, { withCredentials: true }).then(async () => {
                if (this.uploadFile) {
                    const params = new FormData()
                    params.append('logo', this.uploadFile)
                    await this.$axios.$put(`/creator/logo`, params, {
                        headers: {
                            'Content-Type': 'multipart/form-data',
                        },
                        withCredentials: true,
                    })
                }
            })
            this.initCreator = this.creator
            this.dialogVisible = false
            this.notificationVisible = true
            this.$emit('c-change')
        } catch (e) {
            this.errors.push(e.response)
        }
    }

    mounted() {
        // 製作者の初期情報
        this.initCreator = _.cloneDeep(this.creator)
    }

    setInit() {
        this.errors = []
        this.creator = _.cloneDeep(this.initCreator)
        this.uploadFile = null
    }

    @Watch('dialogVisible')
    resetValidation() {
        if (!this.dialogVisible) {
            const refs: any = this.$refs.form
            refs.resetValidation()
        }
    }

    openHandler() {
        this.setInit()
        this.dialogVisible = true
    }

    closeHandler() {
        this.setInit()
        this.dialogVisible = false
    }
}
</script>

<style lang="stylus" scoped>
.c-creator-edit
    .logo-content
        .logo-image
            margin 0 auto
            border-radius 50%
            object-fit cover
    .description-title
        color $title-primary-color
    .description-content
        .description-text
            padding 24px 0
            white-space pre-wrap
            word-break break-all
            +sm()
                font-size 3vw
    .head-title
        color $title-primary-color
        text-align center
        font-weight $title-font-weight
        font-size $font-xxlarge
        .head-title-content
            margin-top 20px
    .sns-content
        position relative
        display flex
        justify-content center
        .sns-item
            margin 0 10px
            text-align center
            .sns-icon
                position absolute
                top -23px
        .sns-name
            position absolute
            bottom 27px
            left 50%
            margin-top 10px
            color $white-color
            font-weight $title-font-weight
            font-size 12px
            transform translateX(-50%)
    .sales-site-content
        text-align center
        .sales-site-buttons
            margin 10px 0
            width 450px
            color $white-color
            font-size $font-large
            +md()
                width 45vw
            +sm()
                width 80%
    .edit-button
        margin-bottom 32px
        text-align center
</style>
