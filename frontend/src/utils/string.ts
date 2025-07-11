/**
 * 文字列の先頭のみ大文字に変換
 * @param {string} str 対象の文字列
 * @return {string} 変換された文字列を返す
 */
export const capitalize = function (str: string) {
    if (typeof str !== 'string' || !str) return str
    return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase()
}
