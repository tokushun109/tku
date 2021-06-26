<template>
    <c-page class="page-login" title="ログイン">
        <c-form bordered slim>
            <c-error :errors.sync="errors" />
            <c-input-label label="メールアドレス" required>
                <c-input :model.sync="form.email" />
            </c-input-label>
            <c-input-label label="パスワード" required>
                <c-input :model.sync="form.password" password />
            </c-input-label>
            <div class="form-actions">
                <c-button label="ログイン" primary @c-click="onSubmit" />
            </div>
        </c-form>
    </c-page>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { IError, ILoginForm, BadRequest } from '~/types'

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

    errors: Array<IError> = []

    async onSubmit() {
        try {
            this.errors = []
            // バリデーション
            if (this.form.email.length === 0) {
                throw new BadRequest('メールアドレスが入力されていません')
            }
            if (this.form.password.length === 0) {
                throw new BadRequest('パスワードが入力されていません')
            }
            await this.$store.dispatch('user/loginUser', this.form)
            this.$router.replace('/admin')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
