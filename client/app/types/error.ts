export interface IAxiosError {
    name: string
    status: number
    data: string
}

export interface IDefaultError {
    statusCode: number
    path: string
    message: string
}

export type IError = IAxiosError | IDefaultError

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
    let error: IAxiosError = new Error('')
    error = {
        name: errorResponse.statusText,
        status: errorResponse.status,
        data,
    }
    return error
}
