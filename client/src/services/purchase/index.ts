import { TransformToPurchaseList } from "./transforms";

export const getPurchasesByUserId = async(userId: number) => {
    return await fetch(`./../../mocks/user/${userId}/whishlist.json`)
    .then(data => data.json())
    .then(json => TransformToPurchaseList(json));
}