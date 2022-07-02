<template>
    <div class="site-menu">
        <v-btn fab x-large class="toggle-button" @click="toggleMenu">
            <c-icon :type="IconType.Menu.name" x-large @c-click="toggleMenu" />
        </v-btn>
        <v-dialog v-model="menuVisible" fullscreen hide-overlay transition="dialog-top-transition" scrollable>
            <v-sheet :color="ColorType.Grey" class="menu-area">
                <v-btn fab x-large class="toggle-button" @click="toggleMenu">
                    <c-icon :type="IconType.Close.name" x-large @c-click="toggleMenu" />
                </v-btn>
                <v-container class="menu-item">
                    <v-row>
                        <v-col v-for="(item, index) in menuItems" :key="index" cols="6">
                            <v-card height="500" elevation="20" :to="`/${item.link}`" nuxt class="menu-item-card" @click="toggleMenu">
                                <v-card-title class="menu-item-card-title">
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
        <v-bottom-navigation class="sm text-center" color="primary" grow app>
            <v-btn v-for="(item, index) in menuItems" :key="index" nuxt :to="`/${item.link}`">
                <span>{{ item.name }}</span>
                <v-icon>{{ item.icon }}</v-icon>
            </v-btn>
        </v-bottom-navigation>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { IconType, ITable, ColorType } from '~/types'
@Component({})
export default class SiteMenu extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType

    menuItems: Array<ITable> = [
        { name: 'CREATOR', link: 'creator', icon: 'mdi-lead-pencil' },
        { name: 'PRODUCTS', link: 'product', icon: 'mdi-view-module' },
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
    .sm
        display none
        +sm()
            display block

.menu-area
    position relative
    z-index 5
    +sm()
        display none
    .toggle-button
        position absolute
        top 60px
        right 60px
    .menu-item
        position relative
        top 20%
        .menu-item-card
            border-radius $image-border-radius
            text-align center
            .menu-item-card-title
                text-align center
</style>
