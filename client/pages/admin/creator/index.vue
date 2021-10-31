<template>
    <v-main>
        <v-container class="grey py-8 px-6" fluid>
            <v-sheet width="640" height="800" class="pa-4 lighten-4 mx-auto">
                <div class="mx-auto">
                    <v-avatar color="grey darken-1" size="200"></v-avatar>
                    <div>john@vuetifyjs.com</div>
                </div>
            </v-sheet>
        </v-container>
    </v-main>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator, ISite } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCreatorIndex extends Vue {
    // 製作者
    creator: ICreator | null = null

    // SNSのリスト
    snsList: Array<ISite> | null = []

    // 販売サイトのリスト
    salesSites: Array<ISite> | null = []

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
}
</script>

<style lang="stylus"></style>
