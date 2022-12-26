<template>
    <div class="error-wrapper">
        <v-sheet color="transparent" class="site-title-area">
            <v-card flat width="300" color="transparent" class="site-title text-h1" to="/"> tocoriri </v-card>
        </v-sheet>
        <v-container>
            <v-container class="error-content">
                <div class="error-message">
                    <div class="error-code">{{ error.statusCode }}</div>
                    <div v-if="error.statusCode === 404">お探しのページは<br class="sm" />見つかりませんでした</div>
                    <div v-else>{{ error.message }}</div>
                </div>
            </v-container>
        </v-container>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IDefaultError } from '~/types'

@Component({})
export default class Error extends Vue {
    @Prop({ type: Object, default: () => {} }) error!: IDefaultError
}
</script>

<style lang="stylus" scoped>
.error-wrapper
    text-align center
    .site-title-area
        position relative
        text-align center
        +sm()
            display none
        .site-title
            margin 0 auto
            color $site-title-text-color
            font-family $title-font-face !important
    .error-content
        height 75vh
        .error-message
            padding 20vh 0
            color $site-title-text-color
            font-size $font-xxlarge
            +sm()
                font-size $font-xlarge
            .sm
                display none
                +sm()
                    display block
            .error-code
                font-size 50px
</style>
