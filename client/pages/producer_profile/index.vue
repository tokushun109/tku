<template>
    <div>
        <p>{{ producerProfile }}</p>
        <p>{{ sns }}</p>
        <p>{{ skillMarket }}</p>
        <p>{{ salesSite }}</p>
    </div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProducerProfile, ISns, ISkillMarket, ISalesSite } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageProducerProfileIndex extends Vue {
    producerProfile: IProducerProfile | null = null
    sns: Array<ISns> | null = []
    skillMarket: Array<ISkillMarket> | null = []
    salesSite: Array<ISalesSite> | null = []
    async asyncData({ app }: Context) {
        try {
            const producerProfile = await app.$axios.$get(`/producer_profile/`)
            const sns = await app.$axios.$get(`/sns/`)
            const skillMarket = await app.$axios.$get(`/skill_market/`)
            const salesSite = await app.$axios.$get(`/sales_site/`)
            return { producerProfile, sns, skillMarket, salesSite }
        } catch (e) {
            return { producerProfile: null }
        }
    }
}
</script>

<style lang="stylus"></style>
