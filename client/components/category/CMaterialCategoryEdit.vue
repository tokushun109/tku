<template>
    <div class="c-material-category-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="materialCategoryModel.uuid === '' ? '新しい材料カテゴリーを登録' : materialCategoryModel.name + 'を編集'"
            class="c-material-category-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-error :errors.sync="errors" />
            <c-form bordered>
                <c-input-label label="材料カテゴリー名" required>
                    <c-input :model.sync="materialCategoryModel.name" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import { BadRequest, IError, ICategory, CategoryType, ICategoryModelValidation } from '~/types'

@Component({})
export default class CMaterialCategoryEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') materialCategoryModel!: ICategory

    errors: Array<IError> = []

    validation: ICategoryModelValidation = {
        name: false,
    }

    validationReset() {
        this.errors = []
        this.validation.name = false
    }

    // 入力時のバリデーション
    @Watch('materialCategoryModel', { deep: true })
    checkValidation() {
        this.validationReset()
        if (this.materialCategoryModel.name.length > 20 && !this.validation.name) {
            this.errors.push(new BadRequest('材料カテゴリー名は20文字以内で入力してください'))
            this.validation.name = true
        }
    }

    async saveHandler() {
        try {
            this.errors = []
            // 送信時のバリデーション
            if (this.materialCategoryModel.name.length === 0) {
                throw new BadRequest('材料カテゴリー名が入力されていません')
            }
            await this.$axios.$post(`/material_category`, this.materialCategoryModel)
            this.dialogVisible = false
            this.$emit('create', CategoryType.Material)
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
