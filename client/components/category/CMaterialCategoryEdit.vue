<template>
    <div class="c-material-category-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="350px"
            :title="materialCategoryModel.uuid === '' ? '新しい材料カテゴリーを登録' : materialCategoryModel.name + 'を編集'"
            class="c-material-category-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-form bordered>
                <c-input-label label="材料カテゴリー名" required>
                    <c-input :model.sync="materialCategoryModel.name" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { IMaterialCategory } from '~/types'

@Component({})
export default class CMaterialCategoryEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') materialCategoryModel!: IMaterialCategory

    async saveHandler() {
        await this.$axios.$post(`/material_category`, this.materialCategoryModel).catch(() => {})
        this.dialogVisible = false
        this.$emit('create', 'material_category')
    }
}
</script>

<style lang="stylus"></style>
