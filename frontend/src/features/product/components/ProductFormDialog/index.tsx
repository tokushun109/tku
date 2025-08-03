'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useEffect, useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'

import { getCategories } from '@/apis/category'
import { getSalesSiteList } from '@/apis/salesSite'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'
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
    isOpen: boolean
    isSubmitting: boolean
    onClose: () => void
    onSubmit: (_data: IProductForm) => Promise<void>
    submitError: string | null
    updateItem: IProduct | null
}

export const ProductFormDialog = ({ isOpen, isSubmitting, onClose, onSubmit, submitError, updateItem }: Props) => {
    const [categories, setCategories] = useState<IClassification[]>([])
    const [targets, setTargets] = useState<IClassification[]>([])
    const [tags, setTags] = useState<IClassification[]>([])
    const [salesSites, setSalesSites] = useState<ISite[]>([])
    const [isLoadingData, setIsLoadingData] = useState<boolean>(false)
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

    // データの取得
    useEffect(() => {
        const fetchData = async () => {
            try {
                setIsLoadingData(true)
                const [categoriesData, targetsData, tagsData, salesSitesData] = await Promise.all([
                    getCategories({ mode: 'all' }),
                    getTargets({ mode: 'all' }),
                    getTags(),
                    getSalesSiteList(),
                ])
                setCategories(categoriesData)
                setTargets(targetsData)
                setTags(tagsData)
                setSalesSites(salesSitesData)
            } catch (error) {
                console.error('データの取得に失敗しました:', error)
            } finally {
                setIsLoadingData(false)
            }
        }

        if (isOpen) {
            fetchData()
        }
    }, [isOpen])

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
                disabled: isSubmitting || isLoadingData,
            }}
            isOpen={isOpen}
            onClose={handleClose}
            title={`商品を${isEdit ? '編集' : '追加'}`}
            wide
        >
            {submitError && <Message type={MessageType.Error}>{submitError}</Message>}
            {isLoadingData ? (
                <div className={styles['loading']}>読み込み中...</div>
            ) : (
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
                        <div className={styles['price-row']}>
                            <div className={styles['price-input']}>
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
                            </div>
                            <div className={styles['price-display']}>
                                <span className={styles['price-text']}>¥{watchedPrice ? watchedPrice.toLocaleString() : '0'}円</span>
                            </div>
                        </div>
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
                                <select className={styles['select']} onChange={(e) => setSelectedSalesSite(e.target.value)} value={selectedSalesSite}>
                                    <option value="">選択してください</option>
                                    {salesSites.map((site) => (
                                        <option key={site.uuid} value={site.uuid}>
                                            {site.name}
                                        </option>
                                    ))}
                                </select>
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
                                            <a
                                                className={styles['site-detail-link']}
                                                href={detail.detailUrl}
                                                rel="noopener noreferrer"
                                                target="_blank"
                                            >
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
                            <label className={styles['checkbox-label']}>
                                <input {...register('isActive')} className={styles['checkbox']} type="checkbox" />
                                <span className={styles['checkbox-text']}>販売中</span>
                            </label>
                            <label className={styles['checkbox-label']}>
                                <input {...register('isRecommend')} className={styles['checkbox']} type="checkbox" />
                                <span className={styles['checkbox-text']}>おすすめに設定</span>
                            </label>
                        </div>
                    </div>
                </Form>
            )}
        </Dialog>
    )
}
