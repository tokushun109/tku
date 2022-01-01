<template>
    <v-bottom-navigation :value="value" color="primary" grow app>
        <v-btn v-for="(item, index) in menuItems" :key="index" nuxt :to="`/${item.link}`">
            <span>{{ item.name }}</span>
            <v-icon>{{ item.icon }}</v-icon>
        </v-btn>
    </v-bottom-navigation>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { ITable } from '~/types'
@Component({})
export default class SiteMenuFooter extends Vue {
    menuItems: Array<ITable> = [
        { name: 'Artist', link: 'creator', icon: 'mdi-lead-pencil' },
        { name: 'Items', link: 'product', icon: 'mdi-view-module' },
    ]

    value: string = 'recent'
}
</script>

<style lang="stylus"></style>
