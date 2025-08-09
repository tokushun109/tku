'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useEffect, useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'

import { Checkbox } from '@/components/bases/Checkbox'
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

import styles from './styles.module.scss'
import { ProductSchema } from '../../product/schema'

import type { IProduct, IProductForm } from '../../type'

interface Props {
    categories: IClassification[]
    isOpen: boolean
    isSubmitting: boolean
    onClose: () => void
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
    salesSites,
    submitError,
    tags,
    targets,
    updateItem,
}: Props) => {
    const [selectedSalesSite, setSelectedSalesSite] = useState<string>('')
    const [siteDetailUrl, setSiteDetailUrl] = useState<string>('')
    const [selectedTags, setSelectedTags] = useState<string[]>([])
    const [formSiteDetails, setFormSiteDetails] = useState<Array<{ detailUrl: string; salesSiteName: string; salesSiteUuid: string }>>([])
    const [uploadImages, setUploadImages] = useState<File[]>([])
    const [imageItems, setImageItems] = useState<ImageItem[]>([])
    const [isImageOrderChanged, setIsImageOrderChanged] = useState<boolean>(false)

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
        setUploadImages([])
        setImageItems([])
        setIsImageOrderChanged(false)
        onClose()
    }

    const handleAddSiteDetail = () => {
        if (selectedSalesSite && siteDetailUrl) {
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

    const isEdit = updateItem !== null

    return (
        <Dialog
            confirmOption={{
                label: isSubmitting ? '送信中...' : isEdit ? '更新' : '追加',
                onClick: handleSubmit(handleFormSubmit),
                disabled: isSubmitting,
            }}
            isOpen={isOpen}
            onClose={handleClose}
            title={`商品を${isEdit ? '編集' : '追加'}`}
            wide
        >
            {submitError && <Message type={MessageType.Error}>{submitError}</Message>}
            <Form noValidate onSubmit={handleSubmit(handleFormSubmit)}>
                <div className={styles['form-row']}>
                    <Input
                        {...register('name')}
                        error={errors.name?.message}
                        id="name"
                        label="商品名"
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
                        <label className={styles['label']}>販売サイト</label>
                        <div className={styles['site-detail-input']}>
                            <SelectForm
                                id="salesSite"
                                onChange={(value) => setSelectedSalesSite(value)}
                                options={salesSiteOptions}
                                placeholder="販売サイトを選択してください"
                                value={selectedSalesSite}
                            />
                            <input
                                className={styles['url-input']}
                                disabled={!selectedSalesSite}
                                onChange={(e) => setSiteDetailUrl(e.target.value)}
                                onKeyDown={(e) => {
                                    if (e.key === 'Enter') {
                                        e.preventDefault()
                                        handleAddSiteDetail()
                                    }
                                }}
                                placeholder="URLを入力してEnterで追加"
                                type="url"
                                value={siteDetailUrl}
                            />
                        </div>
                        {formSiteDetails.length > 0 && (
                            <div className={styles['site-detail-list']}>
                                {formSiteDetails.map((detail, index) => (
                                    <div className={styles['site-detail-item']} key={index}>
                                        <a className={styles['site-detail-link']} href={detail.detailUrl} rel="noopener noreferrer" target="_blank">
                                            {detail.salesSiteName}
                                        </a>
                                        <button className={styles['remove-button']} onClick={() => handleRemoveSiteDetail(index)} type="button">
                                            ×
                                        </button>
                                    </div>
                                ))}
                            </div>
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
        </Dialog>
    )
}
