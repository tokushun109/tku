<template>
    <v-container class="c-product-list">
        <v-sheet class="product-list-area">
            <v-container>
                <div class="product-list-header">
                    <h3 class="title product-list-title">商品</h3>
                    <v-spacer />
                    <c-icon :type="IconType.New.name" @c-click="clickHandler(ExecutionType.Create)" />
                </div>
                <v-divider />
                <v-list class="product-list-content">
                    <c-message v-if="listItems.length === 0" class="message"> 登録されていません </c-message>
                    <v-row>
                        <v-col v-for="listItem in listItems" :key="listItem.uuid" cols="12" sm="6" md="4">
                            <c-product-card :list-item="listItem" admin @c-click="clickHandler" />
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
                    <v-row class="price-input">
                        <v-col cols="6">
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
                        <v-col class="price-unit" cols="6">
                            <p class="price-unit-content">{{ modalItem.price | priceFormat }}円</p>
                        </v-col>
                    </v-row>
                    <v-file-input v-model="uploadFiles" label="商品画像" prepend-icon="mdi-camera" multiple outlined />
                    <c-image-list
                        title="現在の登録"
                        :registered-list="registeredList"
                        :preview-list="previewList"
                        @c-delete-image-handler="deleteImageHandler"
                    />
                    <v-select
                        v-model="modalItem.category"
                        :items="categories"
                        item-text="name"
                        return-object
                        height="56"
                        chips
                        label="カテゴリー"
                        outlined
                    />
                    <v-select
                        v-model="modalItem.tags"
                        :items="tags"
                        item-text="name"
                        return-object
                        chips
                        height="56"
                        multiple
                        label="タグ"
                        outlined
                    />
                    <v-row dense>
                        <v-col cols="12" sm="6">
                            <v-select
                                v-model="previewSiteDetail.salesSite"
                                :items="salesSites"
                                height="54"
                                item-text="name"
                                return-object
                                chips
                                label="販売サイト"
                                outlined
                                hint="選択後にURLを入力してください"
                            />
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-text-field
                                v-model="previewSiteDetail.detailUrl"
                                :append-icon="IconType.Plus.icon"
                                label="URL"
                                :disabled="!previewSiteDetail.salesSite.name"
                                hint="Enterで販売サイトを追加"
                                @keydown.enter="AddSiteDetail"
                            />
                        </v-col>
                    </v-row>
                    <div v-if="modalItem.siteDetails.length > 0" class="site-detail-preview">
                        <v-chip
                            v-for="(siteDetail, index) in modalItem.siteDetails"
                            :key="index"
                            close
                            :color="ColorType.Grey"
                            :text-color="ColorType.White"
                            :href="siteDetail.detailUrl"
                            target="_blank"
                            rel="noopener noreferrer"
                            @click:close="deleteSiteDetail(index)"
                        >
                            {{ siteDetail.salesSite.name }}
                        </v-chip>
                    </div>
                    <div class="active-check-box">
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
    IClassification,
    IError,
    IProduct,
    ISite,
    ISiteDetail,
    ColorType,
    ExecutionType,
    IconType,
    ImageType,
    TExecutionType,
    TImageType,
} from '~/types'
import { maxPrice, min20, newProduct, newSiteDetail, price, required } from '~/methods'

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

    previewSiteDetail: ISiteDetail = newSiteDetail()

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

    @Watch('dialogVisible')
    resetValidation() {
        if (!this.dialogVisible && this.executionType !== ExecutionType.Delete) {
            const refs: any = this.$refs.form
            refs.resetValidation()
        }
    }

    AddSiteDetail() {
        if (this.previewSiteDetail.salesSite && this.previewSiteDetail.detailUrl) {
            this.modalItem.siteDetails.push(this.previewSiteDetail)
            this.previewSiteDetail = newSiteDetail()
        }
    }

    setInit() {
        this.modalItem = newProduct()
        this.uploadFiles = []
    }

    setItem(item: IProduct) {
        this.modalItem = _.cloneDeep(item)
        this.uploadFiles = []
    }

    clickHandler(executionType: TExecutionType, item: IProduct | null = null) {
        this.errors = []
        if (executionType === ExecutionType.Create) {
            this.setInit()
            this.executionType = executionType
            this.dialogVisible = true
        } else if (item && executionType === ExecutionType.Detail) {
            this.$router.push(`/admin/product/${item.uuid}`)
        } else {
            this.setItem(item!)
            this.executionType = executionType
            this.dialogVisible = true
        }
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

    deleteSiteDetail(index: number) {
        this.modalItem.siteDetails.splice(index, 1)
    }
}
</script>

<style lang="stylus" scoped>
.c-product-list
    .product-list-area
        padding 16px
        .product-list-header
            display flex
            .product-list-title
                color $title-text-color
        .product-list-content
            .message
                margin-top 16px

.price-input
    .price-unit
        text-align right
        .price-unit-content
            padding-top 12px

.site-detail-preview
    border 1px dashed $light-dark-color
    border-radius 3px
    text-align left
    .site-detail-chip
        margin 4px

.active-check-box
    display flex
    justify-content center
</style>
