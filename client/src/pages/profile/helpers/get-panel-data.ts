import {
    PURCHASES,
    CUSTOMS,
    WHISHLIST,
    ADDRESSES,
    CARDS
} from '../constants';
import { CardProps } from '../../../components/card';

const getUserPurchases = (userId: String): [] => {
    fetch('./../../../mocks/user_purchases.json')
    .then(res => res.json())
    .then(data => console.log(data))
};

const getUserCustoms = (userId: String) => {
    fetch('./../../../mocks/user_customs.json')
    .then(res => res.json())
    .then(data => console.log(data));
};

const getUserWhishlist = (userId: String) => {
    fetch('./../../../mocks/user_whishlist.json')
    .then(res => res.json())
    .then(data => console.log(data));
};

const getUserAddresses = (userId: String) => {
    fetch('./../../../mocks/user_addresses.json')
    .then(res => res.json())
    .then(data => console.log(data));
};

const getUserCards = (userId: String) => {
    fetch('./../../../mocks/user_credit_cards.json')
    .then(res => res.json())
    .then(data => console.log(data));
};

const getPanelData = (selectedPanel : String, userId: String) => {
    switch (selectedPanel) {
        case PURCHASES:
            return getUserPurchases(userId);
        case CUSTOMS:
            return getUserCustoms(userId);
        case WHISHLIST:
            return getUserWhishlist(userId);
        case ADDRESSES:
            return getUserAddresses(userId);
        case CARDS:
            return getUserCards(userId);
        default:
            return null;
    }
}

export default getPanelData;
