<template>
    <v-sheet> </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator, ISite } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageCreatorIndex extends Vue {
    creator: ICreator | null = null
    sns: Array<ISite> | null = []
    async asyncData({ app }: Context) {
        try {
            const creator = await app.$axios.$get(`/creator`)
            const sns = await app.$axios.$get(`/sns`)
            return { creator, sns }
        } catch (e) {
            return { creator: null, sns: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
