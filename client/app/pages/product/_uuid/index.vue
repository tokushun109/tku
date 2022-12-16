<template>
    <c-product-detail :product="product" />
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProduct } from '~/types'

@Component({})
export default class PageProductDetail extends Vue {
    product: IProduct | null = null
    async asyncData({ app, params }: Context) {
        try {
            const product = await app.$axios.$get(`/product/${params.uuid}`)
            return { product }
        } catch (e) {
            return { product: null }
        }
    }

    head() {
        if (!this.product) {
            return
        }
        const title = `tocoriri | ${this.product.name}`
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
