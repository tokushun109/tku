export interface IServerError {
    name: string
    status: number
    data: string
}

export interface IClientError {
    statusCode: number
    path?: string
    message: string
}

export type IError = IServerError | IClientError

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

// サーバーからのエラーをクライアントエラーに変換する
export function serverToClientError(errorResponse: IServerError): IClientError {
    const error: IClientError = {
        statusCode: errorResponse.status,
        message: errorResponse.data,
    }
    return error
}

// バックエンドから返ってきたエラーをカスタマイズする
export function errorCustomize(errorResponse: IErrorResponse, data: string = errorResponse.data) {
    let error: IServerError = new Error('')
    error = {
        name: errorResponse.statusText,
        status: errorResponse.status,
        data,
    }
    return error
}
