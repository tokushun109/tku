<template>
    <v-list class="c-select-search">
        <v-list-group v-model="isExpand" :prepend-icon="mdiChevronDown" no-action dense>
            <template #activator>
                <v-list-item-content>
                    <v-list-item-title class="c-select-search__title">{{ groupName }} - {{ targetName }} </v-list-item-title>
                </v-list-item-content>
            </template>
            <v-list-item
                v-for="(item, index) in selectItems"
                :key="item.uuid"
                :input-value="item.isActive"
                :color="ColorType.Accent"
                @click="selectItem(index)"
            >
                <v-list-item-title class="c-select-search__item" v-text="item.name"></v-list-item-title>
            </v-list-item>
        </v-list-group>
    </v-list>
</template>

<script lang="ts">
import { mdiChevronDown } from '@mdi/js'
import { Component, Vue, Prop, PropSync } from 'nuxt-property-decorator'
import { ColorType, IClassification } from '~/types'

interface ISelectSearchItem extends IClassification {
    isActive: boolean
}

@Component({})
export default class CSelectSearch extends Vue {
    @PropSync('targetContent') content!: string
    @Prop({ type: Array, default: () => [] }) items!: Array<IClassification>
    @Prop({ type: String, default: '' }) groupName!: string

    ColorType: typeof ColorType = ColorType
    mdiChevronDown = mdiChevronDown

    isExpand: boolean = false
    selectItems: Array<ISelectSearchItem> = []

    get targetName(): string {
        for (const selectItem of this.selectItems) {
            if (selectItem.isActive) {
                return selectItem.name
            }
        }
        return 'all'
    }

    initSelectSearchItems() {
        this.selectItems = [{ name: 'All', uuid: 'all', isActive: true }]
        for (const item of this.items) {
            const newSelectItem: ISelectSearchItem = {
                ...item,
                isActive: false,
            }
            this.selectItems.push(newSelectItem)
        }
    }

    mounted() {
        this.initSelectSearchItems()
    }

    selectItem(index: number) {
        this.selectItems.forEach((item, i) => {
            if (i === index) {
                item.isActive = true
                this.content = item.uuid
                this.$emit('c-select-search')
                setTimeout(() => {
                    this.isExpand = false
                }, 200)
            } else {
                item.isActive = false
            }
        })
    }
}
</script>

<style lang="stylus" scoped>
::v-deep .v-list-item__icon
    margin-right 10px !important

::v-deep .v-list-group__header__append-icon
    display none

::v-deep .v-list-group--no-action > .v-list-group__items > .v-list-item
    padding-left 50px

.c-select-search
    padding 0
    min-width 350px
    +sm()
        max-width inherit
        width 100%
    .v-list-item__icon
        margin-right 10px !important
    &__title
    &__item
        color $text-color
        +sm()
            font-size 12px
</style>
