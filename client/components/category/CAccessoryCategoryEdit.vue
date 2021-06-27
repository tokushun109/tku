<template>
    <div class="c-accessory-category-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="accessoryCategoryModel.uuid === '' ? '新しいアクセサリーカテゴリーを登録' : accessoryCategoryModel.name + 'を編集'"
            class="c-accessory-category-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-error :errors.sync="errors" />
            <c-form bordered>
                <c-input-label label="アクセサリーカテゴリー名" required>
                    <c-input :model.sync="accessoryCategoryModel.name" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import { BadRequest, ICategory, ICategoryModelValidation, IError } from '~/types'

@Component({})
export default class CAccessoryCategoryEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') accessoryCategoryModel!: ICategory

    errors: Array<IError> = []

    validation: ICategoryModelValidation = {
        name: false,
    }

    validationReset() {
        this.errors = []
        this.validation.name = false
    }

    // 入力時のバリデーション
    @Watch('accessoryCategoryModel', { deep: true })
    checkValidation() {
        this.validationReset()
        if (this.accessoryCategoryModel.name.length > 20 && !this.validation.name) {
            this.errors.push(new BadRequest('アクセサリーカテゴリー名は20文字以内で入力してください'))
            this.validation.name = true
        }
    }

    async saveHandler() {
        try {
            this.errors = []
            // 送信時のバリデーション
            if (this.accessoryCategoryModel.name.length === 0) {
                throw new BadRequest('材料カテゴリー名が入力されていません')
            }
            await this.$axios.$post(`/accessory_category`, this.accessoryCategoryModel)
            this.dialogVisible = false
            this.$emit('create', 'accessory_category')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
