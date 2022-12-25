<template>
    <div class="site-layout">
        <!-- md幅以上 -->
        <v-btn fab x-large class="toggle-button" @click="toggleMenu">
            <client-only>
                <c-icon :type="IconType.Menu.name" x-large @c-click="toggleMenu" />
            </client-only>
        </v-btn>
        <v-sheet v-if="isRoot" color="transparent" class="site-title-area">
            <v-card flat width="300" color="transparent" class="site-title text-h1" to="/"> tocoriri </v-card>
            <div class="site-sub-title text-h5">Cotton lace × Macrame</div>
        </v-sheet>
        <v-dialog v-model="menuVisible" fullscreen hide-overlay transition="dialog-top-transition" scrollable>
            <v-sheet :color="ColorType.Grey" class="menu-area">
                <v-btn fab x-large class="toggle-button" @click="toggleMenu">
                    <c-icon :type="IconType.Close.name" x-large @c-click="toggleMenu" />
                </v-btn>
                <v-sheet color="transparent" class="site-title-area">
                    <v-card color="transparent" width="300" class="site-title text-h1" to="/" flat nuxt @click="toggleMenu"> tocoriri </v-card>
                </v-sheet>
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
import { mdiDiamond, mdiEmail, mdiHumanGreetingVariant } from '@mdi/js'
import { Component, Vue } from 'nuxt-property-decorator'
import { IconType, ITable, ColorType } from '~/types'
@Component({})
export default class SiteLayout extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType

    menuItems: Array<ITable> = [
        { name: 'CREATOR', link: 'creator', icon: mdiHumanGreetingVariant },
        { name: 'PRODUCTS', link: 'product', icon: mdiDiamond },
        { name: 'CONTACT', link: 'contact', icon: mdiEmail },
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
        text-align center
        +sm()
            display none
        .site-title
            margin 0 auto
            color $site-title-text-color
            font-family $title-font-face !important
        .site-sub-title
            margin-bottom 40px
            color $text-color
            font-size $font-xxlarge
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
    padding-top 65px
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
            color $white-color
            font-family $title-font-face !important
    .menu-item
        position relative
        top 10%
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
                +md()
                    font-size 3.5vw
</style>
