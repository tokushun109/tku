<template>
    <div class="site-menu light-green lighten-4">
        <v-btn fab x-large class="toggle-button" @click="toggleMenu">
            <c-icon :type="IconType.Menu.name" x-large @c-click="toggleMenu" />
        </v-btn>
        <v-dialog v-model="menuVisible" fullscreen hide-overlay transition="dialog-top-transition" scrollable>
            <v-sheet :color="ColorType.Grey" class="menu-area">
                <v-btn fab x-large class="toggle-button" @click="toggleMenu">
                    <c-icon :type="IconType.Close.name" x-large @c-click="toggleMenu" />
                </v-btn>
                <v-container class="menu-item">
                    <v-list nav>
                        <v-list-item v-for="(item, index) in menuItems" :key="index" nuxt :to="`/${item.link}`" @click="toggleMenu">
                            <v-list-item-icon>
                                <v-icon x-large>{{ item.icon }}</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title> {{ item.name }} </v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                    </v-list>
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
        position absolute
        top 30%
</style>
