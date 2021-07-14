<template>
    <c-page>
        <div class="admin-category-list">
            <c-button primary @c-click="accssoryCategoryCreateHandler">新規追加</c-button>
            <c-accessory-category-edit
                :visible.sync="accessoryDialogVisible"
                :model.sync="accessoryCategoryModel"
                @close="accessoryDialogToggle"
                @create="loadingCategory($event)"
            />
            <ul v-for="accessoryCategory in accessoryCategories" :key="accessoryCategory.uuid">
                <c-column>
                    <li>{{ accessoryCategory.name }}</li>
                    <div class="button-wrapper">
                        <c-button class="edit-button" label="修正" primary @c-click="accssoryCategoryEditHandler(accessoryCategory)" />
                        <c-button class="delete-button" label="削除" @c-click="accssoryCategoryDeleteHandler(accessoryCategory)" />
                    </div>
                </c-column>
            </ul>
            <c-button primary @c-click="materialCategoryCreateHandler">新規追加</c-button>
            <c-material-category-edit
                :visible.sync="materialDialogVisible"
                :model.sync="materialCategoryModel"
                @close="materialDialogToggle"
                @create="loadingCategory($event)"
            />
            <ul v-for="materialCategory in materialCategories" :key="materialCategory.uuid">
                <c-column>
                    <li>{{ materialCategory.name }}</li>
                    <div class="button-wrapper">
                        <c-button class="edit-button" label="修正" primary @c-click="materialCategoryEditHandler(materialCategory)" />
                        <c-button class="delete-button" label="削除" @c-click="materialCategoryDeleteHandler(materialCategory)" />
                    </div>
                </c-column>
            </ul>
        </div>
    </c-page>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import _ from 'lodash'
import { ICategory, newCategory, CategoryType } from '~/types'
@Component({
    head: {
        title: '商品一覧',
    },
})
export default class PageAdminCategoryIndex extends Vue {
    // アクセサリーカテゴリー一覧
    accessoryCategories: Array<ICategory> = []

    // 材料カテゴリー一覧
    materialCategories: Array<ICategory> = []

    // アクセサリーカテゴリー用のmodalの表示切り替え
    accessoryDialogVisible: boolean = false

    // 材料カテゴリー用のmodalの表示切り替え
    materialDialogVisible: boolean = false

    // form用のaccessoryCategoryModel
    accessoryCategoryModel: ICategory = newCategory()

    // form用のmaterialCategoryModel
    materialCategoryModel: ICategory = newCategory()

    // アクセサリーカテゴリーダイアログの切り替え
    accessoryDialogToggle() {
        this.accessoryDialogVisible = !this.accessoryDialogVisible
        // 編集画面から切り替えるときはmodelの中身を空にする
    }

    // 材料カテゴリーダイアログの切り替え
    materialDialogToggle() {
        this.materialDialogVisible = !this.materialDialogVisible
    }

    // アクセサリーカテゴリの新規追加
    accssoryCategoryCreateHandler() {
        this.accessoryCategoryModel = newCategory()
        this.accessoryDialogToggle()
    }

    // アクセサリーカテゴリの更新
    accssoryCategoryEditHandler(accessoryCategory: ICategory) {
        this.accessoryCategoryModel = _.cloneDeep(accessoryCategory)
        this.accessoryDialogToggle()
    }

    // アクセサリーカテゴリーの削除
    async accssoryCategoryDeleteHandler(accessoryCategory: ICategory) {
        if (confirm(`${accessoryCategory.name}を削除します。よろしいですか？`)) {
            await this.$axios.$delete(`/accessory_category/${accessoryCategory.uuid}`)
            this.loadingCategory(CategoryType.Accessory)
        }
    }

    // 材料カテゴリーの新規追加
    materialCategoryCreateHandler() {
        this.materialCategoryModel = newCategory()
        this.materialDialogToggle()
    }

    // 材料カテゴリーの更新
    materialCategoryEditHandler(materialCategory: ICategory) {
        this.materialCategoryModel = _.cloneDeep(materialCategory)
        this.materialDialogToggle()
    }

    // 材料カテゴリーの削除
    async materialCategoryDeleteHandler(materialCategory: ICategory) {
        if (confirm(`${materialCategory.name}を削除します。よろしいですか？`)) {
            await this.$axios.$delete(`/material_category/${materialCategory.uuid}`)
            this.loadingCategory(CategoryType.Material)
        }
    }

    async loadingCategory(type: string) {
        if (type === CategoryType.Accessory) {
            this.accessoryCategories = await this.$axios.$get(`/accessory_category`)
            this.accessoryCategoryModel = newCategory()
        } else if (type === CategoryType.Material) {
            this.materialCategories = await this.$axios.$get(`/material_category`)
            this.materialCategoryModel = newCategory()
        }
    }

    async asyncData({ app }: Context) {
        try {
            const accessoryCategories = await app.$axios.$get(`/accessory_category`)
            const materialCategories = await app.$axios.$get(`/material_category`)
            return { accessoryCategories, materialCategories }
        } catch (e) {
            return { accessoryCategories: [], materialCategories: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
