<template>
    <div>
        <v-app class="default-bg">
            <template v-if="!isAdmin">
                <v-main>
                    <site-layout />
                    <Nuxt />
                    <site-footer />
                </v-main>
            </template>
            <template v-else>
                <v-main>
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

    head() {
        const url = process.env.DOMAIN_URL || 'https://tocoriri.com' + this.$route.path
        return {
            meta: [
                {
                    hid: 'og:url',
                    property: 'og:url',
                    content: url,
                },
            ],
        }
    }
}
</script>

<style lang="stylus" scoped>
* :not(.v-icon)
    font-family $font-face !important
    .default-bg
        background-color $primary-bg-color
</style>
