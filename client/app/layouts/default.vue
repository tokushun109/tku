<template>
    <div>
        <v-app class="default">
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
        const url = (process.env.DOMAIN_URL || 'https://tocoriri.com') + this.$route.path
        const robots = this.isAdmin ? 'noindex' : ''
        return {
            meta: [
                {
                    hid: 'og:url',
                    property: 'og:url',
                    content: url,
                },
                {
                    hid: 'robots',
                    name: 'robots',
                    content: robots,
                },
            ],
            link: [
                {
                    rel: 'canonical',
                    href: url,
                },
            ],
        }
    }
}
</script>

<style lang="stylus" scoped>
.default
    background-color $primary-bg-color

* :not(.v-icon)
    font-family $font-face !important
</style>
