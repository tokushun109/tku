import ErrorPage from './error'

const NotFoundPage = () => {
    return <ErrorPage errorMessage="お探しのページは見つかりませんでした" statusCode={404} />
}

export default NotFoundPage
