import { afterEach, describe, it, expect, vi } from 'vitest'

import { getApiBaseUrl } from '@/apis/baseUrl'
import { postContact } from '@/apis/contact'
import { ApiError } from '@/utils/error'

describe('contact API', () => {
    afterEach(() => {
        vi.restoreAllMocks()
    })

    describe('postContact', () => {
        const mockContactForm = {
            name: 'テスト太郎',
            email: 'test@example.com',
            content: 'これはテストメッセージです。',
        }

        const createJSONResponse = (body: unknown, init?: ConstructorParameters<typeof Response>[1]) =>
            new Response(JSON.stringify(body), {
                headers: {
                    'Content-Type': 'application/json',
                },
                status: 200,
                ...init,
            })

        it('正常にお問い合わせを送信する', async () => {
            const fetchSpy = vi.spyOn(globalThis, 'fetch').mockResolvedValue(
                createJSONResponse({
                    message: 'お問い合わせを受け付けました',
                }),
            )

            const result = await postContact({ form: mockContactForm })

            const [requestUrl, requestInit] = fetchSpy.mock.calls[0]

            expect(requestUrl).toBe(`${getApiBaseUrl()}/contact`)
            expect(JSON.parse(requestInit?.body as string)).toEqual(mockContactForm)
            expect(result).toEqual({
                message: 'お問い合わせを受け付けました',
            })
        })

        it('APIエラーの場合、ApiErrorが投げられる', async () => {
            vi.spyOn(globalThis, 'fetch').mockResolvedValue(new Response(null, { status: 500, statusText: 'Internal Server Error' }))

            await expect(postContact({ form: mockContactForm })).rejects.toThrow(ApiError)
        })

        it('ネットワークエラーの場合、一般的なエラーが投げられる', async () => {
            vi.spyOn(globalThis, 'fetch').mockRejectedValue(new Error('Network Error'))

            await expect(postContact({ form: mockContactForm })).rejects.toThrow('お問い合わせの送信に失敗しました')
        })

        it('正しいContent-Typeヘッダーでリクエストされる', async () => {
            const fetchSpy = vi.spyOn(globalThis, 'fetch').mockResolvedValue(createJSONResponse({ message: 'success' }))

            await postContact({ form: mockContactForm })

            const [, requestInit] = fetchSpy.mock.calls[0]
            const headers = requestInit?.headers as Record<string, string> | undefined

            expect(headers?.['Content-Type']).toBe('application/json')
        })

        it('正しいHTTPメソッドでリクエストされる', async () => {
            const fetchSpy = vi.spyOn(globalThis, 'fetch').mockResolvedValue(createJSONResponse({ message: 'success' }))

            await postContact({ form: mockContactForm })

            const [, requestInit] = fetchSpy.mock.calls[0]

            expect(requestInit?.method).toBe('POST')
        })

        it('正しいエンドポイントにリクエストされる', async () => {
            const fetchSpy = vi.spyOn(globalThis, 'fetch').mockResolvedValue(createJSONResponse({ message: 'success' }))

            await postContact({ form: mockContactForm })

            const [requestUrl] = fetchSpy.mock.calls[0]

            expect(new URL(requestUrl as string).pathname).toBe(new URL(`${getApiBaseUrl()}/contact`).pathname)
        })

        it('400エラーの場合、適切にApiErrorが投げられる', async () => {
            vi.spyOn(globalThis, 'fetch').mockResolvedValue(new Response(null, { status: 400, statusText: 'Bad Request' }))

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

            const fetchSpy = vi.spyOn(globalThis, 'fetch').mockResolvedValue(createJSONResponse({ message: 'success' }))

            const result = await postContact({ form: emptyForm })

            const [, requestInit] = fetchSpy.mock.calls[0]

            expect(JSON.parse(requestInit?.body as string)).toEqual(emptyForm)
            expect(result).toEqual({
                message: 'success',
            })
        })
    })
})
