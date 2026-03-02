import { describe, expect, it } from 'vitest'

import { formatDateToJST } from '@/utils/date'

describe('date utils', () => {
    describe('formatDateToJST', () => {
        it('UTC の日時文字列を日本時間で整形する', () => {
            expect(formatDateToJST('2026-03-02T03:00:00Z')).toBe('2026/03/02 12:00')
        })

        it('同じ実時刻なら入力のタイムゾーンが違っても同じ日本時間になる', () => {
            expect(formatDateToJST('2026-03-02T12:00:00+09:00')).toBe('2026/03/02 12:00')
        })
    })
})
