const jstDateTimeFormatter = new Intl.DateTimeFormat('ja-JP', {
    timeZone: 'Asia/Tokyo',
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
})

export const formatDateToJST = (dateString: string): string => {
    return jstDateTimeFormatter.format(new Date(dateString))
}
