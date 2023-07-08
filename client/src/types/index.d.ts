type Product = {
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

type Address = {
    name: string;
    street: string;
    street_number: string;
    state: string;
    country: string;
    zipcode: string;
}

type CardProps = {
    title: string,
    description: string,
    img: string,
    link: string,
    onClose?: React.MouseEventHandler, 
};