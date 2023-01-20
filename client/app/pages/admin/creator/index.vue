<template>
    <v-sheet class="page-admin-creator">
        <c-creator-edit :item.sync="creator" admin @c-change="loadingCreator" />
    </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator, IError, ISite } from '~/types'
import { newCreator } from '~/methods'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCreatorIndex extends Vue {
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

    async loadingCreator() {
        this.creator = await this.$axios.$get(`/creator`)
    }

    async test() {
        const test = await this.$axios.$get(`/csv/product`, { withCredentials: true })
        const blob = new Blob([test], { type: 'text/csv' }) // 配列に上記の文字列(str)を設定
        const link = document.createElement('a')
        link.href = URL.createObjectURL(blob)
        link.download = 'template.csv'
        link.click()
    }
}
</script>

<style lang="stylus" scoped></style>
