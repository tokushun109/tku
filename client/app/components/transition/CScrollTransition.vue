<template>
    <div ref="content">
        <transition :name="transition">
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
    @Prop({ type: String, default: 'fade' }) transition!: string
    isDisplay: boolean = false

    mounted() {
        window.addEventListener('scroll', this.handleScroll)
    }

    destroyed() {
        window.removeEventListener('scroll', this.handleScroll)
    }

    handleScroll() {
        if (!this.isDisplay) {
            const content = this.$refs.content as Element
            const top = content.getBoundingClientRect().top
            this.isDisplay = top < window.innerHeight - 50
        }
    }
}
</script>

<style lang="stylus" scoped>
.fade-enter-active
    animation fadeUp 1s

.shake-enter-active
    animation shake 1.5s
    animation-delay 2s
</style>
