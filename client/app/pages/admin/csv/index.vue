<template>
    <v-container class="page-admin-csv">
        <v-sheet class="csv-area">
            <div class="csv-header">
                <h3 class="title csv-title">商品レコード</h3>
                <v-spacer />
            </div>
            <v-divider class="divider" />
            <div class="csv-content">
                <div class="csv-buttons">
                    <v-btn :color="ColorType.Primary" class="button" @click="downloadHandler">ダウンロード</v-btn>
                    <v-btn :color="ColorType.Primary" class="button">アップロード</v-btn>
                </div>
            </div>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { ColorType } from '~/types'

@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageAdminCsvIndex extends Vue {
    ColorType: typeof ColorType = ColorType

    async downloadHandler() {
        const csvText = await this.$axios.$get(`/csv/product`, { withCredentials: true })
        const blob = new Blob([csvText], { type: 'text/csv' })
        const link = document.createElement('a')
        link.href = URL.createObjectURL(blob)
        link.download = '商品レコード.csv'
        link.click()
    }
}
</script>

<style lang="stylus" scoped>
.page-admin-csv
    .csv-area
        padding 16px
        .csv-header
            .csv-title
                color $title-primary-color
        .csv-content
            padding-top 20px
            width 100%
            .csv-buttons
                display flex
                margin 0 auto
                width fit-content
                .button
                    margin 20px
                    color $white-color
</style>
