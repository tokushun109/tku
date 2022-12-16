<template>
    <v-container class="page-creator">
        <v-container>
            <h2 class="page-title text-sm-h3 text-h4">CREATOR</h2>
        </v-container>
        <v-sheet>
            <v-container>
                <c-creator-edit :item.sync="creator" />
            </v-container>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICreator, ISite } from '~/types'
@Component({})
export default class PageCreatorIndex extends Vue {
    creator: ICreator | null = null
    snsList: Array<ISite> | null = []
    async asyncData({ app }: Context) {
        try {
            const creator: ICreator = await app.$axios.$get(`/creator`)
            const snsList: Array<ISite> = await app.$axios.$get(`/sns`)
            return { creator, snsList }
        } catch (e) {
            return { creator: null, snsList: [] }
        }
    }

    head() {
        if (!this.creator) {
            return
        }
        const title = '製作者紹介 | tocoriri'
        const description = this.creator.introduction.replace(/\r?\n/g, '')
        const image = this.creator.apiPath
        return {
            title,
            meta: [
                {
                    hid: 'description',
                    name: 'description',
                    content: description,
                },
                {
                    hid: 'og:title',
                    property: 'og:title',
                    content: title,
                },
                {
                    hid: 'og:description',
                    property: 'og:description',
                    content: description,
                },
                {
                    hid: 'og:type',
                    property: 'og:type',
                    content: 'article',
                },
                {
                    hid: 'og:image',
                    property: 'og:image',
                    content: image,
                },
            ],
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
        display none
</style>
