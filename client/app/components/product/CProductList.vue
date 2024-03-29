<template>
    <v-container class="c-product-list">
        <v-sheet class="product-list-area">
            <v-container>
                <div class="product-list-header">
                    <h3 class="title product-list-title">商品</h3>
                    <v-spacer />
                    <c-icon :type="IconType.New.name" @c-click="clickHandler(ExecutionType.Create)" />
                </div>
                <v-divider class="divider" />
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
                <div class="radio-button-area">
                    <v-radio-group v-if="executionType === ExecutionType.Create" v-model="createProductType" row>
                        <v-radio :off-icon="mdiCheckboxBlankOutline" :on-icon="mdiCheckboxMarked" label="手動で入力" value="input" />
                        <v-radio :off-icon="mdiCheckboxBlankOutline" :on-icon="mdiCheckboxMarked" label="Creemaから複製" value="duplicate" />
                    </v-radio-group>
                </div>
                <v-form
                    v-if="(executionType === ExecutionType.Create && createProductType === 'input') || executionType === ExecutionType.Edit"
                    ref="form"
                    v-model="valid"
                    lazy-validation
                >
                    <v-text-field v-model="modalItem.name" :rules="nameRules" label="商品名(必須)" outlined counter="50" />
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
                    <v-file-input v-model="uploadFiles" label="商品画像" :prepend-icon="mdiCamera" multiple outlined @change="orderInit" />
                    <c-image-list
                        title="現在の登録"
                        :registered-list="registeredList"
                        :preview-list="previewList"
                        :max-order="maxOrder"
                        @c-delete-image-handler="deleteImageHandler"
                        @c-order-image-handler="orderImageHandler"
                    />
                    <v-btn v-if="isChangedOrder" class="order-reset" @click="orderInit">表示順リセット</v-btn>
                    <v-autocomplete
                        v-model="modalItem.category"
                        :items="categories"
                        item-text="name"
                        return-object
                        height="56"
                        chips
                        label="カテゴリー"
                        outlined
                    />
                    <v-autocomplete
                        v-model="modalItem.target"
                        :items="targets"
                        item-text="name"
                        return-object
                        height="56"
                        chips
                        label="ターゲット"
                        outlined
                    />
                    <v-autocomplete
                        v-model="modalItem.tags"
                        :search-input.sync="searchText"
                        :items="tags"
                        item-text="name"
                        return-object
                        chips
                        multiple
                        label="タグ"
                        outlined
                        @change="searchText = ''"
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
                            class="site-detail-chip"
                            close
                            :close-icon="mdiCloseCircle"
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
                    <div class="check-box-area">
                        <v-checkbox
                            v-model="modalItem.isRecommend"
                            class="check-box"
                            :off-icon="mdiCheckboxBlankOutline"
                            :on-icon="mdiCheckboxMarked"
                            label="おすすめに設定"
                            dense
                        />
                        <v-checkbox
                            v-model="modalItem.isActive"
                            class="check-box"
                            :off-icon="mdiCheckboxBlankOutline"
                            :on-icon="mdiCheckboxMarked"
                            label="販売中"
                            dense
                        />
                    </div>
                </v-form>
                <v-form v-else-if="executionType === ExecutionType.Create && createProductType === 'duplicate'">
                    <v-text-field v-model="creemaUrl" :rules="urlRules" label="URL(必須)" outlined counter="50" />
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
import { mdiCamera, mdiCheckboxMarked, mdiCheckboxBlankOutline, mdiCloseCircle } from '@mdi/js'
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
    IImagePathOrder,
    BadRequest,
} from '~/types'
import { maxPrice, min50, newProduct, newSiteDetail, nonDoubleByte, nonSpace, price, required } from '~/methods'

interface IIndexOrder {
    [key: number]: number
}

interface IProductImageParams {
    isChanged: boolean
    order: IIndexOrder
}

@Component({})
export default class CProductList extends Vue {
    @PropSync('items') listItems!: Array<IProduct>
    @Prop({ type: Array, default: [] }) categories!: Array<IClassification>
    @Prop({ type: Array, default: [] }) targets!: Array<IClassification>
    @Prop({ type: Array, default: [] }) tags!: Array<IClassification>
    @Prop({ type: Array, default: [] }) salesSites!: Array<ISite>
    @Prop({ type: String, default: '' }) type!: string

    mdiCamera = mdiCamera
    mdiCheckboxMarked = mdiCheckboxMarked
    mdiCheckboxBlankOutline = mdiCheckboxBlankOutline
    mdiCloseCircle = mdiCloseCircle
    ColorType: typeof ColorType = ColorType
    IconType: typeof IconType = IconType
    ExecutionType: typeof ExecutionType = ExecutionType
    executionType: TExecutionType = ExecutionType.Create

    products: Array<IProduct> = []
    modalItem: IProduct = newProduct()
    creemaUrl: string = ''

    // 商品を作成するときに手動入力か複製するか
    createProductType: 'input' | 'duplicate' = 'input'
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

    nameRules = [required, min50]

    priceRules = [required, price, maxPrice]

    urlRules = [required, nonDoubleByte, nonSpace]

    searchText: string = ''

    // 既存登録リスト
    get registeredList(): Array<IImagePathOrder> {
        const registeredList: Array<IImagePathOrder> = []
        this.modalItem.productImages.forEach((productImage, index) => {
            const imageOrder: number | null = this.registeredFileOrder[index] ? this.registeredFileOrder[index] : null
            registeredList[index] = { path: productImage.apiPath, order: imageOrder, type: ImageType.Registered }
        })
        return registeredList
    }

    // プレビューリスト
    get previewList(): Array<IImagePathOrder> {
        const previewList: Array<IImagePathOrder> = []
        this.uploadFiles.forEach((file, index) => {
            const imageOrder: number | null = this.uploadFileOrder[index] ? this.uploadFileOrder[index] : null
            const url = URL.createObjectURL(file)
            previewList[index] = { path: url, order: imageOrder, type: ImageType.Preview }
        })
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

    get isChangedOrder() {
        return this.order !== this.maxOrder
    }

    @Watch('dialogVisible')
    resetValidation() {
        if (!this.dialogVisible && this.executionType !== ExecutionType.Delete && this.createProductType === 'input') {
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
        const selectedLength = this.maxOrder - this.order
        const totalImageLength = this.modalItem.productImages.length + this.uploadFiles.length
        if (this.isChangedOrder && selectedLength < totalImageLength) {
            this.errors.push(new BadRequest('全ての画像の並び替えが選択されていません'))
            return
        }
        if (this.isChangedOrder) {
            this.modalItem.productImages.forEach((image, index) => {
                image.order = this.registeredFileOrder[index]
            })
        }
        if (this.executionType === ExecutionType.Create) {
            try {
                if (this.createProductType === 'input') {
                    const createProduct = await this.$axios.$post(`/product`, this.modalItem, { withCredentials: true })
                    // 画像を選択していたら、アップロードを行う
                    if (this.uploadFiles.length > 0) {
                        const params = new FormData()
                        const orderParams: IProductImageParams = {
                            isChanged: this.isChangedOrder,
                            order: this.uploadFileOrder,
                        }
                        params.append('order', JSON.stringify(orderParams))
                        this.uploadFiles.forEach((file, index) => {
                            params.append(`file${index}`, file)
                        })
                        await this.$axios.$post(`/product/${createProduct.uuid}/product_image`, params, {
                            headers: {
                                'Content-Type': 'multipart/form-data',
                            },
                            withCredentials: true,
                        })
                    }
                } else if (this.createProductType === 'duplicate') {
                    if (!this.creemaUrl.includes('creema')) {
                        this.errors.push(new BadRequest('urlにcreemaが含まれていません'))
                        return
                    }
                    await this.$axios.$post(
                        `/product/duplicate`,
                        { url: this.creemaUrl },
                        {
                            headers: {
                                'Content-Type': 'multipart/form-data',
                            },
                            withCredentials: true,
                        }
                    )
                }
                this.$emit('c-change')
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e.response)
            }
        } else if (this.executionType === ExecutionType.Edit) {
            try {
                await this.$axios.$put(`/product/${this.modalItem.uuid}`, this.modalItem, { withCredentials: true })
                // 画像を選択していたら、アップロードを行う
                if (this.uploadFiles.length > 0) {
                    const params = new FormData()
                    const orderParams: IProductImageParams = {
                        isChanged: this.isChangedOrder,
                        order: this.uploadFileOrder,
                    }
                    params.append('order', JSON.stringify(orderParams))
                    this.uploadFiles.forEach((file, index) => {
                        params.append(`file${index}`, file)
                    })
                    await this.$axios.$post(`/product/${this.modalItem.uuid}/product_image`, params, {
                        headers: {
                            'Content-Type': 'multipart/form-data',
                        },
                        withCredentials: true,
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
                await this.$axios.$delete(`/product/${this.modalItem.uuid}`, { withCredentials: true })
                this.$emit('c-change')
                this.notificationVisible = true
                this.dialogVisible = false
            } catch (e) {
                this.errors.push(e.response)
            }
        }
        this.orderInit()
    }

    deleteImageHandler(index: number, imageType: TImageType) {
        if (imageType === ImageType.Registered) {
            this.modalItem.productImages.splice(index, 1)
            this.$delete(this.registeredFileOrder, index)
        } else {
            this.uploadFiles.splice(index, 1)
            this.$delete(this.uploadFileOrder, index)
        }
        this.orderInit()
    }

    order: number = 100
    maxOrder: 100 = 100
    registeredFileOrder: IIndexOrder = {}
    uploadFileOrder: IIndexOrder = {}
    orderImageHandler(index: number, imageType: TImageType) {
        if (imageType === ImageType.Registered) {
            if (this.registeredFileOrder[index]) {
                this.addOrder(this.registeredFileOrder[index])
                this.$delete(this.registeredFileOrder, index)
                this.order += 1
            } else {
                this.$set(this.registeredFileOrder, index, this.order)
                this.order -= 1
            }
        } else if (imageType === ImageType.Preview) {
            if (this.uploadFileOrder[index]) {
                this.addOrder(this.uploadFileOrder[index])
                this.$delete(this.uploadFileOrder, index)
                this.order += 1
            } else {
                this.$set(this.uploadFileOrder, index, this.order)
                this.order -= 1
            }
        }
    }

    addOrder(baseOrder: number) {
        for (const index in this.registeredFileOrder) {
            if (this.registeredFileOrder[index] <= baseOrder) {
                this.registeredFileOrder[index] += 1
            }
        }
        for (const index in this.uploadFileOrder) {
            if (this.uploadFileOrder[index] <= baseOrder) {
                this.uploadFileOrder[index] += 1
            }
        }
    }

    orderInit() {
        this.order = 100
        this.registeredFileOrder = {}
        this.uploadFileOrder = {}
    }

    deleteSiteDetail(index: number) {
        this.modalItem.siteDetails.splice(index, 1)
    }
}
</script>

<style lang="stylus" scoped>
::v-deep .v-list-item__action
    display none

.c-product-list
    .product-list-area
        padding 16px
        .product-list-header
            display flex
            .product-list-title
                color $title-primary-color
        .product-list-content
            .message
                margin-top 16px

.radio-button-area
    display flex
    justify-content center

.price-input
    .price-unit
        text-align right
        .price-unit-content
            padding-top 12px

.order-reset
    margin-bottom 16px

.site-detail-preview
    border 1px dashed $light-dark-color
    border-radius 3px
    text-align left
    .site-detail-chip
        margin 8px 0 8px 8px

.check-box-area
    display flex
    justify-content center
    .check-box
        margin 20px
</style>
