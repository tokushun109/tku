<template>
    <div class="c-skill-market-edit">
        <c-dialog
            :visible.sync="dialogVisible"
            width="1200px"
            height="450px"
            :title="skiliMarketModel.uuid === '' ? '新しいスキルマーケットを登録' : skiliMarketModel.name + 'を編集'"
            class="c-skill-market-edit-modeal"
            @close="$emit('close')"
            @confirm="saveHandler()"
        >
            <c-form bordered>
                <c-input-label label="スキルマーケット名" required>
                    <c-input :model.sync="skiliMarketModel.name" />
                </c-input-label>
                <c-input-label label="url" required>
                    <c-input :model.sync="skiliMarketModel.url" />
                </c-input-label>
            </c-form>
        </c-dialog>
    </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'nuxt-property-decorator'
import { ISkillMarket } from '~/types'

@Component({})
export default class CSkillMarketEdit extends Vue {
    @PropSync('visible') dialogVisible!: boolean
    @PropSync('model') skiliMarketModel!: ISkillMarket

    async saveHandler() {
        await this.$axios.$post(`/skill_market`, this.skiliMarketModel).catch(() => {})
        this.dialogVisible = false
        this.$emit('create')
    }
}
</script>

<style lang="stylus"></style>
