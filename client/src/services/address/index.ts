import { TransformToAddressList } from './transforms';

export const getAddressesByUserId = async(userId: number) => {
    return await fetch(`./../../mocks/user/${userId}/addresses.json`)
    .then(data => {
        console.log(data.json())
        return data.json()})
    .then(json => TransformToAddressList(json))
}