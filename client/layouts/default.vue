<template>
    <div>
        <Nuxt />
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({
    head: {
        title: 'head title',
    },
})
export default class LayoutDefault extends Vue {}
</script>

<style lang="stylus"></style>
