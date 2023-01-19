<template>
    <v-sheet>
        <c-contact-list :contact-list="formatContactList" />
    </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { dateFormat } from '~/methods'
import { IClientError, IContact, serverToClientError } from '~/types'
@Component({
    head: {
        title: 'お問い合わせ',
    },
})
export default class PageAdminContactIndex extends Vue {
    // お問合わせリスト
    contactList: Array<IContact> = []
    error: IClientError | null = null

    // createdAtをフォーマットしたリスト
    get formatContactList(): Array<IContact> {
        for (const contact of this.contactList) {
            contact.formatCreatedAt = dateFormat(contact.createdAt!)
        }
        return this.contactList
    }

    async asyncData({ app }: Context) {
        try {
            const contactList = await app.$axios.$get(`/contact`, { withCredentials: true })
            return { contactList }
        } catch (e) {
            const error = serverToClientError(e.response)
            return { contactList: [], error }
        }
    }

    mounted() {
        if (this.error) {
            this.$nuxt.error(this.error)
        }
    }
}
</script>

<style lang="stylus"></style>
