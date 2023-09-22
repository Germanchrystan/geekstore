import { TransformToProductList } from './transforms';

export const getWhishlistByUserId = (userId: number) => {
    console.log("WHISHLIST");
    return fetch(`./../../mocks/user/${userId}/whishlist.json`)
    .then(data => data.json())
    .then(json => TransformToProductList(json));
}

