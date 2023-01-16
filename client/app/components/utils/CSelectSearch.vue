<template>
    <div class="c-select-search">
        <v-tabs>
            <v-tab v-for="item in selectSearchItems" :key="item.uuid">{{ item.name }}</v-tab>
        </v-tabs>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IClassification } from '~/types'

@Component({})
export default class CSelectSearch extends Vue {
    @Prop({ type: Array, default: () => [] }) items!: Array<IClassification>
    selectSearchItems: Array<IClassification> = [{ name: 'all', uuid: 'all' }, ...this.items]
}
</script>

<style lang="stylus" scoped>
.c-select-search
    margin-bottom 16px
</style>
