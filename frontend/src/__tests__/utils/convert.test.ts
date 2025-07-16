import { describe, it, expect } from 'vitest'

import { numToPrice } from '@/utils/convert'

describe('convert utils', () => {
    describe('numToPrice', () => {
        it('正の整数を価格形式に変換する', () => {
            expect(numToPrice(1000)).toBe('¥1,000')
            expect(numToPrice(1500)).toBe('¥1,500')
            expect(numToPrice(2000)).toBe('¥2,000')
            expect(numToPrice(10000)).toBe('¥10,000')
        })

        it('大きな数値を価格形式に変換する', () => {
            expect(numToPrice(100000)).toBe('¥100,000')
            expect(numToPrice(1000000)).toBe('¥1,000,000')
            expect(numToPrice(1234567)).toBe('¥1,234,567')
        })

        it('小さな数値を価格形式に変換する', () => {
            expect(numToPrice(0)).toBe('¥0')
            expect(numToPrice(1)).toBe('¥1')
            expect(numToPrice(10)).toBe('¥10')
            expect(numToPrice(100)).toBe('¥100')
        })

        it('小数点を含む数値を価格形式に変換する', () => {
            expect(numToPrice(1000.5)).toBe('¥1,000.5')
            expect(numToPrice(1234.56)).toBe('¥1,234.56')
            expect(numToPrice(999.99)).toBe('¥999.99')
        })

        it('負の数値を価格形式に変換する', () => {
            expect(numToPrice(-1000)).toBe('¥-1,000')
            expect(numToPrice(-1500)).toBe('¥-1,500')
            expect(numToPrice(-123456)).toBe('¥-123,456')
        })

        it('三桁区切りが正しく適用される', () => {
            expect(numToPrice(999)).toBe('¥999') // 3桁未満は区切りなし
            expect(numToPrice(1000)).toBe('¥1,000') // 4桁で区切り
            expect(numToPrice(9999)).toBe('¥9,999')
            expect(numToPrice(10000)).toBe('¥10,000')
            expect(numToPrice(99999)).toBe('¥99,999')
            expect(numToPrice(100000)).toBe('¥100,000')
        })

        it('非常に大きな数値を処理する', () => {
            expect(numToPrice(999999999)).toBe('¥999,999,999')
            expect(numToPrice(1000000000)).toBe('¥1,000,000,000')
        })

        it('Number.MAX_SAFE_INTEGER近くの値を処理する', () => {
            const largeNum = 9007199254740991 // Number.MAX_SAFE_INTEGER
            expect(numToPrice(largeNum)).toBe('¥9,007,199,254,740,991')
        })

        it('非常に小さな正の数値を処理する', () => {
            expect(numToPrice(0.01)).toBe('¥0.01')
            expect(numToPrice(0.1)).toBe('¥0.1')
            expect(numToPrice(0.001)).toBe('¥0.001')
        })

        it('境界値のテスト', () => {
            expect(numToPrice(999.999)).toBe('¥999.999')
            expect(numToPrice(1000.001)).toBe('¥1,000.001')
            expect(numToPrice(-999.999)).toBe('¥-999.999')
            expect(numToPrice(-1000.001)).toBe('¥-1,000.001')
        })
    })
})
