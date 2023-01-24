<template>
    <v-sheet class="page-admin-seo">
        <c-seo-edit :item.sync="creator" admin @c-change="loadingCreator" />
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
export default class PageAdminSeoIndex extends Vue {
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
            return { creator }
        } catch (e) {
            return { creator: null }
        }
    }

    async loadingCreator() {
        this.creator = await this.$axios.$get(`/creator`)
    }
}
</script>

<style lang="stylus" scoped></style>
