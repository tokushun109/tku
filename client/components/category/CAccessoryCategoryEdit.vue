<template>
    <div class="c-accessory-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="300px"
            :title="accessoryCategoryModel.name === '' ? '新しいアクセサリーカテゴリーを登録' : accessoryCategoryModel.name + 'を編集'"
            class="c-accessory-edit-modeal"
            @close="$emit('close')"
        >
            <c-form bordered>
                <c-input :model.sync="accessoryCategoryModel.name" />
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
}
</script>

<style lang="stylus"></style>
