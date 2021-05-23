<template>
    <div class="c-product-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="productModel.uuid === '' ? '新しい商品を登録' : productModel.name + 'を編集'"
            class="c-product-edit-modeal"
            @close="$emit('close')"
        >
            <c-form bordered>
                <c-input-label label="商品名">
                    <c-input :model.sync="productModel.name" />
                </c-input-label>
                <c-input-label label="商品説明">
                    <c-input :model.sync="productModel.description" multiline />
                </c-input-label>
                <c-input-label label="商品画像">
                    <c-file-upload @c-file-uploaded="fileUploadHandler($event)" />
                </c-input-label>
                <c-input-label label="アクセサリーカテゴリー">
                    <c-dropdown name="accessory-category" :items="accessoryCategories" :model.sync="productModel.accessoryCategory" property="name" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { IAccessoryCategory, IMaterialCategory, IProduct } from '~/types'

@Component({})
export default class CProductEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') productModel!: IProduct

    accessoryCategories: Array<IAccessoryCategory> = []
    materialCategories: Array<IMaterialCategory> = []

    async mounted() {
        this.accessoryCategories = await this.$axios.$get(`/accessory_category`)
    }

    fileUploadHandler(files: Array<File>) {
        console.log(files)
    }
}
</script>

<style lang="stylus"></style>
