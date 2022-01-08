<template>
    <div>
        <v-app>
            <template v-if="!isAdmin">
                <v-main class="light-green lighten-4">
                    <site-layout />
                    <Nuxt />
                </v-main>
            </template>
            <template v-else>
                <v-main class="grey lighten-4">
                    <admin-header />
                    <Nuxt />
                    <admin-footer />
                </v-main>
            </template>
        </v-app>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({})
export default class LayoutDefault extends Vue {
    // urlにadminが含まれているかを確認
    get isAdmin() {
        return this.$route.path.includes('admin')
    }
}
</script>

<style lang="stylus"></style>
