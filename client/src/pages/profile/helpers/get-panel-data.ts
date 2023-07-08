import {
    PURCHASES,
    CUSTOMS,
    WHISHLIST,
    ADDRESSES,
    CARDS
} from '../constants';
import { getWhishlist } from '../../../services/user/whishlist';

const getUserPurchases = (userId: number) => {

};

const getUserCustoms = (userId: number) => {

};

const getUserWhishlist = (userId: number) => getWhishlist(userId);

const getUserAddresses = (userId: number) => {

};

const getUserCards = (userId: number) => {

};

const getPanelData = (selectedPanel : string, userId: number) => {
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
