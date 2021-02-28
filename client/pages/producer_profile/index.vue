<template>
    <c-page>
        <div>
            <p>{{ producerProfile }}</p>
            <p>{{ sns }}</p>
            <p>{{ salesSite }}</p>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProducerProfile, ISns } from '~/types'
@Component({
    head: {
        title: '製作者紹介',
    },
})
export default class PageProducerProfileIndex extends Vue {
    producerProfile: IProducerProfile | null = null
    sns: Array<ISns> | null = []
    async asyncData({ app }: Context) {
        try {
            const producerProfile = await app.$axios.$get(`/producer_profile/`)
            const sns = await app.$axios.$get(`/sns/`)
            return { producerProfile, sns }
        } catch (e) {
            return { producerProfile: null, sns: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
