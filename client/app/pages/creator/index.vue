<template>
    <v-container class="page-creator">
        <v-container>
            <h2 class="page-title text-sm-h3 text-h4">CREATOR</h2>
        </v-container>
        <v-sheet>
            <v-container>
                <c-creator-edit :item.sync="creator" :sns-list="snsList" />
            </v-container>
        </v-sheet>
    </v-container>
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
    snsList: Array<ISite> | null = []
    async asyncData({ app }: Context) {
        try {
            const creator = await app.$axios.$get(`/creator`)
            const snsList = await app.$axios.$get(`/sns`)
            return { creator, snsList }
        } catch (e) {
            return { creator: null, snsList: [] }
        }
    }
}
</script>

<style lang="stylus" scoped>
.page-title
    margin-bottom 20px
    color $site-title-text-color
    text-align center
    +sm()
        margin-bottom auto
</style>
