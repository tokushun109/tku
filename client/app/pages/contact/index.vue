<template>
    <c-layout-container class="page-contact" narrow>
        <div class="page-title-container">
            <h1 class="page-title text-sm-h3 text-h4">Contact</h1>
        </div>
        <div v-if="!isSent" class="contact-introduction">
            <p>お問い合わせ・<br class="sm" />ご意見・ご相談はこちらから</p>
        </div>
        <v-sheet>
            <v-container>
                <v-sheet class="content-form-wrapper">
                    <template v-if="!isSent">
                        <v-container>
                            <c-error :errors.sync="errors" />
                            <v-form ref="form" v-model="valid" lazy-validation>
                                <div class="require-form">
                                    <v-chip v-if="!contact.name" class="require-chip">必須</v-chip>
                                    <v-text-field v-model="contact.name" :rules="nameRules" label="お名前" outlined counter="20" />
                                </div>
                                <v-text-field v-model="contact.company" :rules="companyRules" label="会社名" outlined counter="20" />
                                <v-text-field
                                    v-model="contact.phoneNumber"
                                    :rules="phoneRules"
                                    label="電話番号(-を入れずに入力)"
                                    validate-on-blur
                                    outlined
                                />
                                <div class="require-form">
                                    <v-chip v-if="!contact.email" class="require-chip">必須</v-chip>
                                    <v-text-field v-model="contact.email" :rules="emailRules" label="メールアドレス" outlined validate-on-blur />
                                </div>
                                <div class="require-form">
                                    <v-chip v-if="!contact.content" class="require-chip">必須</v-chip>
                                    <v-textarea v-model="contact.content" :rules="contentRules" label="お問い合わせ内容" outlined />
                                </div>
                                <div class="text-center">
                                    <v-btn color="primary" :disabled="!valid" @click="confirmHandler">送信する</v-btn>
                                </div>
                            </v-form>
                        </v-container>
                    </template>
                    <v-container v-else class="content-message-wrapper">
                        <strong>お問い合わせを送信しました</strong>
                    </v-container>
                </v-sheet>
            </v-container>
        </v-sheet>
        <c-breadcrumbs :items="breadCrumbs" />
    </c-layout-container>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { min20, min50, newContact, required, validEmail, validPhoneNumber } from '~/methods'
import { IBreadCrumb, IContact, ICreator, IError } from '~/types'
@Component({})
export default class PageContactIndex extends Vue {
    creator: ICreator | null = null

    contact: IContact = newContact()

    isSent: boolean = false

    valid: boolean = false

    errors: Array<IError> = []

    nameRules = [required, min20]

    companyRules = [min20]

    phoneRules = [validPhoneNumber]

    emailRules = [required, min50, validEmail]

    contentRules = [required]

    breadCrumbs: Array<IBreadCrumb> = [
        { text: 'トップページ', href: '/' },
        { text: 'お問い合わせ', disabled: true },
    ]

    async asyncData({ app }: Context) {
        try {
            const creator: ICreator = await app.$axios.$get(`/creator`)

            return { creator }
        } catch (e) {
            return { creator: null }
        }
    }

    head() {
        const title = 'お問い合わせ | とこりり'
        const description = 'マクラメ編みのアクセサリーショップ【とこりり】へのお問い合わせ・ご意見・ご相談はこちらから'
        const image = this.creator && this.creator.apiPath ? this.creator.apiPath : ''
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
                    content: 'article',
                },
                {
                    hid: 'og:image',
                    property: 'og:image',
                    content: image,
                },
            ],
        }
    }

    async confirmHandler() {
        this.errors = []
        const refs: any = this.$refs.form
        await refs.validate()
        try {
            if (!this.valid) {
                return
            }
            await this.$axios.$post(`/contact`, this.contact)
            this.isSent = true
            await setTimeout(() => {
                return this.$router.push('/')
            }, 3000)
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus" scoped>
.page-title-container
    text-align center
    +sm()
        display none

.contact-introduction
    text-align center
    .sm
        display none
        +sm()
            display block

.content-form-wrapper
    .require-form
        position relative
        .require-chip
            position absolute
            top 13px
            right 20px
            background-color $danger-bg-color
            color $danger-color
    .content-message-wrapper
        text-align center
</style>
