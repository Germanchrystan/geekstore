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

const getUserWhishlist = async(userId: number) => await getWhishlistByUserId(userId);
const getUserPurchases = async(userId: number) => await getPurchasesByUserId(userId);
const getUserCustoms = async(userId: number) => await getCustomsByUserId(userId);
const getUserAddresses = async(userId: number) => await  getAddressesByUserId(userId)
const getUserCards = async(userId: number) => await getCreditCardsByUserId(userId);

const getPanelData = (selectedPanel : string, userId: number) => {
    switch (selectedPanel) {
        case PURCHASES:
            console.log(getUserPurchases(userId));
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
