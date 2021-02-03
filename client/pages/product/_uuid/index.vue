<template>
    <div>{{ product }}</div>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IProduct } from '~/types'

@Component({
    head: {
        title: '商品詳細',
    },
})
export default class PageProductDetail extends Vue {
    product: IProduct | null = null
    async asyncData({ app, params }: Context) {
        try {
            const product = await app.$axios.$get(`/product/${params.uuid}`)
            return { product }
        } catch (e) {
            return { product: '' }
        }
    }
}
</script>

<style lang="stylus"></style>
