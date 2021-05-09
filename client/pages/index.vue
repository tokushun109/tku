<template>
    <c-page>
        <c-top-main :creator="creator" />
    </c-page>
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
