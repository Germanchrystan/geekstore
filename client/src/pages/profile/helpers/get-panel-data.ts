import {
    PURCHASES,
    CUSTOMS,
    WHISHLIST,
    ADDRESSES,
    CARDS
} from '../constants';
import { getWhishlist } from '@services/user/whishlist';
import { getAddresses } from '@services/user/address';
import { getPurchases } from '@services/user/purchase';
import { getCreditCards } from '@services/user/card';
import { getCustoms } from '@services/user/custom';

const getUserWhishlist = (userId: number) => getWhishlist(userId);
const getUserPurchases = (userId: number) => getPurchases(userId);
const getUserCustoms = (userId: number) => getCustoms(userId);
const getUserAddresses = (userId: number) => getAddresses(userId)
const getUserCards = (userId: number) => getCreditCards(userId);

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
