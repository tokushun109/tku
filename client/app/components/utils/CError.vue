<template>
    <v-container v-if="errors && errors.length > 0">
        <c-message :color="ColorType.Accent">
            <ul class="errors">
                <li v-for="(error, index) in errors" :key="index">
                    <span v-if="'data' in error" v-dompurify-html="error.data.replace(/\n/g, '<br />')"></span>
                    <span v-else-if="'message' in error" v-dompurify-html="error.message.replace(/\n/g, '<br />')"></span>
                </li>
            </ul>
        </c-message>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { ColorType, IError } from '~/types'

@Component
export default class CError extends Vue {
    @Prop({ type: Array, default: () => [] }) errors?: Array<IError>
    ColorType: typeof ColorType = ColorType
}
</script>

<style lang="stylus" scoped></style>
