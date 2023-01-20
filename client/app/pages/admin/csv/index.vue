<template>
    <v-container class="page-admin-csv">
        <v-sheet class="csv-area">
            <div class="csv-header">
                <h3 class="title csv-title">商品レコード</h3>
                <v-spacer />
            </div>
            <v-divider class="divider" />
            <div class="csv-content">
                <div class="csv-buttons">
                    <v-btn :color="ColorType.Primary" class="button" @click="downloadHandler">ダウンロード</v-btn>
                    <v-btn :color="ColorType.Primary" class="button" @click="uploadDialogHandler">アップロード</v-btn>
                </div>
            </div>
        </v-sheet>
        <c-dialog :visible.sync="dialogVisible" title="CSVのアップロード" @confirm="confirmHandler" @close="closeHandler">
            <template #content>
                <c-error :errors.sync="errors" />
                <v-form ref="form" lazy-validation>
                    <v-file-input v-model="uploadFile" accept="text/csv" label="CSVファイル" :prepend-icon="mdiFileDocument" outlined />
                </v-form>
            </template>
        </c-dialog>
        <c-notification :visible.sync="notificationVisible">アップロードを完了しました</c-notification>
    </v-container>
</template>

<script lang="ts">
import { mdiFileDocument } from '@mdi/js'
import { Component, Vue } from 'nuxt-property-decorator'
import { BadRequest, ColorType, IError } from '~/types'

@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCsvIndex extends Vue {
    ColorType: typeof ColorType = ColorType
    mdiFileDocument = mdiFileDocument

    uploadFile: File | null = null
    dialogVisible: boolean = false
    notificationVisible: boolean = false
    errors: Array<IError> = []

    async downloadHandler() {
        const csvText = await this.$axios.$get(`/csv/product`, { withCredentials: true })
        const blob = new Blob([csvText], { type: 'text/csv' })
        const link = document.createElement('a')
        link.href = URL.createObjectURL(blob)
        link.download = '商品レコード.csv'
        link.click()
    }

    uploadDialogHandler() {
        this.dialogVisible = true
    }

    async confirmHandler() {
        this.errors = []
        if (!this.uploadFile) {
            this.errors.push(new BadRequest('CSVファイルが添付されていません'))
            return
        }
        try {
            const params = new FormData()
            params.append(`csv`, this.uploadFile)
            await this.$axios.$post('csv/product', params, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
                withCredentials: true,
            })
            this.notificationVisible = true
            this.dialogVisible = false
            this.uploadFile = null
        } catch (e) {
            this.errors.push(e.response)
        }
    }

    closeHandler() {
        this.dialogVisible = false
    }
}
</script>

<style lang="stylus" scoped>
.page-admin-csv
    .csv-area
        padding 16px
        .csv-header
            .csv-title
                color $title-primary-color
        .csv-content
            padding-top 20px
            width 100%
            .csv-buttons
                display flex
                margin 0 auto
                width fit-content
                .button
                    margin 20px
                    color $white-color
</style>
