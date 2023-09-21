export interface tokenData {
    accessToken: string
    selectedProfile: {
        name: string
    }
}

export interface Api<T> {
    code: number
    msg: string
    data: T
}

export type ApiErr = Api<unknown> 