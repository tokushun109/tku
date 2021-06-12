<template>
    <c-page>
        <div>
            <p>{{ creator }}</p>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator, ISalesSite, ISns } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCreatorIndex extends Vue {
    // 製作者
    creator: ICreator | null = null

    // SNSのリスト
    snsList: Array<ISns> | null = []

    // 販売サイトのリスト
    salesSites: Array<ISalesSite> | null = []

    async asyncData({ app }: Context) {
        try {
            const creator = await app.$axios.$get(`/creator`)
            const snsList = await app.$axios.$get(`/sns`)
            const salesSites = await app.$axios.$get(`/sns`)
            return { creator, snsList, salesSites }
        } catch (e) {
            return { creator: null, snsList: [], salesSites: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
