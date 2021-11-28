<template>
    <v-container>
        <v-sheet class="pa-4 lighten-4">
            <v-container>
                <div class="d-flex">
                    <h3 class="title green--text text--darken-3">商品</h3>
                    <v-spacer />
                    <c-icon :type="IconType.New.name" @c-click="openHandler(ExecutionType.Create)" />
                </div>
                <v-divider />
                <v-list>
                    <c-message v-if="listItems.length === 0" class="mt-4"> 登録されていません </c-message>
                    <v-row>
                        <v-col v-for="listItem in listItems" :key="listItem.uuid" cols="12" sm="6" md="4">
                            <v-list-item>
                                <v-card width="100%" color="light-green lighten-5">
                                    <v-card-text>
                                        <div class="my-4">
                                            <div class="d-flex">
                                                <h3 class="green--text text--darken-3">{{ listItem.name }}</h3>
                                                <v-spacer />
                                                <div>
                                                    <c-icon :type="IconType.Edit.name" @c-click="openHandler(ExecutionType.Edit, listItem)" />
                                                    <c-icon :type="IconType.Delete.name" @c-click="openHandler(ExecutionType.Delete, listItem)" />
                                                </div>
                                            </div>
                                            <v-divider />
                                        </div>
                                        <v-carousel
                                            v-if="listItem.productImages.length > 0"
                                            :show-arrows="listItem.productImages.length > 1"
                                            height="auto"
                                            hide-delimiters
                                        >
                                            <v-carousel-item v-for="image in listItem.productImages" :key="image.uuid">
                                                <v-img :src="image.apiPath" :alt="image.uuid" />
                                            </v-carousel-item>
                                        </v-carousel>
                                        <v-carousel v-else :show-arrows="false" height="auto" hide-delimiters>
                                            <v-carousel-item>
                                                <v-img src="/img/product/no-image.png" />
                                            </v-carousel-item>
                                        </v-carousel>
                                    </v-card-text>
                                </v-card>
                            </v-list-item>
                        </v-col>
                    </v-row>
                </v-list>
            </v-container>
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
                    <v-text-field v-model="modalItem.name" :rules="nameRules" label="商品名(必須)" outlined counter="20" />
                    <v-textarea v-model="modalItem.description" label="商品説明" outlined />
                    <v-file-input v-model="uploadFiles" label="商品画像" prepend-icon="mdi-camera" multiple outlined />
                    <c-image-list
                        title="現在の登録"
                        :registered-list="registeredList"
                        :preview-list="previewList"
                        @c-delete-image-handler="deleteImageHandler"
                    />
                    <v-select
                        v-model="modalItem.accessoryCategory"
                        :items="accessoryCategories"
                        item-text="name"
                        return-object
                        chips
                        label="アクセサリーカテゴリー"
                        outlined
                    />
                    <v-select
                        v-model="modalItem.materialCategories"
                        :items="materialCategories"
                        item-text="name"
                        return-object
                        chips
                        multiple
                        label="材料カテゴリー"
                        outlined
                    />
                    <v-select
                        v-model="modalItem.salesSites"
                        :items="salesSites"
                        item-text="name"
                        return-object
                        chips
                        multiple
                        label="販売サイト"
                        outlined
                    />
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
import { ExecutionType, ICategory, IconType, IError, ImageType, IProduct, ISite, newProduct, TExecutionType, TImageType } from '~/types'
import { ConfirmState } from '~/store'
import { min20, required } from '~/methods'
@Component({})
export default class CProductList extends Vue {
    @PropSync('items') listItems!: Array<IProduct>
    @Prop({ type: Array, default: [] }) accessoryCategories!: Array<ICategory>
    @Prop({ type: Array, default: [] }) materialCategories!: Array<ICategory>
    @Prop({ type: Array, default: [] }) salesSites!: Array<ISite>
    @Prop({ type: String, default: '' }) type!: string

    IconType: typeof IconType = IconType
    ExecutionType: typeof ExecutionType = ExecutionType
    executionType: TExecutionType = ExecutionType.Create

    products: Array<IProduct> = []
    modalItem: IProduct = newProduct()

    // アップロードする商品画像
    uploadFiles: Array<File> = []
    // ダイアログの表示
    dialogVisible: boolean = false
    // 通知の表示
    notificationVisible: boolean = false
    // 確認ダイアログの表示
    confirmVisible: boolean = true

    valid: boolean = true

    errors: Array<IError> = []

    nameRules = [required, min20]

    // 既存登録リスト
    get registeredList(): Array<string> {
        return this.modalItem.productImages.map((i) => i.apiPath)
    }

    // プレビューリスト
    get previewList(): Array<string> {
        const previewList = []
        for (const file of this.uploadFiles) {
            const url = URL.createObjectURL(file)
            previewList.push(url)
        }
        return previewList
    }

    get modalTitle(): string {
        let title = ''
        for (const type in ExecutionType) {
            if (this.executionType === ExecutionType[type]) {
                title = `商品の${ExecutionType[type]}`
            }
        }
        return title
    }

    get notificationMessage(): string {
        let message = ''
        for (const type in ExecutionType) {
            if (this.executionType === ExecutionType[type]) {
                message = `商品を${ExecutionType[type]}しました`
            }
        }
        return message
    }

    setInit() {
        this.modalItem = newProduct()
        this.uploadFiles = []
    }

    setItem(item: IProduct) {
        this.modalItem = _.cloneDeep(item)
        this.uploadFiles = []
    }

    @Watch('dialogVisible')
    resetValidation() {
        if (!this.dialogVisible && this.executionType !== ExecutionType.Delete) {
            const refs: any = this.$refs.form
            refs.resetValidation()
        }
    }

    openHandler(executionType: TExecutionType, item: IProduct | null = null) {
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
        if (this.executionType === ExecutionType.Create) {
            try {
                const createProduct = await this.$axios.$post(`/product`, this.modalItem)
                // 画像を選択していたら、アップロードを行う
                if (this.uploadFiles.length > 0) {
                    const params = new FormData()
                    this.uploadFiles.forEach((file, index) => {
                        params.append(`file${index}`, file)
                    })
                    await this.$axios.$post(`/product/${createProduct.uuid}/product_image`, params, {
                        headers: {
                            'Content-Type': 'multipart/form-data',
                        },
                    })
                }
                this.$emit('c-change')
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e.response)
            }
        } else if (this.executionType === ExecutionType.Edit) {
            // try {
            //     await this.$axios.$put(`/product/${this.modalItem.uuid}`, this.modalItem)
            //     this.$emit('c-change')
            //     this.notificationVisible = true
            //     this.dialogVisible = false
            // } catch (e) {
            //     this.errors.push(e)
            // }
        } else if (this.executionType === ExecutionType.Delete) {
            try {
                await this.$axios.$delete(`/product/${this.modalItem.uuid}`)
                this.$emit('c-change')
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e)
            }
        }
    }

    async deleteImageHandler(index: number, imageType: TImageType) {
        if (imageType === ImageType.Registered) {
            await this.$store.dispatch('confirm', {
                confirmMessage: '登録画像を削除してもよろしいですか？',
                confirmAction: () => {
                    this.$axios.$delete(`product/${this.modalItem.uuid}/product_image/${this.modalItem.productImages[index].uuid}`)
                },
                cancelAction: () => {},
            } as ConfirmState)
        } else {
            this.uploadFiles.splice(index, 1)
        }
    }
}
</script>

<style lang="stylus">
.v-image
    aspect-ratio 16 / 9
</style>
