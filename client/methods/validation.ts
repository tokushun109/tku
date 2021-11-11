// 必須項目
export function required(value: string) {
    return !!value || '入力されていません'
}

// 20文字以内
export function min20(value: string) {
    return value.length <= 20 || `20文字以内で入力してください`
}

// 全角文字を含まない
export function nonDoubleByte(value: string) {
    return !value.match(/^[^\x01-\x7E\xA1-\xDF]+$/) || `URLに全角文字が含まれています`
}

// スペースを含まない
export function nonSpace(value: string) {
    return !value.match(/\s+/) || `URLにスペースが含まれています`
}
