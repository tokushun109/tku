export function dateFormat(dateString: string, format: string = ''): string {
    if (!format) {
        format = 'YYYY/MM/DD hh:mm:ss'
    }

    const date = new Date(dateString)
    format = format.replace(/YYYY/g, date.getFullYear().toString())
    format = format.replace(/MM/g, ('0' + (date.getMonth() + 1)).toString().slice(-2))
    format = format.replace(/DD/g, ('0' + date.getDate()).toString().slice(-2))
    format = format.replace(/hh/g, ('0' + date.getHours()).toString().slice(-2))
    format = format.replace(/mm/g, ('0' + date.getMinutes()).toString().slice(-2))
    format = format.replace(/ss/g, ('0' + date.getSeconds()).toString().slice(-2))
    return format
}
