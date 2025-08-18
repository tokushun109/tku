import fetch from 'node-fetch'

export const handler = async () => {
    try {
        await fetch('https://tocoriri.com/health_check', {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })
        return {
            statusCode: 200,
            body: JSON.stringify({
                message: 'health_checkにアクセスしました',
                result: null,
            }),
        }
    } catch (e) {
        return {
            statusCode: 500,
            body: JSON.stringify({
                message: 'health_checkのアクセスに失敗しました',
                result: null,
            }),
        }
    }
}
