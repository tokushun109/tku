<template>
    <c-layout-container normal class="c-product-detail-page-wrapper">
        <c-product-detail v-if="product" :product="product" />
        <c-breadcrumbs :items="breadCrumbs" />
    </c-layout-container>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IBreadCrumb, IClientError, IProduct, serverToClientError } from '~/types'

@Component({})
export default class PageProductDetail extends Vue {
    product: IProduct | null = null
    error: IClientError | null = null

    breadCrumbs: Array<IBreadCrumb> = [
        { text: 'トップページ', href: '/' },
        { text: '商品一覧', href: '/product' },
        { text: this.product ? this.product.name : '', disabled: true },
    ]

    async asyncData({ app, params }: Context) {
        try {
            const product = await app.$axios.$get(`/product/${params.uuid}`)
            return { product }
        } catch (e) {
            const error = serverToClientError(e.response)
            return { product: null, error }
        }
    }

    mounted() {
        if (this.error) {
            this.$nuxt.error(this.error)
        }
    }

    head() {
        if (!this.product) {
            return
        }
        const title = `${this.product.name} | とこりり`
        const description = this.product.description.replace(/\r?\n/g, '')
        const image = this.product.productImages[0].apiPath
        return {
            title,
            meta: [
                {
                    hid: 'description',
                    name: 'description',
                    content: description,
                },
                {
                    hid: 'og:title',
                    property: 'og:title',
                    content: title,
                },
                {
                    hid: 'og:description',
                    property: 'og:description',
                    content: description,
                },
                {
                    hid: 'og:type',
                    property: 'og:type',
                    content: 'product',
                },
                {
                    hid: 'og:image',
                    property: 'og:image',
                    content: image,
                },
            ],
        }
    }
}
</script>

<style lang="stylus"></style>
