import { ISiteDetail } from '../types/site'
import { IProduct, IClassification, ISite, ICreator, IContact } from '~/types'
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
        price: 1,
        category: newClassification(),
        target: newClassification(),
        tags: [],
        productImages: [],
        siteDetails: [],
        isRecommend: false,
        isActive: true,
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
        icon: '',
    }
}

export function newSiteDetail(): ISiteDetail {
    return {
        uuid: '',
        detailUrl: '',
        salesSite: newSite(),
    }
}

export function newContact(): IContact {
    return {
        name: '',
        company: '',
        phoneNumber: '',
        email: '',
        content: '',
    }
}
