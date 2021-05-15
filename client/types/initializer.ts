import { IProduct } from './product'

export function newProduct(): IProduct {
    return {
        uuid: '',
        name: '',
        description: '',
        accessoryCategory: null,
        materialCategories: [],
        productImage: null,
        salesSites: [],
    }
}
