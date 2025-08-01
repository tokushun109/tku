/**
 * CSV関連のAPI呼び出し
 */

/**
 * 商品レコードCSVをダウンロード
 */
export const downloadProductCsv = async (): Promise<void> => {
    try {
        const response = await fetch(`${process.env.API_URL}/csv/product`, {
            method: 'GET',
            credentials: 'include',
        })

        if (!response.ok) {
            throw new Error('CSVダウンロードに失敗しました')
        }

        const csvText = await response.text()
        const blob = new Blob([csvText], { type: 'text/csv' })
        const link = document.createElement('a')
        link.href = URL.createObjectURL(blob)
        link.download = '商品レコード.csv'
        link.click()
    } catch (error) {
        console.error('CSV download error:', error)
        throw new Error('CSVダウンロードに失敗しました')
    }
}

/**
 * 商品レコードCSVをアップロード
 * @param file - アップロードするCSVファイル
 */
export const uploadProductCsv = async (file: File): Promise<void> => {
    try {
        const formData = new FormData()
        formData.append('csv', file)

        const response = await fetch(`${process.env.API_URL}/csv/product`, {
            method: 'POST',
            credentials: 'include',
            body: formData,
        })

        if (!response.ok) {
            throw new Error('CSVアップロードに失敗しました')
        }
    } catch (error) {
        console.error('CSV upload error:', error)
        throw new Error('CSVアップロードに失敗しました')
    }
}
