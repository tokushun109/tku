<template>
    <c-page>
        <div>
            <c-button primary @c-click="accessoryDialogToggle">新規追加</c-button>
            <c-accessory-category-edit
                :visible.sync="accessoryDialogVisible"
                :model.sync="accessoryCategoryModel"
                @close="accessoryDialogToggle"
                @create="loadingCategory($event)"
            />
            <ul v-for="(accessoryCategory, index) in accessoryCategories" :key="index">
                <li>{{ accessoryCategory }}</li>
            </ul>
            <c-button primary @c-click="materialDialogToggle">新規追加</c-button>
            <c-material-category-edit
                :visible.sync="materialDialogVisible"
                :model.sync="materialCategoryModel"
                @close="materialDialogToggle"
                @create="loadingCategory($event)"
            />
            <ul v-for="(materialCategory, index) in materialCategories" :key="index">
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

    async loadingCategory(mode: string) {
        if (mode === 'accessory_category') {
            this.accessoryCategories = await this.$axios.$get(`/accessory_category`)
            this.accessoryCategoryModel = newAccessoryCategory()
        } else if (mode === 'material_category') {
            this.materialCategories = await this.$axios.$get(`/material_category`)
            this.materialCategoryModel = newMaterialCategory()
        }
    }
}
</script>

<style lang="stylus"></style>
