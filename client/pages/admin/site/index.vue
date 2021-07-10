<template>
    <c-page>
        <div class="admin-sns-list">
            <c-button primary @c-click="snsDialogToggle">新規追加</c-button>
            <c-sns-edit :visible.sync="snsDialogVisible" :model.sync="snsModel" @close="snsDialogToggle" @create="loadingSite($event)" />
            <ul v-for="sns in snsList" :key="sns.uuid">
                <c-column>
                    <li>{{ sns.name }}</li>
                    <div class="button-wrapper">
                        <c-button class="delete-button" label="削除" @c-click="snsDeleteHandler(sns)" />
                    </div>
                </c-column>
            </ul>
        </div>
        <div class="admin-sales-site-list">
            <c-button primary @c-click="salesSiteDialogtoggle">新規追加</c-button>
            <c-sales-site-edit
                :visible.sync="salesSiteDialogVisible"
                :model.sync="salesSiteModel"
                @close="salesSiteDialogtoggle"
                @create="loadingSite($event)"
            />
            <ul v-for="salesSite in salesSites" :key="salesSite.uuid">
                <c-column>
                    <li>{{ salesSite.name }}</li>
                    <div class="button-wrapper">
                        <c-button class="delete-button" label="削除" @c-click="salesSiteDeleteHandler(salesSite)" />
                    </div>
                </c-column>
            </ul>
        </div>
        <div class="admin-skill-market-list">
            <c-button primary @c-click="skillMarketDialogToggle">新規追加</c-button>
            <c-skill-market-edit
                :visible.sync="skillMarketDialogVisible"
                :model.sync="skilMarketModel"
                @close="skillMarketDialogToggle"
                @create="loadingSite($event)"
            />
            <ul v-for="skillMarket in skillMarkets" :key="skillMarket.uuid">
                <c-column>
                    <li>{{ skillMarket.name }}</li>
                    <div class="button-wrapper">
                        <c-button class="delete-button" label="削除" @c-click="skillMarketDeleteHandler(skillMarket)" />
                    </div>
                </c-column>
            </ul>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISite, newSite, SiteType } from '~/types'
@Component({
    head: {
        title: 'サイト一覧',
    },
})
export default class PageAdminSiteIndex extends Vue {
    // SNS一覧
    snsList: Array<ISite> = []
    // 販売サイト一覧
    salesSites: Array<ISite> = []
    // スキルマーケット一覧
    skillMarkets: Array<ISite> = []

    // SNS用のmodalの表示切り替え
    snsDialogVisible: boolean = false
    // 販売サイト用の表示切り替え
    salesSiteDialogVisible: boolean = false
    // スキルマーケット用の表示切り替え
    skillMarketDialogVisible: boolean = false

    // form用のsalesSiteModel
    snsModel: ISite = newSite()
    // form用のsalesSiteModel
    salesSiteModel: ISite = newSite()
    // form用のskillMarketModel
    skilMarketModel: ISite = newSite()

    // SNSダイアログの切り替え
    snsDialogToggle() {
        this.snsDialogVisible = !this.snsDialogVisible
    }

    // 販売サイトダイアログの切り替え
    salesSiteDialogtoggle() {
        this.salesSiteDialogVisible = !this.salesSiteDialogVisible
    }

    // スキルマーケットダイアログの切り替え
    skillMarketDialogToggle() {
        this.skillMarketDialogVisible = !this.skillMarketDialogVisible
    }

    // SNSの削除
    async snsDeleteHandler(sns: ISite) {
        if (confirm(`${sns.name}を削除します。よろしいですか？`)) {
            await this.$axios.$delete(`/sns/${sns.uuid}`)
            this.loadingSite(SiteType.Sns)
        }
    }

    // 販売サイトの削除
    async salesSiteDeleteHandler(salesSite: ISite) {
        if (confirm(`${salesSite.name}を削除します。よろしいですか？`)) {
            await this.$axios.$delete(`/sales_site/${salesSite.uuid}`)
            this.loadingSite(SiteType.SalesSite)
        }
    }

    // スキルマーケットの削除
    async skillMarketDeleteHandler(skillMarket: ISite) {
        if (confirm(`${skillMarket.name}を削除します。よろしいですか？`)) {
            await this.$axios.$delete(`/skill_market/${skillMarket.uuid}`)
            this.loadingSite(SiteType.SkillMarket)
        }
    }

    async loadingSite(type: string) {
        if (type === SiteType.Sns) {
            this.snsList = await this.$axios.$get(`/sns`)
            this.snsModel = newSite()
        } else if (type === SiteType.SkillMarket) {
            this.skillMarkets = await this.$axios.$get(`/skill_market`)
            this.skilMarketModel = newSite()
        } else if (type === SiteType.SalesSite) {
            this.salesSites = await this.$axios.$get(`/sales_site`)
            this.salesSiteModel = newSite()
        }
    }

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
}
</script>

<style lang="stylus"></style>
