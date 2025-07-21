import { getCategories } from '@/apis/category'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'

import { ClassificationTemplate } from './template'

const ClassificationPage = async () => {
    try {
        const [categories, targets, tags] = await Promise.all([getCategories({ mode: 'all' }), getTargets({ mode: 'all' }), getTags()])

        return <ClassificationTemplate categories={categories} tags={tags} targets={targets} />
    } catch (error) {
        console.error('データの取得に失敗しました:', error)
        return <ClassificationTemplate categories={[]} tags={[]} targets={[]} />
    }
}

export default ClassificationPage
