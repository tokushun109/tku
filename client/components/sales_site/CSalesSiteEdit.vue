<template>
    <div class="c-sales-site-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="salesSiteModel.uuid === '' ? '新しい販売サイトを登録' : salesSiteModel.name + 'を編集'"
            class="c-sales-site-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-error :errors.sync="errors" />
            <c-form bordered>
                <c-input-label label="販売サイト名" required>
                    <c-input :model.sync="salesSiteModel.name" />
                </c-input-label>
                <c-input-label label="URL" required>
                    <c-input :model.sync="salesSiteModel.url" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import { BadRequest, IError, ISite, ISiteModelValidation } from '~/types'

@Component({})
export default class CSalesSiteEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') salesSiteModel!: ISite

    errors: Array<IError> = []

    validation: ISiteModelValidation = {
        name: false,
        url: false,
    }

    validationReset() {
        this.errors = []
        this.validation.name = false
        this.validation.url = false
    }

    // 入力時のバリデーション
    @Watch('salesSiteModel', { deep: true })
    checkValidation() {
        this.validationReset()
        if (this.salesSiteModel.name.length > 20 && !this.validation.name) {
            this.errors.push(new BadRequest('販売サイト名は20文字以内で入力してください'))
            this.validation.name = true
        }
        if (this.salesSiteModel.url.match(/^[^\x01-\x7E\xA1-\xDF]+$/) && !this.validation.url) {
            this.errors.push(new BadRequest('URLに全角文字が含まれています'))
            this.validation.url = true
        }
        if (this.salesSiteModel.url.match(/\s+/) && !this.validation.url) {
            this.errors.push(new BadRequest('URLにスペースが含まれています'))
            this.validation.url = true
        }
    }

    async saveHandler() {
        try {
            this.errors = []
            // 送信時のバリデーション
            if (this.salesSiteModel.name.length === 0) {
                throw new BadRequest('販売サイト名が入力されていません')
            }
            if (this.salesSiteModel.url.length === 0) {
                throw new BadRequest('URLが入力されていません')
            }
            await this.$axios.$post(`/sales_site`, this.salesSiteModel)
            this.dialogVisible = false
            this.$emit('create', 'salesSite')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
