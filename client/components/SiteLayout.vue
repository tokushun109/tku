<template>
    <div class="site-menu light-green lighten-4">
        <!-- md幅以上 -->
        <v-btn fab x-large class="toggle-button" @click="toggleMenu">
            <c-icon :type="IconType.Menu.name" x-large @c-click="toggleMenu" />
        </v-btn>
        <v-sheet height="190" color="transparent" class="site-title-area">
            <v-card flat color="transparent" class="site-title mx-auto text-h1 grey--text text--darken-1" to="/">tku</v-card>
        </v-sheet>
        <v-dialog v-model="menuVisible" fullscreen hide-overlay transition="dialog-top-transition" scrollable>
            <v-sheet :color="ColorType.Grey" class="menu-area">
                <v-btn fab x-large class="toggle-button" @click="toggleMenu">
                    <c-icon :type="IconType.Close.name" x-large @c-click="toggleMenu" />
                </v-btn>
                <v-sheet height="190" color="transparent" class="site-title-area">
                    <v-card color="transparent" class="site-title mx-auto text-h1 white--text" to="/" flat nuxt @click="toggleMenu">tku</v-card>
                </v-sheet>
                <v-container class="menu-item">
                    <v-row>
                        <v-col v-for="(item, index) in menuItems" :key="index" cols="6">
                            <v-card height="500" elevation="20" :to="item.link" nuxt rounded class="text-center rounded-xl" @click="toggleMenu">
                                <v-card-title class="justify-center">
                                    <v-avatar size="300">
                                        <v-icon size="300">{{ item.icon }}</v-icon>
                                    </v-avatar>
                                </v-card-title>
                                <v-card-text>
                                    <div class="text-h1">{{ item.name }}</div>
                                </v-card-text>
                            </v-card>
                        </v-col>
                        <v-spacer />
                    </v-row>
                </v-container>
            </v-sheet>
        </v-dialog>

        <!-- sm幅以下 -->
        <div class="sm text-center">
            <v-app-bar app>
                <v-card color="transparent" class="site-title mx-auto text-h4 grey--text text--darken-1" to="/" flat nuxt>tku</v-card>
            </v-app-bar>
            <v-bottom-navigation color="primary" grow app>
                <v-btn v-for="(item, index) in menuItems" :key="index" nuxt :to="`/${item.link}`">
                    <span>{{ item.name }}</span>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-btn>
            </v-bottom-navigation>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { IconType, ITable, ColorType } from '~/types'
@Component({})
export default class SiteLayout extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType

    menuItems: Array<ITable> = [
        { name: 'About', link: 'creator', icon: 'mdi-lead-pencil' },
        { name: 'Items', link: 'product', icon: 'mdi-view-module' },
    ]

    menuVisible: boolean = false

    toggleMenu() {
        this.menuVisible = !this.menuVisible
    }
}
</script>

<style lang="stylus" scoped>
.site-menu
    position relative
    .toggle-button
        position absolute
        top 60px
        right 60px
        z-index 5
        +sm()
            display none
    .site-title-area
        position relative
        +sm()
            display none
        .site-title
            position absolute
            top 50%
            left 50%
            font-family 'Lobster' !important
            transform translateY(-50%) translateX(-50%)
    .sm
        display none
        +sm()
            display block
        .site-title
            font-family 'Lobster' !important

.menu-area
    position relative
    z-index 5
    +sm()
        display none
    .toggle-button
        position absolute
        top 60px
        right 60px
        z-index 6
    .site-title-area
        position relative
        +sm()
            display none
        .site-title
            position absolute
            top 50%
            left 50%
            font-family 'Lobster' !important
            transform translateY(-50%) translateX(-50%)
    .menu-item
        position relative
        top 10%
</style>
