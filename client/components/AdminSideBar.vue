<template>
    <v-navigation-drawer clipped app>
        <!-- <v-list>
                <v-list-item v-for="[icon, text] in links" :key="icon" link>
                    <v-list-item-icon>
                        <v-icon>{{ icon }}</v-icon>
                    </v-list-item-icon>

                    <v-list-item-content>
                        <v-list-item-title>{{ text }}</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list> -->
        <v-list>
            <v-list-item v-for="(table, index) in tables" :key="index">
                <v-list-item-content>
                    <v-list-item-title>
                        <v-btn color="secondary" text block>
                            <nuxt-link :to="`/admin/${table.key}`">{{ table.name }}</nuxt-link>
                        </v-btn>
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
        <v-dialog v-if="!$store.getters['user/isGuest']" v-model="dialog" width="400">
            <template #activator="{ on, attrs }">
                <v-btn class="mt-10" color="secondary" text v-bind="attrs" block v-on="on">ログアウト</v-btn>
            </template>
            <v-card>
                <v-card-title class="text-h5 justify-center"></v-card-title>
                <v-card-text class="d-flex justify-center">ログアウトします。よろしいですか？</v-card-text>
                <v-divider />
                <v-card-actions class="d-flex justify-center">
                    <v-btn color="primary" @click="logoutHandler">はい</v-btn>
                    <v-btn @click="dialog = false">いいえ</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-navigation-drawer>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable } from '~/types'
@Component({})
export default class AdminSideBar extends Vue {
    tables: Array<ITable> = [
        { name: '製作者', key: 'creator' },
        { name: '商品', key: 'product' },
        { name: 'カテゴリー', key: 'category' },
        { name: 'サイト', key: 'site' },
    ]

    dialog: boolean = false

    async logoutHandler() {
        await this.$store.dispatch('user/logoutUser')
        this.$router.replace('/admin/user/login')
    }
}
</script>

<style lang="stylus"></style>
