<template>
    <c-page>
        <c-top-main :producer="producer" />
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
export default class PageTop extends Vue {
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
