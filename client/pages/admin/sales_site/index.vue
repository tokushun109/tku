<template>
    <c-page>
        <div>
            <c-button primary @c-click="toggle">新規追加</c-button>
            <c-sales-site-edit :visible.sync="dialogVisible" :model.sync="salesSiteModel" @close="toggle" @create="loadingSaleSite()" />
            <ul v-for="salesSite in salesSites" :key="salesSite.uuid">
                <li>{{ salesSite }}</li>
            </ul>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISalesSite, newSalesSite } from '~/types'
@Component({
    head: {
        title: '販売サイト一覧',
    },
})
export default class PageAdminSalesSiteIndex extends Vue {
    salesSites: Array<ISalesSite> = []
    // modalの表示切り替え
    dialogVisible: boolean = false
    // form用のsalesSiteModel
    salesSiteModel: ISalesSite = newSalesSite()
    async asyncData({ app }: Context) {
        try {
            const salesSites = await app.$axios.$get(`/sales_site`)
            return { salesSites }
        } catch (e) {
            return { salesSites: [] }
        }
    }

    async loadingSaleSite() {
        this.salesSites = await this.$axios.$get(`/sales_site`)
        this.salesSiteModel = newSalesSite()
    }

    // ボタンの切り替え
    toggle() {
        this.dialogVisible = !this.dialogVisible
    }
}
</script>

<style lang="stylus"></style>
