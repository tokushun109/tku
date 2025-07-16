import { describe, it, expect } from 'vitest'

import { convertObjectToURLSearchParams } from '@/utils/request'

describe('request utils', () => {
    describe('convertObjectToURLSearchParams', () => {
        it('空のオブジェクトを変換する', () => {
            const result = convertObjectToURLSearchParams({})
            expect(result).toBeInstanceOf(URLSearchParams)
            expect(result.toString()).toBe('')
        })

        it('単一のキーと値を変換する', () => {
            const result = convertObjectToURLSearchParams({ key: 'value' })
            expect(result.get('key')).toBe('value')
            expect(result.toString()).toBe('key=value')
        })

        it('複数のキーと値を変換する', () => {
            const result = convertObjectToURLSearchParams({
                name: 'John',
                age: '30',
                city: 'Tokyo',
            })
            expect(result.get('name')).toBe('John')
            expect(result.get('age')).toBe('30')
            expect(result.get('city')).toBe('Tokyo')
        })

        it('数値を文字列として変換する', () => {
            const result = convertObjectToURLSearchParams({
                count: 42,
                price: 1500,
            })
            expect(result.get('count')).toBe('42')
            expect(result.get('price')).toBe('1500')
        })

        it('boolean値を文字列として変換する', () => {
            const result = convertObjectToURLSearchParams({
                isActive: true,
                isVisible: false,
            })
            expect(result.get('isActive')).toBe('true')
            expect(result.get('isVisible')).toBe('false')
        })

        it('nullとundefinedを文字列として変換する', () => {
            const result = convertObjectToURLSearchParams({
                nullValue: null,
                undefinedValue: undefined,
            })
            expect(result.get('nullValue')).toBe('null')
            expect(result.get('undefinedValue')).toBe('undefined')
        })

        it('空文字列を処理する', () => {
            const result = convertObjectToURLSearchParams({
                emptyString: '',
                normalString: 'value',
            })
            expect(result.get('emptyString')).toBe('')
            expect(result.get('normalString')).toBe('value')
        })

        it('特殊文字を含む値を処理する', () => {
            const result = convertObjectToURLSearchParams({
                specialChars: 'hello world!',
                encoded: 'hello%20world',
                japanese: 'こんにちは',
            })
            expect(result.get('specialChars')).toBe('hello world!')
            expect(result.get('encoded')).toBe('hello%20world')
            expect(result.get('japanese')).toBe('こんにちは')
        })

        it('URLSearchParamsのメソッドが正しく動作する', () => {
            const result = convertObjectToURLSearchParams({
                param1: 'value1',
                param2: 'value2',
                param3: 'value3',
            })

            expect(result.has('param1')).toBe(true)
            expect(result.has('param4')).toBe(false)
            expect(result.get('param1')).toBe('value1')
            expect(result.get('param4')).toBe(null)
        })

        it('同じキーを複数回設定した場合、最後の値が使われる', () => {
            const obj = { key: 'value1' }
            const result = convertObjectToURLSearchParams(obj)
            result.set('key', 'value2')
            expect(result.get('key')).toBe('value2')
        })

        it('配列を文字列として変換する', () => {
            const result = convertObjectToURLSearchParams({
                arrayValue: ['item1', 'item2', 'item3'],
            })
            expect(result.get('arrayValue')).toBe('item1,item2,item3')
        })

        it('オブジェクトを文字列として変換する', () => {
            const result = convertObjectToURLSearchParams({
                objectValue: { nested: 'value' },
            })
            expect(result.get('objectValue')).toBe('[object Object]')
        })

        it('日付オブジェクトを文字列として変換する', () => {
            const date = new Date('2023-01-01T00:00:00.000Z')
            const result = convertObjectToURLSearchParams({
                dateValue: date,
            })
            expect(result.get('dateValue')).toBe(date.toString())
        })

        it('APIクエリパラメータの典型的な使用例', () => {
            const result = convertObjectToURLSearchParams({
                category: 'electronics',
                minPrice: 100,
                maxPrice: 500,
                inStock: true,
                sort: 'price',
                order: 'asc',
            })

            expect(result.get('category')).toBe('electronics')
            expect(result.get('minPrice')).toBe('100')
            expect(result.get('maxPrice')).toBe('500')
            expect(result.get('inStock')).toBe('true')
            expect(result.get('sort')).toBe('price')
            expect(result.get('order')).toBe('asc')
        })

        it('検索フォームの典型的な使用例', () => {
            const result = convertObjectToURLSearchParams({
                q: 'search term',
                page: 1,
                limit: 20,
                filter: 'active',
            })

            expect(result.get('q')).toBe('search term')
            expect(result.get('page')).toBe('1')
            expect(result.get('limit')).toBe('20')
            expect(result.get('filter')).toBe('active')
        })

        it('toStringメソッドが正しいクエリ文字列を生成する', () => {
            const result = convertObjectToURLSearchParams({
                a: '1',
                b: '2',
                c: '3',
            })

            const queryString = result.toString()
            expect(queryString).toBe('a=1&b=2&c=3')
        })

        it('URLエンコードが正しく処理される', () => {
            const result = convertObjectToURLSearchParams({
                'special chars': 'hello world & more',
                unicode: 'こんにちは世界',
            })

            const queryString = result.toString()
            expect(queryString).toContain('special+chars=hello+world+%26+more')
            expect(queryString).toContain('unicode=%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF%E4%B8%96%E7%95%8C')
        })

        it('大きなオブジェクトを処理する', () => {
            const largeObj: Record<string, string> = {}
            for (let i = 0; i < 100; i++) {
                largeObj[`key${i}`] = `value${i}`
            }

            const result = convertObjectToURLSearchParams(largeObj)
            expect(result.get('key0')).toBe('value0')
            expect(result.get('key50')).toBe('value50')
            expect(result.get('key99')).toBe('value99')
        })

        it('返されたURLSearchParamsが変更可能である', () => {
            const result = convertObjectToURLSearchParams({ initial: 'value' })

            result.set('new', 'added')
            result.delete('initial')

            expect(result.has('initial')).toBe(false)
            expect(result.get('new')).toBe('added')
        })
    })
})
