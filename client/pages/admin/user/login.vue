<template>
    <v-container class="page-admin-login">
        <v-sheet class="admin-login-area" max-width="800">
            <c-error :errors.sync="errors" />
            <h3 class="login-title">ログイン</h3>
            <v-form ref="form" v-model="valid" lazy-validation>
                <v-text-field v-model="form.email" :rules="rules" label="email(必須)" outlined />
                <v-text-field v-model="form.password" type="password" :rules="rules" label="パスワード(必須)" outlined />
                <v-btn color="primary" :disabled="!valid" @click="onSubmit">確定</v-btn>
            </v-form>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { required } from '~/methods'
import { IError, ILoginForm } from '~/types'

@Component({
    head: {
        titleTemplate: 'ログイン | admin',
    },
})
export default class PageAdminUserLogin extends Vue {
    form: ILoginForm = {
        email: '',
        password: '',
    }

    valid: boolean = true

    rules = [required]

    errors: Array<IError> = []

    async onSubmit() {
        try {
            this.errors = []
            await this.$store.dispatch('user/loginUser', this.form)
            this.$router.replace('/admin/product')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus" scoped>
.page-admin-login
    .admin-login-area
        margin 0 auto
        padding 16px
        text-align center
        .login-title
            margin-bottom 16px
            color $title-text-color
</style>
