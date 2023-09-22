import {
    PURCHASES,
    CUSTOMS,
    WHISHLIST,
    ADDRESSES,
    CARDS
} from '../constants';
import { getWhishlistByUserId } from '@services/whishlist'; //'@services/user/whishlist';
import { getAddressesByUserId } from '@services/address';
import { getPurchasesByUserId } from '@services/purchase';
import { getCreditCardsByUserId } from '@services/credit-card';
import { getCustomsByUserId } from '@services/custom';

const getUserWhishlist = (userId: number) => getWhishlistByUserId(userId);
const getUserPurchases = (userId: number) => getPurchasesByUserId(userId);
const getUserCustoms = (userId: number) => getCustomsByUserId(userId);
const getUserAddresses = (userId: number) => getAddressesByUserId(userId)
const getUserCards = (userId: number) => getCreditCardsByUserId(userId);

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
