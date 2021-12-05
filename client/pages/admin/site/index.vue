<template>
    <v-main class="grey lighten-4">
        <v-row>
            <v-col cols="12" sm="12" md="4">
                <c-site-list :type="SiteType.Sns.name" :items="snsList" @c-change="loadingSite" />
            </v-col>
            <v-col cols="12" sm="12" md="4">
                <c-site-list :type="SiteType.SalesSite.name" :items="salesSites" @c-change="loadingSite" />
            </v-col>
            <v-col cols="12" sm="12" md="4">
                <c-site-list :type="SiteType.SkillMarket.name" :items="skillMarkets" @c-change="loadingSite" />
            </v-col>
        </v-row>
    </v-main>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISite, SiteType } from '~/types'
@Component({
    head: {
        title: 'サイト一覧',
    },
})
export default class PageAdminSiteIndex extends Vue {
    SiteType: typeof SiteType = SiteType

    // SNS一覧
    snsList: Array<ISite> = []
    // 販売サイト一覧
    salesSites: Array<ISite> = []
    // スキルマーケット一覧
    skillMarkets: Array<ISite> = []

    async asyncData({ app }: Context) {
        try {
            const snsList = await app.$axios.$get(`/sns`)
            const skillMarkets = await app.$axios.$get(`/skill_market`)
            const salesSites = await app.$axios.$get(`/sales_site`)
            return { snsList, skillMarkets, salesSites }
        } catch (e) {
            return { snsList: [], skillMarkets: [], salesSites: [] }
        }
    }

    async loadingSite(type: string) {
        if (type === SiteType.Sns.name) {
            this.snsList = await this.$axios.$get(`/sns`)
        } else if (type === SiteType.SkillMarket.name) {
            this.skillMarkets = await this.$axios.$get(`/skill_market`)
        } else if (type === SiteType.SalesSite.name) {
            this.salesSites = await this.$axios.$get(`/sales_site`)
        }
    }
}
</script>

<style lang="stylus"></style>
