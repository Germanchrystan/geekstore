export interface ProductCard {
    id: number;
    name: string;
    price: number;
    img: string;
    discount: {
        percentage: number;
        discounterPrice: number;
        finishDate: number;
    }
}

export interface AddressCard {
    id: number;
    name: string;
    street: string;
    street_number: string;
    state: string;
    country: string;
    zipcode: string;
}

export interface PurchaseCard {
    id: number;
    status: string;
    total: number;
    address_name: string;
}

export interface CustomProductCard {
    id: number;
    status: string;
    name: string;
    img: string;
    size: string;
    price: number;
}

export interface CreditCardList {
    id: number;
    company: string;
    last_code_number: number;
}

export interface CardProps {
    title: string,
    description: string,
    img: string,
    link: string,
    onClose?: React.MouseEventHandler, 
};

export {}