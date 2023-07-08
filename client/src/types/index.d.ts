type ProductCard = {
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

type AddressCard = {
    id: number;
    name: string;
    street: string;
    street_number: string;
    state: string;
    country: string;
    zipcode: string;
}

type PurchaseCard = {
    id: number;
    status: string;
    total: number;
    address_name: string;
}

type CustomProductCard = {
    id: number;
    status: string;
    name: string;
    img: string;
    size: string;
    price: number;
}

type CreditCardList = {
    id: number;
    company: string;
    last_code_number: number;
}

type CardProps = {
    title: string,
    description: string,
    img: string,
    link: string,
    onClose?: React.MouseEventHandler, 
};