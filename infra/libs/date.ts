// 日付をYYYYMMddhhmmssの文字列で取得
export const getDateString = (date: Date | null = null): string => {
    const d = date || new Date()

    return (
        d.getFullYear() +
        String(d.getMonth() + 1).padStart(2, '0') +
        String(d.getDate()).padStart(2, '0') +
        String(d.getHours()).padStart(2, '0') +
        String(d.getMinutes()).padStart(2, '0') +
        String(d.getSeconds()).padStart(2, '0')
    )
}
