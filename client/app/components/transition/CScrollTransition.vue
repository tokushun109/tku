<template>
    <div ref="content">
        <transition name="fade">
            <div v-if="isDisplay">
                <slot />
            </div>
        </transition>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component
export default class CScrollAppear extends Vue {
    @Prop({ type: Boolean, default: false }) init!: boolean
    isDisplay: boolean = false

    mounted() {
        if (this.init) {
            this.isDisplay = true
        }
        window.addEventListener('scroll', this.handleScroll)
    }

    destroyed() {
        window.removeEventListener('scroll', this.handleScroll)
    }

    handleScroll() {
        if (!this.isDisplay) {
            const content = this.$refs.content as Element
            const top = content.getBoundingClientRect().top
            this.isDisplay = top < window.innerHeight + 100
        }
    }
}
</script>

<style lang="stylus" scoped>
.fade-enter-active
    animation fadeUp 1s

.fade-leave-active
    animation fadeUp 1s reverse
</style>
