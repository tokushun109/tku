<template>
    <div class="site-layout">
        <!-- md幅以上 -->
        <v-btn fab x-large class="toggle-button" @click="toggleMenu">
            <client-only>
                <c-icon :type="IconType.Menu.name" x-large @c-click="toggleMenu" />
            </client-only>
        </v-btn>
        <div v-if="isRoot" class="site-title-area">
            <nuxt-link to="/">
                <img class="site-title" src="/img/logo/tocoriri_logo.png" alt="アクセサリーショップ とこりり" />
            </nuxt-link>
        </div>
        <v-dialog v-model="menuVisible" fullscreen hide-overlay transition="dialog-top-transition" scrollable>
            <v-sheet :color="ColorType.Grey" class="menu-area">
                <v-btn fab x-large class="toggle-button" @click="toggleMenu">
                    <c-icon :type="IconType.Close.name" x-large @c-click="toggleMenu" />
                </v-btn>
                <div class="site-title-area" @click="toggleMenu">
                    <nuxt-link to="/">
                        <img class="site-title" src="/img/logo/tocoriri_logo_white.png" alt="とこりり メニュー" />
                    </nuxt-link>
                </div>
                <v-container class="menu-item">
                    <v-row>
                        <v-col v-for="(item, index) in menuItems" :key="index" cols="4">
                            <v-card height="100%" elevation="20" :to="`/${item.link}`" nuxt class="menu-card" @click="toggleMenu">
                                <v-card-title class="menu-card-icon">
                                    <v-avatar size="90%">
                                        <v-icon size="90%">{{ item.icon }}</v-icon>
                                    </v-avatar>
                                </v-card-title>
                                <v-card-text class="menu-card-name">
                                    {{ item.name }}
                                </v-card-text>
                            </v-card>
                        </v-col>
                        <v-spacer />
                    </v-row>
                </v-container>
            </v-sheet>
        </v-dialog>

        <!-- sm幅以下 -->
        <div class="sm">
            <v-app-bar dense class="site-header">
                <v-card color="transparent" class="site-title text-h4" to="/" flat nuxt>tocoriri</v-card>
            </v-app-bar>
        </div>
    </div>
</template>

<script lang="ts">
import { mdiDiamond, mdiEmail, mdiFaceWoman } from '@mdi/js'
import { Component, Vue } from 'nuxt-property-decorator'
import { IconType, ITable, ColorType } from '~/types'
@Component({})
export default class SiteLayout extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType

    menuItems: Array<ITable> = [
        { name: 'About', link: 'about', icon: mdiFaceWoman },
        { name: 'Product', link: 'product', icon: mdiDiamond },
        { name: 'Contact', link: 'contact', icon: mdiEmail },
    ]

    menuVisible: boolean = false

    get isRoot(): boolean {
        return this.$route.path === '/'
    }

    toggleMenu() {
        this.menuVisible = !this.menuVisible
    }
}
</script>

<style lang="stylus" scoped>
.site-layout
    position relative
    .toggle-button
        position fixed
        top 40px
        right 40px
        z-index 5
        +sm()
            display none
    .site-title-area
        position relative
        padding-top 20px
        text-align center
        +sm()
            display none
        .site-title
            margin 0 auto
            width 400px
            height 200px
            object-fit cover
    .sm
        display none
        text-align center
        +sm()
            display block
            .site-header
                z-index 10 !important
                .site-title
                    margin 0 auto
                    color $text-color
                    font-family $title-font-face !important

.menu-area
    position relative
    z-index 5
    padding-top 20px
    +sm()
        display none
    .toggle-button
        position absolute
        top 40px
        right 40px
        z-index 6
    .site-title-area
        position relative
        text-align center
        +sm()
            display none
        .site-title
            margin 0 auto
            width 400px
            height 200px
            object-fit cover
    .menu-item
        position relative
        .menu-card
            padding 0 0 20px
            border-radius $image-border-radius
            text-align center
            transition all 0.2s
            &:hover
                background-color #DCEDC8
                cursor pointer
                transform translateY(-10px)
            .menu-card-icon
                justify-content center
            .menu-card-name
                font-size 35px
                font-family $title-font-face !important
                +md()
                    font-size 3.5vw
</style>
