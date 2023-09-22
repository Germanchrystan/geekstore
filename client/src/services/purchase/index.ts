import { TransformToPurchaseList } from "./transforms";

export const getPurchasesByUserId = (userId: number) => {
    return fetch(`./../../mocks/user/${userId}/whishlist.json`)
    .then(data => data.json())
    .then(json => TransformToPurchaseList(json));
}