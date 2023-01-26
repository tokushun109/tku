<template>
    <transition name="fade">
        <div class="c-scroll-down">
            <p class="c-scroll-down__message">Scroll<span class="c-scroll-down__message__arrow">→</span></p>
        </div>
    </transition>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { ColorType } from '~/types'

@Component
export default class CScrollDown extends Vue {
    ColorType: typeof ColorType = ColorType
}
</script>

<style lang="stylus" scoped>
.c-scroll-down
    position absolute
    top 90vh
    right 50px
    animation upDown 1s ease-in-out infinite
    +sm()
        top 65vh
        right 25px
    &__message
        position relative
        color $primary
        letter-spacing 0.1em
        -ms-writing-mode tb-rl
        -webkit-writing-mode vertical-rl
        writing-mode vertical-rl
        +sm()
            right -5px
            font-size $font-small
        &__arrow
            margin-top 5px

.fade-leave-active
    animation fadeIn 1s reverse
</style>
