export interface IError {
    name: string
    status: number
    data?: string
    message?: string
}
class Error {
    public name: string = 'Error'
    public status: number = 0
    public data: string

    constructor(initMessage: string) {
        this.data = initMessage
    }

    toString() {
        return `${this.name} ${this.data} (${this.status})`
    }
}

export class BadRequest extends Error {
    public name: string = 'BadRequest'
    public status: number = 400
}

export class Unauthorized extends Error {
    public name: string = 'Unauthorized'
    public status: number = 401
}

export class ApplicationError extends Error {
    public name: string = 'ApplicationError'
    public status: number = 403
}
export class InternalServerError extends Error {
    public name: string = 'InternalServerError'
    public status: number = 500
}

interface IErrorResponse {
    statusText: string
    status: number
    data: string
}

// バックエンドから返ってきたエラーをカスタマイズする
export function errorCustomize(errorResponse: IErrorResponse, data: string = errorResponse.data) {
    let error: IError = new Error('')
    error = {
        name: errorResponse.statusText,
        status: errorResponse.status,
        data,
    }
    return error
}
