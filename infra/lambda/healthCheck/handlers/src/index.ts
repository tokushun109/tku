import fetch from 'node-fetch'

export const handler = async () => {

    try {
        const res = await fetch('https://api.tocoriri.com/api/health_check', {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })
        const result = await res.json() as { success?: boolean }
        if (result.success === undefined) {
            throw new Error('api取得失敗')
        }
        return {
            statusCode: 200,
            body: {
                message: 'api取得OK',
                result
            },
        }
    } catch (e) {
        return {
            statusCode: 500,
            body: {
                message: 'apiの取得に失敗しました',
                result: null
            },
        }
    }
}
