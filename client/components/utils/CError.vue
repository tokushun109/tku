<template>
    <section v-if="errors && errors.length > 0" class="c-error c-form-item">
        <ul class="errors">
            <li v-for="(error, index) in errors" :key="index" class="error">
                <span v-dompurify-html="error.message.replace(/\n/g, '<br />')"></span>
            </li>
        </ul>
    </section>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IError } from '~/types'

@Component
export default class CError extends Vue {
    @Prop({ type: Array, default: [] }) errors?: Array<IError>
}
</script>

<style lang="stylus">
.c-error
    padding 16px
    border 1px solid $danger-border-color
    border-radius 4px
    background-color $danger-bg-color
    color $danger-color
    font-size 14px
    line-height 170%
    .errors
        @extend .zero-space
        list-style none
        .error
            @extend .zero-space
            font-weight bold
            em
                font-style normal
                font-size 80%
</style>
