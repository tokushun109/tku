<template>
    <div class="c-sns-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="450px"
            :title="snsModel.uuid === '' ? '新しいSNSを登録' : snsModel.name + 'を編集'"
            class="c-sns-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-form bordered>
                <c-input-label label="SNS名" required>
                    <c-input :model.sync="snsModel.name" />
                </c-input-label>
                <c-input-label label="url" required>
                    <c-input :model.sync="snsModel.url" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { ISns } from '~/types'

@Component({})
export default class CSnsEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') snsModel!: ISns

    async saveHandler() {
        await this.$axios.$post(`/sns`, this.snsModel).catch(() => {})
        this.dialogVisible = false
        this.$emit('create')
    }
}
</script>

<style lang="stylus"></style>
