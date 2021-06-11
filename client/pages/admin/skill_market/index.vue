<template>
    <c-page>
        <div class="admin-skill-market-list">
            <c-button primary @c-click="toggle">新規追加</c-button>
            <c-skill-market-edit :visible.sync="dialogVisible" :model.sync="skilMarketModel" @close="toggle" @create="loadingSkillMarket()" />
            <ul v-for="skillMarket in skillMarkets" :key="skillMarket.uuid">
                <li>{{ skillMarket }}</li>
            </ul>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISkillMarket, newSkillMarket } from '~/types'
@Component({
    head: {
        title: 'スキルマーケット一覧',
    },
})
export default class PageAdminSkillMarketIndex extends Vue {
    // スキルマーケット一覧
    skillMarkets: Array<ISkillMarket> = []

    // modalの表示切り替え
    dialogVisible: boolean = false

    // form用のsalesSiteModel
    skilMarketModel: ISkillMarket = newSkillMarket()

    // ボタンの切り替え
    toggle() {
        this.dialogVisible = !this.dialogVisible
    }

    async loadingSkillMarket() {
        this.skillMarkets = await this.$axios.$get(`/skill_market`)
        this.skilMarketModel = newSkillMarket()
    }

    async asyncData({ app }: Context) {
        try {
            const skillMarkets = await app.$axios.$get(`/skill_market`)
            return { skillMarkets }
        } catch (e) {
            return { skillMarkets: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
