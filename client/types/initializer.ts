import { IProduct, ICategory, ISite } from '~/types'

export function newProduct(): IProduct {
    return {
        uuid: '',
        name: '',
        description: '',
        accessoryCategory: null,
        materialCategories: [],
        productImages: [],
        salesSites: [],
    }
}

export function newCategory(): ICategory {
    return {
        uuid: '',
        name: '',
    }
}

export function newSite(): ISite {
    return {
        uuid: '',
        name: '',
        url: '',
    }
}
