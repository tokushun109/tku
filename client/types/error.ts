export interface IError {
    name: string
    statusCode: number
    message: string
}
class Error {
    public name: string = 'Error'
    public statusCode: number = 0
    public message: string

    constructor(initMessage: string) {
        this.message = initMessage
    }

    toString() {
        return `${this.name} ${this.message} (${this.statusCode})`
    }
}

export class BadRequest extends Error {
    public name: string = 'BadRequest'
    public statusCode: number = 400
}

export class Unauthorized extends Error {
    public name: string = 'Unauthorized'
    public statusCode: number = 401
}

export class ApplicationError extends Error {
    public name: string = 'ApplicationError'
    public statusCode: number = 403
}
export class InternalServerError extends Error {
    public name: string = 'InternalServerError'
    public statusCode: number = 500
}

interface IErrorResponse {
    statusText: string
    status: number
    data: string
}

// バックエンドから返ってきたエラーをカスタマイズする
export function errorCustomize(errorResponse: IErrorResponse, message: string = errorResponse.data) {
    let error: IError = new Error('')
    error = {
        name: errorResponse.statusText,
        statusCode: errorResponse.status,
        message,
    }
    return error
}
