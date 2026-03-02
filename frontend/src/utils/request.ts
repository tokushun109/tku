/** オブジェクトをURLSearchParamsに変換 */
export const convertObjectToURLSearchParams = (params: object): URLSearchParams => {
    const result = new URLSearchParams()

    for (const [key, value] of Object.entries(params)) {
        if (value === undefined || value === null) continue

        result.set(key, `${value}`)
    }

    return result
}
