'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useEffect, useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'

import { Checkbox } from '@/components/bases/Checkbox'
import { Dialog } from '@/components/bases/Dialog'
import { Form } from '@/components/bases/Form'
import { Input } from '@/components/bases/Input'
import { Message, MessageType } from '@/components/bases/Message'
import { MultiSelectForm, MultiSelectFormOption } from '@/components/bases/MultiSelectForm'
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
        }
    }, [updateItem, isOpen, reset])

    const handleClose = () => {
        reset()
        setSelectedTags([])
        setFormSiteDetails([])
        setSelectedSalesSite('')
        setSiteDetailUrl('')
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

    const handleFormSubmit: SubmitHandler<IProductForm> = async (_data) => {
        const formData = {
            ..._data,
            tagUuids: selectedTags,
            siteDetails: formSiteDetails.map((detail) => ({
                salesSiteUuid: detail.salesSiteUuid,
                detailUrl: detail.detailUrl,
            })),
        }
        await onSubmit(formData)
        reset()
        setSelectedTags([])
        setFormSiteDetails([])
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
                        rows={4}
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
