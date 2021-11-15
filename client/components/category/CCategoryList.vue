<template>
    <v-container>
        <v-sheet class="pa-4 lighten-4">
            <div class="d-flex">
                <h3 class="title green--text text--darken-3">{{ categoryTypeValue }}</h3>
                <v-spacer />
                <c-dialog-2
                    :visible.sync="createDialogVisible"
                    :title="`${categoryTypeValue}の追加`"
                    width="800"
                    @confirm="saveHandler"
                    @close="closeHandler"
                >
                    <template #trigger>
                        <c-icon type="new" @c-click="setInit" />
                    </template>
                    <template #content>
                        <v-text-field v-model="modalItem.name" :rules="nameRules" label="カテゴリー名" outlined counter="20" />
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
                                :title="`${categoryTypeValue}の追加`"
                                width="800"
                                @confirm="editHandler"
                                @close="closeHandler"
                            >
                                <template #trigger>
                                    <c-icon type="edit" @c-click="setItem(listItem)" />
                                </template>
                                <template #content>
                                    <v-text-field v-model="modalItem.name" :rules="nameRules" label="カテゴリー名" outlined counter="20" />
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
import { CategoryType, ICategory, IError, ISite, newCategory } from '~/types'
import { min20, nonDoubleByte, nonSpace, required } from '~/methods'
@Component({})
export default class CCategoryList extends Vue {
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

    modalItem: ICategory = newCategory()

    errors: Array<IError> = []

    nameRules = [required, min20]

    urlRules = [nonDoubleByte, nonSpace]

    get categoryTypeValue(): string {
        let categoryType = ''
        for (const type in CategoryType) {
            if (this.type === CategoryType[type].name) {
                categoryType = CategoryType[type].value
            }
        }
        return categoryType
    }

    async saveHandler() {
        try {
            if (this.categoryTypeValue === CategoryType.Accessory.value) {
                await this.$axios.$post(`/accessory_category`, this.modalItem)
                this.$emit('c-change', CategoryType.Accessory.name)
            } else if (this.categoryTypeValue === CategoryType.Material.value) {
                await this.$axios.$post(`/material_category`, this.modalItem)
                this.$emit('c-change', CategoryType.Material.name)
            }
            this.createNotificationVisible = true
            this.createDialogVisible = false
        } catch (e) {
            this.errors.push(e)
        }
    }

    async editHandler() {
        try {
            if (this.categoryTypeValue === CategoryType.Accessory.value) {
                await this.$axios.$put(`/accessory_category/${this.modalItem}`, this.modalItem)
                this.$emit('c-change', CategoryType.Accessory.name)
            } else if (this.categoryTypeValue === CategoryType.Material.value) {
                await this.$axios.$put(`/material_category/${this.modalItem}`, this.modalItem)
                this.$emit('c-change', CategoryType.Material.name)
            }
            this.editNotificationVisible = true
            this.editDialogVisible = false
        } catch (e) {
            this.errors.push(e)
        }
    }

    async deleteHandler() {
        try {
            if (this.categoryTypeValue === CategoryType.Accessory.value) {
                await this.$axios.$delete(`/accessory_category/${this.modalItem.uuid}`)
                this.$emit('c-change', CategoryType.Accessory.name)
            } else if (this.categoryTypeValue === CategoryType.Material.value) {
                await this.$axios.$delete(`/material_category/${this.modalItem.uuid}`)
                this.$emit('c-change', CategoryType.Material.name)
            }
            this.deleteNotificationVisible = true
            this.deleteDialogVisible = false
        } catch (e) {
            this.errors.push(e)
        }
    }

    setInit() {
        this.modalItem = newCategory()
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

<style lang="stylus"></style>
