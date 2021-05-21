import { IAccessoryCategory, IMaterialCategory } from './category'
import { IProduct } from './product'

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

export function newAccessoryCategory(): IAccessoryCategory {
    return {
        uuid: '',
        name: '',
    }
}

export function newMaterialCategory(): IMaterialCategory {
    return {
        uuid: '',
        name: '',
    }
}
