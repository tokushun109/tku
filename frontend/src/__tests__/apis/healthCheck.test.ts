import { http, HttpResponse } from 'msw'
import { describe, it, expect, beforeEach } from 'vitest'

import { healthCheck } from '@/apis/healthCheck'
import { ApiError } from '@/utils/error'

import { server } from '../mocks/server'

describe('healthCheck API', () => {
    beforeEach(() => {
        server.resetHandlers()
    })

    describe('healthCheck', () => {
        it('正常にヘルスチェックを実行する', async () => {
            server.use(
                http.get('http://localhost:8080/health_check', () => {
                    return HttpResponse.json({
                        message: 'API server is healthy',
                    })
                }),
            )

            const result = await healthCheck()

            expect(result).toEqual({
                message: 'API server is healthy',
            })
        })

        it('メッセージなしでもレスポンスを処理する', async () => {
            server.use(
                http.get('http://localhost:8080/health_check', () => {
                    return HttpResponse.json({})
                }),
            )

            const result = await healthCheck()

            expect(result).toEqual({})
        })

        it('APIエラーの場合、ApiErrorが投げられる', async () => {
            server.use(
                http.get('http://localhost:8080/health_check', () => {
                    return new HttpResponse(null, { status: 500 })
                }),
            )

            await expect(healthCheck()).rejects.toThrow(ApiError)
        })

        it('ネットワークエラーの場合、一般的なエラーが投げられる', async () => {
            server.use(
                http.get('http://localhost:8080/health_check', () => {
                    return HttpResponse.error()
                }),
            )

            await expect(healthCheck()).rejects.toThrow('APIサーバーのヘルスチェックに失敗しました')
        })

        it('正しいHTTPメソッドでリクエストされる', async () => {
            let capturedMethod: string | undefined

            server.use(
                http.get('http://localhost:8080/health_check', ({ request }) => {
                    capturedMethod = request.method
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            await healthCheck()

            expect(capturedMethod).toBe('GET')
        })

        it('正しいエンドポイントにリクエストされる', async () => {
            let capturedPath: string | undefined

            server.use(
                http.get('http://localhost:8080/health_check', ({ request }) => {
                    capturedPath = new URL(request.url).pathname
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            await healthCheck()

            expect(capturedPath).toBe('/health_check')
        })

        it('正しいContent-Typeヘッダーでリクエストされる', async () => {
            let capturedHeaders: Headers | undefined

            server.use(
                http.get('http://localhost:8080/health_check', ({ request }) => {
                    capturedHeaders = request.headers
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            await healthCheck()

            expect(capturedHeaders?.get('content-type')).toBe('application/json')
        })

        it('503エラーの場合、適切にApiErrorが投げられる', async () => {
            server.use(
                http.get('http://localhost:8080/health_check', () => {
                    return new HttpResponse(null, { status: 503 })
                }),
            )

            try {
                await healthCheck()
            } catch (error) {
                expect(error).toBeInstanceOf(ApiError)
                expect((error as ApiError).statusCode).toBe(503)
            }
        })

        it('404エラーの場合、適切にApiErrorが投げられる', async () => {
            server.use(
                http.get('http://localhost:8080/health_check', () => {
                    return new HttpResponse(null, { status: 404 })
                }),
            )

            try {
                await healthCheck()
            } catch (error) {
                expect(error).toBeInstanceOf(ApiError)
                expect((error as ApiError).statusCode).toBe(404)
            }
        })
    })
})
