<template>
    <transition v-if="isRoot" name="fade">
        <div v-if="isDisplay" class="c-top-loading">
            <div class="c-top-loading__logo">
                <img src="img/logo/tocoriri_logo_white.png" class="fadeUp" />
            </div>
        </div>
    </transition>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component({})
export default class CTopLoading extends Vue {
    isDisplay: boolean = true

    @Prop({ type: Boolean, default: false }) isRoot!: boolean

    mounted() {
        setTimeout(() => {
            this.isDisplay = false
        }, 1000)
    }
}
</script>

<style lang="stylus">
.c-top-loading
    position fixed
    z-index 999
    width 100%
    height 100%
    background $secondary
    text-align center
    &__logo
        position absolute
        top 50%
        left 50%
        transform translate(-50%, -50%)
        img
            width 400px
            height 200px
            object-fit cover
        .fadeUp
            opacity 0
            animation-name fadeUpAnime
            animation-duration 0.5s
            animation-fill-mode forwards

@keyframes fadeUpAnime
    from
        opacity 0
        transform translateY(50px)
    to
        opacity 1
        transform translateY(0)

.fade-leave-active
    animation fadeOut 0.7s

@keyframes fadeOut
    from
        opacity 1
    to
        opacity 0
</style>
