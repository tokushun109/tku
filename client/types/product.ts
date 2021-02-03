export interface IProduct {
    uuid: string
    name: string
    description: string
    accessoryCategory: number
    materialCategory: Array<number>
    productImage: string
    salesSite: Array<number>
}
