<template>
    <div class="c-sns-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            :title="snsModel.uuid === '' ? '新しいSNSを登録' : snsModel.name + 'を編集'"
            class="c-sns-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-error :errors.sync="errors" />
            <c-form bordered>
                <c-input-label label="SNS名" required>
                    <c-input :model.sync="snsModel.name" />
                </c-input-label>
                <c-input-label label="URL" required>
                    <c-input :model.sync="snsModel.url" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import { BadRequest, IError, ISns } from '~/types'

interface ISnsModelValidation {
    name: boolean
    url: boolean
}

@Component({})
export default class CSnsEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') snsModel!: ISns

    errors: Array<IError> = []

    validation: ISnsModelValidation = {
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
        if (this.snsModel.name.length > 20 && !this.validation.name) {
            this.errors.push(new BadRequest('SNS名は20文字以内で入力してください'))
            this.validation.name = true
        }
        if (this.snsModel.url.match(/^[^\x01-\x7E\xA1-\xDF]+$/) && !this.validation.url) {
            this.errors.push(new BadRequest('URLに全角文字が含まれています'))
            this.validation.url = true
        }
        if (this.snsModel.url.match(/\s+/) && !this.validation.url) {
            this.errors.push(new BadRequest('URLにスペースが含まれています'))
            this.validation.url = true
        }
    }

    async saveHandler() {
        try {
            this.errors = []
            // 送信時のバリデーション
            if (this.snsModel.name.length === 0) {
                throw new BadRequest('SNS名が入力されていません')
            }
            if (this.snsModel.url.length === 0) {
                throw new BadRequest('URLが入力されていません')
            }
            await this.$axios.$post(`/sns`, this.snsModel)
            this.dialogVisible = false
            this.$emit('create')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
