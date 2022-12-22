<template>
    <div class="site-footer">
        <v-container class="site-footer-content">
            <v-sheet color="transparent">
                <div class="copy-right">©︎2022 tocoriri</div>
            </v-sheet>
        </v-container>
        <!-- sm幅以下 -->
        <div class="sm">
            <v-bottom-navigation color="primary" grow app>
                <v-btn v-for="(item, index) in menuItems" :key="index" :color="ColorType.White" nuxt :to="`/${item.link}`">
                    <span>{{ item.name }}</span>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-btn>
            </v-bottom-navigation>
        </div>
    </div>
</template>

<script lang="ts">
import { mdiDiamond, mdiEmail, mdiHumanGreetingVariant } from '@mdi/js'
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable, ColorType } from '~/types'

@Component({})
export default class SiteFooter extends Vue {
    ColorType: typeof ColorType = ColorType

    menuItems: Array<ITable> = [
        { name: 'CREATOR', link: 'creator', icon: mdiHumanGreetingVariant },
        { name: 'PRODUCTS', link: 'product', icon: mdiDiamond },
        { name: 'CONTACT', link: 'contact', icon: mdiEmail },
    ]

    menuVisible: boolean = false

    get isRoot(): boolean {
        return this.$route.path === '/'
    }
}
</script>

<style lang="stylus" scoped>
.site-footer
    .site-footer-content
        display flex
        justify-content center
        .copy-right
            font-size 10px
    .sm
        display none
        +sm()
            display block
</style>
