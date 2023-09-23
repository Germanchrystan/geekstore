import { TransformToCardList } from './transforms';

export const getCreditCardsByUserId = async(userId: number) => {
    return await fetch(`./../../mocks/user/${userId}/cards.json`)
    .then(data => data.json())
    .then(json => TransformToCardList(json))
}