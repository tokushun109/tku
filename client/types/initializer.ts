import { IProduct, IClassification, ISite, ICreator } from '~/types'
export function newCreator(): ICreator {
    return {
        name: '',
        introduction: '',
        logo: '',
        apiPath: '',
    }
}
export function newProduct(): IProduct {
    return {
        uuid: '',
        name: '',
        description: '',
        category: null,
        tags: [],
        productImages: [],
        salesSites: [],
    }
}

export function newClassification(): IClassification {
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
