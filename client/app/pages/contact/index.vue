<template>
    <v-container class="page-contact">
        <v-container class="page-title-container">
            <h2 class="page-title text-sm-h3 text-h4">CONTACT</h2>
        </v-container>
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
                                <v-text-field v-model="contact.name" :rules="nameRules" label="お名前(必須)" outlined counter="20" />
                                <v-text-field v-model="contact.company" :rules="companyRules" label="会社名" outlined counter="20" />
                                <v-text-field
                                    v-model="contact.phoneNumber"
                                    :rules="phoneRules"
                                    label="電話番号(-を入れずに入力)"
                                    validate-on-blur
                                    outlined
                                />
                                <v-text-field v-model="contact.email" :rules="emailRules" label="メールアドレス(必須)" outlined validate-on-blur />
                                <v-textarea v-model="contact.content" :rules="contentRules" label="お問い合わせ内容(必須)" outlined />
                                <div class="text-center">
                                    <v-btn color="primary" :disabled="!valid" @click="confirmHandler">送信する</v-btn>
                                </div>
                            </v-form>
                        </v-container>
                    </template>
                    <v-container v-else class="content-message-wrapper">
                        <strong>お問い合せを送信しました</strong>
                    </v-container>
                </v-sheet>
            </v-container>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { min20, min50, newContact, required, validEmail, validPhoneNumber } from '~/methods'
import { IContact, IError } from '~/types'
@Component({
    head: {
        meta: [
            {
                hid: 'robots',
                name: 'robots',
                content: 'noindex',
            },
        ],
    },
})
export default class PageContactIndex extends Vue {
    contact: IContact = newContact()

    isSent: boolean = false

    valid: boolean = false

    errors: Array<IError> = []

    nameRules = [required, min20]

    companyRules = [min20]

    phoneRules = [validPhoneNumber]

    emailRules = [required, min50, validEmail]

    contentRules = [required]

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
    .page-title
        margin-bottom 20px
        color $site-title-text-color
        text-align center

.contact-introduction
    text-align center
    .sm
        display none
        +sm()
            display block

.content-form-wrapper
    .content-message-wrapper
        text-align center
</style>