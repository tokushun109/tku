<template>
    <v-footer class="admin-footer" color="primary" dark app>tku</v-footer>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class AdminFooter extends Vue {}
</script>

<style lang="stylus" scoped>
.admin-footer
    display flex
    justify-content center
</style>
