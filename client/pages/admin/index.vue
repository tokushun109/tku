<template>
    <c-page></c-page>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable } from '~/types'
@Component({
    head: {
        title: 'テーブル一覧',
    },
})
export default class PageAdminIndex extends Vue {}
</script>

<style lang="stylus"></style>
