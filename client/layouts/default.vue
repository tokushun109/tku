<template>
    <div>
        <v-app>
            <template v-if="!isAdmin">
                <site-layout />
                <Nuxt />
            </template>
            <template v-else>
                <admin-header />
                <Nuxt />
                <admin-footer />
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
