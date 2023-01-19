<template>
    <v-container class="c-site-list">
        <v-sheet class="list-header-wrapper">
            <div class="list-header">
                <h3 class="title list-title">{{ siteTypeValue }}</h3>
                <v-spacer />
                <c-icon :type="IconType.New.name" @c-click="openHandler(ExecutionType.Create)" />
            </div>
            <v-divider class="divider" />
            <v-list class="list-content-wrapper" dense>
                <c-message v-if="listItems.length === 0" class="no-record-message"> 登録されていません </c-message>
                <v-list-item v-for="listItem in listItems" v-else :key="listItem.uuid" class="list-content">
                    <v-list-item-content>
                        <v-list-item-title class="list-content-title">
                            <div>{{ listItem.name }}</div>
                            <v-spacer />
                            <c-icon :type="IconType.Edit.name" @c-click="openHandler(ExecutionType.Edit, listItem)" />
                            <c-icon :type="IconType.Delete.name" @c-click="openHandler(ExecutionType.Delete, listItem)" />
                        </v-list-item-title>
                        <v-divider class="divider" />
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-sheet>
        <c-dialog :visible.sync="dialogVisible" :title="modalTitle" :confirm-button-disabled="!valid" @confirm="confirmHandler" @close="closeHandler">
            <template #content>
                <c-error :errors.sync="errors" />
                <v-form
                    v-if="executionType === ExecutionType.Create || executionType === ExecutionType.Edit"
                    ref="form"
                    v-model="valid"
                    lazy-validation
                >
                    <v-text-field v-model="modalItem.name" :rules="nameRules" label="サイト名(必須)" outlined counter="20" />
                    <v-text-field v-model="modalItem.url" :rules="urlRules" label="URL" outlined />
                    <v-container v-if="modalItem.icon" class="icon-preview">
                        <p>アイコンプレビュー</p>
                        <v-btn fab :color="ColorType.Orange" :href="modalItem.url" x-large>
                            <v-icon :color="ColorType.White">{{ modalItem.icon }}</v-icon>
                        </v-btn>
                    </v-container>
                </v-form>
                <p v-else-if="executionType === ExecutionType.Delete">削除してもよろしいですか？</p>
            </template>
        </c-dialog>
        <c-notification :visible.sync="notificationVisible">{{ notificationMessage }}</c-notification>
    </v-container>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue, Watch } from 'nuxt-property-decorator'
import _ from 'lodash'
import { IError, ISite, ExecutionType, IconType, SiteType, TExecutionType, ColorType } from '~/types'
import { min20, newSite, nonDoubleByte, nonSpace, required } from '~/methods'
@Component({})
export default class CSiteList extends Vue {
    @PropSync('items') listItems!: Array<ISite>
    @Prop({ type: String, default: '' }) type!: string

    IconType: typeof IconType = IconType
    ExecutionType: typeof ExecutionType = ExecutionType
    executionType: TExecutionType = ExecutionType.Create
    SiteType: typeof SiteType = SiteType
    ColorType: typeof ColorType = ColorType

    // ダイアログの表示
    dialogVisible: boolean = false

    // 通知の表示
    notificationVisible: boolean = false

    modalItem: ISite = newSite()

    valid: boolean = true

    errors: Array<IError> = []

    nameRules = [required, min20]

    urlRules = [required, nonDoubleByte, nonSpace]

    get siteTypeValue(): string {
        let siteType = ''
        for (const type in SiteType) {
            if (this.type === SiteType[type].name) {
                siteType = SiteType[type].value
            }
        }
        return siteType
    }

    get modalTitle(): string {
        let title = ''
        for (const type in ExecutionType) {
            if (this.executionType === ExecutionType[type]) {
                title = `${this.siteTypeValue}の${ExecutionType[type]}`
            }
        }
        return title
    }

    get notificationMessage(): string {
        let message = ''
        for (const type in ExecutionType) {
            if (this.executionType === ExecutionType[type]) {
                message = `${this.modalItem.name}を${ExecutionType[type]}しました`
            }
        }
        return message
    }

    setInit() {
        this.modalItem = newSite()
    }

    setItem(item: ISite) {
        this.modalItem = _.cloneDeep(item)
    }

    @Watch('dialogVisible')
    resetValidation() {
        if (!this.dialogVisible && this.executionType !== ExecutionType.Delete) {
            const refs: any = this.$refs.form
            refs.resetValidation()
        }
    }

    openHandler(executionType: TExecutionType, item: ISite | null = null) {
        this.errors = []
        if (executionType === ExecutionType.Create) {
            this.setInit()
        } else {
            this.setItem(item!)
        }
        this.executionType = executionType
        this.dialogVisible = true
    }

    closeHandler() {
        this.dialogVisible = false
    }

    async confirmHandler() {
        this.errors = []
        if (this.executionType === ExecutionType.Create) {
            try {
                if (this.siteTypeValue === SiteType.Sns.value) {
                    await this.$axios.$post(`/sns`, this.modalItem)
                    this.$emit('c-change', SiteType.Sns.name)
                } else if (this.siteTypeValue === SiteType.SalesSite.value) {
                    await this.$axios.$post(`/sales_site`, this.modalItem, { withCredentials: true })
                    this.$emit('c-change', SiteType.SalesSite.name)
                } else if (this.siteTypeValue === SiteType.SkillMarket.value) {
                    await this.$axios.$post(`/skill_market`, this.modalItem)
                    this.$emit('c-change', SiteType.SkillMarket.name)
                }
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e.response)
            }
        } else if (this.executionType === ExecutionType.Edit) {
            try {
                if (this.siteTypeValue === SiteType.Sns.value) {
                    await this.$axios.$put(`/sns/${this.modalItem.uuid}`, this.modalItem)
                    this.$emit('c-change', SiteType.Sns.name)
                } else if (this.siteTypeValue === SiteType.SalesSite.value) {
                    await this.$axios.$put(`/sales_site/${this.modalItem.uuid}`, this.modalItem, { withCredentials: true })
                    this.$emit('c-change', SiteType.SalesSite.name)
                } else if (this.siteTypeValue === SiteType.SkillMarket.value) {
                    await this.$axios.$put(`/skill_market/${this.modalItem.uuid}`, this.modalItem)
                    this.$emit('c-change', SiteType.SkillMarket.name)
                }
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e.response)
            }
        } else if (this.executionType === ExecutionType.Delete) {
            try {
                if (this.siteTypeValue === SiteType.Sns.value) {
                    await this.$axios.$delete(`/sns/${this.modalItem.uuid}`)
                    this.$emit('c-change', SiteType.Sns.name)
                } else if (this.siteTypeValue === SiteType.SalesSite.value) {
                    await this.$axios.$delete(`/sales_site/${this.modalItem.uuid}`, { withCredentials: true })
                    this.$emit('c-change', SiteType.SalesSite.name)
                } else if (this.siteTypeValue === SiteType.SkillMarket.value) {
                    await this.$axios.$delete(`/skill_market/${this.modalItem.uuid}`)
                    this.$emit('c-change', SiteType.SkillMarket.name)
                }
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e.response)
            }
        }
    }
}
</script>

<style lang="stylus" scoped>
.c-site-list
    .list-header-wrapper
        padding 16px
        .list-header
            display flex
            .list-title
                color $title-primary-color
    .list-content-wrapper
        .no-record-message
            margin-top 16px
        .list-content
            .list-content-title
                display flex

.icon-preview
    margin-bottom 32px
</style>
