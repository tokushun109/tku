/** オブジェクトをURLSearchParamsに変換 */
export const convertObjectToURLSearchParams = (params: object): URLSearchParams => {
    const result = new URLSearchParams()

    for (const [key, value] of Object.entries(params)) {
        result.set(key, value)
    }

    return result
}
