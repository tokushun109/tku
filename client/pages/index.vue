<template>
    <v-main class="light-green lighten-4"> </v-main>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator } from '~/types'
@Component({
    head: {
        title: 'tku',
    },
})
export default class PageTop extends Vue {
    creator: ICreator | null = null
    async asyncData({ app }: Context) {
        try {
            const creator = await app.$axios.$get(`creator`)
            return { creator }
        } catch (e) {
            return { creator: null }
        }
    }
}
</script>

<style lang="stylus"></style>
