export const handler = async () => {
    return {
        statusCode: 200,
        body: {
            message: 'バリデーションOK',
            params: 'test',
        },
    }
}
