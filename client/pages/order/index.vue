<template>
    <v-sheet class="light-green lighten-4"> </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISite } from '~/types'
@Component({
    head: {
        title: '仕事の依頼',
    },
})
export default class PageOrderIndex extends Vue {
    skillMarket: Array<ISite> | null = []
    async asyncData({ app }: Context) {
        try {
            const skillMarket = await app.$axios.$get(`/skill_market`)
            return { skillMarket }
        } catch (e) {
            return { skillMarket: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
