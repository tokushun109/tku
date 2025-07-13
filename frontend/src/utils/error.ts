// クライアアンドのエラー
export interface IClientError {
    message: string
    path?: string
    statusCode: number
}

// apiに関するエラー
export class ApiError extends Error {
    public readonly statusCode
    constructor(res: Response) {
        super(res.statusText)
        this.statusCode = res.status
    }
}
