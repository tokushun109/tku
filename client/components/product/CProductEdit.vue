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
            <c-error :errors.sync="errors" />
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
import { Component, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import { BadRequest, IError, IProduct, ICategory, ISite } from '~/types'

interface IProductValidation {
    name: boolean
}

@Component({})
export default class CSalesSiteEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') productModel!: IProduct

    accessoryCategories: Array<ICategory> = []
    materialCategories: Array<ICategory> = []
    salesSites: Array<ISite> = []
    uploadFiles: Array<File> = []

    errors: Array<IError> = []

    validation: IProductValidation = {
        name: false,
    }

    validationReset() {
        this.errors = []
        this.validation.name = false
    }

    // 入力時のバリデーション
    @Watch('productModel', { deep: true })
    checkValidation() {
        this.validationReset()
        if (this.productModel.name.length > 20 && !this.validation.name) {
            this.errors.push(new BadRequest('商品名は20文字以内で入力してください'))
            this.validation.name = true
        }
    }

    async mounted() {
        this.accessoryCategories = await this.$axios.$get(`/accessory_category`)
        this.materialCategories = await this.$axios.$get(`/material_category`)
        this.salesSites = await this.$axios.$get(`/sales_site`)
    }

    async saveHandler() {
        try {
            this.errors = []
            // 送信時のバリデーション
            if (this.productModel.name.length === 0) {
                throw new BadRequest('商品名が入力されていません')
            }
            const createProduct = await this.$axios.$post(`/product`, this.productModel)
            // 画像を選択していたら、アップロードを行う
            if (this.uploadFiles.length > 0) {
                const params = new FormData()
                this.uploadFiles.forEach((file, index) => {
                    params.append(`file${index}`, file)
                })
                await this.$axios.$post(`/product/${createProduct.uuid}/product_image`, params, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                })
            }
            this.dialogVisible = false
            this.$emit('create')
        } catch (e) {
            this.errors.push(e)
        }
    }

    fileUploadHandler(files: Array<File>) {
        this.uploadFiles = files
    }
}
</script>

<style lang="stylus"></style>
