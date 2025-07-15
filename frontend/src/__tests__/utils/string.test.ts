import { describe, it, expect } from 'vitest'

import { capitalize } from '@/utils/string'

describe('string utils', () => {
    describe('capitalize', () => {
        it('通常の文字列を正しく変換する', () => {
            expect(capitalize('hello')).toBe('Hello')
            expect(capitalize('world')).toBe('World')
            expect(capitalize('javascript')).toBe('Javascript')
        })

        it('大文字の文字列を正しく変換する', () => {
            expect(capitalize('HELLO')).toBe('Hello')
            expect(capitalize('WORLD')).toBe('World')
            expect(capitalize('JAVASCRIPT')).toBe('Javascript')
        })

        it('混合文字列を正しく変換する', () => {
            expect(capitalize('hELLO')).toBe('Hello')
            expect(capitalize('wORLD')).toBe('World')
            expect(capitalize('jAvAsCrIpT')).toBe('Javascript')
        })

        it('単一文字の文字列を正しく変換する', () => {
            expect(capitalize('a')).toBe('A')
            expect(capitalize('z')).toBe('Z')
            expect(capitalize('A')).toBe('A')
            expect(capitalize('Z')).toBe('Z')
        })

        it('数字で始まる文字列を正しく変換する', () => {
            expect(capitalize('1hello')).toBe('1hello')
            expect(capitalize('9world')).toBe('9world')
            expect(capitalize('123abc')).toBe('123abc')
        })

        it('空文字列を正しく処理する', () => {
            expect(capitalize('')).toBe('')
        })

        it('null値を正しく処理する', () => {
            expect(capitalize(null as any)).toBe(null)
        })

        it('undefined値を正しく処理する', () => {
            expect(capitalize(undefined as any)).toBe(undefined)
        })

        it('非文字列型を正しく処理する', () => {
            expect(capitalize(123 as any)).toBe(123)
            expect(capitalize(true as any)).toBe(true)
            expect(capitalize([] as any)).toEqual([])
            expect(capitalize({} as any)).toEqual({})
        })

        it('特殊文字で始まる文字列を正しく処理する', () => {
            expect(capitalize('!hello')).toBe('!hello')
            expect(capitalize('@world')).toBe('@world')
            expect(capitalize('#javascript')).toBe('#javascript')
            expect(capitalize('$money')).toBe('$money')
        })

        it('スペースで始まる文字列を正しく処理する', () => {
            expect(capitalize(' hello')).toBe(' hello')
            expect(capitalize('  world')).toBe('  world')
            expect(capitalize('\thello')).toBe('\thello')
            expect(capitalize('\nhello')).toBe('\nhello')
        })

        it('複数単語の文字列を正しく処理する（先頭文字のみ大文字）', () => {
            expect(capitalize('hello world')).toBe('Hello world')
            expect(capitalize('HELLO WORLD')).toBe('Hello world')
            expect(capitalize('hello WORLD')).toBe('Hello world')
        })

        it('日本語文字列を正しく処理する', () => {
            expect(capitalize('こんにちは')).toBe('こんにちは')
            expect(capitalize('Hello世界')).toBe('Hello世界')
            expect(capitalize('hELLO世界')).toBe('Hello世界')
        })

        it('絵文字を含む文字列を正しく処理する', () => {
            expect(capitalize('😀hello')).toBe('😀hello')
            expect(capitalize('hello😀')).toBe('Hello😀')
            expect(capitalize('Hello😀world')).toBe('Hello😀world')
        })

        it('典型的な名前の処理', () => {
            expect(capitalize('john')).toBe('John')
            expect(capitalize('mary')).toBe('Mary')
            expect(capitalize('smith')).toBe('Smith')
            expect(capitalize('JOHN')).toBe('John')
            expect(capitalize('MARY')).toBe('Mary')
        })

        it('アルファベット以外の文字が先頭の場合', () => {
            expect(capitalize('áhello')).toBe('Áhello')
            expect(capitalize('üworld')).toBe('Üworld')
            expect(capitalize('çhello')).toBe('Çhello')
        })

        it('長い文字列を正しく処理する', () => {
            const longString = 'a'.repeat(1000)
            const result = capitalize(longString)
            expect(result.charAt(0)).toBe('A')
            expect(result.length).toBe(1000)
            expect(result.slice(1)).toBe('a'.repeat(999))
        })

        it('特殊な文字列パターン', () => {
            expect(capitalize('a-b-c')).toBe('A-b-c')
            expect(capitalize('a_b_c')).toBe('A_b_c')
            expect(capitalize('a.b.c')).toBe('A.b.c')
            expect(capitalize('a,b,c')).toBe('A,b,c')
        })

        it('最初の文字がすでに大文字の場合', () => {
            expect(capitalize('Hello')).toBe('Hello')
            expect(capitalize('World')).toBe('World')
            expect(capitalize('JavaScript')).toBe('Javascript')
        })

        it('文字列の残りの部分が正しく小文字に変換される', () => {
            expect(capitalize('hELLO wORLD')).toBe('Hello world')
            expect(capitalize('jAvAsCrIpT iS cOoL')).toBe('Javascript is cool')
            expect(capitalize('TEST STRING')).toBe('Test string')
        })

        it('境界値のテスト', () => {
            expect(capitalize('A')).toBe('A')
            expect(capitalize('a')).toBe('A')
            expect(capitalize('AA')).toBe('Aa')
            expect(capitalize('aa')).toBe('Aa')
            expect(capitalize('aA')).toBe('Aa')
            expect(capitalize('Aa')).toBe('Aa')
        })
    })
})
