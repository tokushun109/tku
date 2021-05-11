<template>
    <div class="c-product-edit">
        <c-modal
            :visible.sync="modalVisible"
            width="1200px"
            :title="editProductModel === null ? '新しい商品を登録' : editProductModel.name + 'を編集'"
            class="c-product-edit-modeal"
            @close="$emit('close')"
        >
        </c-modal>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { IProduct } from '~/types'

@Component({})
export default class CProductEdit extends Vue {
    @PropSync('visible') modalVisible!: boolean
    @PropSync('model') editProductModel!: IProduct | null
}
</script>

<style lang="stylus"></style>
