import { TransformToCardList } from '@transforms/cards';

export const getCreditCards = (userId: number) => {
    return fetch(`./../../mocks/user/${userId}/cards.json`)
    .then(data => data.json())
    .then(json => TransformToCardList(json))
}