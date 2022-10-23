export interface IUser {
    uuid: string
    name: string
    email: string
    isAdmin: boolean
}

export interface ILoginForm {
    email: string
    password: string
}
export interface ISession {
    uuid: string
}
