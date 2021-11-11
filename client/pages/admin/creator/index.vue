<template>
    <v-main class="grey lighten-4">
        <v-container>
            <v-sheet class="pa-4 lighten-4">
                <h3 class="title green--text text--darken-3">サイトロゴ</h3>
                <v-divider />
                <div class="text-center">
                    <v-avatar color="grey darken-1" class="my-4" size="240">
                        <img v-if="creator.apiPath" :src="creator.apiPath" style="object-fit: cover" alt="ロゴ画像" />
                    </v-avatar>
                </div>
                <h3 class="title green--text text--darken-3">紹介文</h3>
                <v-divider />
                <div class="my-4">
                    <pre>{{ creator.introduction }}</pre>
                </div>
                <c-dialog-2 :visible.sync="dialogVisible" title="製作者の編集" width="800" @confirm="saveHandler" @close="closeHandler">
                    <template #trigger>
                        <v-btn color="primary" @click="setInit"><c-icon type="edit" />編集</v-btn>
                    </template>
                    <template #content>
                        <v-file-input v-model="uploadFile" label="ロゴ画像" outlined />
                        <v-textarea v-model="creator.introduction" label="紹介文" outlined />
                    </template>
                </c-dialog-2>
            </v-sheet>
        </v-container>
        <c-notification :visible.sync="notificationVisible">製作者を更新しました</c-notification>
    </v-main>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import _ from 'lodash'
import { ICreator, IError, ISite, newCreator } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCreatorIndex extends Vue {
    // 製作者
    creator: ICreator = newCreator()

    // 最作者の初期情報
    initCreator: ICreator = newCreator()

    // アップロードするロゴファイル
    uploadFile: File | null = null

    // SNSのリスト
    snsList: Array<ISite> | null = []

    // 販売サイトのリスト
    salesSites: Array<ISite> | null = []

    // 製作者編集ダイアログの表示
    dialogVisible: boolean = false

    // 通知の表示
    notificationVisible: boolean = false

    errors: Array<IError> = []
    async asyncData({ app }: Context) {
        try {
            const creator = await app.$axios.$get(`/creator`)
            const snsList = await app.$axios.$get(`/sns`)
            const salesSites = await app.$axios.$get(`/sales_site`)
            return { creator, snsList, salesSites }
        } catch (e) {
            return { creator: null, snsList: [], salesSites: [] }
        }
    }

    async loadingCreator() {
        this.creator = await this.$axios.$get(`/creator`)
    }

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
            await this.loadingCreator()
            this.initCreator = this.creator
            this.dialogVisible = false
            this.notificationVisible = true
        } catch (e) {
            this.errors.push(e)
        }
    }

    mounted() {
        // 製作者の初期情報
        this.initCreator = _.cloneDeep(this.creator)
    }

    setInit() {
        this.creator = _.cloneDeep(this.initCreator)
    }

    closeHandler() {
        this.dialogVisible = false
    }
}
</script>

<style lang="stylus"></style>
