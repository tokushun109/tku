'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { Close } from '@mui/icons-material'
import { useEffect, useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'

import { Button } from '@/components/bases/Button'
import { Checkbox } from '@/components/bases/Checkbox'
import { Chip, ChipSize } from '@/components/bases/Chip'
import { Dialog } from '@/components/bases/Dialog'
import { Form } from '@/components/bases/Form'
import { ImagePreviewList, ImageItem } from '@/components/bases/ImagePreviewList'
import { Input } from '@/components/bases/Input'
import { Message, MessageType } from '@/components/bases/Message'
import { MultiSelectForm, MultiSelectFormOption } from '@/components/bases/MultiSelectForm'
import { MultipleImageInput } from '@/components/bases/MultipleImageInput'
import { SelectForm, SelectFormOption } from '@/components/bases/SelectForm'
import { TextArea } from '@/components/bases/TextArea'
import { IClassification } from '@/features/classification/type'
import { ISite } from '@/features/site/type'
import { ColorType, FontSizeType } from '@/types'

import styles from './styles.module.scss'
import { CreemaDuplicateSchema, ProductSchema } from '../../product/schema'

import type { ICreemaDuplicateForm } from '../../product/type'
import type { IProduct, IProductForm } from '../../type'

const CreateProductType = {
    Input: 'input',
    Duplicate: 'duplicate',
} as const
type CreateProductType = (typeof CreateProductType)[keyof typeof CreateProductType]

interface Props {
    categories: IClassification[]
    isOpen: boolean
    isSubmitting: boolean
    onClose: () => void
    onCreemaDuplicate: (_data: ICreemaDuplicateForm) => Promise<void>
    onSubmit: (_data: IProductForm) => Promise<void>
    salesSites: ISite[]
    submitError: string | null
    tags: IClassification[]
    targets: IClassification[]
    updateItem: IProduct | null
}

export const ProductFormDialog = ({
    categories,
    isOpen,
    isSubmitting,
    onClose,
    onSubmit,
    onCreemaDuplicate,
    salesSites,
    submitError,
    tags,
    targets,
    updateItem,
}: Props) => {
    const [selectedSalesSite, setSelectedSalesSite] = useState<string>('')
    const [siteDetailUrl, setSiteDetailUrl] = useState<string>('')
    const [siteDetailUrlError, setSiteDetailUrlError] = useState<string>('')
    const [selectedTags, setSelectedTags] = useState<string[]>([])
    const [formSiteDetails, setFormSiteDetails] = useState<Array<{ detailUrl: string; salesSiteName: string; salesSiteUuid: string }>>([])
    const [uploadImages, setUploadImages] = useState<File[]>([])
    const [imageItems, setImageItems] = useState<ImageItem[]>([])
    const [isImageOrderChanged, setIsImageOrderChanged] = useState<boolean>(false)
    const [createProductType, setCreateProductType] = useState<CreateProductType>(CreateProductType.Duplicate)

    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
        watch,
        setValue,
    } = useForm<IProductForm>({
        resolver: zodResolver(ProductSchema),
        defaultValues: {
            name: '',
            description: '',
            price: 0,
            isActive: true,
            isRecommend: false,
            categoryUuid: '',
            targetUuid: '',
            tagUuids: [],
            siteDetails: [],
            uploadImages: [],
        },
    })

    // 複製フォーム用のuseForm
    const {
        register: registerCreema,
        handleSubmit: handleSubmitCreema,
        formState: { errors: errorsCreema },
        reset: resetCreema,
    } = useForm<ICreemaDuplicateForm>({
        resolver: zodResolver(CreemaDuplicateSchema),
        defaultValues: {
            creemaUrl: '',
        },
    })

    const watchedPrice = watch('price')

    // SelectForm用のオプション配列を作成
    const categoryOptions: SelectFormOption[] = categories.map((category) => ({
        label: category.name,
        value: category.uuid,
    }))

    const targetOptions: SelectFormOption[] = targets.map((target) => ({
        label: target.name,
        value: target.uuid,
    }))

    const tagOptions: MultiSelectFormOption[] = tags.map((tag) => ({
        label: tag.name,
        value: tag.uuid,
    }))

    const salesSiteOptions: SelectFormOption[] = salesSites
        .filter((site) => site.uuid)
        .map((site) => ({
            label: site.name,
            value: site.uuid as string,
        }))

    // updateItemが変更されたときにフォームをリセット
    useEffect(() => {
        if (isOpen) {
            const tagUuids = updateItem?.tags?.map((tag) => tag.uuid) || []
            const siteDetails =
                updateItem?.siteDetails?.map((detail) => ({
                    salesSiteUuid: detail.salesSite.uuid || '',
                    detailUrl: detail.detailUrl,
                })) || []

            reset({
                name: updateItem?.name || '',
                description: updateItem?.description || '',
                price: updateItem?.price || 0,
                isActive: updateItem?.isActive ?? true,
                isRecommend: updateItem?.isRecommend ?? false,
                categoryUuid: updateItem?.category?.uuid || '',
                targetUuid: updateItem?.target?.uuid || '',
                tagUuids,
                siteDetails,
                uploadImages: [],
            })

            // ローカル状態も更新
            setSelectedTags(tagUuids)
            setFormSiteDetails(
                updateItem?.siteDetails?.map((detail) => ({
                    salesSiteUuid: detail.salesSite.uuid || '',
                    detailUrl: detail.detailUrl,
                    salesSiteName: detail.salesSite.name,
                })) || [],
            )

            // 既存画像をImageItemsに変換
            const existingImages: ImageItem[] =
                updateItem?.productImages?.map((image, _index) => ({
                    id: `existing-${image.uuid}`,
                    src: image.apiPath,
                    isNewUpload: false,
                    order: image.order,
                })) || []
            setImageItems(existingImages)
        }
    }, [updateItem, isOpen, reset])

    const handleClose = () => {
        reset()
        setSelectedTags([])
        setFormSiteDetails([])
        setSelectedSalesSite('')
        setSiteDetailUrl('')
        setSiteDetailUrlError('')
        setUploadImages([])
        setImageItems([])
        setIsImageOrderChanged(false)
        setCreateProductType(CreateProductType.Duplicate)
        resetCreema()
        onClose()
    }

    const validateSiteDetailUrl = (url: string): string => {
        if (!url) {
            return 'URLを入力してください'
        }

        try {
            new URL(url)
        } catch {
            return '正しいURLを入力してください'
        }

        // 同じサイトで既に登録済みかチェック
        const isDuplicate = formSiteDetails.some((detail) => detail.salesSiteUuid === selectedSalesSite && detail.detailUrl === url)
        if (isDuplicate) {
            return '同じサイトで既に登録済みのURLです'
        }

        return ''
    }

    const handleSiteDetailUrlChange = (value: string) => {
        setSiteDetailUrl(value)
        if (value && selectedSalesSite) {
            const error = validateSiteDetailUrl(value)
            setSiteDetailUrlError(error)
        } else {
            setSiteDetailUrlError('')
        }
    }

    const handleAddSiteDetail = () => {
        if (selectedSalesSite && siteDetailUrl) {
            const error = validateSiteDetailUrl(siteDetailUrl)
            if (error) {
                setSiteDetailUrlError(error)
                return
            }

            const salesSite = salesSites.find((site) => site.uuid === selectedSalesSite)
            if (salesSite) {
                const newSiteDetail = {
                    salesSiteUuid: selectedSalesSite,
                    detailUrl: siteDetailUrl,
                    salesSiteName: salesSite.name,
                }
                setFormSiteDetails([...formSiteDetails, newSiteDetail])
                setSelectedSalesSite('')
                setSiteDetailUrl('')
                setSiteDetailUrlError('')
            }
        }
    }

    const handleRemoveSiteDetail = (index: number) => {
        setFormSiteDetails(formSiteDetails.filter((_, i) => i !== index))
    }

    const handleImageUpload = (files: File[]) => {
        setUploadImages(files)
        setValue('uploadImages', files)

        // 新規画像をImageItemsに追加
        const newImageItems: ImageItem[] = files.map((file, index) => ({
            id: `new-${Date.now()}-${index}`,
            src: URL.createObjectURL(file),
            isNewUpload: true,
        }))

        // 既存画像と新規画像を統合
        const existingImages = imageItems.filter((item) => !item.isNewUpload)
        setImageItems([...existingImages, ...newImageItems])
    }

    const handleImageDelete = (id: string) => {
        const updatedItems = imageItems.filter((item) => item.id !== id)
        setImageItems(updatedItems)

        // 新規画像の場合はuploadImagesからも削除
        if (id.startsWith('new-')) {
            const imageIndex = imageItems.findIndex((item) => item.id === id)
            if (imageIndex !== -1) {
                const newUploadImages = uploadImages.filter((_, _index) => {
                    const newImageStartIndex = imageItems.filter((item) => !item.isNewUpload).length
                    return _index !== imageIndex - newImageStartIndex
                })
                setUploadImages(newUploadImages)
                setValue('uploadImages', newUploadImages)
            }
        }
    }

    const handleImageReorder = (dragId: string, hoverId: string) => {
        const dragIndex = imageItems.findIndex((item) => item.id === dragId)
        const hoverIndex = imageItems.findIndex((item) => item.id === hoverId)

        if (dragIndex === -1 || hoverIndex === -1) return

        const newItems = [...imageItems]
        const draggedItem = newItems[dragIndex]
        newItems.splice(dragIndex, 1)
        newItems.splice(hoverIndex, 0, draggedItem)

        // 順序を更新
        const updatedItems = newItems.map((item, index) => ({
            ...item,
            order: index + 1,
        }))

        setImageItems(updatedItems)
        setIsImageOrderChanged(true) // 並び替えが発生したフラグを立てる
    }

    const handleFormSubmit: SubmitHandler<IProductForm> = async (data) => {
        const formData = {
            ...data,
            tagUuids: selectedTags,
            siteDetails: formSiteDetails.map((detail) => ({
                salesSiteUuid: detail.salesSiteUuid,
                detailUrl: detail.detailUrl,
            })),
            uploadImages,
            imageItems, // 並び替え後の画像リスト
            isImageOrderChanged, // 並び替えフラグ
        }
        await onSubmit(formData)
        reset()
        setSelectedTags([])
        setFormSiteDetails([])
        setUploadImages([])
        setImageItems([])
        setIsImageOrderChanged(false)
    }

    const handleCreemaDuplicateSubmit: SubmitHandler<ICreemaDuplicateForm> = async (data) => {
        await onCreemaDuplicate(data)
        resetCreema()
    }

    const isEdit = updateItem !== null

    return (
        <Dialog
            confirmOption={{
                label: isSubmitting ? '送信中...' : isEdit ? '更新' : createProductType === CreateProductType.Duplicate ? '複製' : '追加',
                onClick:
                    createProductType === CreateProductType.Duplicate && !isEdit
                        ? handleSubmitCreema(handleCreemaDuplicateSubmit)
                        : handleSubmit(handleFormSubmit),
                disabled: isSubmitting,
            }}
            isOpen={isOpen}
            onClose={handleClose}
            title={`商品を${isEdit ? '編集' : '追加'}`}
            wide
        >
            {submitError && <Message type={MessageType.Error}>{submitError}</Message>}

            {/* 新規作成時のみ表示される選択フォーム */}
            {!isEdit && (
                <div className={styles['form-row']}>
                    <div className={styles['checkbox-group']}>
                        <Checkbox
                            checked={createProductType === CreateProductType.Duplicate}
                            label="Creemaから複製"
                            onChange={(checked) => setCreateProductType(checked ? CreateProductType.Duplicate : CreateProductType.Input)}
                        />
                        <Checkbox
                            checked={createProductType === CreateProductType.Input}
                            label="手動で入力"
                            onChange={(checked) => setCreateProductType(checked ? CreateProductType.Input : CreateProductType.Duplicate)}
                        />
                    </div>
                </div>
            )}

            {/* 複製フォーム */}
            {createProductType === CreateProductType.Duplicate && !isEdit ? (
                <Form noValidate onSubmit={handleSubmitCreema(handleCreemaDuplicateSubmit)}>
                    <div className={styles['form-row']}>
                        <Input
                            {...registerCreema('creemaUrl')}
                            error={errorsCreema.creemaUrl?.message}
                            id="creemaUrl"
                            label="Creema URL"
                            placeholder="CreemaのURLを入力してください"
                            required
                            type="url"
                        />
                    </div>
                </Form>
            ) : (
                /* 通常の商品フォーム */
                <Form noValidate onSubmit={handleSubmit(handleFormSubmit)}>
                    <div className={styles['form-row']}>
                        <Input
                            {...register('name')}
                            error={errors.name?.message}
                            id="name"
                            label="商品名"
                            maxLength={255}
                            placeholder="商品名を入力してください"
                            required
                            type="text"
                        />
                    </div>

                    <div className={styles['form-row']}>
                        <TextArea
                            {...register('description')}
                            error={errors.description?.message}
                            id="description"
                            label="商品説明"
                            placeholder="商品説明を入力してください（任意）"
                            rows={8}
                        />
                    </div>

                    <div className={styles['form-row']}>
                        <Input
                            {...register('price', { valueAsNumber: true })}
                            error={errors.price?.message}
                            id="price"
                            label="税込価格"
                            max={1000000}
                            min={1}
                            placeholder="0"
                            required
                            type="number"
                        />
                        <p className={styles['price-value']}>¥{watchedPrice ? watchedPrice.toLocaleString() : '0'}</p>
                    </div>

                    <div className={styles['form-row']}>
                        <SelectForm
                            error={errors.categoryUuid?.message}
                            id="categoryUuid"
                            label="カテゴリー"
                            onChange={(value) => setValue('categoryUuid', value)}
                            options={categoryOptions}
                            placeholder="カテゴリーを選択してください"
                            value={watch('categoryUuid')}
                        />
                    </div>

                    <div className={styles['form-row']}>
                        <SelectForm
                            error={errors.targetUuid?.message}
                            id="targetUuid"
                            label="ターゲット"
                            onChange={(value) => setValue('targetUuid', value)}
                            options={targetOptions}
                            placeholder="ターゲットを選択してください"
                            value={watch('targetUuid')}
                        />
                    </div>

                    <div className={styles['form-row']}>
                        <MultiSelectForm
                            id="tagUuids"
                            label="タグ"
                            onChange={setSelectedTags}
                            options={tagOptions}
                            placeholder="タグを選択してください"
                            value={selectedTags}
                        />
                    </div>

                    <div className={styles['form-row']}>
                        <MultipleImageInput
                            helperText="画像ファイルを選択してください（複数選択可能）"
                            id="uploadImages"
                            label="商品画像"
                            onChange={handleImageUpload}
                            value={uploadImages}
                        />
                        <ImagePreviewList images={imageItems} onDelete={handleImageDelete} onReorder={handleImageReorder} title="現在の登録" />
                    </div>

                    <div className={styles['form-row']}>
                        <div className={styles['form-field']}>
                            <label className={styles['label']}>販売ページ</label>
                            <div className={styles['site-detail-input']}>
                                <SelectForm
                                    id="salesSite"
                                    onChange={(value) => {
                                        setSelectedSalesSite(value)
                                        setSiteDetailUrlError('')
                                    }}
                                    options={salesSiteOptions}
                                    placeholder="販売サイトを選択してください"
                                    value={selectedSalesSite}
                                />
                            </div>
                            {selectedSalesSite && (
                                <div className={styles['url-input-row']}>
                                    <div className={styles['url-input-wrapper']}>
                                        <input
                                            className={`${styles['url-input']} ${siteDetailUrlError ? styles['error'] : ''}`}
                                            onChange={(e) => handleSiteDetailUrlChange(e.target.value)}
                                            onKeyDown={(e) => {
                                                if (e.key === 'Enter') {
                                                    e.preventDefault()
                                                    handleAddSiteDetail()
                                                }
                                            }}
                                            placeholder="URLを入力してください"
                                            type="url"
                                            value={siteDetailUrl}
                                        />
                                        {siteDetailUrlError && <span className={styles['field-error']}>{siteDetailUrlError}</span>}
                                    </div>
                                    <Button
                                        colorType={ColorType.Primary}
                                        disabled={!siteDetailUrl || !!siteDetailUrlError}
                                        onClick={handleAddSiteDetail}
                                        type="button"
                                    >
                                        URLを追加
                                    </Button>
                                </div>
                            )}
                            {formSiteDetails.length > 0 && (
                                <>
                                    <label className={styles['registered-sites-label']}>現在の登録</label>
                                    <div className={styles['site-detail-list']}>
                                        {formSiteDetails.map((detail, index) => (
                                            <div className={styles['selected-chip-container']} key={index}>
                                                <Chip
                                                    color={ColorType.Secondary}
                                                    fontColor="#ffffff"
                                                    fontSize={FontSizeType.SmMd}
                                                    size={ChipSize.Small}
                                                >
                                                    <a
                                                        className={styles['chip-link']}
                                                        href={detail.detailUrl}
                                                        rel="noopener noreferrer"
                                                        target="_blank"
                                                    >
                                                        {detail.salesSiteName}
                                                    </a>
                                                    <button
                                                        className={styles['chip-close-button']}
                                                        onClick={(e) => {
                                                            e.preventDefault()
                                                            handleRemoveSiteDetail(index)
                                                        }}
                                                        type="button"
                                                    >
                                                        <Close />
                                                    </button>
                                                </Chip>
                                            </div>
                                        ))}
                                    </div>
                                </>
                            )}
                        </div>
                    </div>

                    <div className={styles['form-row']}>
                        <div className={styles['checkbox-group']}>
                            <Checkbox {...register('isActive')} label="販売中" />
                            <Checkbox {...register('isRecommend')} label="おすすめに設定" />
                        </div>
                    </div>
                </Form>
            )}
        </Dialog>
    )
}
