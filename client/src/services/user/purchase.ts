import { TransformToPurchaseList } from "@transforms/purchase";

export const getPurchases = (userId: number) => {
    return fetch(`./../../mocks/user/${userId}/whishlist.json`)
    .then(data => data.json())
    .then(json => TransformToPurchaseList(json));
}