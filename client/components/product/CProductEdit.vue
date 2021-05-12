<template>
    <div class="c-product-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="productModel === null ? '新しい商品を登録' : productModel.name + 'を編集'"
            class="c-product-edit-modeal"
            @close="$emit('close')"
        >
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { IProduct } from '~/types'

@Component({})
export default class CProductEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') productModel!: IProduct | null
}
</script>

<style lang="stylus"></style>
