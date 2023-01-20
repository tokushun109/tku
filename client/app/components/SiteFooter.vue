<template>
    <div class="site-footer">
        <v-container class="site-footer-content">
            <v-sheet>
                <div class="copy-right">©︎2022 とこりり</div>
            </v-sheet>
        </v-container>
        <!-- sm幅以下 -->
        <div class="sm">
            <v-bottom-navigation :color="ColorType.Primary" height="50px" grow app>
                <v-btn v-for="(item, index) in menuItems" :key="index" :color="ColorType.White" nuxt :to="`/${item.link}`">
                    <span class="menu-name">{{ item.name }}</span>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-btn>
            </v-bottom-navigation>
        </div>
    </div>
</template>

<script lang="ts">
import { mdiDiamond, mdiEmail, mdiFaceWoman } from '@mdi/js'
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable, ColorType } from '~/types'

@Component({})
export default class SiteFooter extends Vue {
    ColorType: typeof ColorType = ColorType

    menuItems: Array<ITable> = [
        { name: 'About', link: 'about', icon: mdiFaceWoman },
        { name: 'Product', link: 'product', icon: mdiDiamond },
        { name: 'Contact', link: 'contact', icon: mdiEmail },
    ]

    menuVisible: boolean = false

    get isRoot(): boolean {
        return this.$route.path === '/'
    }
}
</script>

<style lang="stylus" scoped>
.site-footer
    background-color $primary
    .site-footer-content
        display flex
        justify-content center
        .copy-right
            background-color $primary
            color $white-color
            font-size 10px
    .sm
        display none
        +sm()
            display block
            .v-item-group.v-bottom-navigation .v-btn
                width 50% !important
                height 100% !important
        .menu-name
            font-family $title-font-face !important
</style>
