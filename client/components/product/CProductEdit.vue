<template>
    <div class="c-product-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="productModel.uuid === '' ? '新しい商品を登録' : productModel.name + 'を編集'"
            class="c-product-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-form bordered>
                <c-input-label label="商品名" required>
                    <c-input :model.sync="productModel.name" />
                </c-input-label>
                <c-input-label label="商品説明">
                    <c-input :model.sync="productModel.description" multiline />
                </c-input-label>
                <c-input-label label="商品画像">
                    <c-file-upload @c-file-uploaded="fileUploadHandler($event)" />
                </c-input-label>
                <c-input-label label="アクセサリーカテゴリー">
                    <c-dropdown name="accessory-category" :items="accessoryCategories" :model.sync="productModel.accessoryCategory" />
                </c-input-label>
                <c-input-label label="材料カテゴリー">
                    <c-dropdown name="material-category" :items="materialCategories" :model.sync="productModel.materialCategories" multiple />
                </c-input-label>
                <c-input-label label="販売サイト">
                    <c-dropdown name="sales-site" :items="salesSites" :model.sync="productModel.salesSites" multiple />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { IAccessoryCategory, IMaterialCategory, IProduct, ISalesSite } from '~/types'

@Component({})
export default class CSalesSiteEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') productModel!: IProduct

    accessoryCategories: Array<IAccessoryCategory> = []
    materialCategories: Array<IMaterialCategory> = []
    salesSites: Array<ISalesSite> = []

    async mounted() {
        this.accessoryCategories = await this.$axios.$get(`/accessory_category`)
        this.materialCategories = await this.$axios.$get(`/material_category`)
        this.salesSites = await this.$axios.$get(`/sales_site`)
    }

    async saveHandler() {
        await this.$axios.$post(`/product`, this.productModel).catch(() => {})
        this.dialogVisible = false
        this.$emit('create')
    }

    fileUploadHandler(files: Array<File>) {
        console.log(files)
    }
}
</script>

<style lang="stylus"></style>
