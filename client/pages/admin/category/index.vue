<template>
    <v-main class="grey lighten-4">
        <v-row>
            <v-col cols="12" sm="12" md="6">
                <c-classification-list type="category" :items="categories" @c-change="loadingCategory" />
            </v-col>
            <v-col cols="12" sm="12" md="6">
                <c-classification-list type="tag" :items="tags" @c-change="loadingCategory" />
            </v-col>
        </v-row>
    </v-main>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { ICategory, CategoryType } from '~/types'
@Component({
    head: {
        title: '商品一覧',
    },
})
export default class PageAdminCategoryIndex extends Vue {
    // カテゴリー一覧
    categories: Array<ICategory> = []

    // タグ一覧
    tags: Array<ICategory> = []

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
