<template>
    <v-container>
        <v-sheet class="pa-4 lighten-4">
            <div class="d-flex">
                <h3 class="title green--text text--darken-3">{{ siteTypeValue }}</h3>
                <v-spacer />
                <c-dialog-2
                    :visible.sync="createDialogVisible"
                    :title="`${siteTypeValue}の追加`"
                    width="800"
                    @confirm="saveHandler"
                    @close="closeHandler"
                >
                    <template #trigger>
                        <c-icon type="new" @c-click="setInit" />
                    </template>
                    <template #content>
                        <v-text-field v-model="modalItem.name" :rules="nameRules" label="サイト名" outlined counter="20" />
                        <v-text-field v-model="modalItem.url" :rules="urlRules" label="URL" outlined />
                    </template>
                </c-dialog-2>
            </div>
            <v-divider />
            <v-list dense>
                <c-message v-if="listItems.length === 0" class="mt-4"> 登録されていません </c-message>
                <v-list-item v-for="listItem in listItems" v-else :key="listItem.uuid">
                    <v-list-item-content>
                        <v-list-item-title class="d-flex">
                            <div>{{ listItem.name }}</div>
                            <v-spacer />
                            <c-dialog-2
                                :visible.sync="editDialogVisible"
                                :title="`${siteTypeValue}の追加`"
                                width="800"
                                @confirm="editHandler"
                                @close="closeHandler"
                            >
                                <template #trigger>
                                    <c-icon type="edit" @c-click="setItem(listItem)" />
                                </template>
                                <template #content>
                                    <v-text-field v-model="modalItem.name" :rules="nameRules" label="サイト名" outlined counter="20" />
                                    <v-text-field v-model="modalItem.url" :rules="urlRules" label="URL" outlined />
                                </template>
                            </c-dialog-2>
                            <c-dialog-2
                                :visible.sync="deleteDialogVisible"
                                :title="`${listItem.name}の更新`"
                                width="400"
                                @confirm="deleteHandler"
                                @close="closeHandler"
                            >
                                <template #trigger>
                                    <c-icon type="delete" @c-click="setItem(listItem)" />
                                </template>
                                <template #content>
                                    <p>削除してもよろしいですか？</p>
                                </template>
                            </c-dialog-2>
                        </v-list-item-title>
                        <v-divider />
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-sheet>
        <c-notification :visible.sync="createNotificationVisible">{{ `${modalItem.name}を作成しました` }}</c-notification>
        <c-notification :visible.sync="editNotificationVisible">{{ `${modalItem.name}を更新しました` }}</c-notification>
        <c-notification :visible.sync="deleteNotificationVisible">{{ `${modalItem.name}を削除しました` }}</c-notification>
    </v-container>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue } from 'nuxt-property-decorator'
import _ from 'lodash'
import { IError, ISite, newSite, SiteType } from '~/types'
import { min20, nonDoubleByte, nonSpace, required } from '~/methods'
@Component({})
export default class CSiteList extends Vue {
    @PropSync('items') listItems!: Array<ISite>
    @Prop({ type: String, default: '' }) type!: string

    // 新規作成ダイアログの表示
    createDialogVisible: boolean = false

    // 更新ダイアログの表示
    editDialogVisible: boolean = false

    // 削除ダイアログの表示
    deleteDialogVisible: boolean = false

    // 新規作成通知の表示
    createNotificationVisible: boolean = false

    // 更新通知の表示
    editNotificationVisible: boolean = false

    // 削除通知の表示
    deleteNotificationVisible: boolean = false

    modalItem: ISite = newSite()

    errors: Array<IError> = []

    nameRules = [required, min20]

    urlRules = [nonDoubleByte, nonSpace]

    get siteTypeValue(): string {
        let siteType = ''
        for (const type in SiteType) {
            if (this.type === SiteType[type].name) {
                siteType = SiteType[type].value
            }
        }
        return siteType
    }

    async saveHandler() {
        try {
            if (this.siteTypeValue === SiteType.Sns.value) {
                await this.$axios.$post(`/sns`, this.modalItem)
                this.$emit('c-change', SiteType.Sns.name)
            } else if (this.siteTypeValue === SiteType.SalesSite.value) {
                await this.$axios.$post(`/sales_site`, this.modalItem)
                this.$emit('c-change', SiteType.SalesSite.name)
            } else if (this.siteTypeValue === SiteType.SkillMarket.value) {
                await this.$axios.$post(`/skill_market`, this.modalItem)
                this.$emit('c-change', SiteType.SkillMarket.name)
            }
            this.createNotificationVisible = true
            this.createDialogVisible = false
        } catch (e) {
            this.errors.push(e)
        }
    }

    async editHandler() {
        try {
            if (this.siteTypeValue === SiteType.Sns.value) {
                await this.$axios.$put(`/sns/${this.modalItem.uuid}`, this.modalItem)
                this.$emit('c-change', SiteType.Sns.name)
            } else if (this.siteTypeValue === SiteType.SalesSite.value) {
                await this.$axios.$put(`/sales_site/${this.modalItem.uuid}`, this.modalItem)
                this.$emit('c-change', SiteType.SalesSite.name)
            } else if (this.siteTypeValue === SiteType.SkillMarket.value) {
                await this.$axios.$put(`/skill_market/${this.modalItem.uuid}`, this.modalItem)
                this.$emit('c-change', SiteType.SkillMarket.name)
            }
            this.editNotificationVisible = true
            this.editDialogVisible = false
        } catch (e) {
            this.errors.push(e)
        }
    }

    async deleteHandler() {
        try {
            if (this.siteTypeValue === SiteType.Sns.value) {
                await this.$axios.$delete(`/sns/${this.modalItem.uuid}`)
                this.$emit('c-change', SiteType.Sns.name)
            } else if (this.siteTypeValue === SiteType.SalesSite.value) {
                await this.$axios.$delete(`/sales_site/${this.modalItem.uuid}`)
                this.$emit('c-change', SiteType.SalesSite.name)
            } else if (this.siteTypeValue === SiteType.SkillMarket.value) {
                await this.$axios.$delete(`/skill_market/${this.modalItem.uuid}`)
                this.$emit('c-change', SiteType.SkillMarket.name)
            }
            this.deleteNotificationVisible = true
            this.deleteDialogVisible = false
        } catch (e) {
            this.errors.push(e)
        }
    }

    setInit() {
        this.modalItem = newSite()
    }

    setItem(item: ISite) {
        this.modalItem = _.cloneDeep(item)
    }

    closeHandler() {
        this.createDialogVisible = false
        this.editDialogVisible = false
        this.deleteDialogVisible = false
    }
}
</script>

<style lang="stylus">
.v-application.error--text
    color black !important
</style>
