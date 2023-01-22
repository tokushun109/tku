<template>
    <v-row>
        <v-col cols="12" md="4">
            <c-classification-list :type="ClassificationType.Category.name" :items="categories" @c-change="loadingClassification" />
        </v-col>
        <v-col cols="12" md="4">
            <c-classification-list :type="ClassificationType.Target.name" :items="targets" @c-change="loadingClassification" />
        </v-col>
        <v-col cols="12" md="4">
            <c-classification-list :type="ClassificationType.Tag.name" :items="tags" @c-change="loadingClassification" />
        </v-col>
    </v-row>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IClassification, ClassificationType, IGetClassificationParams, IClientError, serverToClientError } from '~/types'
@Component({
    head: {
        title: '分類一覧',
    },
})
export default class PageAdminClassificationIndex extends Vue {
    ClassificationType: typeof ClassificationType = ClassificationType

    // カテゴリー一覧
    categories: Array<IClassification> = []

    // ターゲット一覧
    targets: Array<IClassification> = []

    // タグ一覧
    tags: Array<IClassification> = []

    error: IClientError | null = null

    async asyncData({ app }: Context) {
        try {
            const params: IGetClassificationParams = {
                mode: 'all',
            }
            const categories = await app.$axios.$get(`/category`, { params })
            const targets = await app.$axios.$get(`/target`, { params })
            const tags = await app.$axios.$get(`/tag`)
            return { categories, targets, tags }
        } catch (e) {
            const error = serverToClientError(e.response)
            return { categories: [], targets: [], tags: [], error }
        }
    }

    mounted() {
        if (this.error) {
            this.$nuxt.error(this.error)
        }
    }

    async loadingClassification(type: string) {
        const params: IGetClassificationParams = {
            mode: 'all',
        }
        if (type === ClassificationType.Category.name) {
            this.categories = await this.$axios.$get(`/category`, { params })
        } else if (type === ClassificationType.Target.name) {
            this.targets = await this.$axios.$get(`/target`, { params })
        } else if (type === ClassificationType.Tag.name) {
            this.tags = await this.$axios.$get(`/tag`)
        }
    }
}
</script>

<style lang="stylus"></style>
