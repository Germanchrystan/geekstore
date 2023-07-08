import { TransformToCustomList } from '@transforms/custom';

export const getCustoms = (userId: number) => {
    return fetch(`./../../mocks/user/${userId}/cards.json`)
    .then(data => data.json())
    .then(json => TransformToCustomList(json))
}