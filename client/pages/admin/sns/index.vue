<template>
    <c-page>
        <div class="admin-sns-list">
            <c-button primary @c-click="toggle">新規追加</c-button>
            <c-sns-edit :visible.sync="dialogVisible" :model.sync="snsModel" @close="toggle" @create="loadingSns()" />
            <ul v-for="sns in snsList" :key="sns.uuid">
                <li>{{ sns }}</li>
            </ul>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISns, newSns } from '~/types'
@Component({
    head: {
        title: 'SNS一覧',
    },
})
export default class PageAdminSnsIndex extends Vue {
    // SNS一覧
    snsList: Array<ISns> = []

    // modalの表示切り替え
    dialogVisible: boolean = false

    // form用のsalesSiteModel
    snsModel: ISns = newSns()

    // ボタンの切り替え
    toggle() {
        this.dialogVisible = !this.dialogVisible
    }

    async loadingSns() {
        this.snsList = await this.$axios.$get(`/sns`)
        this.snsModel = newSns()
    }

    async asyncData({ app }: Context) {
        try {
            const snsList = await app.$axios.$get(`/sns`)
            return { snsList }
        } catch (e) {
            return { snsList: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
