<template>
    <div class="c-sales-site-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="450px"
            :title="salesSiteModel.uuid === '' ? '新しい販売サイトを登録' : salesSiteModel.name + 'を編集'"
            class="c-sales-site-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-form bordered>
                <c-input-label label="販売サイト名" required>
                    <c-input :model.sync="salesSiteModel.name" />
                </c-input-label>
                <c-input-label label="url" required>
                    <c-input :model.sync="salesSiteModel.url" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { ISalesSite } from '~/types'

@Component({})
export default class CProductEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') salesSiteModel!: ISalesSite

    async saveHandler() {
        await this.$axios.$post(`/sales_site`, this.salesSiteModel).catch(() => {})
        this.dialogVisible = false
        this.$emit('create')
    }
}
</script>

<style lang="stylus"></style>
