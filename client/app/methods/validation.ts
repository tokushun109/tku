// 必須項目
export function required(value: string) {
    return !!value || '入力されていません'
}

// 20文字以内
export function min20(value: string) {
    return value.length <= 20 || `20文字以内で入力してください`
}

// 50文字以内
export function min50(value: string) {
    return value.length <= 50 || `0文字以内で入力してください`
}

// 全角文字を含まない
export function nonDoubleByte(value: string) {
    return !value.match(/^[^\x01-\x7E\xA1-\xDF]+$/) || `URLに全角文字が含まれています`
}

// スペースを含まない
export function nonSpace(value: string) {
    return !value.match(/\s+/) || `URLにスペースが含まれています`
}

// 金額
export function price(value: number) {
    return !!value.toString().match(/^([1-9]\d*|0)(\.\d+)?$/) || `金額を入力してください`
}

// 最大金額
export function maxPrice(value: number) {
    return !(value > 1000000) || `最大金額を超えています`
}

// 電話番号
export function validPhoneNumber(value: string) {
    if (!value) {
        return true
    }
    return !!value.match(/^0[0-9]{9,10}$/) || `入力内容を確認してください`
}

// メールアドレス
export function validEmail(value: string) {
    if (!value) {
        return true
    }
    return !!value.match(/^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/) || `入力内容を確認してください`
}
