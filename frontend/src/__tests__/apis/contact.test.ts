import { http, HttpResponse } from 'msw'
import { describe, it, expect, beforeEach } from 'vitest'

import { postContact } from '@/apis/contact'
import { ApiError } from '@/utils/error'

import { server } from '../mocks/server'

describe('contact API', () => {
    beforeEach(() => {
        server.resetHandlers()
    })

    describe('postContact', () => {
        const mockContactForm = {
            name: 'テスト太郎',
            email: 'test@example.com',
            content: 'これはテストメッセージです。',
        }

        it('正常にお問い合わせを送信する', async () => {
            server.use(
                http.post('http://localhost:8080/contact', async ({ request }) => {
                    const body = await request.json()

                    // リクエストボディが正しいかチェック
                    expect(body).toEqual(mockContactForm)

                    return HttpResponse.json({
                        message: 'お問い合わせを受け付けました',
                    })
                }),
            )

            const result = await postContact({ form: mockContactForm })

            expect(result).toEqual({
                message: 'お問い合わせを受け付けました',
            })
        })

        it('APIエラーの場合、ApiErrorが投げられる', async () => {
            server.use(
                http.post('http://localhost:8080/contact', () => {
                    return new HttpResponse(null, { status: 500 })
                }),
            )

            await expect(postContact({ form: mockContactForm })).rejects.toThrow(ApiError)
        })

        it('ネットワークエラーの場合、一般的なエラーが投げられる', async () => {
            server.use(
                http.post('http://localhost:8080/contact', () => {
                    return HttpResponse.error()
                }),
            )

            await expect(postContact({ form: mockContactForm })).rejects.toThrow('お問い合わせの送信に失敗しました')
        })

        it('正しいContent-Typeヘッダーでリクエストされる', async () => {
            let capturedHeaders: Headers | undefined

            server.use(
                http.post('http://localhost:8080/contact', ({ request }) => {
                    capturedHeaders = request.headers
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            await postContact({ form: mockContactForm })

            expect(capturedHeaders?.get('content-type')).toBe('application/json')
        })

        it('正しいHTTPメソッドでリクエストされる', async () => {
            let capturedMethod: string | undefined

            server.use(
                http.post('http://localhost:8080/contact', ({ request }) => {
                    capturedMethod = request.method
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            await postContact({ form: mockContactForm })

            expect(capturedMethod).toBe('POST')
        })

        it('正しいエンドポイントにリクエストされる', async () => {
            let capturedPath: string | undefined

            server.use(
                http.post('http://localhost:8080/contact', ({ request }) => {
                    capturedPath = new URL(request.url).pathname
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            await postContact({ form: mockContactForm })

            expect(capturedPath).toBe('/contact')
        })

        it('400エラーの場合、適切にApiErrorが投げられる', async () => {
            server.use(
                http.post('http://localhost:8080/contact', () => {
                    return new HttpResponse(null, { status: 400 })
                }),
            )

            try {
                await postContact({ form: mockContactForm })
            } catch (error) {
                expect(error).toBeInstanceOf(ApiError)
                expect((error as ApiError).statusCode).toBe(400)
            }
        })

        it('空のフォームデータでもリクエストが送信される', async () => {
            const emptyForm = {
                name: '',
                email: '',
                content: '',
            }

            server.use(
                http.post('http://localhost:8080/contact', async ({ request }) => {
                    const body = await request.json()
                    expect(body).toEqual(emptyForm)
                    return HttpResponse.json({ message: 'success' })
                }),
            )

            const result = await postContact({ form: emptyForm })

            expect(result).toEqual({
                message: 'success',
            })
        })
    })
})
