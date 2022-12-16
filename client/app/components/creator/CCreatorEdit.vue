<template>
    <v-sheet class="c-creator-edit">
        <v-container>
            <v-sheet class="creator-wrapper">
                <v-container class="logo-content">
                    <v-img v-if="creator.apiPath" :src="creator.apiPath" max-width="400" aspect-ratio="1" class="logo-image" alt="ロゴ画像" />
                </v-container>
                <v-container class="description-content">
                    <pre class="description-text">{{ creator.introduction }}</pre>
                </v-container>
                <v-container v-if="!admin" class="sns-content">
                    <div v-for="sns in snsList" :key="sns.name" class="sns-item">
                        <v-btn fab :color="ColorType.Orange" :href="sns.url" x-large>
                            <v-icon large :color="ColorType.White">{{ sns.icon }}</v-icon>
                        </v-btn>
                        <div class="sns-name">
                            <strong>{{ sns.name }}</strong>
                        </div>
                    </div>
                </v-container>
                <div v-if="admin" class="edit-button">
                    <v-btn color="primary" @click="openHandler"><c-icon :type="IconType.Edit.name" @c-click="openHandler" />編集</v-btn>
                </div>
            </v-sheet>
        </v-container>
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
import { mdiCamera, mdiInstagram } from '@mdi/js'
import { Component, Vue, Watch, Prop, PropSync } from 'nuxt-property-decorator'
import _ from 'lodash'
import { ICreator, IError, ISite, IconType, ColorType } from '~/types'
import { newCreator } from '~/methods'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class CCreatorEdit extends Vue {
    mdiCamera = mdiCamera
    IconType: typeof IconType = IconType
    ColorType: typeof ColorType = ColorType
    snsList: Array<ISite> = [
        {
            name: 'instagram',
            url: 'https://instagram.com/tku_accessory?igshid=YmMyMTA2M2Y=',
            icon: mdiInstagram,
        },
    ]

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
            await this.$axios.$put(`/creator`, this.creator, {}).then(async () => {
                if (this.uploadFile) {
                    const params = new FormData()
                    params.append('logo', this.uploadFile)
                    await this.$axios.$put(`/creator/logo`, params, {
                        headers: {
                            'Content-Type': 'multipart/form-data',
                        },
                    })
                }
            })
            this.initCreator = this.creator
            this.dialogVisible = false
            this.notificationVisible = true
            this.$emit('c-change')
        } catch (e) {
            this.errors.push(e)
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
    .creator-wrapper
        padding 16px
        .logo-content
            .logo-image
                margin 0 auto
                border-radius 50%
                object-fit cover
        .description-title
            color $title-text-color
        .description-content
            text-align center
            .description-text
                white-space pre-wrap
                word-break break-all
                +sm()
                    font-size 3vw
        .sns-content
            display flex
            justify-content center
            .sns-item
                margin 0 10px
                text-align center
            .sns-name
                margin-top 10px
        .edit-button
            text-align center
</style>
