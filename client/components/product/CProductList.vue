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
                                <c-product-card :list-item="listItem" admin @c-open="openHandler" />
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
                    <v-row>
                        <v-col cols="7">
                            <v-text-field
                                v-model.number="modalItem.price"
                                :rules="priceRules"
                                label="税込価格(必須)"
                                outlined
                                :min="1"
                                :max="1000000"
                                type="number"
                            />
                        </v-col>
                        <v-col class="text-right" cols="5">
                            <p class="pt-3">{{ modalItem.price | priceFormat }}円</p>
                        </v-col>
                    </v-row>
                    <v-file-input v-model="uploadFiles" label="商品画像" prepend-icon="mdi-camera" multiple outlined />
                    <c-image-list
                        title="現在の登録"
                        :registered-list="registeredList"
                        :preview-list="previewList"
                        @c-delete-image-handler="deleteImageHandler"
                    />
                    <v-select v-model="modalItem.category" :items="categories" item-text="name" return-object chips label="カテゴリー" outlined />
                    <v-select v-model="modalItem.tags" :items="tags" item-text="name" return-object chips multiple label="タグ" outlined />
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
                    <div class="d-flex justify-center outlined">
                        <v-checkbox v-model="modalItem.isActive" label="販売中" dense />
                    </div>
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
import {
    ColorType,
    ExecutionType,
    IClassification,
    IconType,
    IError,
    ImageType,
    IProduct,
    ISite,
    newProduct,
    TExecutionType,
    TImageType,
} from '~/types'
import { maxPrice, min20, price, required } from '~/methods'
@Component({})
export default class CProductList extends Vue {
    @PropSync('items') listItems!: Array<IProduct>
    @Prop({ type: Array, default: [] }) categories!: Array<IClassification>
    @Prop({ type: Array, default: [] }) tags!: Array<IClassification>
    @Prop({ type: Array, default: [] }) salesSites!: Array<ISite>
    @Prop({ type: String, default: '' }) type!: string

    ColorType: typeof ColorType = ColorType
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

    priceRules = [required, price, maxPrice]

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

    getColor(listItem: IProduct): string {
        return listItem.isActive ? 'light-green lighten-5' : 'grey lighten-3'
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
        this.errors = []
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
            try {
                await this.$axios.$put(`/product/${this.modalItem.uuid}`, this.modalItem)
                // 画像を選択していたら、アップロードを行う
                if (this.uploadFiles.length > 0) {
                    const params = new FormData()
                    this.uploadFiles.forEach((file, index) => {
                        params.append(`file${index}`, file)
                    })
                    await this.$axios.$post(`/product/${this.modalItem.uuid}/product_image`, params, {
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

    deleteImageHandler(index: number, imageType: TImageType) {
        if (imageType === ImageType.Registered) {
            this.modalItem.productImages.splice(index, 1)
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
