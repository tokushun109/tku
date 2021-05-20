<template>
    <c-page>
        <div>
            <c-button primary @c-click="accessoryDialogToggle">新規追加</c-button>
            <c-accessory-category-edit :visible.sync="accessoryDialogVisible" :model.sync="accessoryCategoryModel" @close="accessoryDialogToggle" />
            <ul v-for="accessoryCategory in accessoryCategories" :key="accessoryCategory.name">
                <li>{{ accessoryCategory }}</li>
            </ul>
            <c-button primary @c-click="materialDialogToggle">新規追加</c-button>
            <c-material-category-edit :visible.sync="materialDialogVisible" :model.sync="materialCategoryModel" @close="materialDialogToggle" />
            <ul v-for="materialCategory in materialCategories" :key="materialCategory.name">
                <li>{{ materialCategory }}</li>
            </ul>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IAccessoryCategory, IMaterialCategory, newAccessoryCategory, newMaterialCategory } from '~/types'
@Component({
    head: {
        title: '商品一覧',
    },
})
export default class PageAdminCategoryIndex extends Vue {
    accessoryCategories: Array<IAccessoryCategory> = []
    materialCategories: Array<IMaterialCategory> = []
    // アクセサリーカテゴリー用のmodalの表示切り替え
    accessoryDialogVisible: boolean = false
    // 材料カテゴリー用のmodalの表示切り替え
    materialDialogVisible: boolean = false
    // form用のaccessoryCategoryModel
    accessoryCategoryModel: IAccessoryCategory = newAccessoryCategory()
    // form用のmaterialCategoryModel
    materialCategoryModel: IMaterialCategory = newMaterialCategory()
    async asyncData({ app }: Context) {
        try {
            const accessoryCategories = await app.$axios.$get(`/accessory_category`)
            const materialCategories = await app.$axios.$get(`/material_category`)
            return { accessoryCategories, materialCategories }
        } catch (e) {
            return { accessoryCategories: [], materialCategories: [] }
        }
    }

    // モーダルの切り替え
    accessoryDialogToggle() {
        this.accessoryDialogVisible = !this.accessoryDialogVisible
    }

    // モーダルの切り替え
    materialDialogToggle() {
        this.materialDialogVisible = !this.materialDialogVisible
    }
}
</script>

<style lang="stylus"></style>
