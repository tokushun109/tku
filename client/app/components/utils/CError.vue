<template>
    <v-container v-if="errors && errors.length > 0">
        <c-message color="error">
            <ul class="errors">
                <li v-for="(error, index) in errors" :key="index" class="error">
                    <span v-dompurify-html="error.data.replace(/\n/g, '<br />')"></span>
                </li>
            </ul>
        </c-message>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IError } from '~/types'

@Component
export default class CError extends Vue {
    @Prop({ type: Array, default: [] }) errors?: Array<IError>
}
</script>

<style lang="stylus" scoped></style>
