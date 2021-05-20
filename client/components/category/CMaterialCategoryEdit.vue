<template>
    <div class="c-material-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="300px"
            :title="materialCategoryModel.name === '' ? '新しい材料カテゴリーを登録' : materialCategoryModel.name + 'を編集'"
            class="c-material-edit-modeal"
            @close="$emit('close')"
        >
            <c-form bordered>
                <c-input :model.sync="materialCategoryModel.name" />
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
}
</script>

<style lang="stylus"></style>
