<template>
    <c-layout-container normal class="page-creator">
        <v-container class="page-title-container">
            <h2 class="page-title text-sm-h3 text-h4">About</h2>
        </v-container>
        <v-sheet>
            <c-top-logo></c-top-logo>
            <c-creator-edit v-if="creator" :item.sync="creator" :sales-site-list="salesSiteList" />
        </v-sheet>
        <c-breadcrumbs :items="breadCrumbs" />
    </c-layout-container>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IBreadCrumb, ICreator, ISite } from '~/types'
@Component({})
export default class PageAboutIndex extends Vue {
    creator: ICreator | null = null
    snsList: Array<ISite> | null = []
    salesSiteList: Array<ISite> | null = []

    breadCrumbs: Array<IBreadCrumb> = [
        { text: 'トップページ', href: '/' },
        { text: '製作者紹介', disabled: true },
    ]

    async asyncData({ app }: Context) {
        try {
            const creator: ICreator = await app.$axios.$get(`/creator`)
            const snsList: Array<ISite> = await app.$axios.$get(`/sns`)
            const salesSiteList: Array<ISite> = await app.$axios.$get(`/sales_site`)
            return { creator, snsList, salesSiteList }
        } catch (e) {
            return { creator: null, snsList: [], salesSiteList: [] }
        }
    }

    head() {
        if (!this.creator) {
            return
        }
        const title = '製作者紹介 | とこりり'
        const description = this.creator.introduction.replace(/\r?\n/g, '')
        const image = this.creator.apiPath
        return {
            title,
            meta: [
                {
                    hid: 'description',
                    name: 'description',
                    content: description,
                },
                {
                    hid: 'og:title',
                    property: 'og:title',
                    content: title,
                },
                {
                    hid: 'og:description',
                    property: 'og:description',
                    content: description,
                },
                {
                    hid: 'og:type',
                    property: 'og:type',
                    content: 'article',
                },
                {
                    hid: 'og:image',
                    property: 'og:image',
                    content: image,
                },
            ],
        }
    }
}
</script>

<style lang="stylus" scoped>
.page-title-container
    +sm()
        display none
    .page-title
        color $site-title-text-color
        text-align center
        font-family $title-font-face !important
</style>
