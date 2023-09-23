import { TransformToProductList } from './transforms';

export const getWhishlistByUserId = async(userId: number) => {
    console.log("WHISHLIST");
    return await fetch(`./../../mocks/user/${userId}/whishlist.json`)
    .then(data => data.json())
    .then(json => TransformToProductList(json));
}

