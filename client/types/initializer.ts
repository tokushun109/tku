import { IAccessoryCategory, IMaterialCategory } from './category'
import { IProduct } from './product'

export function newProduct(): IProduct {
    return {
        uuid: '',
        name: '',
        description: '',
        accessoryCategory: newAccessoryCategory(),
        materialCategories: new Array(newMaterialCategory()),
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
