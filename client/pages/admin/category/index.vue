<template>
    <v-main class="grey lighten-4">
        <v-row>
            <v-col cols="12" sm="12" md="6">
                <c-category-list type="accessory" :items="accessoryCategories" @c-change="loadingCategory" />
            </v-col>
            <v-col cols="12" sm="12" md="6">
                <c-category-list type="material" :items="materialCategories" @c-change="loadingCategory" />
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
    // アクセサリーカテゴリー一覧
    accessoryCategories: Array<ICategory> = []

    // タグ一覧
    materialCategories: Array<ICategory> = []

    async asyncData({ app }: Context) {
        try {
            const accessoryCategories = await app.$axios.$get(`/accessory_category`)
            const materialCategories = await app.$axios.$get(`/material_category`)
            return { accessoryCategories, materialCategories }
        } catch (e) {
            return { accessoryCategories: [], materialCategories: [] }
        }
    }

    async loadingCategory(type: string) {
        if (type === CategoryType.Accessory.name) {
            this.accessoryCategories = await this.$axios.$get(`/accessory_category`)
        } else if (type === CategoryType.Material.name) {
            this.materialCategories = await this.$axios.$get(`/material_category`)
        }
    }
}
</script>

<style lang="stylus"></style>
