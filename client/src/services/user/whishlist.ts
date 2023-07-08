import { TransformToProductList } from '../../transforms/product';

export const getWhishlist = (userId: number) => {
    return fetch("./../../mocks/user/1/whishlist.json")
    .then(data => data.json())
    .then(json => TransformToProductList(json))
}

