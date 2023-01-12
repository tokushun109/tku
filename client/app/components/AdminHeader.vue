<template>
    <div class="admin-header">
        <v-app-bar app :color="ColorType.Green" dark clipped-left>
            <v-app-bar-nav-icon v-if="!$store.getters['user/isGuest']" @click="sidebarVisible = !sidebarVisible">
                <client-only>
                    <c-icon :type="IconType.Menu.name" @click.native="sidebarVisible = !sidebarVisible" />
                </client-only>
            </v-app-bar-nav-icon>
            <v-app-bar-title>tku</v-app-bar-title>
            <v-spacer />
            <v-dialog v-if="!$store.getters['user/isGuest']" v-model="dialogVisible" width="400">
                <template #activator="{ on, attrs }">
                    <v-btn outlined v-bind="attrs" v-on="on">ログアウト</v-btn>
                </template>
                <v-card class="sigh-out-area">
                    <v-card-title />
                    <v-card-text class="sigh-out-text">ログアウトします。よろしいですか？</v-card-text>
                    <v-divider class="divider" />
                    <v-card-actions class="sigh-out-actions">
                        <v-btn color="primary" outlined @click="dialogVisible = false">いいえ</v-btn>
                        <v-btn color="primary" @click="logoutHandler">はい</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </v-app-bar>
        <v-navigation-drawer v-if="!$store.getters['user/isGuest']" v-model="sidebarVisible" class="admin-navigation" clipped app>
            <v-container>
                <v-list-item class="navigation-list">
                    <v-list-item-content class="navigation-list-item">
                        <v-list-item-title class="title navigation-title"> 設定 </v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-divider class="divider" />
                <v-list dense nav>
                    <v-list-item v-for="(table, index) in tables" :key="index" nuxt :to="`/admin/${table.link}`">
                        <v-list-item-icon>
                            <v-icon>{{ table.icon }}</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content>
                            <v-list-item-title> {{ table.name }} </v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                </v-list>
            </v-container>
        </v-navigation-drawer>
    </div>
</template>

<script lang="ts">
import { mdiAccount, mdiApplicationOutline, mdiCartVariant, mdiEmail, mdiTagOutline } from '@mdi/js'
import { Component, Vue } from 'nuxt-property-decorator'
import { ColorType, IconType, ITable } from '~/types'
@Component({})
export default class AdminHeader extends Vue {
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType

    tables: Array<ITable> = [
        { name: '製作者', link: 'creator', icon: mdiAccount },
        { name: '商品', link: 'product', icon: mdiCartVariant },
        { name: '分類', link: 'classification', icon: mdiTagOutline },
        { name: 'サイト', link: 'site', icon: mdiApplicationOutline },
        { name: 'お問い合わせ', link: 'contact', icon: mdiEmail },
    ]

    sidebarVisible: boolean = false
    dialogVisible: boolean = false

    async logoutHandler() {
        await this.$store.dispatch('user/logoutUser')
        this.$router.replace('/admin/user/login')
    }
}
</script>

<style lang="stylus" scoped>
.sigh-out-area
    .sigh-out-text
        text-align center
        font-size $font-medium
    .sigh-out-actions
        display flex
        justify-content center

.admin-navigation
    .navigation-list
        .navigation-list-item
            .navigation-title
                color $title-primary-color
</style>
