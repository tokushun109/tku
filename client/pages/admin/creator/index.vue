<template>
    <v-main class="grey lighten-4">
        <v-container>
            <v-sheet class="pa-4 lighten-4">
                <h3 class="title">サイトロゴ</h3>
                <v-divider />
                <div class="text-center"><v-avatar color="grey darken-1" class="my-4" size="240px" /></div>
                <h3 class="title">紹介文</h3>
                <v-divider />
                <div class="my-4">{{ creator.introduction }}</div>
                <v-dialog v-if="!$store.getters['user/isGuest']" v-model="dialogVisible" width="800" height="800">
                    <template #activator="{ on, attrs }">
                        <div class="text-center"><v-btn color="primary" v-bind="attrs" v-on="on">編集</v-btn></div>
                    </template>
                    <v-card>
                        <v-card-title class="text-h5 justify-center blue white--text">製作者の編集</v-card-title>
                        <v-card-text class="pt-5">
                            <v-file-input v-model="creator.uploadFile" label="ロゴ画像" outlined />
                            <v-textarea v-model="creator.introduction" label="紹介文" outlined />
                            <div class="text-center"><v-btn color="primary" @click="saveHandler">登録</v-btn></div>
                        </v-card-text>
                    </v-card>
                </v-dialog>
            </v-sheet>
        </v-container>
    </v-main>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator, IError, ISite, newCreator } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCreatorIndex extends Vue {
    dialogVisible: boolean = false

    // 製作者
    creator: ICreator = newCreator()

    // SNSのリスト
    snsList: Array<ISite> | null = []

    // 販売サイトのリスト
    salesSites: Array<ISite> | null = []

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

    async saveHandler() {
        try {
            this.errors = []
            await this.$axios.$put(`/creator`, this.creator, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            })
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
