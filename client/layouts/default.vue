<template>
    <div>
        <template v-if="!isAdmin">
            <site-header />
            <Nuxt />
        </template>
        <v-app v-else>
            <v-app>
                <admin-site-header />
                <Nuxt />
            </v-app>
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
