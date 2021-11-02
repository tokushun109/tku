<template>
    <div class="admin-header">
        <v-app-bar app color="primary" dark clipped-left>
            <v-app-bar-nav-icon v-if="!$store.getters['user/isGuest']" @click="sidebarVisible = !sidebarVisible"></v-app-bar-nav-icon>
            <v-tool-bar-title class="text-h5">tku</v-tool-bar-title>
            <v-spacer />
            <v-dialog v-if="!$store.getters['user/isGuest']" v-model="dialogVisible" width="400">
                <template #activator="{ on, attrs }">
                    <v-btn outlined v-bind="attrs" v-on="on">ログアウト</v-btn>
                </template>
                <v-card>
                    <v-card-title class="text-h5 justify-center"></v-card-title>
                    <v-card-text class="d-flex justify-center">ログアウトします。よろしいですか？</v-card-text>
                    <v-divider />
                    <v-card-actions class="d-flex justify-center">
                        <v-btn color="primary" @click="logoutHandler">はい</v-btn>
                        <v-btn @click="dialogVisible = false">いいえ</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </v-app-bar>
        <v-navigation-drawer v-if="!$store.getters['user/isGuest']" v-model="sidebarVisible" clipped app>
            <v-container>
                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-title class="title grey--text text--darken-2"> 設定 </v-list-item-title>
                    </v-list-item-content>
                </v-list-item>

                <v-divider />

                <v-list dense nav>
                    <v-list-item v-for="(table, index) in tables" :key="index" :to="`/admin/${table.link}`">
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
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable } from '~/types'
@Component({})
export default class AdminHeader extends Vue {
    tables: Array<ITable> = [
        { name: '製作者', link: 'creator', icon: 'mdi-account' },
        { name: '商品', link: 'product', icon: 'mdi-cart-variant' },
        { name: 'カテゴリー', link: 'category', icon: 'mdi-tag-outline' },
        { name: 'サイト', link: 'site', icon: 'mdi-application-outline' },
    ]

    sidebarVisible: boolean = false
    dialogVisible: boolean = false

    async logoutHandler() {
        await this.$store.dispatch('user/logoutUser')
        this.$router.replace('/admin/user/login')
    }
}
</script>

<style lang="stylus"></style>
