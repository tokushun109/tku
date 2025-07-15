import { describe, it, expect } from 'vitest'

import { formatPrice } from '@/utils/price'

describe('price utils', () => {
    describe('formatPrice', () => {
        it('正の整数を日本語の3桁区切り形式に変換する', () => {
            expect(formatPrice(1000)).toBe('1,000')
            expect(formatPrice(1500)).toBe('1,500')
            expect(formatPrice(2000)).toBe('2,000')
            expect(formatPrice(10000)).toBe('10,000')
        })

        it('大きな数値を3桁区切り形式に変換する', () => {
            expect(formatPrice(100000)).toBe('100,000')
            expect(formatPrice(1000000)).toBe('1,000,000')
            expect(formatPrice(1234567)).toBe('1,234,567')
            expect(formatPrice(12345678)).toBe('12,345,678')
        })

        it('小さな数値を正しく変換する', () => {
            expect(formatPrice(0)).toBe('0')
            expect(formatPrice(1)).toBe('1')
            expect(formatPrice(10)).toBe('10')
            expect(formatPrice(100)).toBe('100')
            expect(formatPrice(999)).toBe('999')
        })

        it('負の数値を3桁区切り形式に変換する', () => {
            expect(formatPrice(-1000)).toBe('-1,000')
            expect(formatPrice(-1500)).toBe('-1,500')
            expect(formatPrice(-123456)).toBe('-123,456')
            expect(formatPrice(-1234567)).toBe('-1,234,567')
        })

        it('小数点を含む数値を3桁区切り形式に変換する', () => {
            expect(formatPrice(1000.5)).toBe('1,000.5')
            expect(formatPrice(1234.56)).toBe('1,234.56')
            expect(formatPrice(999.99)).toBe('999.99')
            expect(formatPrice(1234567.89)).toBe('1,234,567.89')
        })

        it('3桁区切りの境界値を正しく処理する', () => {
            expect(formatPrice(999)).toBe('999') // 3桁未満は区切りなし
            expect(formatPrice(1000)).toBe('1,000') // 4桁で区切り
            expect(formatPrice(9999)).toBe('9,999')
            expect(formatPrice(10000)).toBe('10,000')
            expect(formatPrice(99999)).toBe('99,999')
            expect(formatPrice(100000)).toBe('100,000')
        })

        it('非常に大きな数値を処理する', () => {
            expect(formatPrice(999999999)).toBe('999,999,999')
            expect(formatPrice(1000000000)).toBe('1,000,000,000')
            expect(formatPrice(1234567890)).toBe('1,234,567,890')
        })

        it('Number.MAX_SAFE_INTEGER近くの値を処理する', () => {
            const largeNum = 9007199254740991 // Number.MAX_SAFE_INTEGER
            expect(formatPrice(largeNum)).toBe('9,007,199,254,740,991')
        })

        it('非常に小さな正の数値を処理する', () => {
            expect(formatPrice(0.01)).toBe('0.01')
            expect(formatPrice(0.1)).toBe('0.1')
            expect(formatPrice(0.001)).toBe('0.001')
            // 0.0001は日本語ロケールで0に丸められる
            expect(formatPrice(0.0001)).toBe('0')
        })

        it('非常に小さな負の数値を処理する', () => {
            expect(formatPrice(-0.01)).toBe('-0.01')
            expect(formatPrice(-0.1)).toBe('-0.1')
            expect(formatPrice(-0.001)).toBe('-0.001')
        })

        it('境界値のテスト', () => {
            expect(formatPrice(999.999)).toBe('999.999')
            expect(formatPrice(1000.001)).toBe('1,000.001')
            expect(formatPrice(-999.999)).toBe('-999.999')
            expect(formatPrice(-1000.001)).toBe('-1,000.001')
        })

        it('複数の小数点桁数を処理する', () => {
            expect(formatPrice(1234.1)).toBe('1,234.1')
            expect(formatPrice(1234.12)).toBe('1,234.12')
            expect(formatPrice(1234.123)).toBe('1,234.123')
            // 日本語ロケールでは小数点が丸められる場合がある
            expect(formatPrice(1234.1234)).toBe('1,234.123')
        })

        it('ゼロと非常に小さな値の境界', () => {
            expect(formatPrice(0)).toBe('0')
            expect(formatPrice(0.0)).toBe('0')
            expect(formatPrice(-0)).toBe('-0')
            expect(formatPrice(-0.0)).toBe('-0')
        })

        it('商品価格の典型的な値をテストする', () => {
            // 一般的な商品価格の例
            expect(formatPrice(500)).toBe('500')
            expect(formatPrice(1500)).toBe('1,500')
            expect(formatPrice(2980)).toBe('2,980')
            expect(formatPrice(5000)).toBe('5,000')
            expect(formatPrice(9800)).toBe('9,800')
            expect(formatPrice(15000)).toBe('15,000')
            expect(formatPrice(29800)).toBe('29,800')
            expect(formatPrice(50000)).toBe('50,000')
        })

        it('日本語ロケールが正しく適用される', () => {
            // 日本語ロケールでは3桁区切りはカンマ
            expect(formatPrice(1000)).toBe('1,000')
            expect(formatPrice(1000000)).toBe('1,000,000')

            // 小数点は英数字のピリオド
            expect(formatPrice(1000.5)).toBe('1,000.5')
        })
    })
})
