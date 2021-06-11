import { IAccessoryCategory, IMaterialCategory, IProduct, ISalesSite, ISkillMarket, ISns } from '~/types'

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

export function newSalesSite(): ISalesSite {
    return {
        uuid: '',
        name: '',
        url: '',
    }
}

export function newSkillMarket(): ISkillMarket {
    return {
        uuid: '',
        name: '',
        url: '',
    }
}

export function newSns(): ISns {
    return {
        uuid: '',
        name: '',
        url: '',
    }
}
