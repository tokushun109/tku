import { describe, it, expect } from 'vitest'

import { ApiError, type IClientError } from '@/utils/error'

describe('error utils', () => {
    describe('IClientError interface', () => {
        it('必須プロパティを持つClientErrorオブジェクトを作成できる', () => {
            const error: IClientError = {
                message: 'エラーメッセージ',
                statusCode: 400,
            }

            expect(error.message).toBe('エラーメッセージ')
            expect(error.statusCode).toBe(400)
            expect(error.path).toBeUndefined()
        })

        it('オプションのpathプロパティを持つClientErrorオブジェクトを作成できる', () => {
            const error: IClientError = {
                message: 'エラーメッセージ',
                statusCode: 404,
                path: '/api/products',
            }

            expect(error.message).toBe('エラーメッセージ')
            expect(error.statusCode).toBe(404)
            expect(error.path).toBe('/api/products')
        })

        it('様々なstatusCodeを設定できる', () => {
            const error400: IClientError = {
                message: 'Bad Request',
                statusCode: 400,
            }

            const error401: IClientError = {
                message: 'Unauthorized',
                statusCode: 401,
            }

            const error500: IClientError = {
                message: 'Internal Server Error',
                statusCode: 500,
            }

            expect(error400.statusCode).toBe(400)
            expect(error401.statusCode).toBe(401)
            expect(error500.statusCode).toBe(500)
        })
    })

    describe('ApiError class', () => {
        it('Responseオブジェクトから正しくApiErrorを作成する', () => {
            const mockResponse = {
                status: 404,
                statusText: 'Not Found',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error).toBeInstanceOf(ApiError)
            expect(error).toBeInstanceOf(Error)
            expect(error.message).toBe('Not Found')
            expect(error.statusCode).toBe(404)
        })

        it('500エラーレスポンスからApiErrorを作成する', () => {
            const mockResponse = {
                status: 500,
                statusText: 'Internal Server Error',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error.message).toBe('Internal Server Error')
            expect(error.statusCode).toBe(500)
        })

        it('401エラーレスポンスからApiErrorを作成する', () => {
            const mockResponse = {
                status: 401,
                statusText: 'Unauthorized',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error.message).toBe('Unauthorized')
            expect(error.statusCode).toBe(401)
        })

        it('400エラーレスポンスからApiErrorを作成する', () => {
            const mockResponse = {
                status: 400,
                statusText: 'Bad Request',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error.message).toBe('Bad Request')
            expect(error.statusCode).toBe(400)
        })

        it('statusCodeプロパティがreadonlyである', () => {
            const mockResponse = {
                status: 404,
                statusText: 'Not Found',
            } as Response

            const error = new ApiError(mockResponse)

            // TypeScriptのreadonly制約はランタイムには影響しないため、値は変更される
            // しかし、TypeScriptレベルでは制約されている
            expect(() => {
                // @ts-expect-error readonlyプロパティへの代入をテスト
                error.statusCode = 500
            }).not.toThrow()

            // 実際にはJavaScriptレベルでは変更されるが、TypeScriptで制約されている
            expect(error.statusCode).toBe(500)
        })

        it('エラーメッセージが継承される', () => {
            const mockResponse = {
                status: 403,
                statusText: 'Forbidden',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error.message).toBe('Forbidden')
            expect(error.toString()).toBe('Error: Forbidden')
        })

        it('空のstatusTextでもエラーオブジェクトが作成される', () => {
            const mockResponse = {
                status: 200,
                statusText: '',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error.message).toBe('')
            expect(error.statusCode).toBe(200)
        })

        it('カスタムstatusTextからエラーを作成する', () => {
            const mockResponse = {
                status: 422,
                statusText: 'Unprocessable Entity',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error.message).toBe('Unprocessable Entity')
            expect(error.statusCode).toBe(422)
        })

        it('throwできる', () => {
            const mockResponse = {
                status: 404,
                statusText: 'Not Found',
            } as Response

            expect(() => {
                throw new ApiError(mockResponse)
            }).toThrow(ApiError)

            expect(() => {
                throw new ApiError(mockResponse)
            }).toThrow('Not Found')
        })

        it('try-catch文でキャッチできる', () => {
            const mockResponse = {
                status: 404,
                statusText: 'Not Found',
            } as Response

            try {
                throw new ApiError(mockResponse)
            } catch (error) {
                expect(error).toBeInstanceOf(ApiError)
                expect((error as ApiError).statusCode).toBe(404)
                expect((error as ApiError).message).toBe('Not Found')
            }
        })

        it('instanceof演算子が正しく動作する', () => {
            const mockResponse = {
                status: 404,
                statusText: 'Not Found',
            } as Response

            const error = new ApiError(mockResponse)

            expect(error instanceof ApiError).toBe(true)
            expect(error instanceof Error).toBe(true)
            expect(error instanceof Object).toBe(true)
        })
    })
})
