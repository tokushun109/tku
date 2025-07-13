import ErrorPage from './error'

const NotFoundPage: React.FC = () => {
    return <ErrorPage errorMessage="お探しのページは見つかりませんでした" statusCode={404} />
}

export default NotFoundPage
