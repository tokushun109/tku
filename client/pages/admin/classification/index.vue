<template>
    <v-row>
        <v-col cols="12" md="6">
            <c-classification-list :type="CategoryType.Category.name" :items="categories" @c-change="loadingCategory" />
        </v-col>
        <v-col cols="12" md="6">
            <c-classification-list :type="CategoryType.Tag.name" :items="tags" @c-change="loadingCategory" />
        </v-col>
    </v-row>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IClassification, CategoryType } from '~/types'
@Component({
    head: {
        title: '分類一覧',
    },
})
export default class PageAdminClassificationIndex extends Vue {
    CategoryType: typeof CategoryType = CategoryType

    // カテゴリー一覧
    categories: Array<IClassification> = []

    // タグ一覧
    tags: Array<IClassification> = []

    async asyncData({ app }: Context) {
        try {
            const categories = await app.$axios.$get(`/category`)
            const tags = await app.$axios.$get(`/tag`)
            return { categories, tags }
        } catch (e) {
            return { categories: [], tags: [] }
        }
    }

    async loadingCategory(type: string) {
        if (type === CategoryType.Category.name) {
            this.categories = await this.$axios.$get(`/category`)
        } else if (type === CategoryType.Tag.name) {
            this.tags = await this.$axios.$get(`/tag`)
        }
    }
}
</script>

<style lang="stylus"></style>
