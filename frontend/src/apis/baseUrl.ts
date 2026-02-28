const defaultApiBaseUrl = 'http://localhost:8080/api'

export const getApiBaseUrl = (): string => {
    if (typeof window !== 'undefined') {
        return process.env.BROWSER_BASE_URL || process.env.API_BASE_URL || defaultApiBaseUrl
    }

    return process.env.API_BASE_URL || process.env.BROWSER_BASE_URL || defaultApiBaseUrl
}
