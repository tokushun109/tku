'use client'

import { useRef, useState } from 'react'
import { toast } from 'sonner'

import { getProductsByCategory } from '@/apis/product'
import { SelectOption } from '@/components/bases/Select'
import { INITIAL_CATEGORY_PRODUCT_LIMIT, MORE_CATEGORY_PRODUCT_LIMIT } from '@/features/product/constants'
import { IProductsByCategory } from '@/features/product/type'

type UseProductTemplateProps = {
    initialProductsByCategory: IProductsByCategory[]
}

interface IFilteredCondition {
    category: SelectOption | undefined
    target: SelectOption | undefined
}

export const useProductTemplate = ({ initialProductsByCategory }: UseProductTemplateProps) => {
    const [filteredCondition, setFilteredCondition] = useState<IFilteredCondition>({ category: undefined, target: undefined })
    const [productsByCategory, setProductsByCategory] = useState<IProductsByCategory[]>(initialProductsByCategory)
    const [isFetchingProducts, setIsFetchingProducts] = useState<boolean>(false)
    const [loadingCategoryUUIDs, setLoadingCategoryUUIDs] = useState<string[]>([])
    const latestRequestIdRef = useRef<number>(0)

    const visibleProductsByCategory = productsByCategory.filter((v) => v.products.length > 0)

    const fetchProducts = async (condition: IFilteredCondition) => {
        const requestId = latestRequestIdRef.current + 1
        latestRequestIdRef.current = requestId

        setIsFetchingProducts(true)
        setLoadingCategoryUUIDs([])

        try {
            const nextProductsByCategory = await getProductsByCategory({
                category: condition.category?.value || 'all',
                limit: INITIAL_CATEGORY_PRODUCT_LIMIT,
                target: condition.target?.value || 'all',
            })

            if (latestRequestIdRef.current !== requestId) return

            setProductsByCategory(nextProductsByCategory)
        } catch (error) {
            if (latestRequestIdRef.current !== requestId) return

            console.error('商品の一覧取得に失敗しました:', error)
            toast.error('商品の一覧取得に失敗しました。もう一度お試しください。')
        } finally {
            if (latestRequestIdRef.current === requestId) {
                setIsFetchingProducts(false)
            }
        }
    }

    const onSelect = (option: SelectOption | undefined, key: keyof IFilteredCondition) => {
        const nextFilteredCondition = { ...filteredCondition, [key]: option }

        setFilteredCondition(nextFilteredCondition)
        void fetchProducts(nextFilteredCondition)
    }

    const onClickMore = async (categoryUUID: string) => {
        if (isFetchingProducts || loadingCategoryUUIDs.includes(categoryUUID)) return

        const targetCategory = productsByCategory.find((v) => v.category.uuid === categoryUUID)
        if (!targetCategory?.pageInfo.hasMore || !targetCategory.pageInfo.nextCursor) return

        setLoadingCategoryUUIDs((prev) => (prev.includes(categoryUUID) ? prev : [...prev, categoryUUID]))

        try {
            const nextProductsByCategory = await getProductsByCategory({
                category: categoryUUID,
                cursor: targetCategory.pageInfo.nextCursor,
                limit: MORE_CATEGORY_PRODUCT_LIMIT,
                target: filteredCondition.target?.value || 'all',
            })
            const nextCategoryProducts = nextProductsByCategory[0]
            if (!nextCategoryProducts) return

            setProductsByCategory((prev) =>
                prev.map((categoryProducts) => {
                    if (categoryProducts.category.uuid !== categoryUUID) return categoryProducts

                    const existingProductUUIDs = new Set(categoryProducts.products.map((product) => product.uuid))
                    const newProducts = nextCategoryProducts.products.filter((product) => {
                        if (existingProductUUIDs.has(product.uuid)) return false

                        existingProductUUIDs.add(product.uuid)
                        return true
                    })
                    const mergedProducts = [...categoryProducts.products, ...newProducts]

                    return {
                        ...categoryProducts,
                        pageInfo: nextCategoryProducts.pageInfo,
                        products: mergedProducts,
                    }
                }),
            )
        } catch (error) {
            console.error('商品の追加読み込みに失敗しました:', error)
            toast.error('商品の追加読み込みに失敗しました。もう一度お試しください。')
        } finally {
            setLoadingCategoryUUIDs((prev) => prev.filter((v) => v !== categoryUUID))
        }
    }

    return {
        isFetchingProducts,
        loadingCategoryUUIDs,
        onClickMore,
        onSelect,
        visibleProductsByCategory,
    }
}
