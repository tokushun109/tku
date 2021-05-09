<template>
    <c-page>
        <div>
            <p>{{ skillMarket }}</p>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ISkillMarket } from '~/types'
@Component({
    head: {
        title: '仕事の依頼',
    },
})
export default class PageOrderIndex extends Vue {
    skillMarket: Array<ISkillMarket> | null = []
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
