/**
 * 価格をフォーマットする（3桁区切り）
 */
export const formatPrice = (price: number): string => {
    return price.toLocaleString('ja-JP')
}
