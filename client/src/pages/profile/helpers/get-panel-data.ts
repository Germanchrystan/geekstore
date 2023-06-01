import {
    PURCHASES,
    CUSTOMS,
    WHISHLIST,
    ADDRESSES,
    CARDS
} from '../constants';
import { CardProps } from '../../../components/card';

const getUserPurchases = (userId: String): [] => {

};

const getUserCustoms = (userId: String) => {

};

const getUserWhishlist = (userId: String) => {

};

const getUserAddresses = (userId: String) => {

};

const getUserCards = (userId: String) => {

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
