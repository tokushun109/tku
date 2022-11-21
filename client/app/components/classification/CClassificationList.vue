<template>
    <v-container class="c-classification-list">
        <v-sheet class="list-header-wrapper">
            <div class="list-header">
                <h3 class="title list-title">{{ categoryTypeValue }}</h3>
                <v-spacer />
                <c-icon :type="IconType.New.name" @c-click="openHandler(ExecutionType.Create)" />
            </div>
            <v-divider />
            <v-list class="list-content-wrapper" dense>
                <c-message v-if="listItems.length === 0" class="mt-4"> 登録されていません </c-message>
                <v-list-item v-for="listItem in listItems" v-else :key="listItem.uuid" class="list-content">
                    <v-list-item-content>
                        <v-list-item-title class="list-content-title">
                            <div>{{ listItem.name }}</div>
                            <v-spacer />
                            <c-icon :type="IconType.Edit.name" @c-click="openHandler(ExecutionType.Edit, listItem)" />
                            <c-icon :type="IconType.Delete.name" @c-click="openHandler(ExecutionType.Delete, listItem)" />
                        </v-list-item-title>
                        <v-divider />
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
                    <v-text-field v-model="modalItem.name" :rules="nameRules" label="カテゴリー名(必須)" outlined counter="20" />
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
import { IClassification, IError, CategoryType, ExecutionType, IconType, TExecutionType } from '~/types'
import { min20, newClassification, required } from '~/methods'
@Component({})
export default class CClassificationList extends Vue {
    @PropSync('items') listItems!: Array<IClassification>
    @Prop({ type: String, default: '' }) type!: string

    IconType: typeof IconType = IconType
    ExecutionType: typeof ExecutionType = ExecutionType
    executionType: TExecutionType = ExecutionType.Create

    // ダイアログの表示
    dialogVisible: boolean = false

    // 通知の表示
    notificationVisible: boolean = false

    modalItem: IClassification = newClassification()

    valid: boolean = true

    errors: Array<IError> = []

    nameRules = [required, min20]

    get categoryTypeValue(): string {
        let categoryType = ''
        for (const type in CategoryType) {
            if (this.type === CategoryType[type].name) {
                categoryType = CategoryType[type].value
            }
        }
        return categoryType
    }

    get modalTitle(): string {
        let title = ''
        for (const type in ExecutionType) {
            if (this.executionType === ExecutionType[type]) {
                title = `${this.categoryTypeValue}の${ExecutionType[type]}`
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
        this.modalItem = newClassification()
    }

    setItem(item: IClassification) {
        this.modalItem = _.cloneDeep(item)
    }

    @Watch('dialogVisible')
    resetValidation() {
        if (!this.dialogVisible && this.executionType !== ExecutionType.Delete) {
            const refs: any = this.$refs.form
            refs.resetValidation()
        }
    }

    openHandler(executionType: TExecutionType, item: IClassification | null = null) {
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
                if (this.categoryTypeValue === CategoryType.Category.value) {
                    await this.$axios.$post(`/category`, this.modalItem)
                    this.$emit('c-change', CategoryType.Category.name)
                } else if (this.categoryTypeValue === CategoryType.Tag.value) {
                    await this.$axios.$post(`/tag`, this.modalItem)
                    this.$emit('c-change', CategoryType.Tag.name)
                }
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e)
            }
        } else if (this.executionType === ExecutionType.Edit) {
            try {
                if (this.categoryTypeValue === CategoryType.Category.value) {
                    await this.$axios.$put(`/category/${this.modalItem.uuid}`, this.modalItem)
                    this.$emit('c-change', CategoryType.Category.name)
                } else if (this.categoryTypeValue === CategoryType.Tag.value) {
                    await this.$axios.$put(`/tag/${this.modalItem.uuid}`, this.modalItem)
                    this.$emit('c-change', CategoryType.Tag.name)
                }
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e)
            }
        } else if (this.executionType === ExecutionType.Delete) {
            try {
                if (this.categoryTypeValue === CategoryType.Category.value) {
                    await this.$axios.$delete(`/category/${this.modalItem.uuid}`)
                    this.$emit('c-change', CategoryType.Category.name)
                } else if (this.categoryTypeValue === CategoryType.Tag.value) {
                    await this.$axios.$delete(`/tag/${this.modalItem.uuid}`)
                    this.$emit('c-change', CategoryType.Tag.name)
                }
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e)
            }
        }
    }
}
</script>

<style lang="stylus" scoped>
.c-classification-list
    .list-header-wrapper
        padding 16px
        .list-header
            display flex
            .list-title
                color $title-text-color
    .list-content-wrapper
        .list-content
            .list-content-title
                display flex
</style>