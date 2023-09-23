import { TransformToCustomList } from './transforms';

export const getCustomsByUserId = async(userId: number) => {
    return await fetch(`./../../mocks/user/${userId}/cards.json`)
    .then(data => data.json())
    .then(json => TransformToCustomList(json))
}