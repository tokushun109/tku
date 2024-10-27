<template>
    <div class="maintenance-wrapper">
        <div class="site-title-area">
            <nuxt-link to="#">
                <img class="site-title" src="/img/logo/tocoriri_logo.png" alt="アクセサリーショップ とこりり" />
            </nuxt-link>
        </div>
        <v-container>
            <v-container class="maintenance-content">
                <div class="maintenance-message">
                    <p>ただいまメンテナンス中です</p>
                    <p>しばらく経ってからお試しください</p>
                </div>
            </v-container>
        </v-container>
    </div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'

@Component({})
export default class PageMaintenance extends Vue {
    async asyncData({ app, redirect }: Context) {
        try {
            await app.$axios.$get(`/health_check`)
            // メンテナンス中ではない場合は、トップページにリダイレクト
            redirect('/')
        } catch {}
    }

    head() {
        const title = 'メンテナンス中です | とこりり'
        return {
            title,
        }
    }
}
</script>

<style lang="stylus" scoped>
.maintenance-wrapper
    text-align center
    .site-title-area
        position relative
        padding-top 35px
        text-align center
        +sm()
            display none
        .site-title
            margin 0 auto
            width 370px
            height 150px
            object-fit cover
    .maintenance-content
        height 75vh
        .maintenance-message
            padding 20vh 0
            color $title-primary-color
            font-size $font-xxlarge
            .maintenance-code
                font-size 50px
</style>
