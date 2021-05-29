<template>
    <div class="c-accessory-category-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="350px"
            :title="accessoryCategoryModel.uuid === '' ? '新しいアクセサリーカテゴリーを登録' : accessoryCategoryModel.name + 'を編集'"
            class="c-accessory-category-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-form bordered>
                <c-input-label label="アクセサリーカテゴリー名" required>
                    <c-input :model.sync="accessoryCategoryModel.name" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { IAccessoryCategory } from '~/types'

@Component({})
export default class CAccessoryCategoryEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') accessoryCategoryModel!: IAccessoryCategory

    async saveHandler() {
        await this.$axios.$post(`/accessory_category`, this.accessoryCategoryModel).catch(() => {})
        this.dialogVisible = false
        this.$emit('create', 'accessory_category')
    }
}
</script>

<style lang="stylus"></style>
