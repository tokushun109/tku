<template>
    <v-app-bar app>
        <v-app-bar-nav-icon v-if="!$store.getters['user/isGuest']" app></v-app-bar-nav-icon>
        <v-toolbar-title>
            <nuxt-link to="/admin">tku</nuxt-link>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-dialog v-if="!$store.getters['user/isGuest']" v-model="visibleSync" width="400">
            <template #activator="{ on, attrs }">
                <v-btn color="secondary" text v-bind="attrs" v-on="on">ログアウト</v-btn>
            </template>
            <v-card>
                <v-card-title class="text-h5 justify-center"></v-card-title>
                <v-card-text class="d-flex justify-center">ログアウトします。よろしいですか？</v-card-text>
                <v-divider />
                <v-card-actions class="d-flex justify-center">
                    <v-btn color="primary" @click="logoutHandler">はい</v-btn>
                    <v-btn @click="visibleSync = false">いいえ</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-app-bar>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class AdminSiteHeader extends Vue {
    visibleSync: boolean = false

    async logoutHandler() {
        await this.$store.dispatch('user/logoutUser')
        this.$router.replace('/admin/user/login')
    }
}
</script>

<style lang="stylus"></style>
