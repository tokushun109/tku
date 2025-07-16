import { screen, waitFor } from '@testing-library/react'
import { expect } from 'vitest'

// ローディング状態の確認
export const waitForLoadingToDisappear = async () => {
    await waitFor(() => {
        expect(screen.queryByText('Loading...')).not.toBeInTheDocument()
    })
}

// エラーメッセージの確認
export const expectErrorMessage = async (message: string) => {
    await waitFor(() => {
        expect(screen.getByText(message)).toBeInTheDocument()
    })
}

// 商品情報の確認
export const expectProductInfo = async (productName: string, price: number) => {
    await waitFor(() => {
        expect(screen.getByText(productName)).toBeInTheDocument()
        expect(screen.getByText(`¥${price.toLocaleString()}`)).toBeInTheDocument()
    })
}

// 画像の確認
export const expectImageToBeVisible = async (alt: string) => {
    await waitFor(() => {
        const image = screen.getByAltText(alt)
        expect(image).toBeInTheDocument()
        expect(image).toHaveAttribute('src')
    })
}

// フォーム送信の確認
export const expectFormSubmission = async (buttonText: string) => {
    await waitFor(() => {
        const submitButton = screen.getByRole('button', { name: buttonText })
        expect(submitButton).toBeInTheDocument()
        expect(submitButton).not.toBeDisabled()
    })
}

// ナビゲーションの確認
export const expectNavigationLinks = async (links: string[]) => {
    await waitFor(() => {
        links.forEach((linkText) => {
            expect(screen.getByText(linkText)).toBeInTheDocument()
        })
    })
}

// カテゴリー表示の確認
export const expectCategoryDisplay = async (categoryName: string, productCount: number) => {
    await waitFor(() => {
        expect(screen.getByText(categoryName)).toBeInTheDocument()
        const categorySection = screen.getByTestId(`category-${categoryName}`)
        expect(categorySection).toBeInTheDocument()

        // 商品数の確認（必要に応じて調整）
        if (productCount > 0) {
            expect(categorySection).not.toBeEmptyDOMElement()
        }
    })
}

// 検索結果の確認
export const expectSearchResults = async (expectedCount: number) => {
    await waitFor(() => {
        const results = screen.getAllByTestId(/product-item/)
        expect(results).toHaveLength(expectedCount)
    })
}

// カルーセル/スライドショーの確認
export const expectCarouselToBeVisible = async () => {
    await waitFor(() => {
        const carousel = screen.getByTestId('carousel')
        expect(carousel).toBeInTheDocument()

        // 最初のスライドが表示されているか確認
        const firstSlide = screen.getByTestId('carousel-item-0')
        expect(firstSlide).toBeInTheDocument()
    })
}

// モーダルの確認
export const expectModalToBeOpen = async (modalTitle: string) => {
    await waitFor(() => {
        expect(screen.getByText(modalTitle)).toBeInTheDocument()
        expect(screen.getByRole('dialog')).toBeInTheDocument()
    })
}

// パンくずリストの確認
export const expectBreadcrumbs = async (breadcrumbs: string[]) => {
    await waitFor(() => {
        breadcrumbs.forEach((breadcrumb) => {
            expect(screen.getByText(breadcrumb)).toBeInTheDocument()
        })
    })
}
