<template>
    <c-page>
        <img v-if="producer" width="30%" :src="producer.logo" alt="producerLogo" />
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProducerProfile } from '~/types'
@Component({
    head: {
        title: 'tku',
    },
})
export default class PageIndex extends Vue {
    producer: IProducerProfile | null = null
    async asyncData({ app }: Context) {
        try {
            const producer = await app.$axios.$get(`producer_profile`)
            return { producer }
        } catch (e) {
            return { producer: null }
        }
    }
}
</script>

<style lang="stylus"></style>
