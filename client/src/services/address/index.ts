import { TransformToAddressList } from './transforms';

export const getAddressesByUserId = (userId: number) => {
    return fetch(`./../../mocks/user/${userId}/addresses.json`)
    .then(data => data.json())
    .then(json => TransformToAddressList(json))
}