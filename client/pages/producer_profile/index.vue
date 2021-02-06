<template>
    <div>{{ producerProfile }}</div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProducerProfile } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageProducerProfileIndex extends Vue {
    producerProfile: IProducerProfile | null = null
    async asyncData({ app }: Context) {
        try {
            const producerProfile = await app.$axios.$get(`/producer_profile/`)
            return { producerProfile }
        } catch (e) {
            return { producerProfile: null }
        }
    }
}
</script>

<style lang="stylus"></style>
