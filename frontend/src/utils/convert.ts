// 数値をカンマ区切りの価格形式に変換する
export const numToPrice = (num: number): string => `¥${num.toLocaleString()}`
