<template>
    <c-page>
        <ul v-for="table in tables" :key="table">
            <li>
                <nuxt-link :to="`/admin/${table.key}`">{{ table.name }}</nuxt-link>
            </li>
        </ul>
    </c-page>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable } from '~/types'
@Component({
    head: {
        title: 'テーブル一覧',
    },
})
export default class PageAdminIndex extends Vue {
    tables: Array<ITable> = [
        { name: '製作者', key: 'creator' },
        { name: '商品', key: 'product' },
        { name: 'カテゴリー', key: 'category' },
        { name: '販売サイト', key: 'sales_site' },
        { name: 'スキルマーケット', key: 'skill_market' },
        { name: 'SNS', key: 'sns' },
    ]
}
</script>

<style lang="stylus"></style>
