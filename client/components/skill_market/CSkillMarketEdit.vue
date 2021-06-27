<template>
    <div class="c-skill-market-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="skiliMarketModel.uuid === '' ? '新しいスキルマーケットを登録' : skiliMarketModel.name + 'を編集'"
            class="c-skill-market-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-error :errors.sync="errors" />
            <c-form bordered>
                <c-input-label label="スキルマーケット名" required>
                    <c-input :model.sync="skiliMarketModel.name" />
                </c-input-label>
                <c-input-label label="URL" required>
                    <c-input :model.sync="skiliMarketModel.url" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import { BadRequest, IError, ISiteModelValidation, ISite } from '~/types'

@Component({})
export default class CSkillMarketEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') skiliMarketModel!: ISite

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
    @Watch('snsModel', { deep: true })
    checkValidation() {
        this.validationReset()
        if (this.skiliMarketModel.name.length > 20 && !this.validation.name) {
            this.errors.push(new BadRequest('スキルマーケット名は20文字以内で入力してください'))
            this.validation.name = true
        }
        if (this.skiliMarketModel.url.match(/^[^\x01-\x7E\xA1-\xDF]+$/) && !this.validation.url) {
            this.errors.push(new BadRequest('URLに全角文字が含まれています'))
            this.validation.url = true
        }
        if (this.skiliMarketModel.url.match(/\s+/) && !this.validation.url) {
            this.errors.push(new BadRequest('URLにスペースが含まれています'))
            this.validation.url = true
        }
    }

    async saveHandler() {
        try {
            this.errors = []
            // 送信時のバリデーション
            if (this.skiliMarketModel.name.length === 0) {
                throw new BadRequest('スキルマーケット名が入力されていません')
            }
            if (this.skiliMarketModel.url.length === 0) {
                throw new BadRequest('URLが入力されていません')
            }
            await this.$axios.$post(`/skill_market`, this.skiliMarketModel)
            this.dialogVisible = false
            this.$emit('create', 'skillMarket')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
